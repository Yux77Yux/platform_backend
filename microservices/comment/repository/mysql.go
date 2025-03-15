package repository

import (
	"context"
	"database/sql"
	"fmt"
	"math"
	"strings"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	generated "github.com/Yux77Yux/platform_backend/generated/comment"
	common "github.com/Yux77Yux/platform_backend/generated/common"
	"github.com/Yux77Yux/platform_backend/microservices/auth/tools"
	errMap "github.com/Yux77Yux/platform_backend/pkg/error"
)

// POST
func BatchInsert(ctx context.Context, comments []*generated.Comment) (int64, error) {
	count := len(comments)
	// 构建用于构建 IN 子句的占位符部分
	queryCommentCount := make([]string, count) // 使用切片存储占位符
	queryContentCount := make([]string, count) // 使用切片存储占位符
	for i := 0; i < count; i++ {
		queryCommentCount[i] = "(?,?,?,?,?,?)" // 每对 id 和 user_id 用 (?,?,...) 来占位
		queryContentCount[i] = "(?,?,?)"       // 每对 id 和 user_id 用 (?,?,...) 来占位
	}
	var (
		queryComment = fmt.Sprintf(`
				INSERT INTO db_comment_1.Comment (
					root,
					parent,
					dialog,
					creation_id,
					user_id,
					created_at)
				VALUES%s`, strings.Join(queryCommentCount, ","))
		queryContent = fmt.Sprintf(`
				INSERT INTO db_comment_1.CommentContent (
					comment_id,
					content,
					media)
				VALUES%s`, strings.Join(queryContentCount, ","))
		queryArea = `
		INSERT INTO db_comment_area_1.CommentArea (creation_id, total_comments)
		VALUES (?,?)
		ON DUPLICATE KEY UPDATE total_comments = total_comments + VALUES(total_comments)`
		CommentValues = make([]interface{}, 0, count*6)
		ContentValues = make([]interface{}, 0, count*3)
		creationId    = comments[0].GetCreationId()
	)

	// 格式化输入
	for _, comment := range comments {
		CommentValues = append(CommentValues,
			comment.GetRoot(),
			comment.GetParent(),
			comment.GetDialog(),
			comment.GetCreationId(),
			comment.GetUserId(),
			comment.GetCreatedAt().AsTime(),
		)
	}

	tx, err := db.BeginTransaction()
	if err != nil {
		return -1, err
	}

	// 在发生 panic 时自动回滚事务，以确保数据库的状态不会因为程序异常而不一致
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("transaction failed because %v", r)
			if errSecond := db.RollbackTransaction(tx); errSecond != nil {
				err = fmt.Errorf("%w and %w", err, errSecond)
			}
			tools.LogError("", "db recover", err)
		}
	}()

	var rowsAffected int64
	select {
	case <-ctx.Done():
		if err := db.RollbackTransaction(tx); err != nil {
			return -1, err
		}
		return -1, errMap.GetStatusError(err)
	default:
		// 执行拿到id
		ids, err := tx.ExecContext(
			ctx,
			queryComment,
			CommentValues...,
		)
		if err != nil {
			if errSecond := db.RollbackTransaction(tx); errSecond != nil {
				tools.LogError("", "db roolback", errSecond)
			}

			return -1, err
		}

		// 获取最后插入 ID 和插入的总记录数
		lastInsertID, err := ids.LastInsertId()
		if err != nil {
			if errSecond := db.RollbackTransaction(tx); errSecond != nil {
				tools.LogError("", "db roolback", errSecond)
			}

			return -1, err
		}
		rowsAffected, err = ids.RowsAffected()
		if err != nil {
			if errSecond := db.RollbackTransaction(tx); errSecond != nil {
				tools.LogError("", "db roolback", errSecond)
			}

			return -1, err
		}

		countInt64 := int64(count)
		if countInt64 != rowsAffected {
			return -1, errMap.GetStatusError(err)
		}

		// 映射 comment_id
		for i := int64(0); i < rowsAffected; i++ {
			commentID := lastInsertID + i
			ContentValues = append(ContentValues,
				commentID, comments[i].GetContent(), comments[i].GetMedia())
		}

		_, err = tx.ExecContext(
			ctx,
			queryContent,
			ContentValues...,
		)
		if err != nil {
			if errSecond := db.RollbackTransaction(tx); errSecond != nil {
				tools.LogError("", "db roolback", errMap.MapMySQLErrorToStatus(errSecond))
			}

			return -1, err
		}

		_, err = tx.ExecContext(
			ctx,
			queryArea,
			creationId,
			count,
		)
		if err != nil {
			if errSecond := db.RollbackTransaction(tx); errSecond != nil {
				tools.LogError("", "db roolback", errSecond)
			}

			return -1, err
		}

		if err = db.CommitTransaction(tx); err != nil {
			return -1, err
		}
	}

	return rowsAffected, nil
}

// GET
// (creationId, userId, error)
func GetCreationIdInTransaction(ctx context.Context, comment_id int32) (int64, int64, error) {
	const (
		query = `
			SELECT 
				creation_id,
				user_id
			FROM
				db_comment_1.Comment
			WHERE
				id = ?`
	)

	var creationId int64 = -1
	var userId int64 = -1

	select {
	case <-ctx.Done():
		return -1, -1, errMap.GetStatusError(ctx.Err())
	default:
		// 查统计
		err := db.QueryRowContext(
			ctx,
			query,
			comment_id,
		).Scan(&creationId, &userId)

		if err != nil {
			return -1, -1, errMap.MapMySQLErrorToStatus(err)
		}
	}
	return creationId, userId, nil
}

const (
	TOP_LIMIT    = 25
	SECOND_LIMIT = 10
)

// 初始化一级评论
func GetInitialTopCommentsInTransaction(ctx context.Context, creation_id int64) (*generated.CommentArea, []*generated.TopComment, int32, error) {
	const (
		LIMIT            = TOP_LIMIT
		queryTopComments = `
			SELECT count(*) 
			FROM db_comment_1.Comment 
			WHERE creation_id = ? 
			AND root = 0 
			AND status = 'PUBLISHED'`

		queryArea = `
			SELECT 
				total_comments,
				areas_status
			FROM
				db_comment_area_1.CommentArea
			WHERE
				creation_id = ?`
	)

	var (
		query = fmt.Sprintf(`
			SELECT 
    			c.id,
    			c.root,
    			c.parent,
    			c.dialog,
    			c.user_id,
    			c.created_at,
    			cc.content,
    			cc.media,
				(SELECT count(*) FROM db_comment_1.Comment b WHERE b.root = c.id AND b.status = 'PUBLISHED') AS subCount
			FROM 
    			db_comment_1.Comment c
			LEFT JOIN 
    			db_comment_1.CommentContent cc 
			ON 
				c.id = cc.comment_id
			WHERE 
    			c.creation_id = ?
			AND 
				c.root = 0
			AND 
				c.status = 'PUBLISHED'
			ORDER BY c.created_at DESC
			LIMIT %d`, LIMIT)

		total  int32  = -1
		status string = ""
		count  int32

		comments = make([]*generated.TopComment, 0, LIMIT)
	)

	select {
	case <-ctx.Done():
		return nil, nil, -1, errMap.GetStatusError(ctx.Err())
	default:
		// 查统计
		err := db.QueryRowContext(ctx,
			queryArea,
			creation_id,
		).Scan(&total, &status)
		if err != nil {
			return nil, nil, -1, errMap.MapMySQLErrorToStatus(err)
		}

		err = db.QueryRowContext(ctx,
			queryTopComments,
			creation_id,
		).Scan(&count)
		if err != nil {
			return nil, nil, -1, errMap.MapMySQLErrorToStatus(err)
		}

		// 非公开则直接返回
		if status != generated.CommentArea_DEFAULT.String() {
			return &generated.CommentArea{
				AreaStatus: generated.CommentArea_Status(generated.CommentArea_Status_value[status]),
			}, nil, -0, nil
		}

		// 查评论
		rows, err := db.QueryContext(
			ctx,
			query,
			creation_id,
		)
		if err != nil {
			return nil, nil, -1, err
		}
		defer rows.Close()

		for rows.Next() {
			var (
				id         int32
				root       int32
				parent     int32
				dialog     int32
				user_id    int64
				created_at time.Time
				content    sql.NullString
				media      sql.NullString
				subCount   int32
			)

			err = rows.Scan(&id, &root, &parent, &dialog, &user_id, &created_at, &content, &media, &subCount)
			if err != nil {
				return nil, nil, -1, errMap.MapMySQLErrorToStatus(err)
			}
			comments = append(comments, &generated.TopComment{
				Comment: &generated.Comment{
					CommentId:  id,
					Root:       root,
					Parent:     parent,
					Dialog:     dialog,
					UserId:     user_id,
					CreationId: creation_id,
					CreatedAt:  timestamppb.New(created_at),
					Content:    content.String,
					Media:      media.String,
				},
				SubCount: subCount,
			})
		}
	}

	// 计算页数返回
	pageCount := int32(math.Ceil(float64(count) / float64(LIMIT)))
	return &generated.CommentArea{
		AreaStatus:    generated.CommentArea_Status(generated.CommentArea_Status_value[status]),
		TotalComments: total,
	}, comments, pageCount, nil
}

func GetTopCommentsInTransaction(ctx context.Context, creation_id int64, pageNumber int32) ([]*generated.TopComment, error) {
	const LIMIT = TOP_LIMIT
	var (
		offset = (pageNumber - 1) * LIMIT
		query  = fmt.Sprintf(`
			SELECT 
    			c.id,
    			c.root,
    			c.parent,
    			c.dialog,
    			c.user_id,
    			c.created_at,
    			cc.content,
    			cc.media,
				(SELECT count(*) FROM db_comment_1.Comment b WHERE b.root = c.id AND b.status = 'PUBLISHED') AS subCount
			FROM 
    			db_comment_1.Comment c
			LEFT JOIN 
    			db_comment_1.CommentContent cc 
			ON 
				c.id = cc.comment_id
			WHERE 
    			c.creation_id = ?
			AND 
				c.root = 0
			AND 
				c.status = 'PUBLISHED'
			ORDER BY c.created_at DESC
			LIMIT %d 
			OFFSET ?`, LIMIT)

		comments = make([]*generated.TopComment, 0, LIMIT)
	)

	select {
	case <-ctx.Done():
		return nil, errMap.GetStatusError(ctx.Err())
	default:
		// 查评论
		rows, err := db.QueryContext(ctx,
			query,
			creation_id,
			offset,
		)
		if err != nil {
			return nil, errMap.MapMySQLErrorToStatus(err)
		}
		defer rows.Close()

		for rows.Next() {
			var (
				id         int32
				root       int32
				parent     int32
				dialog     int32
				user_id    int64
				created_at time.Time
				content    sql.NullString
				media      sql.NullString
				subCount   int32
			)

			err = rows.Scan(&id, &root, &parent, &dialog, &user_id, &created_at, &content, &media, &subCount)
			if err != nil {
				return nil, errMap.MapMySQLErrorToStatus(err)
			}
			comments = append(comments, &generated.TopComment{
				Comment: &generated.Comment{
					CommentId:  id,
					Root:       root,
					Parent:     parent,
					Dialog:     dialog,
					UserId:     user_id,
					CreationId: creation_id,
					CreatedAt:  timestamppb.New(created_at),
					Content:    content.String,
					Media:      media.String,
				},
				SubCount: subCount,
			})
		}
	}

	return comments, nil
}

func GetSecondCommentsInTransaction(ctx context.Context, creation_id int64, root, pageNumber int32) ([]*generated.SecondComment, error) {
	const LIMIT = SECOND_LIMIT
	var (
		offset = (pageNumber - 1) * LIMIT
		query  = fmt.Sprintf(`
			SELECT 
    			c.id,
    			c.root,
    			c.parent,
    			c.dialog,
    			c.user_id,
    			c.created_at,
    			cc.content,
    			cc.media,
				(SELECT b.user_id FROM db_comment_1.Comment b WHERE b.id = c.parent AND b.status = 'PUBLISHED') AS reply_user_id
			FROM 
    			db_comment_1.Comment c
			LEFT JOIN 
    			db_comment_1.CommentContent cc 
			ON 
				c.id = cc.comment_id
			WHERE 
    			c.creation_id = ?
			AND 
				c.root = ?
			AND 
				c.status = 'PUBLISHED'
			LIMIT %d 
			OFFSET ?`, LIMIT)
	)

	var (
		comments []*generated.SecondComment
	)

	select {
	case <-ctx.Done():
		return nil, errMap.GetStatusError(ctx.Err())
	default:
		// 查评论
		rows, err := db.QueryContext(ctx,
			query,
			creation_id,
			root,
			offset,
		)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {
			var (
				id            int32
				root          int32
				parent        int32
				dialog        int32
				user_id       int64
				created_at    time.Time
				content       sql.NullString
				media         sql.NullString
				reply_user_id sql.NullInt64
			)

			if err := rows.Scan(&id, &root, &parent, &dialog, &user_id, &created_at, &content, &media, &reply_user_id); err != nil {
				return nil, errMap.MapMySQLErrorToStatus(err)
			}
			comments = append(comments, &generated.SecondComment{
				Comment: &generated.Comment{
					CommentId:  id,
					Root:       root,
					Parent:     parent,
					Dialog:     dialog,
					UserId:     user_id,
					CreationId: creation_id,
					CreatedAt:  timestamppb.New(created_at),
					Content:    content.String,
					Media:      media.String,
				},
				ReplyUserId: reply_user_id.Int64,
			})
		}
	}

	return comments, nil
}

func GetReplyCommentsInTransaction(ctx context.Context, user_id int64, page int32) ([]*generated.Comment, error) {
	const LIMIT = TOP_LIMIT
	var (
		offset = (page - 1) * LIMIT
		query  = fmt.Sprintf(`
			SELECT 
    			c.id,
    			c.root,
    			c.parent,
    			c.dialog,
    			c.user_id,
    			c.created_at,
    			cc.content,
    			cc.media
			FROM 
    			db_comment_1.Comment c
			LEFT JOIN 
    			db_comment_1.CommentContent cc 
			ON
				c.id = cc.comment_id
			WHERE 
				c.user_id = ?
			ORDER BY c.created_at DESC
			LIMIT %d 
			OFFSET ?`, LIMIT)
	)

	var (
		comments []*generated.Comment
	)

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		// 查评论
		rows, err := db.QueryContext(ctx,
			query,
			user_id,
			offset,
		)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {
			var (
				id          int32
				root        int32
				parent      int32
				dialog      int32
				user_id     int64
				creation_id int64
				created_at  time.Time
				content     sql.NullString
				media       sql.NullString
			)

			err := rows.Scan(&id, &root, &parent, &dialog, &user_id, &created_at, &content, &media)
			if err != nil {
				return nil, errMap.MapMySQLErrorToStatus(err)
			}
			comments = append(comments, &generated.Comment{
				CommentId:  id,
				Root:       root,
				Parent:     parent,
				Dialog:     dialog,
				UserId:     user_id,
				CreationId: creation_id,
				CreatedAt:  timestamppb.New(created_at),
				Content:    content.String,
				Media:      media.String,
			})
		}
	}

	return comments, nil
}

func GetCommentInfo(ctx context.Context, comments []*common.AfterAuth) ([]*common.AfterAuth, error) {
	count := len(comments)
	// 构建用于构建 IN 子句的占位符部分
	queryCount := make([]string, count) // 使用切片存储占位符
	for i := 0; i < count; i++ {
		queryCount[i] = "?"
	}

	var (
		queryComment = fmt.Sprintf(`
				SELECT 
					id,
					creation_id
				FROM db_comment_1.CommentContent
				WHERE id 
				IN (%s)`, strings.Join(queryCount, ","))
		values = make([]interface{}, 0, count)
		result = make([]*common.AfterAuth, 0, count)
	)

	for i := 0; i < count; i++ {
		values[i] = comments[i]
	}

	select {
	case <-ctx.Done():
		return nil, errMap.MapMySQLErrorToStatus(ctx.Err())
	default:
		// 查评论
		rows, err := db.QueryContext(
			ctx,
			queryComment,
			values...,
		)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {
			var (
				id         int32
				creationId int64
			)

			err := rows.Scan(&id, &creationId)
			if err != nil {
				return nil, errMap.MapMySQLErrorToStatus(err)
			}
			result = append(result, &common.AfterAuth{
				CommentId:  id,
				CreationId: creationId,
			})
		}
	}

	return result, nil
}

func GetComments(ctx context.Context, ids []int32) ([]*generated.Comment, error) {
	count := len(ids)
	// 构建用于构建 IN 子句的占位符部分
	sqlStr := make([]string, count) // 使用切片存储占位符
	for i := 0; i < count; i++ {
		sqlStr[i] = "?"
	}

	values := make([]any, count)
	for i := 0; i < count; i++ {
		values[i] = ids[i]
	}

	query := fmt.Sprintf(`
			SELECT 
    			c.id,
    			c.root,
    			c.parent,
    			c.dialog,
    			c.user_id,
    			c.creation_id,
    			c.created_at,
    			cc.content,
    			cc.media
			FROM 
    			db_comment_1.Comment c
			LEFT JOIN 
    			db_comment_1.CommentContent cc 
			ON
				c.id = cc.comment_id
			WHERE c.id IN (%s)`, strings.Join(sqlStr, ","))

	comments := make([]*generated.Comment, 0, count)
	select {
	case <-ctx.Done():
		return nil, errMap.GetStatusError(ctx.Err())
	default:
		// 查评论
		rows, err := db.QueryContext(
			ctx,
			query,
			values...,
		)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {
			var (
				id          int32
				root        int32
				parent      int32
				dialog      int32
				user_id     int64
				creation_id int64
				created_at  time.Time
				content     sql.NullString
				media       sql.NullString
			)

			if err := rows.Scan(&id, &root, &parent, &dialog, &user_id, &creation_id, &created_at, &content, &media); err != nil {
				return nil, errMap.MapMySQLErrorToStatus(err)
			}
			comments = append(comments, &generated.Comment{
				CommentId:  id,
				Root:       root,
				Parent:     parent,
				Dialog:     dialog,
				UserId:     user_id,
				CreationId: creation_id,
				CreatedAt:  timestamppb.New(created_at),
				Content:    content.String,
				Media:      media.String,
			})
		}
	}

	return comments, nil
}

// UPDATE
func BatchUpdateDeleteStatus(ctx context.Context, comments []*common.AfterAuth) (int64, error) {
	count := len(comments)
	// 构建用于构建 IN 子句的占位符部分
	queryCount := make([]string, count) // 使用切片存储占位符
	for i := 0; i < count; i++ {
		queryCount[i] = "?"
	}
	var rowsAffected int64

	join := strings.Join(queryCount, ",")
	var (
		queryComment = fmt.Sprintf(`
				UPDATE db_comment_1.Comment 
				SET status = "DELETED"
				WHERE id 
				IN (%s)
				OR root 
				IN (%s)`, join, join)
		queryArea = `
				UPDATE db_comment_area_1.CommentArea 
				SET
					total_comments = total_comments - ?
				WHERE creation_id = ?`
		values     = make([]interface{}, count*2)
		creationId = comments[0].GetCreationId()
	)

	// 格式化输入
	for i := 0; i < count; i++ {
		values[i] = comments[i].GetCommentId()
		values[count+i] = comments[i].GetCommentId()
	}

	tx, err := db.BeginTransaction()
	if err != nil {
		return -1, err
	}

	// 在发生 panic 时自动回滚事务，以确保数据库的状态不会因为程序异常而不一致
	defer func() {
		if r := recover(); r != nil {
			if errSecond := db.RollbackTransaction(tx); errSecond != nil {
				err = fmt.Errorf("%w and %w", err, errSecond)
			}
			tools.LogError("", "db recover", err)
		}
	}()

	select {
	case <-ctx.Done():
		if errSecond := db.RollbackTransaction(tx); errSecond != nil {
			tools.LogError("", "db recover", errSecond)
		}
		return -1, errMap.GetStatusError(ctx.Err())
	default:
		result, err := tx.ExecContext(
			ctx,
			queryComment,
			values...,
		)
		if err != nil {
			if errSecond := db.RollbackTransaction(tx); errSecond != nil {
				tools.LogError("", "db rollback", errSecond)
			}

			return -1, err
		}

		rowsAffected, err = result.RowsAffected()
		if err != nil {
			if errSecond := db.RollbackTransaction(tx); errSecond != nil {
				tools.LogError("", "db rollback", errSecond)
			}

			return -1, errMap.MapMySQLErrorToStatus(err)
		}
		_, err = tx.ExecContext(
			ctx,
			queryArea,
			rowsAffected,
			creationId,
		)
		if err != nil {
			if errSecond := db.RollbackTransaction(tx); errSecond != nil {
				tools.LogError("", "db rollback", errSecond)
			}

			return -1, err
		}

		if err = db.CommitTransaction(tx); err != nil {
			return -1, err
		}
	}
	return rowsAffected, nil
}

package repository

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	generated "github.com/Yux77Yux/platform_backend/generated/comment"
)

// POST
func BatchInsert(comments []*generated.Comment) error {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	count := len(comments)
	// 构建用于构建 IN 子句的占位符部分
	queryCommentCount := make([]string, count) // 使用切片存储占位符
	queryContentCount := make([]string, count) // 使用切片存储占位符
	for i := 0; i < count; i++ {
		queryCommentCount[i] = "(?,?,?,?,?)" // 每对 id 和 user_id 用 (?,?,...) 来占位
		queryContentCount[i] = "(?,?,?)"     // 每对 id 和 user_id 用 (?,?,...) 来占位
	}
	var (
		queryComment = fmt.Sprintf(`
				INSERT INTO db_comment_1.comment (
					root,
					parent,
					dialog,
					creation_id,
					user_id)
				VALUES%s`, strings.Join(queryCommentCount, ","))
		queryContent = fmt.Sprintf(`
				INSERT INTO db_comment_1.comment (
					comment_id,
					content,
					media)
				VALUES%s`, strings.Join(queryContentCount, ","))
		queryArea = `
				UPDATE db_comment_areas_1.CommentAreas 
				SET
					total_comments = total_comments + ?,
				WHERE creation_id = ?`
		CommentValues = make([]interface{}, 0, count*5)
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
			comment.GetUserId())
	}

	tx, err := db.BeginTransaction()
	if err != nil {
		return err
	}

	// 在发生 panic 时自动回滚事务，以确保数据库的状态不会因为程序异常而不一致
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("transaction failed because %v", r)
			if errSecond := db.RollbackTransaction(tx); errSecond != nil {
				err = fmt.Errorf("%w and %w", err, errSecond)
			}
		}
	}()

	select {
	case <-ctx.Done():
		err = fmt.Errorf("exec timeout :%w", ctx.Err())
		if errSecond := db.RollbackTransaction(tx); errSecond != nil {
			err = fmt.Errorf("%w and %w", err, errSecond)
		}

		return err
	default:
		// 执行拿到id
		ids, err := tx.ExecContext(
			ctx,
			queryComment,
			CommentValues...,
		)
		if err != nil {
			err = fmt.Errorf("batchInsert transaction exec failed during queryComment because %v", err)
			if errSecond := db.RollbackTransaction(tx); errSecond != nil {
				err = fmt.Errorf("%w and %w", err, errSecond)
			}

			return err
		}

		// 获取最后插入 ID 和插入的总记录数
		lastInsertID, err := ids.LastInsertId()
		if err != nil {
			err = fmt.Errorf("batchInsert transaction exec failed during LastInsertId because  %v", err)
			if errSecond := db.RollbackTransaction(tx); errSecond != nil {
				err = fmt.Errorf("%w and %w", err, errSecond)
			}

			return err
		}
		rowsAffected, err := ids.RowsAffected()
		if err != nil {
			err = fmt.Errorf("batchInsert transaction exec failed during RowsAffected because  %v", err)
			if errSecond := db.RollbackTransaction(tx); errSecond != nil {
				err = fmt.Errorf("%w and %w", err, errSecond)
			}

			return err
		}

		countInt64 := int64(count)
		if countInt64 != rowsAffected {
			return fmt.Errorf("count not match the rowsAffected")
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
			err = fmt.Errorf("batchInsert transaction exec failed during queryContent because %v", err)
			if errSecond := db.RollbackTransaction(tx); errSecond != nil {
				err = fmt.Errorf("%w and %w", err, errSecond)
			}

			return err
		}

		_, err = tx.ExecContext(
			ctx,
			queryArea,
			count,
			creationId,
		)
		if err != nil {
			err = fmt.Errorf("batchInsert transaction exec failed during queryArea because %v", err)
			if errSecond := db.RollbackTransaction(tx); errSecond != nil {
				err = fmt.Errorf("%w and %w", err, errSecond)
			}

			return err
		}

		if err = db.CommitTransaction(tx); err != nil {
			return err
		}
	}

	return nil
}

// GET
func GetPublisherIdInTransaction(comment_id int32) (int64, error) {
	const (
		query = `
			SELECT 
				user_id
			FROM
				db_comments_1.Comments
			WHERE
				id = ?
		`
	)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

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
		}
	}()

	var userId int64 = -1

	select {
	case <-ctx.Done():
		err = fmt.Errorf("exec timeout :%w", ctx.Err())
		if errSecond := db.RollbackTransaction(tx); errSecond != nil {
			err = fmt.Errorf("%w and %w", err, errSecond)
		}

		return -1, err
	default:
		// 查统计
		err = tx.QueryRow(
			query,
			comment_id,
		).Scan(&userId)

		if err != nil {
			err = fmt.Errorf("queryCreation transaction exec failed because %v", err)
			if errSecond := db.RollbackTransaction(tx); errSecond != nil {
				err = fmt.Errorf("%w and %w", err, errSecond)
			}
			return -1, err
		}

		if err = db.CommitTransaction(tx); err != nil {
			return -1, err
		}
	}
	return userId, nil
}

func GetFirstCommentsInTransaction(creation_id int64) (*generated.CommentArea, []*generated.Comment, error) {
	const (
		query = `
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
    			db_comments_1.Comments c
			LEFT JOIN 
    			db_comments_1.CommentContent cc 
			ON 
				c.id = cc.comment_id
			WHERE 
    			c.creation_id = ?
			AND 
				c.root = 0
			LIMIT 50
		`

		queryArea = `
			SELECT 
				total_comments,
				areas_status
			FROM
				db_comment_areas_1.CommentAreas
			WHERE
				creation_id = ?
		`
	)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	tx, err := db.BeginTransaction()
	if err != nil {
		return nil, []*generated.Comment{}, err
	}

	// 在发生 panic 时自动回滚事务，以确保数据库的状态不会因为程序异常而不一致
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("transaction failed because %v", r)
			if errSecond := db.RollbackTransaction(tx); errSecond != nil {
				err = fmt.Errorf("%w and %w", err, errSecond)
			}
		}
	}()

	var (
		total  int32  = -1
		status string = ""

		comments []*generated.Comment
	)

	select {
	case <-ctx.Done():
		err = fmt.Errorf("exec timeout :%w", ctx.Err())
		if errSecond := db.RollbackTransaction(tx); errSecond != nil {
			err = fmt.Errorf("%w and %w", err, errSecond)
		}

		return nil, []*generated.Comment{}, err
	default:
		// 查统计
		err = tx.QueryRow(
			queryArea,
			creation_id,
		).Scan(&total, &status)
		if err != nil {
			err = fmt.Errorf("getFirstCommentInTransaction transaction exec failed during queryArea because %v", err)
			if errSecond := db.RollbackTransaction(tx); errSecond != nil {
				err = fmt.Errorf("%w and %w", err, errSecond)
			}

			return nil, []*generated.Comment{}, err
		}

		// 非公开则直接返回
		if status != "ACTIVE" {
			return &generated.CommentArea{
				AreaStatus: generated.CommentArea_Status(generated.CommentArea_Status_value[status]),
			}, []*generated.Comment{}, nil
		}

		// 查评论
		rows, err := tx.Query(
			query,
			creation_id,
		)
		if err != nil {
			err = fmt.Errorf("queryCreation transaction exec failed during Comments because %v", err)
			if errSecond := db.RollbackTransaction(tx); errSecond != nil {
				err = fmt.Errorf("%w and %w", err, errSecond)
			}

			return nil, []*generated.Comment{}, err
		}

		for rows.Next() {
			var (
				id         int32
				root       int32
				parent     int32
				dialog     int32
				user_id    int64
				created_at time.Time
				content    string
				media      string
			)

			rows.Scan(&id, &root, &parent, &dialog, &user_id, &created_at, &content, &media)
			comments = append(comments, &generated.Comment{
				CommentId:  id,
				Root:       root,
				Parent:     parent,
				Dialog:     dialog,
				UserId:     user_id,
				CreationId: creation_id,
				CreatedAt:  timestamppb.New(created_at),
				Content:    content,
				Media:      media,
			})
		}

		if err = db.CommitTransaction(tx); err != nil {
			return nil, []*generated.Comment{}, err
		}
	}

	return &generated.CommentArea{
		AreaStatus:    generated.CommentArea_Status(generated.CommentArea_Status_value[status]),
		TotalComments: total,
	}, comments, nil
}

func GetTopCommentsInTransaction(creation_id int64, pageNumber int32) ([]*generated.Comment, error) {
	offset := (pageNumber - 1) * 50
	const (
		query = `
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
    			db_comments_1.Comments c
			LEFT JOIN 
    			db_comments_1.CommentContent cc 
			ON 
				c.id = cc.comment_id
			WHERE 
    			c.creation_id = ?
			AND 
				c.root = 0
			LIMIT 50 OFFSET ?
		`
	)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	tx, err := db.BeginTransaction()
	if err != nil {
		return []*generated.Comment{}, err
	}

	// 在发生 panic 时自动回滚事务，以确保数据库的状态不会因为程序异常而不一致
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("transaction failed because %v", r)
			if errSecond := db.RollbackTransaction(tx); errSecond != nil {
				err = fmt.Errorf("%w and %w", err, errSecond)
			}
		}
	}()

	var (
		comments []*generated.Comment
	)

	select {
	case <-ctx.Done():
		err = fmt.Errorf("exec timeout :%w", ctx.Err())
		if errSecond := db.RollbackTransaction(tx); errSecond != nil {
			err = fmt.Errorf("%w and %w", err, errSecond)
		}

		return []*generated.Comment{}, err
	default:
		// 查评论
		rows, err := tx.Query(
			query,
			creation_id,
			offset,
		)
		if err != nil {
			err = fmt.Errorf("queryCreation transaction exec failed because %v", err)
			if errSecond := db.RollbackTransaction(tx); errSecond != nil {
				err = fmt.Errorf("%w and %w", err, errSecond)
			}

			return []*generated.Comment{}, err
		}

		for rows.Next() {
			var (
				id         int32
				root       int32
				parent     int32
				dialog     int32
				user_id    int64
				created_at time.Time
				content    string
				media      string
			)

			rows.Scan(&id, &root, &parent, &dialog, &user_id, &created_at, &content, &media)
			comments = append(comments, &generated.Comment{
				CommentId:  id,
				Root:       root,
				Parent:     parent,
				Dialog:     dialog,
				UserId:     user_id,
				CreationId: creation_id,
				CreatedAt:  timestamppb.New(created_at),
				Content:    content,
				Media:      media,
			})
		}

		if err = db.CommitTransaction(tx); err != nil {
			return []*generated.Comment{}, err
		}
	}

	return comments, nil
}

func GetSecondCommentsInTransaction(creation_id int64, root, pageNumber int32) ([]*generated.Comment, error) {
	offset := (pageNumber - 1) * 50
	const (
		query = `
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
    			db_comments_1.Comments c
			LEFT JOIN 
    			db_comments_1.CommentContent cc 
			ON 
				c.id = cc.comment_id
			WHERE 
    			c.creation_id = ?
			AND 
				c.root = ?
			LIMIT 10 OFFSET ?
		`
	)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	tx, err := db.BeginTransaction()
	if err != nil {
		return []*generated.Comment{}, err
	}

	// 在发生 panic 时自动回滚事务，以确保数据库的状态不会因为程序异常而不一致
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("transaction failed because %v", r)
			if errSecond := db.RollbackTransaction(tx); errSecond != nil {
				err = fmt.Errorf("%w and %w", err, errSecond)
			}
		}
	}()

	var (
		comments []*generated.Comment
	)

	select {
	case <-ctx.Done():
		err = fmt.Errorf("exec timeout :%w", ctx.Err())
		if errSecond := db.RollbackTransaction(tx); errSecond != nil {
			err = fmt.Errorf("%w and %w", err, errSecond)
		}

		return []*generated.Comment{}, err
	default:
		// 查评论
		rows, err := tx.Query(
			query,
			creation_id,
			root,
			offset,
		)
		if err != nil {
			err = fmt.Errorf("queryCreation transaction exec failed because %v", err)
			if errSecond := db.RollbackTransaction(tx); errSecond != nil {
				err = fmt.Errorf("%w and %w", err, errSecond)
			}

			return []*generated.Comment{}, err
		}

		for rows.Next() {
			var (
				id         int32
				root       int32
				parent     int32
				dialog     int32
				user_id    int64
				created_at time.Time
				content    string
				media      string
			)

			rows.Scan(&id, &root, &parent, &dialog, &user_id, &created_at, &content, &media)
			comments = append(comments, &generated.Comment{
				CommentId:  id,
				Root:       root,
				Parent:     parent,
				Dialog:     dialog,
				UserId:     user_id,
				CreationId: creation_id,
				CreatedAt:  timestamppb.New(created_at),
				Content:    content,
				Media:      media,
			})
		}

		if err = db.CommitTransaction(tx); err != nil {
			return []*generated.Comment{}, err
		}
	}

	return comments, nil
}

func GetReplyCommentsInTransaction(user_id int64, page int32) ([]*generated.Comment, error) {
	var (
		offset = (page - 1) * 50
		query  = `
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
    			db_comments_1.Comments c
			LEFT JOIN 
    			db_comments_1.CommentContent cc 
			ON
				c.id = cc.comment_id
			WHERE 
				c.user_id = ?
			ORDER BY c.created_at DESC
			LIMIT 50 
			OFFSET ?`
	)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	tx, err := db.BeginTransaction()
	if err != nil {
		return []*generated.Comment{}, err
	}

	// 在发生 panic 时自动回滚事务，以确保数据库的状态不会因为程序异常而不一致
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("transaction failed because %v", r)
			if errSecond := db.RollbackTransaction(tx); errSecond != nil {
				err = fmt.Errorf("%w and %w", err, errSecond)
			}
		}
	}()

	var (
		comments []*generated.Comment
	)

	select {
	case <-ctx.Done():
		err = fmt.Errorf("exec timeout :%w", ctx.Err())
		if errSecond := db.RollbackTransaction(tx); errSecond != nil {
			err = fmt.Errorf("%w and %w", err, errSecond)
		}

		return []*generated.Comment{}, err
	default:
		// 查评论
		rows, err := tx.Query(
			query,
			user_id,
			offset,
		)
		if err != nil {
			err = fmt.Errorf("queryCreation transaction exec failed because %v", err)
			if errSecond := db.RollbackTransaction(tx); errSecond != nil {
				err = fmt.Errorf("%w and %w", err, errSecond)
			}

			return []*generated.Comment{}, err
		}

		for rows.Next() {
			var (
				id          int32
				root        int32
				parent      int32
				dialog      int32
				user_id     int64
				creation_id int64
				created_at  time.Time
				content     string
				media       string
			)

			rows.Scan(&id, &root, &parent, &dialog, &user_id, &created_at, &content, &media)
			comments = append(comments, &generated.Comment{
				CommentId:  id,
				Root:       root,
				Parent:     parent,
				Dialog:     dialog,
				UserId:     user_id,
				CreationId: creation_id,
				CreatedAt:  timestamppb.New(created_at),
				Content:    content,
				Media:      media,
			})
		}

		if err = db.CommitTransaction(tx); err != nil {
			return []*generated.Comment{}, err
		}
	}

	return comments, nil
}

func GetCommentInfo(comments []*generated.AfterAuth) ([]*generated.AfterAuth, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	count := len(comments)
	// 构建用于构建 IN 子句的占位符部分
	queryCount := make([]string, count) // 使用切片存储占位符
	for i := 0; i < count; i++ {
		queryCount[i] = "(?,?)"
	}

	var (
		queryComment = fmt.Sprintf(`
				SELECT 
					id,
					creation_id
				FROM db_comment_1.comment
				WHERE (id,user_id) 
				IN (%s)`, strings.Join(queryCount, ","))
		values = make([]interface{}, 0, count*2)
		result = make([]*generated.AfterAuth, 0, count)
	)

	for _, comment := range comments {
		values = append(values, comment.GetCommentId(), comment.GetUserId())
	}

	tx, err := db.BeginTransaction()
	if err != nil {
		return []*generated.AfterAuth{}, err
	}

	// 在发生 panic 时自动回滚事务，以确保数据库的状态不会因为程序异常而不一致
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("transaction failed because %v", r)
			if errSecond := db.RollbackTransaction(tx); errSecond != nil {
				err = fmt.Errorf("%w and %w", err, errSecond)
			}
		}
	}()

	select {
	case <-ctx.Done():
		err = fmt.Errorf("exec timeout :%w", ctx.Err())
		if errSecond := db.RollbackTransaction(tx); errSecond != nil {
			err = fmt.Errorf("%w and %w", err, errSecond)
		}

		return []*generated.AfterAuth{}, err
	default:
		// 查评论
		rows, err := tx.QueryContext(
			ctx,
			queryComment,
			values...,
		)
		if err != nil {
			err = fmt.Errorf("queryCreation transaction exec failed because %v", err)
			if errSecond := db.RollbackTransaction(tx); errSecond != nil {
				err = fmt.Errorf("%w and %w", err, errSecond)
			}

			return []*generated.AfterAuth{}, err
		}

		for rows.Next() {
			var (
				id         int32
				creationId int64
			)

			rows.Scan(&id, &creationId)
			result = append(result, &generated.AfterAuth{
				CommentId:  id,
				CreationId: creationId,
			})
		}

		if err = db.CommitTransaction(tx); err != nil {
			return []*generated.AfterAuth{}, err
		}
	}

	return result, nil
}

// UPDATE
func BatchUpdate(comments []*generated.AfterAuth) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	count := len(comments)
	// 构建用于构建 IN 子句的占位符部分
	queryCount := make([]string, count) // 使用切片存储占位符
	for i := 0; i < count; i++ {
		queryCount[i] = "?"
	}

	var (
		queryComment = fmt.Sprintf(`
				UPDATE db_comment_1.comment 
				SET status = "DELETE"
				WHERE id 
				IN (%s)`, strings.Join(queryCount, ","))
		queryArea = `
				UPDATE db_comment_areas_1.CommentAreas 
				SET
					total_comments = total_comments - ?,
				WHERE creation_id = ?`
		values     = make([]interface{}, 0, count)
		creationId = comments[0].GetCreationId()
	)

	// 格式化输入
	for _, comment := range comments {
		values = append(values, comment.GetCommentId())
	}

	tx, err := db.BeginTransaction()
	if err != nil {
		return err
	}

	// 在发生 panic 时自动回滚事务，以确保数据库的状态不会因为程序异常而不一致
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("transaction failed because %v", r)
			if errSecond := db.RollbackTransaction(tx); errSecond != nil {
				err = fmt.Errorf("%w and %w", err, errSecond)
			}
		}
	}()

	select {
	case <-ctx.Done():
		err = fmt.Errorf("exec timeout :%w", ctx.Err())
		if errSecond := db.RollbackTransaction(tx); errSecond != nil {
			err = fmt.Errorf("%w and %w", err, errSecond)
		}

		return err
	default:
		_, err := tx.ExecContext(
			ctx,
			queryComment,
			values...,
		)
		if err != nil {
			err = fmt.Errorf("batchInsert transaction exec failed during queryComment because %v", err)
			if errSecond := db.RollbackTransaction(tx); errSecond != nil {
				err = fmt.Errorf("%w and %w", err, errSecond)
			}

			return err
		}
		_, err = tx.ExecContext(
			ctx,
			queryArea,
			count,
			creationId,
		)
		if err != nil {
			err = fmt.Errorf("batchInsert transaction exec failed during queryArea because %v", err)
			if errSecond := db.RollbackTransaction(tx); errSecond != nil {
				err = fmt.Errorf("%w and %w", err, errSecond)
			}

			return err
		}

		if err = db.CommitTransaction(tx); err != nil {
			return err
		}
	}
	return nil
}

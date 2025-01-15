package repository

import (
	"context"
	"fmt"
	"strconv"
	"time"

	generated "github.com/Yux77Yux/platform_backend/generated/comment"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// POST
func BatchInsert(values map[string]string) error {
	ctx := context.Background()
	var (
		queryComment        = values["queryComment"]
		queryCommentContent = values["queryCommentContent"]
		count               = values["count"]
	)
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
		ids, err := tx.Exec(
			queryComment,
		)
		if err != nil {
			err = fmt.Errorf("batchInsert transaction exec failed because %v", err)
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

		countInt64, err := strconv.ParseInt(count, 10, 64)
		if err != nil {
			return fmt.Errorf("count is not a number")
		}
		if countInt64 != rowsAffected {
			return fmt.Errorf("count not match the rowsAffected")
		}
		idsSlice := make([]int64, 0, countInt64)

		// 映射 comment_id
		for i := int64(0); i < rowsAffected; i++ {
			commentID := lastInsertID + i
			idsSlice = append(idsSlice, commentID)
		}

		_, err = tx.Exec(
			queryCommentContent,
			idsSlice,
		)
		if err != nil {
			err = fmt.Errorf("batchInsert transaction exec failed because %v", err)
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

func GetFirstCommentInTransaction(creation_id int64) (*generated.CommentArea, []*generated.Comment, error) {
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

func GetTopCommentInTransaction(creation_id int64, pageNumber int32) ([]*generated.Comment, error) {
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

func GetSecondCommentInTransaction(creation_id int64, root, pageNumber int32) ([]*generated.Comment, error) {
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

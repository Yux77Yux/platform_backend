package repository

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	generated "github.com/Yux77Yux/platform_backend/generated/creation"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// POST
func CreationAddInTransaction(creation *generated.Creation) error {
	queryCreation := `insert into db_creation_1.Creation 
	(id,
	author_id,
	src,
	thumbnail,
	title,
	bio,
	status,
	duration,
	category_id,
	upload_time
	)
	values(?,?,?,?,?,?,?,?,?,?)`

	// queryCreationEngagement := `insert into db_creation_engagment_1.CreationEngagement
	// (creation_id
	// )
	// values(?)`

	ctx := context.Background()

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

	var (
		id          = creation.GetCreationId()
		author_id   = creation.GetBaseInfo().GetAuthorId()
		src         = creation.GetBaseInfo().GetSrc()
		thumbnail   = creation.GetBaseInfo().GetThumbnail()
		title       = creation.GetBaseInfo().GetTitle()
		bio         = creation.GetBaseInfo().GetBio()
		status      = creation.GetBaseInfo().GetStatus().String()
		duration    = creation.GetBaseInfo().GetDuration()
		category_id = creation.GetBaseInfo().GetCategoryId()
		upload_time = creation.GetUploadTime().AsTime()
	)

	select {
	case <-ctx.Done():
		err = fmt.Errorf("exec timeout :%w", ctx.Err())
		if errSecond := db.RollbackTransaction(tx); errSecond != nil {
			err = fmt.Errorf("%w and %w", err, errSecond)
		}

		return err
	default:
		_, err := tx.Exec(
			queryCreation,
			id,
			author_id,
			src,
			thumbnail,
			title,
			bio,
			status,
			duration,
			category_id,
			upload_time,
		)
		if err != nil {
			err = fmt.Errorf("queryCreation transaction exec failed because %v", err)
			if errSecond := db.RollbackTransaction(tx); errSecond != nil {
				err = fmt.Errorf("%w and %w", err, errSecond)
			}

			return err
		}

		// _, err = tx.Exec(
		// 	queryCreationEngagement,
		// 	id,
		// )
		// if err != nil {
		// 	err = fmt.Errorf("queryCreationEngagement transaction exec failed because %v", err)
		// 	if errSecond := db.RollbackTransaction(tx); errSecond != nil {
		// 		err = fmt.Errorf("%w and %w", err, errSecond)
		// 	}

		// 	return err
		// }

		if err = db.CommitTransaction(tx); err != nil {
			return err
		}
	}

	return nil
}

// GET

// 详细页
func GetDetailInTransaction(creationId int64) (*generated.CreationInfo, error) {
	queryCreation :=
		`SELECT
		author_id,
		src,
		thumbnail,
		title,
		bio,
		status,
		duration,
		category_id,
		upload_time
	FROM db_creation_1.Creation 
	WHERE id = ?
	`

	queryCategory := `SELECT
		parent,
		name,
		description
	FROM db_creation_category_1.Category 
	WHERE id = ?
	`

	queryCreationEngagement := `SELECT
		views,
		likes,
		saves,
		publish_time
	FROM db_creation_1.Creation 
	WHERE creation_id = ?
	`

	ctx := context.Background()

	tx, err := db.BeginTransaction()
	if err != nil {
		return nil, err
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
		author_id   int64
		src         string
		thumbnail   string
		title       string
		bio         string
		status      string
		duration    int32
		category_id int32
		upload_time time.Time

		views        int32
		likes        int32
		saves        int32
		publish_time time.Time

		parent      int32
		name        string
		description string
	)

	select {
	case <-ctx.Done():
		err = fmt.Errorf("exec timeout :%w", ctx.Err())
		if errSecond := db.RollbackTransaction(tx); errSecond != nil {
			err = fmt.Errorf("%w and %w", err, errSecond)
		}

		return nil, err
	default:
		// 查作品信息
		err := tx.QueryRow(
			queryCreation,
			creationId,
		).Scan(
			// 字段读取
			&author_id,
			&src,
			&thumbnail,
			&title,
			&bio,
			&status,
			&duration,
			&category_id,
			&upload_time,
		)
		if err != nil {
			err = fmt.Errorf("queryCreation transaction exec failed because %v", err)
			if errSecond := db.RollbackTransaction(tx); errSecond != nil {
				err = fmt.Errorf("%w and %w", err, errSecond)
			}

			return nil, err
		}

		// 查 统计数
		err = tx.QueryRow(
			queryCreationEngagement,
			creationId,
		).Scan(
			// 字段读取
			&views,
			&likes,
			&saves,
			&publish_time,
		)
		if err != nil {
			err = fmt.Errorf("queryCreationEngagement transaction exec failed because %v", err)
			if errSecond := db.RollbackTransaction(tx); errSecond != nil {
				err = fmt.Errorf("%w and %w", err, errSecond)
			}

			return nil, err
		}

		// 查 分区
		err = tx.QueryRow(
			queryCategory,
			category_id,
		).Scan(
			// 字段读取
			&parent,
			&name,
			&description,
		)
		if err != nil {
			err = fmt.Errorf("queryCreationEngagement transaction exec failed because %v", err)
			if errSecond := db.RollbackTransaction(tx); errSecond != nil {
				err = fmt.Errorf("%w and %w", err, errSecond)
			}

			return nil, err
		}

		if err = db.CommitTransaction(tx); err != nil {
			return nil, err
		}
	}

	// 统合
	creation := &generated.CreationInfo{
		Creation: &generated.Creation{
			CreationId: creationId,
			BaseInfo: &generated.CreationUpload{
				AuthorId:   author_id,
				Src:        src,
				Thumbnail:  thumbnail,
				Title:      title,
				Bio:        bio,
				Status:     generated.CreationStatus(generated.CreationStatus_value[status]),
				CategoryId: category_id,
				Duration:   duration,
			},
			UploadTime: timestamppb.New(upload_time),
		},
		CreationEngagement: &generated.CreationEngagement{
			Views:       views,
			Likes:       likes,
			Saves:       saves,
			PublishTime: timestamppb.New(publish_time),
		},
		Category: &generated.Category{
			CategoryId:  category_id,
			Parent:      parent,
			Name:        name,
			Description: description,
		},
	}

	return creation, nil
}

// 返回作者ID
func GetAuthorIdInTransaction(creationId int64) (int64, error) {
	queryCreation := `
		SELECT author_id
		FROM db_creation_1.Creation 
		WHERE id = ?
	`

	ctx := context.Background()

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

	var (
		author_id int64
	)

	select {
	case <-ctx.Done():
		err = fmt.Errorf("exec timeout :%w", ctx.Err())
		if errSecond := db.RollbackTransaction(tx); errSecond != nil {
			err = fmt.Errorf("%w and %w", err, errSecond)
		}

		return -1, err
	default:
		// 查作品信息
		err := tx.QueryRow(
			queryCreation,
			creationId,
		).Scan(
			// 字段读取
			&author_id,
		)
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
	return author_id, nil
}

// 卡片型

func GetCardInTransaction(ids []int64) ([]*generated.CreationInfo, error) {
	const (
		queryCardCategory = `SELECT
			parent,
			name
		FROM db_creation_category_1.Category 
		WHERE id IN (?)
		`

		queryCardEngagement = `
		SELECT
			views,
			publish_time
		FROM db_creation_1.Creation 
		WHERE creation_id IN (?)
		`

		// 主页,相似列表,分区
		query = `SELECT
			id,
			author_id,
			src,
			thumbnail,
			title,
			duration,
			category_id,
			upload_time
		FROM db_creation_1.Creation 
		WHERE id IN (?)
		`
	)

	ctx := context.Background()

	tx, err := db.BeginTransaction()
	if err != nil {
		return nil, err
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

	// 返回的卡片信息
	cards := make([]*generated.Creation, 0, len(ids))
	// 返回的卡片统计信息
	creationEngagements := make([]*generated.CreationEngagement, 0, len(ids))
	// 返回的卡片分区信息
	categories := make([]*generated.Category, 0, len(ids))

	select {
	case <-ctx.Done():
		err = fmt.Errorf("exec timeout :%w", ctx.Err())
		if errSecond := db.RollbackTransaction(tx); errSecond != nil {
			err = fmt.Errorf("%w and %w", err, errSecond)
		}

		return nil, err
	default:
		// 查作品信息
		// []int64 转 []string
		strIDs := make([]string, len(ids))
		for i, id := range ids {
			strIDs[i] = strconv.FormatInt(id, 10)
		}
		// 拼接
		str := strings.Join(strIDs, ",")

		// 查询的分区id
		categoriesIds := make([]int32, len(ids))

		rows, err := tx.QueryContext(ctx,
			query,
			str,
		)
		if err != nil {
			err = fmt.Errorf("queryCreation transaction exec failed because %v", err)
			if errSecond := db.RollbackTransaction(tx); errSecond != nil {
				err = fmt.Errorf("%w and %w", err, errSecond)
			}

			return nil, err
		}

		for rows.Next() {
			var (
				author_id   int64
				src         string
				thumbnail   string
				title       string
				status      string
				duration    int32
				category_id int32
				upload_time time.Time
			)
			// 从当前行读取值，依次填充到变量中
			err := rows.Scan(&author_id, &src, &thumbnail, &title, &status, &duration, &category_id, &upload_time)
			if err != nil {
				// 如果读取行数据失败，处理错误
				err = fmt.Errorf("failed to scan row because %v", err)
				if errSecond := db.RollbackTransaction(tx); errSecond != nil {
					err = fmt.Errorf("%w and %w", err, errSecond)
				}
				return nil, err
			}

			// 存储 卡片基本信息切片
			creation := &generated.Creation{
				BaseInfo: &generated.CreationUpload{
					AuthorId:   author_id,
					Src:        src,
					Thumbnail:  thumbnail,
					Title:      title,
					Status:     generated.CreationStatus(generated.CreationStatus_value[status]),
					Duration:   duration,
					CategoryId: category_id,
				},
				UploadTime: timestamppb.New(upload_time),
			}
			cards = append(cards, creation)

			// 存储 分区id切片
			categoriesIds = append(categoriesIds, category_id)
		}

		// 检查是否有额外的错误（比如数据读取完成后的关闭错误）
		if err = rows.Err(); err != nil {
			err = fmt.Errorf("rows iteration failed because %v", err)
			if errSecond := db.RollbackTransaction(tx); errSecond != nil {
				err = fmt.Errorf("%w and %w", err, errSecond)
			}
			return nil, err
		}

		// 结束第一次查询
		rows.Close()

		// 查 统计数
		rows, err = tx.QueryContext(ctx,
			queryCardEngagement,
			str,
		)
		if err != nil {
			err = fmt.Errorf("queryCreation transaction exec failed because %v", err)
			if errSecond := db.RollbackTransaction(tx); errSecond != nil {
				err = fmt.Errorf("%w and %w", err, errSecond)
			}

			return nil, err
		}

		for rows.Next() {
			var (
				views        int32
				publish_time time.Time
			)
			// 从当前行读取值，依次填充到变量中
			err := rows.Scan(&views, &publish_time)
			if err != nil {
				// 如果读取行数据失败，处理错误
				err = fmt.Errorf("failed to scan row because %v", err)
				if errSecond := db.RollbackTransaction(tx); errSecond != nil {
					err = fmt.Errorf("%w and %w", err, errSecond)
				}
				return nil, err
			}

			// 存储 作品卡片的统计信息
			creationEngagement := &generated.CreationEngagement{
				Views:       views,
				PublishTime: timestamppb.New(publish_time),
			}
			creationEngagements = append(creationEngagements, creationEngagement)
		}

		// 检查是否有额外的错误（比如数据读取完成后的关闭错误）
		if err = rows.Err(); err != nil {
			err = fmt.Errorf("rows iteration failed because %v", err)
			if errSecond := db.RollbackTransaction(tx); errSecond != nil {
				err = fmt.Errorf("%w and %w", err, errSecond)
			}
			return nil, err
		}

		// 结束第二次查询
		rows.Close()

		// 查 分区信息
		// []int32 转 []string
		categorys := make([]string, len(categoriesIds))
		for i, id := range categoriesIds {
			categorys[i] = strconv.Itoa(int(id))
		}
		// 拼接
		str = strings.Join(categorys, ",")

		rows, err = tx.QueryContext(ctx,
			queryCardCategory,
			str,
		)
		if err != nil {
			err = fmt.Errorf("queryCreation transaction exec failed because %v", err)
			if errSecond := db.RollbackTransaction(tx); errSecond != nil {
				err = fmt.Errorf("%w and %w", err, errSecond)
			}

			return nil, err
		}

		for rows.Next() {
			var (
				parent int32
				name   string
			)
			// 从当前行读取值，依次填充到变量中
			err := rows.Scan(&parent, &name)
			if err != nil {
				// 如果读取行数据失败，处理错误
				err = fmt.Errorf("failed to scan row because %v", err)
				if errSecond := db.RollbackTransaction(tx); errSecond != nil {
					err = fmt.Errorf("%w and %w", err, errSecond)
				}
				return nil, err
			}

			// 存储 作品卡片的统计信息
			category := &generated.Category{
				Parent: parent,
				Name:   name,
			}
			categories = append(categories, category)
		}

		// 检查是否有额外的错误（比如数据读取完成后的关闭错误）
		if err = rows.Err(); err != nil {
			err = fmt.Errorf("rows iteration failed because %v", err)
			if errSecond := db.RollbackTransaction(tx); errSecond != nil {
				err = fmt.Errorf("%w and %w", err, errSecond)
			}
			return nil, err
		}

		// 结束第三次查询
		rows.Close()

		if err = db.CommitTransaction(tx); err != nil {
			return nil, err
		}
	}

	creationInfos := make([]*generated.CreationInfo, 0, len(ids))
	// 统合
	for i := 0; i < len(ids); i = i + 1 {
		creationInfo := &generated.CreationInfo{
			Creation:           cards[i],
			CreationEngagement: creationEngagements[i],
			Category:           categories[i],
		}
		creationInfo.Creation.CreationId = ids[i]
		creationInfos = append(creationInfos, creationInfo)
	}
	return creationInfos, nil
}

// DELETE
func DeleteCreationInTransaction(id int64) error {
	query := `DELETE FROM db_creation_1.Creation 
		WHERE id = ?
	`

	ctx := context.Background()

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
		_, err := tx.Exec(
			query,
			id,
		)
		if err != nil {
			err = fmt.Errorf("delete creation in transaction exec failed because %v", err)
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

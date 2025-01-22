package repository

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	generated "github.com/Yux77Yux/platform_backend/generated/interaction"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// POST
func interactionAddInTransaction(interaction *generated.interaction) error {
	queryinteraction := `insert into db_interaction_1.interaction 
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

	// queryinteractionEngagement := `insert into db_interaction_engagment_1.interactionEngagement
	// (interaction_id
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
		id          = interaction.GetinteractionId()
		author_id   = interaction.GetBaseInfo().GetAuthorId()
		src         = interaction.GetBaseInfo().GetSrc()
		thumbnail   = interaction.GetBaseInfo().GetThumbnail()
		title       = interaction.GetBaseInfo().GetTitle()
		bio         = interaction.GetBaseInfo().GetBio()
		status      = interaction.GetBaseInfo().GetStatus().String()
		duration    = interaction.GetBaseInfo().GetDuration()
		category_id = interaction.GetBaseInfo().GetCategoryId()
		upload_time = interaction.GetUploadTime().AsTime()
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
			queryinteraction,
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
			err = fmt.Errorf("queryinteraction transaction exec failed because %v", err)
			if errSecond := db.RollbackTransaction(tx); errSecond != nil {
				err = fmt.Errorf("%w and %w", err, errSecond)
			}

			return err
		}

		// _, err = tx.Exec(
		// 	queryinteractionEngagement,
		// 	id,
		// )
		// if err != nil {
		// 	err = fmt.Errorf("queryinteractionEngagement transaction exec failed because %v", err)
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
func GetDetailInTransaction(interactionId int64) (*generated.interactionInfo, error) {
	queryinteraction :=
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
	FROM db_interaction_1.interaction 
	WHERE id = ?
	`

	queryCategory := `SELECT
		parent,
		name,
		description
	FROM db_interaction_category_1.Category 
	WHERE id = ?
	`

	queryinteractionEngagement := `SELECT
		views,
		likes,
		saves,
		publish_time
	FROM db_interaction_1.interaction 
	WHERE interaction_id = ?
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
			queryinteraction,
			interactionId,
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
			err = fmt.Errorf("queryinteraction transaction exec failed because %v", err)
			if errSecond := db.RollbackTransaction(tx); errSecond != nil {
				err = fmt.Errorf("%w and %w", err, errSecond)
			}

			return nil, err
		}

		// 查 统计数
		err = tx.QueryRow(
			queryinteractionEngagement,
			interactionId,
		).Scan(
			// 字段读取
			&views,
			&likes,
			&saves,
			&publish_time,
		)
		if err != nil {
			err = fmt.Errorf("queryinteractionEngagement transaction exec failed because %v", err)
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
			err = fmt.Errorf("queryinteractionEngagement transaction exec failed because %v", err)
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
	interaction := &generated.interactionInfo{
		interaction: &generated.interaction{
			interactionId: interactionId,
			BaseInfo: &generated.interactionUpload{
				AuthorId:   author_id,
				Src:        src,
				Thumbnail:  thumbnail,
				Title:      title,
				Bio:        bio,
				Status:     generated.interactionStatus(generated.interactionStatus_value[status]),
				CategoryId: category_id,
				Duration:   duration,
			},
			UploadTime: timestamppb.New(upload_time),
		},
		interactionEngagement: &generated.interactionEngagement{
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

	return interaction, nil
}

// 卡片型

func GetCardInTransaction(ids []int64) ([]*generated.interactionInfo, error) {
	const (
		queryCardCategory = `SELECT
			parent,
			name
		FROM db_interaction_category_1.Category 
		WHERE id IN (?)
		`

		queryCardEngagement = `
		SELECT
			views,
			publish_time
		FROM db_interaction_1.interaction 
		WHERE interaction_id IN (?)
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
		FROM db_interaction_1.interaction 
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
	cards := make([]*generated.interaction, 0, len(ids))
	// 返回的卡片统计信息
	interactionEngagements := make([]*generated.interactionEngagement, 0, len(ids))
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

		rows, err := tx.Query(
			query,
			str,
		)
		if err != nil {
			err = fmt.Errorf("queryinteraction transaction exec failed because %v", err)
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
			interaction := &generated.interaction{
				BaseInfo: &generated.interactionUpload{
					AuthorId:   author_id,
					Src:        src,
					Thumbnail:  thumbnail,
					Title:      title,
					Status:     generated.interactionStatus(generated.interactionStatus_value[status]),
					Duration:   duration,
					CategoryId: category_id,
				},
				UploadTime: timestamppb.New(upload_time),
			}
			cards = append(cards, interaction)

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
		rows, err = tx.Query(
			queryCardEngagement,
			str,
		)
		if err != nil {
			err = fmt.Errorf("queryinteraction transaction exec failed because %v", err)
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
			interactionEngagement := &generated.interactionEngagement{
				Views:       views,
				PublishTime: timestamppb.New(publish_time),
			}
			interactionEngagements = append(interactionEngagements, interactionEngagement)
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

		rows, err = tx.Query(
			queryCardCategory,
			str,
		)
		if err != nil {
			err = fmt.Errorf("queryinteraction transaction exec failed because %v", err)
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

	interactionInfos := make([]*generated.interactionInfo, 0, len(ids))
	// 统合
	for i := 0; i < len(ids); i = i + 1 {
		interactionInfo := &generated.interactionInfo{
			interaction:           cards[i],
			interactionEngagement: interactionEngagements[i],
			Category:              categories[i],
		}
		interactionInfo.interaction.interactionId = ids[i]
		interactionInfos = append(interactionInfos, interactionInfo)
	}
	return interactionInfos, nil
}

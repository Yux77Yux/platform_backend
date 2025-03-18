package repository

import (
	"context"
	"database/sql"
	"fmt"
	"math"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	grpcStatus "google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	generated "github.com/Yux77Yux/platform_backend/generated/creation"
	tools "github.com/Yux77Yux/platform_backend/microservices/creation/tools"
	errMap "github.com/Yux77Yux/platform_backend/pkg/error"
)

type SqlMethodStruct struct {
	db SqlInterface
}

// POST
func (c *SqlMethodStruct) CreationAddInTransaction(ctx context.Context, creation *generated.Creation) error {
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

	queryCreationEngagement := `insert into db_creation_engagment_1.CreationEngagement
	(creation_id
	)
	values(?)`

	tx, err := c.db.BeginTransaction()
	if err != nil {
		return err
	}

	// 在发生 panic 时自动回滚事务，以确保数据库的状态不会因为程序异常而不一致
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("transaction failed because %v", r)
			if errSecond := c.db.RollbackTransaction(tx); errSecond != nil {
				err = fmt.Errorf("%w and %w", err, errSecond)
			}
			tools.LogError("", "c.db recover", err)
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
		if err := c.db.RollbackTransaction(tx); err != nil {
			return err
		}
		return errMap.GetStatusError(err)
	default:
		_, err := tx.ExecContext(
			ctx,
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
			if errSecond := c.db.RollbackTransaction(tx); errSecond != nil {
				tools.LogError("", "c.db roolback", errSecond)
			}

			return err
		}

		_, err = tx.Exec(
			queryCreationEngagement,
			id,
		)
		if err != nil {
			if errSecond := c.db.RollbackTransaction(tx); errSecond != nil {
				tools.LogError("", "c.db roolback", errSecond)
			}

			return err
		}

		if err = c.db.CommitTransaction(tx); err != nil {
			return err
		}
	}

	return nil
}

// GET

// 详细页
func (c *SqlMethodStruct) GetDetailInTransaction(ctx context.Context, creationId int64) (*generated.CreationInfo, error) {
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
	WHERE id = ?`

	queryCategory := `SELECT
		parent,
		name,
		description
	FROM db_creation_category_1.Category 
	WHERE id = ?`

	queryCreationEngagement := `SELECT
		views,
		likes,
		saves,
		publish_time
	FROM db_creation_engagment_1.CreationEngagement 
	WHERE creation_id = ?`

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
		publish_time sql.NullTime

		parent      int32
		name        string
		description string
	)

	select {
	case <-ctx.Done():
		return nil, errMap.GetStatusError(ctx.Err())
	default:
		// 查作品信息
		err := c.db.QueryRowContext(
			ctx,
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
			return nil, errMap.MapMySQLErrorToStatus(err)
		}

		// 查 统计数
		err = c.db.QueryRowContext(
			ctx,
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
			return nil, errMap.MapMySQLErrorToStatus(err)
		}

		// 查 分区
		err = c.db.QueryRowContext(
			ctx,
			queryCategory,
			category_id,
		).Scan(
			// 字段读取
			&parent,
			&name,
			&description,
		)
		if err != nil {
			return nil, errMap.MapMySQLErrorToStatus(err)
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
			PublishTime: timestamppb.New(publish_time.Time),
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
func (c *SqlMethodStruct) GetAuthorIdInTransaction(ctx context.Context, creationId int64) (int64, error) {
	queryCreation := `
		SELECT author_id
		FROM db_creation_1.Creation 
		WHERE id = ?`

	var (
		author_id int64
	)

	select {
	case <-ctx.Done():
		return -1, errMap.GetStatusError(ctx.Err())
	default:
		// 查作品信息
		err := c.db.QueryRowContext(
			ctx,
			queryCreation,
			creationId,
		).Scan(
			// 字段读取
			&author_id,
		)
		if err != nil {
			return 0, errMap.MapMySQLErrorToStatus(err)
		}
	}
	return author_id, nil
}

// Card型
func (c *SqlMethodStruct) GetUserCreations(ctx context.Context, req *generated.GetUserCreationsRequest) ([]*generated.CreationInfo, int32, error) {
	const LIMIT = 10
	authorId := req.GetUserId()
	page := req.GetPage()
	if page == 0 {
		page = 1
	}
	offset := (page - 1) * LIMIT
	status := req.GetStatus().String()

	// 主页,相似列表,分区
	query := `SELECT
			id,
			src,
			thumbnail,
			title,
			duration,
			upload_time
		FROM db_creation_1.Creation 
		WHERE author_id = ?
		AND status = ?
		ORDER BY upload_time,id DESC
		LIMIT ?
		OFFSET ?`

	queryCount := `SELECT
		count(*)
	FROM db_creation_1.Creation 
	WHERE author_id = ?
	AND status = ?`

	sqlStr := make([]string, 0, LIMIT)
	creationInfos := make([]*generated.CreationInfo, 0, 10)
	creationIds := make([]any, 0, 10)
	var count int32
	select {
	case <-ctx.Done():
		return nil, -1, errMap.GetStatusError(ctx.Err())
	default:
		if page <= 1 {
			var num int32
			err := c.db.QueryRowContext(
				ctx,
				queryCount,
				authorId,
				status,
			).Scan(&num)
			if err != nil {
				return nil, -1, errMap.MapMySQLErrorToStatus(err)
			}
			if num <= 0 {
				return nil, 0, nil
			}
			count = int32(math.Ceil(float64(num) / float64(LIMIT)))
		}

		rows, err := c.db.QueryContext(
			ctx,
			query,
			authorId,
			status,
			LIMIT,
			offset,
		)
		if err != nil {
			return nil, -1, err
		}
		defer rows.Close()

		for rows.Next() {
			var (
				creationId  int64
				src         string
				thumbnail   string
				title       string
				duration    int32
				upload_time time.Time
			)
			// 从当前行读取值，依次填充到变量中
			err := rows.Scan(&creationId, &src, &thumbnail, &title, &duration, &upload_time)
			if err != nil {
				return nil, -1, errMap.MapMySQLErrorToStatus(err)
			}

			// 存储 卡片基本信息切片
			creationInfo := &generated.CreationInfo{
				Creation: &generated.Creation{
					CreationId: creationId,
					BaseInfo: &generated.CreationUpload{
						AuthorId:  authorId,
						Src:       src,
						Thumbnail: thumbnail,
						Title:     title,
						Duration:  duration,
					},
					UploadTime: timestamppb.New(upload_time),
				},
			}
			creationIds = append(creationIds, creationId)
			creationInfos = append(creationInfos, creationInfo)
			sqlStr = append(sqlStr, "?")
		}

		// 检查是否有额外的错误（比如数据读取完成后的关闭错误）
		if err = rows.Err(); err != nil {
			return nil, -1, errMap.MapMySQLErrorToStatus(err)
		}

		// 结束第一次查询
		if err = rows.Close(); err != nil {
			return nil, -1, errMap.MapMySQLErrorToStatus(err)
		}

		if len(sqlStr) <= 0 {
			return nil, 0, nil
		}

		queryCardEngagement := fmt.Sprintf(`
		SELECT
			views,
			likes,
			saves,
			publish_time
		FROM db_creation_engagment_1.CreationEngagement
		WHERE creation_id IN (%s)`, strings.Join(sqlStr, ","))

		// 查 统计数
		rows, err = c.db.QueryContext(
			ctx,
			queryCardEngagement,
			creationIds...,
		)
		if err != nil {
			return nil, -1, err
		}
		defer rows.Close()

		i := -1
		for rows.Next() {
			i++
			var (
				views        int32
				likes        int32
				saves        int32
				publish_time sql.NullTime
			)
			// 从当前行读取值，依次填充到变量中
			err := rows.Scan(&views, &likes, &saves, &publish_time)
			if err != nil {
				return nil, -1, errMap.MapMySQLErrorToStatus(err)
			}

			// 存储 作品卡片的统计信息
			creationEngagement := &generated.CreationEngagement{
				CreationId:  creationInfos[i].Creation.CreationId,
				Views:       views,
				PublishTime: timestamppb.New(publish_time.Time),
			}
			creationInfos[i].CreationEngagement = creationEngagement
		}

		// 检查是否有额外的错误（比如数据读取完成后的关闭错误）
		if err = rows.Err(); err != nil {
			return nil, -1, errMap.MapMySQLErrorToStatus(err)
		}
	}

	return creationInfos, count, nil
}

func (c *SqlMethodStruct) GetCreationCardInTransaction(ctx context.Context, ids []int64) ([]*generated.CreationInfo, error) {
	count := len(ids)
	if count <= 0 {
		return nil, nil
	}

	cardsMap := make(map[int64]*generated.CreationInfo)
	// 查作品信息
	// []int64 转 []string
	sqlStrs := make([]string, count)
	values := make([]any, count)
	for i, val := range ids {
		sqlStrs[i] = "?"
		values[i] = val
	}
	// 拼接
	str := strings.Join(sqlStrs, ",")

	queryCardEngagement := fmt.Sprintf(`
		SELECT
			creation_id,
			views,
			publish_time
		FROM db_creation_engagment_1.CreationEngagement 
		WHERE creation_id IN (%s)`, str)

	query := fmt.Sprintf(`SELECT
			id,
			author_id,
			src,
			thumbnail,
			title,
			status,
			duration,
			category_id,
			upload_time
		FROM db_creation_1.Creation 
		WHERE id IN (%s)`, str)

	select {
	case <-ctx.Done():
		return nil, errMap.GetStatusError(ctx.Err())
	default:
		rows, err := c.db.QueryContext(
			ctx,
			query,
			values...,
		)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		valuesC := make([]any, 0, count)
		for rows.Next() {
			var (
				id          int64
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
			err := rows.Scan(&id, &author_id, &src, &thumbnail, &title, &status, &duration, &category_id, &upload_time)
			if err != nil {
				return nil, errMap.MapMySQLErrorToStatus(err)
			}

			c_status := generated.CreationStatus(generated.CreationStatus_value[status])

			// 存储 卡片基本信息切片
			cardsMap[id] = &generated.CreationInfo{}
			cardsMap[id].Creation = &generated.Creation{
				CreationId: id,
				BaseInfo: &generated.CreationUpload{
					AuthorId:   author_id,
					Src:        src,
					Thumbnail:  thumbnail,
					Title:      title,
					Status:     c_status,
					Duration:   duration,
					CategoryId: category_id,
				},
				UploadTime: timestamppb.New(upload_time),
			}
			valuesC = append(valuesC, id)
		}

		// 检查是否有额外的错误（比如数据读取完成后的关闭错误）
		if err = rows.Err(); err != nil {
			return nil, errMap.MapMySQLErrorToStatus(err)
		}

		// 结束第一次查询
		if err = rows.Close(); err != nil {
			return nil, errMap.MapMySQLErrorToStatus(err)
		}

		// 查 统计数
		rows, err = c.db.QueryContext(
			ctx,
			queryCardEngagement,
			valuesC...,
		)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {
			var (
				creation_id  int64
				views        int32
				publish_time sql.NullTime
			)
			// 从当前行读取值，依次填充到变量中
			err := rows.Scan(&creation_id, &views, &publish_time)
			if err != nil {
				return nil, errMap.MapMySQLErrorToStatus(err)
			}

			// 存储 作品卡片的统计信息
			cardsMap[creation_id].CreationEngagement = &generated.CreationEngagement{
				CreationId:  creation_id,
				Views:       views,
				PublishTime: timestamppb.New(publish_time.Time),
			}
		}

		// 检查是否有额外的错误（比如数据读取完成后的关闭错误）
		if err = rows.Err(); err != nil {
			return nil, errMap.MapMySQLErrorToStatus(err)
		}
	}

	creationInfos := make([]*generated.CreationInfo, 0, len(cardsMap))
	// 统合
	for _, info := range cardsMap {
		creationInfos = append(creationInfos, info)
	}
	return creationInfos, nil
}

// DELETE
func (c *SqlMethodStruct) DeleteCreationInTransaction(ctx context.Context, id int64) error {
	query := `DELETE FROM db_creation_1.Creation 
		WHERE id = ?`

	select {
	case <-ctx.Done():
		return errMap.GetStatusError(ctx.Err())
	default:
		_, err := c.db.ExecContext(
			ctx,
			query,
			id,
		)
		if err != nil {
			return err
		}
	}

	return nil
}

// UPDATE
func (c *SqlMethodStruct) UpdateViewsInTransaction(ctx context.Context, creationId int64, changingNum int) error {
	query := `
	UPDATE db_creation_engagment_1.CreationEngagement
	SET views = views + ?
	WHERE creation_id = ?`

	select {
	case <-ctx.Done():
		return errMap.GetStatusError(ctx.Err())
	default:
		_, err := c.db.ExecContext(
			ctx,
			query,
			changingNum,
			creationId,
		)
		if err != nil {
			return errMap.MapMySQLErrorToStatus(err)
		}
	}

	return nil
}

func (c *SqlMethodStruct) UpdateLikesInTransaction(ctx context.Context, creationId int64, changingNum int) error {
	query := `
	UPDATE db_creation_engagment_1.CreationEngagement
	SET likes = likes + ?
	WHERE creation_id = ?`

	select {
	case <-ctx.Done():
		return errMap.GetStatusError(ctx.Err())
	default:
		_, err := c.db.ExecContext(
			ctx,
			query,
			changingNum,
			creationId,
		)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *SqlMethodStruct) UpdateSavesInTransaction(ctx context.Context, creationId int64, changingNum int) error {
	query := `
	UPDATE db_creation_engagment_1.CreationEngagement
	SET saves = saves + ?
	WHERE creation_id = ?`

	select {
	case <-ctx.Done():
		return errMap.GetStatusError(ctx.Err())
	default:
		_, err := c.db.ExecContext(
			ctx,
			query,
			changingNum,
			creationId,
		)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *SqlMethodStruct) UpdateCreationInTransaction(ctx context.Context, creation *generated.CreationUpdated) error {
	var (
		thumbnail = creation.GetThumbnail()
		title     = creation.GetTitle()
		bio       = creation.GetBio()
		src       = creation.GetSrc()
		duration  = creation.GetDuration()
		userId    = creation.GetAuthorId()
		AND       = " "
	)
	const (
		setThumbnail = "thumbnail = ?"
		setTitle     = "title = ?"
		setBio       = "bio = ?"
		setSrc       = "src = ?"
		setDuration  = "duration = ?"
	)

	values := make([]any, 0, 8)
	sqlStr := make([]string, 0, 5)
	if thumbnail != "" {
		sqlStr = append(sqlStr, setThumbnail)
		values = append(values, thumbnail)
	}
	if title != "" {
		sqlStr = append(sqlStr, setTitle)
		values = append(values, title)
	}
	if bio != "" {
		sqlStr = append(sqlStr, setBio)
		values = append(values, bio)
	}
	if src != "" {
		sqlStr = append(sqlStr, setSrc)
		values = append(values, src)
	}
	if duration != 0 {
		sqlStr = append(sqlStr, setDuration)
		values = append(values, duration)
	}

	values = append(values, creation.GetCreationId())
	if userId != -403 {
		AND = " AND author_id = ? "
		values = append(values, userId)
	}

	if len(sqlStr) <= 0 {
		return nil
	}

	query := fmt.Sprintf(`
		UPDATE db_creation_1.Creation
		SET 
			status = 'PENDING',
			%s
		WHERE 
			id = ? 
		%s`, strings.Join(sqlStr, ","), AND)
	affected, err := c.db.ExecContext(ctx, query, values...)
	if err != nil {
		return err
	}
	num, err := affected.RowsAffected()
	if err != nil {
		return errMap.MapMySQLErrorToStatus(err)
	}
	if num <= 0 {
		return grpcStatus.Errorf(codes.NotFound, "not match the author")
	}
	return nil
}

func (c *SqlMethodStruct) UpdateCreationStatusInTransaction(ctx context.Context, creation *generated.CreationUpdateStatus) error {
	var (
		creationId = creation.GetCreationId()
		status     = creation.GetStatus()
		userId     = creation.GetAuthorId()
		AND        = " "
	)

	values := make([]any, 0, 8)

	values = append(values, status.String(), creationId)
	if userId != -403 {
		AND = " AND author_id = ? "
		values = append(values, userId)
	}

	query := fmt.Sprintf(`
		UPDATE db_creation_1.Creation
		SET 
			status = ?
		WHERE 
			id = ? 
		%s`, AND)
	affected, err := c.db.ExecContext(ctx, query, values...)
	if err != nil {
		return err
	}
	num, err := affected.RowsAffected()
	if err != nil {
		return errMap.MapMySQLErrorToStatus(err)
	}
	if num <= 0 {
		return grpcStatus.Errorf(codes.NotFound, fmt.Sprintf("%d not match the %d author", userId, creationId))
	}
	return nil
}

func (c *SqlMethodStruct) PublishCreationInTransaction(ctx context.Context, creationId int64, publishTime *timestamppb.Timestamp) error {
	query := `
		UPDATE db_creation_engagment_1.CreationEngagement
		SET publish_time = CASE 
    		WHEN publish_time IS NULL THEN ? 
    		ELSE publish_time 
		END
		WHERE id = ?`
	_, err := c.db.ExecContext(
		ctx,
		query,
		publishTime.AsTime(),
		creationId,
	)
	if err != nil {
		return err
	}

	return nil
}

func (c *SqlMethodStruct) UpdateCreationCount(ctx context.Context, creationId int64, saveCount, likeCount, viewCount int32) error {
	const (
		saveSql = "saves = saves + ?"
		likeSql = "likes = likes + ?"
		viewSql = "views = views + ?"
	)

	sqlStr := make([]string, 0, 3)
	values := make([]any, 0, 4)
	if saveCount != 0 {
		sqlStr = append(sqlStr, saveSql)
		values = append(values, saveCount)
	}
	if likeCount != 0 {
		sqlStr = append(sqlStr, likeSql)
		values = append(values, likeCount)
	}
	if viewCount != 0 {
		sqlStr = append(sqlStr, viewSql)
		values = append(values, viewCount)
	}
	if len(sqlStr) <= 0 {
		return nil
	}
	values = append(values, creationId)

	query := fmt.Sprintf(`
		UPDATE db_creation_engagment_1.CreationEngagement 
		SET 
			%s
		WHERE creation_id = ?
	`, strings.Join(sqlStr, ","))

	_, err := c.db.ExecContext(
		ctx,
		query,
		values...,
	)
	return err
}

func (c *SqlMethodStruct) SearchCreations(ctx context.Context, title string, page int32) ([]*generated.CreationInfo, int32, error) {
	const LIMIT = 20
	offset := (page - 1) * LIMIT

	// 主页,相似列表,分区
	query := `SELECT
			id,
			src,
			thumbnail,
			duration,
			upload_time
		FROM db_creation_1.Creation 
		WHERE title like %?%
		AND status = 'PUBLISHED' 
		LIMIT ?
		OFFSET ?`

	queryCount := `SELECT
		count(*)
	FROM db_creation_1.Creation 
	WHERE title like %?%
	AND status = 'PUBLISHED' `

	sqlStr := make([]string, 0, LIMIT)
	creationInfos := make([]*generated.CreationInfo, 0, 20)
	creationIds := make([]any, 0, 20)
	var count int32
	select {
	case <-ctx.Done():
		return nil, -1, errMap.GetStatusError(ctx.Err())
	default:
		var num int32
		err := c.db.QueryRowContext(
			ctx,
			queryCount,
		).Scan(&num)
		if err != nil {
			return nil, -1, errMap.MapMySQLErrorToStatus(err)
		}
		if num <= 0 {
			return nil, 0, nil
		}
		count = int32(math.Ceil(float64(num) / float64(LIMIT)))

		rows, err := c.db.QueryContext(
			ctx,
			query,
			title,
			LIMIT,
			offset,
		)
		if err != nil {
			return nil, -1, err
		}
		defer rows.Close()

		for rows.Next() {
			var (
				creationId  int64
				src         string
				thumbnail   string
				authorId    int64
				duration    int32
				upload_time time.Time
			)
			// 从当前行读取值，依次填充到变量中
			err := rows.Scan(&creationId, &src, &thumbnail, &authorId, &duration, &upload_time)
			if err != nil {
				return nil, -1, errMap.MapMySQLErrorToStatus(err)
			}

			// 存储 卡片基本信息切片
			creationInfo := &generated.CreationInfo{
				Creation: &generated.Creation{
					CreationId: creationId,
					BaseInfo: &generated.CreationUpload{
						AuthorId:  authorId,
						Src:       src,
						Thumbnail: thumbnail,
						Title:     title,
						Duration:  duration,
					},
					UploadTime: timestamppb.New(upload_time),
				},
			}
			creationIds = append(creationIds, creationId)
			creationInfos = append(creationInfos, creationInfo)
			sqlStr = append(sqlStr, "?")
		}

		// 检查是否有额外的错误（比如数据读取完成后的关闭错误）
		if err = rows.Err(); err != nil {
			return nil, -1, errMap.MapMySQLErrorToStatus(err)
		}

		// 结束第一次查询
		if err = rows.Close(); err != nil {
			return nil, -1, errMap.MapMySQLErrorToStatus(err)
		}

		if len(sqlStr) <= 0 {
			return nil, 0, nil
		}

		queryCardEngagement := fmt.Sprintf(`
		SELECT
			views,
			likes,
			saves,
			publish_time
		FROM db_creation_engagment_1.CreationEngagement
		WHERE creation_id IN (%s)`, strings.Join(sqlStr, ","))

		// 查 统计数
		rows, err = c.db.QueryContext(
			ctx,
			queryCardEngagement,
			creationIds...,
		)
		if err != nil {
			return nil, -1, err
		}
		defer rows.Close()

		i := -1
		for rows.Next() {
			i++
			var (
				views        int32
				likes        int32
				saves        int32
				publish_time sql.NullTime
			)
			// 从当前行读取值，依次填充到变量中
			err := rows.Scan(&views, &likes, &saves, &publish_time)
			if err != nil {
				return nil, -1, errMap.MapMySQLErrorToStatus(err)
			}

			// 存储 作品卡片的统计信息
			creationEngagement := &generated.CreationEngagement{
				CreationId:  creationInfos[i].Creation.CreationId,
				Views:       views,
				PublishTime: timestamppb.New(publish_time.Time),
			}
			creationInfos[i].CreationEngagement = creationEngagement
		}

		// 检查是否有额外的错误（比如数据读取完成后的关闭错误）
		if err = rows.Err(); err != nil {
			return nil, -1, errMap.MapMySQLErrorToStatus(err)
		}
	}

	return creationInfos, count, nil
}

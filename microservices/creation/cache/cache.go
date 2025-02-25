package cache

import (
	"context"
	"fmt"
	"log"
	"math"
	"strconv"

	generated "github.com/Yux77Yux/platform_backend/generated/creation"
	tools "github.com/Yux77Yux/platform_backend/microservices/creation/tools"
	"github.com/go-redis/redis/v8"
)

// POST
func CreationAddInCache(creationInfo *generated.CreationInfo) error {
	ctx := context.Background()

	creation := creationInfo.GetCreation()

	id := strconv.FormatInt(creation.GetCreationId(), 10)

	resultCh := make(chan error, 1)

	categoryId := creation.GetBaseInfo().GetCategoryId()

	reqFunc := func(CacheClient CacheInterface) {
		err := CacheClient.SetFieldsHash(ctx, "CreationInfo", id,
			"author_id", creation.GetBaseInfo().GetAuthorId(),
			"src", creation.GetBaseInfo().GetSrc(),
			"thumbnail", creation.GetBaseInfo().GetThumbnail(),
			"title", creation.GetBaseInfo().GetTitle(),
			"bio", creation.GetBaseInfo().GetBio(),
			"status", creation.GetBaseInfo().GetStatus().String(),
			"duration", creation.GetBaseInfo().GetDuration(),
			"category_id", categoryId,
			"upload_time", creation.GetUploadTime().AsTime(),

			"views", 0,
			"saves", 0,
			"likes", 0,
			"publish_time", "none",

			"category_name", tools.Categories[categoryId].Name,
			"category_parent", tools.Categories[categoryId].Parent,
		)
		resultCh <- err
	}

	cacheRequestChannel <- reqFunc

	// 使用 select 来监听超时和结果
	select {
	case <-ctx.Done():
		return fmt.Errorf("timeout: %w", ctx.Err())
	case result := <-resultCh:
		if result != nil {
			return result
		}
		return nil
	}
}

// GET
func GetCreationInfoFields(ctx context.Context, creation_id int64, fields []string) (map[string]string, error) {
	resultCh := make(chan struct {
		creationInfo map[string]string
		err          error
	}, 1)

	cacheRequestChannel <- func(CacheClient CacheInterface) {
		if len(fields) == 0 || fields == nil {
			result, err := CacheClient.GetAllHash(ctx, "CreationInfo", strconv.FormatInt(creation_id, 10))
			resultCh <- struct {
				creationInfo map[string]string
				err          error
			}{
				creationInfo: result,
				err:          err,
			}
		} else {
			values, err := CacheClient.GetAnyHash(ctx, "CreationInfo", strconv.FormatInt(creation_id, 10), fields...)
			// 构造结果 map
			result := make(map[string]string, len(fields))

			for i, field := range fields {
				// 类型断言并检查 nil 值
				if values[i] != nil {
					strValue, ok := values[i].(string)
					if !ok {
						err = fmt.Errorf("unexpected value type for field %s", field)
						break
					}
					result[field] = strValue
				}
			}

			resultCh <- struct {
				creationInfo map[string]string
				err          error
			}{
				creationInfo: result,
				err:          err,
			}
		}
	}

	// 使用 select 来监听超时和结果
	select {
	case <-ctx.Done():
		return nil, fmt.Errorf("timeout: %w", ctx.Err())
	case result := <-resultCh:
		if result.err != nil {
			return nil, result.err
		}

		creationInfo := result.creationInfo
		for _, key := range fields {
			if val, ok := creationInfo[key]; !ok || val == "" {
				return nil, fmt.Errorf("error: missing or empty field %s", key)
			}
		}
		return creationInfo, nil
	}
}

// 视频展示页的Redis缓存
func GetCreationInfo(ctx context.Context, creation_id int64) (*generated.CreationInfo, error) {
	results, err := GetCreationInfoFields(ctx, creation_id, nil)
	if err != nil {
		log.Printf("error: GetCreationInfo GetCreationInfoFields %v", err)
		return nil, err
	}
	creationInfo, err := tools.MapCreationInfoByString(results)
	if err != nil {
		return nil, err
	}

	if creationInfo == nil {
		return nil, nil
	}

	return creationInfo, nil
}

func parseIntField(value string, bitSize int) (int64, error) {
	if value == "" {
		return 0, fmt.Errorf("数值字段为空")
	}
	return strconv.ParseInt(value, 10, bitSize)
}
func mapToCreationInfo(results map[string]string, creation_id int64) (*generated.CreationInfo, error) {
	requiredKeys := []string{
		"author_id", "src", "thumbnail", "title", "bio",
		"duration", "views",
	}

	// 确保所有必须字段存在且非空
	for _, key := range requiredKeys {
		if val, ok := results[key]; !ok || val == "" {
			return nil, nil
		}
	}

	var (
		authorIdStr = results["author_id"]
		src         = results["src"]
		thumbnail   = results["thumbnail"]
		title       = results["title"]
		bio         = results["bio"]
		durationStr = results["duration"]
		viewsStr    = results["views"]
	)

	authorId, err := parseIntField(authorIdStr, 64)
	if err != nil {
		log.Printf("error: GetCreationInfo authorIdStr ParseInt %v", err)
		return nil, err
	}

	durationInt, err := strconv.Atoi(durationStr)
	if err != nil {
		log.Printf("error: GetCreationInfo durationStr Atoi %v", err)
		return nil, err
	}
	duration := int32(durationInt)

	viewsInt, err := strconv.Atoi(viewsStr)
	if err != nil {
		log.Printf("error: GetCreationInfo viewsStr Atoi %v", err)
		return nil, err
	}
	views := int32(viewsInt)

	creationInfo := &generated.CreationInfo{
		Creation: &generated.Creation{
			CreationId: creation_id,
			BaseInfo: &generated.CreationUpload{
				AuthorId:  authorId,
				Src:       src,
				Thumbnail: thumbnail,
				Title:     title,
				Bio:       bio,
				Duration:  duration,
			},
		},
		CreationEngagement: &generated.CreationEngagement{
			CreationId: creation_id,
			Views:      views,
		},
	}

	return creationInfo, nil
}

func GetSimilarCreationList(ctx context.Context, creation_id int64) ([]int64, error) {
	strs, err := CacheClient.RevRangeZSet(ctx, "SimilarCreation", strconv.FormatInt(creation_id, 10), 0, 149)
	if err != nil {
		log.Printf("error: GetSpaceCreationList RevRangeZSet %v", err)
		return nil, err
	}

	count := len(strs)
	if count == 0 {
		return nil, nil // 返回空结果
	}

	ids := make([]int64, count)
	for i, str := range strs {
		id, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			return nil, err
		}
		ids[i] = id
	}

	return ids, nil
}

func GetSpaceCreationList(ctx context.Context, user_id int64, page int32, typeStr string) ([]int64, int32, error) {
	const LIMIT = 25
	start := int64((page - 1) * LIMIT)
	stop := start + 24

	pipe := CacheClient.Pipeline()

	strsCmd := pipe.ZRevRange(ctx, fmt.Sprintf("ZSet_SpaceCreation_%s_%d", typeStr, user_id), start, stop)
	countCmd := pipe.ZCard(ctx, fmt.Sprintf("ZSet_SpaceCreation_%s_%d", typeStr, user_id))

	_, err := pipe.Exec(ctx)
	if err != nil {
		return nil, -1, err
	}

	strs, err := strsCmd.Result()
	if err != nil {
		return nil, -1, err
	}

	length := len(strs)
	if length == 0 {
		return nil, 0, nil // 返回空结果
	}

	count, err := countCmd.Result()
	if err != nil {
		return nil, -1, err
	}

	ids := make([]int64, count)
	for i, str := range strs {
		id, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			return nil, -1, err
		}
		ids[i] = id
	}

	pagesNum := int32(math.Ceil(float64(count) / float64(LIMIT)))
	return ids, pagesNum, nil
}

func GetHistoryCreationList(ctx context.Context, user_id int64) ([]int64, error) {
	strs, err := CacheClient.RevRangeZSet(ctx, "Histories", strconv.FormatInt(user_id, 10), 0, 149)
	if err != nil {
		log.Printf("error: GetHistoryCreationList RevRangeZSet %v", err)
		return nil, err
	}

	count := len(strs)
	if count == 0 {
		return nil, nil // 返回空结果
	}

	ids := make([]int64, count)
	for i, str := range strs {
		id, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			return nil, err
		}
		ids[i] = id
	}

	return ids, nil
}

func GetCollectionCreationList(ctx context.Context, user_id int64) ([]int64, error) {
	strs, err := CacheClient.RevRangeZSet(ctx, "Collections", strconv.FormatInt(user_id, 10), 0, 149)
	if err != nil {
		log.Printf("error: GetCollectionCreationList RevRangeZSet %v", err)
		return nil, err
	}

	count := len(strs)
	if count == 0 {
		return nil, nil // 返回空结果
	}

	ids := make([]int64, count)
	for i, str := range strs {
		id, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			return nil, err
		}
		ids[i] = id
	}

	return ids, nil
}

// Collections,History
func GetCreationList(ctx context.Context, creationIds []int64) ([]*generated.CreationInfo, []int64, error) {
	pipe := CacheClient.Pipeline()

	length := len(creationIds)
	infos := make([]*generated.CreationInfo, 0, length)
	notCaches := make([]int64, 0, length)
	cmds := make(map[int64]*redis.StringStringMapCmd)

	for _, id := range creationIds {
		key := fmt.Sprintf("Hash_CreationInfo_%d", id)
		cmds[id] = pipe.HGetAll(ctx, key)
	}

	_, err := pipe.Exec(ctx)
	if err != nil {
		return nil, nil, err
	}

	for id, cmd := range cmds {
		results, err := cmd.Result()
		// 有错误则跳过
		if err != nil {
			notCaches = append(notCaches, id)
			continue
		}

		creationInfo, err := mapToCreationInfo(results, id)
		if err != nil {
			notCaches = append(notCaches, id)
			continue
		}
		infos = append(infos, creationInfo)
	}

	return infos, notCaches, nil
}

// DEL
func DeleteCreation(creation_id int64) error {
	idStr := strconv.FormatInt(creation_id, 10)
	ctx := context.Background()

	resultCh := make(chan error, 1)

	cacheRequestChannel <- func(CacheClient CacheInterface) {
		err := CacheClient.SetFieldHash(ctx, "Hash_CreationInfo", idStr, "status", generated.CreationStatus_DELETE.String())
		resultCh <- err
	}

	// 使用 select 来监听超时和结果
	select {
	case <-ctx.Done():
		return fmt.Errorf("timeout: %w", ctx.Err())
	case err := <-resultCh:
		if err != nil {
			return err
		}

		return nil
	}
}

// UPDATE
func UpdateCreation(creation *generated.CreationUpdated) error {
	var (
		creationId = creation.GetCreationId()
		thumbnail  = creation.GetThumbnail()
		title      = creation.GetTitle()
		bio        = creation.GetBio()
		src        = creation.GetSrc()
		duration   = creation.GetDuration()
	)

	values := make([]any, 0, 5*2)
	if thumbnail != "" {
		values = append(values, "thumbnail", thumbnail)
	}
	if title != "" {
		values = append(values, "title", title)
	}
	if bio != "" {
		values = append(values, "bio", bio)
	}
	if src != "" {
		values = append(values, "src", src)
	}
	if duration != 0 {
		values = append(values, "duration", duration)
	}
	if len(values) <= 0 {
		return nil
	}

	ctx := context.Background()
	err := CacheClient.SetFieldsHash(ctx, "CreationInfo", strconv.FormatInt(creationId, 10), values...)
	return err
}

func UpdateCreationStatus(creation *generated.CreationUpdateStatus) error {
	var (
		creationId = creation.GetCreationId()
		status     = creation.GetStatus()
	)

	values := make([]any, 0, 2)
	values = append(values, "status", status.String())

	ctx := context.Background()
	err := CacheClient.SetFieldsHash(ctx, "CreationInfo", strconv.FormatInt(creationId, 10), values...)
	return err
}

package cache

import (
	"context"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"google.golang.org/protobuf/types/known/timestamppb"

	generated "github.com/Yux77Yux/platform_backend/generated/interaction"
)

// GET

type Result struct {
	result []redis.Z
	err    error
}

type ResultStr struct {
	result []string
	err    error
}

func ToBaseInteraction(results []redis.Z) ([]*generated.Interaction, error) {
	count := len(results)
	res := make([]*generated.Interaction, count)
	for i, val := range results {
		id, err := strconv.ParseInt(val.Member.(string), 10, 64)
		if err != nil {
			return nil, err
		}
		times := int64(math.Round(val.Score))
		timestamp := timestamppb.New(time.Unix(times, 0))
		res[i].Base = &generated.BaseInteraction{
			CreationId: id,
		}
		res[i].SaveAt = timestamp
		res[i].UpdatedAt = timestamp
	}
	return res, nil
}

// 历史记录
func GetHistories(ctx context.Context, userId int64, page int32) ([]*generated.Interaction, error) {
	const scope = 30
	start := int64((page - 1) * scope)
	stop := start + scope

	userIdStr := strconv.FormatInt(userId, 10)
	resultCh := make(chan *Result, 1)

	cacheRequestChannel <- func(CacheClient CacheInterface) {
		results, err := CacheClient.RevRangeZSetWithScore(ctx, "Histories", userIdStr, start, stop)
		resultCh <- &Result{
			result: results,
			err:    err,
		}
	}

	// 使用 select 来监听超时和结果
	select {
	case <-ctx.Done():
		return nil, fmt.Errorf("timeout: %w", ctx.Err())
	case result := <-resultCh:
		if result.err != nil {
			return nil, fmt.Errorf("error: %w", result.err)
		}
		res, err := ToBaseInteraction(result.result)
		if err != nil {
			return nil, fmt.Errorf("error: %w", err)
		}

		for i := range res {
			res[i].Base.UserId = userId
		}
		return res, nil
	}
}

// 收藏夹
func GetCollections(ctx context.Context, userId int64, page int32) ([]*generated.Interaction, error) {
	const scope = 30
	start := int64((page - 1) * scope)
	stop := start + scope
	userIdStr := strconv.FormatInt(userId, 10)
	resultCh := make(chan *Result, 1)

	cacheRequestChannel <- func(CacheClient CacheInterface) {
		results, err := CacheClient.RevRangeZSetWithScore(ctx, "Collections", userIdStr, start, stop)
		resultCh <- &Result{
			result: results,
			err:    err,
		}
	}

	// 使用 select 来监听超时和结果
	select {
	case <-ctx.Done():
		return nil, fmt.Errorf("timeout: %w", ctx.Err())
	case result := <-resultCh:
		if result.err != nil {
			return nil, fmt.Errorf("error: %w", result.err)
		}
		res, err := ToBaseInteraction(result.result)
		if err != nil {
			return nil, fmt.Errorf("error: %w", err)
		}

		for i := range res {
			res[i].Base.UserId = userId
		}
		return res, nil
	}
}

// Like
func GetLikes(userId int64) ([]*generated.BaseInteraction, error) {
	ctx := context.Background()
	userIdStr := strconv.FormatInt(userId, 10)

	resultCh := make(chan *ResultStr, 1)

	cacheRequestChannel <- func(CacheClient CacheInterface) {
		results, err := CacheClient.GetMembersSet(ctx, "Likes", userIdStr)
		resultCh <- &ResultStr{
			result: results,
			err:    err,
		}
	}

	// 使用 select 来监听超时和结果
	select {
	case <-ctx.Done():
		return nil, fmt.Errorf("timeout: %w", ctx.Err())
	case result := <-resultCh:
		if result.err != nil {
			return nil, fmt.Errorf("error: %w", result.err)
		}

		creationIdStrs := result.result
		length := len(creationIdStrs)
		res := make([]*generated.BaseInteraction, length)
		for i, val := range creationIdStrs {
			CreationId, idErr := strconv.ParseInt(val, 10, 64)
			if idErr != nil {
				return nil, idErr
			}
			res[i].CreationId = CreationId
			res[i].UserId = userId
		}

		return res, nil
	}
}

// 观看作品的用户
func GetUsers(creationId int64) ([]int64, error) {
	ctx := context.Background()

	creationIdStr := strconv.FormatInt(creationId, 10)
	resultCh := make(chan *ResultStr, 1)

	cacheRequestChannel <- func(CacheClient CacheInterface) {
		results, err := CacheClient.RevRangeZSet(ctx, "Item_Users", creationIdStr, 0, 199)
		resultCh <- &ResultStr{
			result: results,
			err:    err,
		}
	}

	// 使用 select 来监听超时和结果
	select {
	case <-ctx.Done():
		return nil, fmt.Errorf("timeout: %w", ctx.Err())
	case result := <-resultCh:
		if result.err != nil {
			return nil, fmt.Errorf("error: %w", result.err)
		}
		strs := result.result
		res := make([]int64, len(strs))
		for i, str := range strs {
			id, err := strconv.ParseInt(str, 10, 64)
			if err != nil {
				return nil, err
			}
			res[i] = id
		}
		return res, nil
	}
}

type ActionResult struct {
	err        error
	action_tag int32
}

// 展示页·点赞收藏情况
func GetInteraction(ctx context.Context, interaction *generated.BaseInteraction) (*generated.Interaction, error) {
	userId := interaction.GetUserId()
	creationId := interaction.GetCreationId()

	resultCh := make(chan *ActionResult, 1)
	cacheRequestChannel <- func(CacheClient CacheInterface) {
		pipe := CacheClient.Pipeline()
		isLike := fmt.Sprintf("ZSet_Likes_%d", userId)             // 自己是否点赞
		isCollection := fmt.Sprintf("ZSet_Collections_%d", userId) // 自己是否收藏

		creationIdStr := strconv.FormatInt(creationId, 10)

		// 使用 pipeline 执行多个查询
		zScoreLikeCmd := pipe.ZScore(ctx, isLike, creationIdStr)
		zScoreCollectionCmd := pipe.ZScore(ctx, isCollection, creationIdStr)

		// 执行 pipeline
		_, err := pipe.Exec(ctx)
		if err != nil {
			resultCh <- &ActionResult{
				err: fmt.Errorf("pipeline execution failed: %w", err),
			}
			return
		}

		// 解析返回结果
		likeScore, err := zScoreLikeCmd.Result()
		if err == redis.Nil {
			likeScore = -1 // 代表用户没有点赞
		} else if err != nil {
			resultCh <- &ActionResult{
				err: fmt.Errorf("failed to get like score: %w", err),
			}
			return
		}

		collectionScore, err := zScoreCollectionCmd.Result()
		if err == redis.Nil {
			collectionScore = -1 // 代表用户没有收藏
		} else if err != nil {
			resultCh <- &ActionResult{
				err: fmt.Errorf("failed to get collection score: %w", err),
			}
			return
		}

		action_tag := int32(1)
		if likeScore != -1 {
			action_tag = action_tag | 2
		}
		if collectionScore != -1 {
			action_tag = action_tag | 4
		}
		resultCh <- &ActionResult{
			action_tag: action_tag,
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
		return &generated.Interaction{
			Base:      interaction,
			ActionTag: result.action_tag,
		}, nil
	}
}

// 查看是否过期，是否重新计算

// POST
func SetRecommendBaseUser(id int64, ids []int64) error {
	ctx := context.Background()
	count := len(ids)
	values := make([]any, count)
	for i, val := range ids {
		values[i] = val
	}

	err := CacheClient.AddToSet(ctx, "RecommendBaseUser", strconv.FormatInt(id, 10), values...)
	if err != nil {
		return err
	}
	return nil
}

func SetRecommendBaseItem(id int64, ids []int64) error {
	ctx := context.Background()
	count := len(ids)
	values := make([]any, count)
	for i, val := range ids {
		values[i] = val
	}

	err := CacheClient.AddToSet(ctx, "RecommendBaseItem", strconv.FormatInt(id, 10), values...)
	if err != nil {
		return err
	}
	return nil
}

// POST & UPDATE
// 历史记录 更新时间戳
func UpdateHistories(data []*generated.Interaction) error {
	ctx := context.Background()
	resultCh := make(chan error, 1)
	cacheRequestChannel <- func(CacheClient CacheInterface) {
		pipe := CacheClient.Pipeline()
		for _, option := range data {
			base := option.GetBase()

			userId := base.GetUserId()
			creationId := base.GetCreationId()
			timestampScore := option.GetUpdatedAt().GetSeconds()

			key := fmt.Sprintf("ZSet_Histories_%d", userId)
			pipe.ZAdd(ctx, key, &redis.Z{
				Score:  float64(timestampScore),
				Member: creationId,
			})

			viewKey := fmt.Sprintf("Hash_CreationInfo_%d", creationId)
			pipe.HIncrBy(ctx, viewKey, "views", 1)
		}
		_, err := pipe.Exec(ctx)
		if err != nil {
			resultCh <- fmt.Errorf("pipeline execution failed: %w", err)
		}
	}
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

// 收藏夹
func ModifyCollections(data []*generated.Interaction) error {
	ctx := context.Background()
	resultCh := make(chan error, 1)
	cacheRequestChannel <- func(CacheClient CacheInterface) {
		ctx := context.Background()
		pipe := CacheClient.Pipeline()
		for _, option := range data {
			base := option.GetBase()

			userId := base.GetUserId()
			creationId := base.GetCreationId()
			timestampScore := option.GetUpdatedAt().GetSeconds()

			key := fmt.Sprintf("ZSet_Collections_%d", userId)
			pipe.ZAdd(ctx, key, &redis.Z{
				Score:  float64(timestampScore),
				Member: creationId,
			})

			collectionKey := fmt.Sprintf("Hash_CreationInfo_%d", creationId)
			pipe.HIncrBy(ctx, collectionKey, "saves", 1)
		}
		_, err := pipe.Exec(ctx)
		if err != nil {
			resultCh <- fmt.Errorf("pipeline execution failed: %w", err)
		}
	}
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

// 点赞
func ModifyLike(data []*generated.BaseInteraction) error {
	ctx := context.Background()
	resultCh := make(chan error, 1)
	cacheRequestChannel <- func(CacheClient CacheInterface) {
		ctx := context.Background()
		pipe := CacheClient.Pipeline()
		for _, base := range data {

			userId := base.GetUserId()
			creationId := base.GetCreationId()
			key := fmt.Sprintf("Set_Likes_%d", userId)
			pipe.SAdd(ctx, key, creationId)

			likeKey := fmt.Sprintf("Hash_CreationInfo_%d", creationId)
			pipe.HIncrBy(ctx, likeKey, "likes", 1)
		}
		_, err := pipe.Exec(ctx)
		if err != nil {
			resultCh <- fmt.Errorf("pipeline execution failed: %w", err)
		}
	}
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

// DELETE

func DelHistories(data []*generated.BaseInteraction) error {
	ctx := context.Background()
	resultCh := make(chan error, 1)
	cacheRequestChannel <- func(CacheClient CacheInterface) {
		ctx := context.Background()
		pipe := CacheClient.Pipeline()
		for _, base := range data {

			userId := base.GetUserId()
			creationId := base.GetCreationId()

			key := fmt.Sprintf("ZSet_Histories_%d", userId)
			pipe.ZRem(ctx, key, creationId)
		}
		_, err := pipe.Exec(ctx)
		if err != nil {
			resultCh <- fmt.Errorf("pipeline execution failed: %w", err)
		}
	}
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

func DelCollections(data []*generated.BaseInteraction) error {
	ctx := context.Background()
	resultCh := make(chan error, 1)
	cacheRequestChannel <- func(CacheClient CacheInterface) {
		ctx := context.Background()
		pipe := CacheClient.Pipeline()
		for _, base := range data {

			userId := base.GetUserId()
			creationId := base.GetCreationId()

			key := fmt.Sprintf("ZSet_Collections_%d", userId)
			pipe.ZRem(ctx, key, creationId)

			collectionKey := fmt.Sprintf("Hash_CreationInfo_%d", creationId)
			pipe.HIncrBy(ctx, collectionKey, "saves", -1)
		}
		_, err := pipe.Exec(ctx)
		if err != nil {
			resultCh <- fmt.Errorf("pipeline execution failed: %w", err)
		}
	}
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

func DelLike(data []*generated.BaseInteraction) error {
	ctx := context.Background()
	resultCh := make(chan error, 1)
	cacheRequestChannel <- func(CacheClient CacheInterface) {
		ctx := context.Background()
		pipe := CacheClient.Pipeline()
		for _, base := range data {

			userId := base.GetUserId()
			creationId := base.GetCreationId()

			key := fmt.Sprintf("Set_Likes_%d", creationId)
			pipe.ZRem(ctx, key, userId)

			likeKey := fmt.Sprintf("Hash_CreationInfo_%d", creationId)
			pipe.HIncrBy(ctx, likeKey, "likes", 1)
		}
		_, err := pipe.Exec(ctx)
		if err != nil {
			resultCh <- fmt.Errorf("pipeline execution failed: %w", err)
		}
	}
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

// Scan
// 拿到别人的历史记录
func ScanZSetsByHistories() ([]string, error) {
	ctx := context.Background()

	results, _, err := CacheClient.ScanZSet(ctx, "Histories", "*", 0, 2500)
	if err != nil {
		return nil, err
	}

	length := len(results)
	idStrs := make([]string, length)
	for i, val := range results {
		idStr := strings.Split(val, "_")
		idStrs[i] = idStr[len(idStr)-1]
	}

	return idStrs, nil
}

func ScanZSetsByCreationId() ([]string, error) {
	ctx := context.Background()

	results, _, err := CacheClient.ScanZSet(ctx, "Item_Users", "*", 0, 2500)
	if err != nil {
		return nil, err
	}

	length := len(results)
	idStrs := make([]string, length)
	for i, val := range results {
		idStr := strings.Split(val, "_")
		idStrs[i] = idStr[len(idStr)-1]
	}

	return idStrs, nil
}

func GetAllInteractions(idStrs []string) (map[int64]map[int64]float64, error) {
	const (
		viewWeight = 1
	)
	ctx := context.Background()
	pipe := CacheClient.Pipeline()

	// 用来存储 pipeline 请求的结果
	historyCmds := make([]*redis.StringSliceCmd, len(idStrs))

	// 依次遍历用户 ID，把请求加入 pipeline
	for i, str := range idStrs {
		historyKey := fmt.Sprintf("ZSet_Histories_%s", str) // 观看记录 (ZSet)

		// 用 ZRange 取 ZSet，避免用 SMembers 读错数据类型
		historyCmds[i] = pipe.ZRevRange(ctx, historyKey, 0, 199)
	}

	// 统一执行 Pipeline
	_, err := pipe.Exec(ctx)
	if err != nil {
		log.Printf("error: pipeline Exec %v", err)
		return nil, err
	}

	// 先初始化 map，避免 nil map 导致 panic
	histories := make(map[int64]map[int64]float64)
	// 解析 pipeline 结果
	for i, str := range idStrs {
		vSet, err := historyCmds[i].Result()
		if err != nil {
			log.Printf("error: ZSet_Histories %v", err)
			return nil, err
		}

		userWeight := make(map[int64]float64)
		// 计算观看的权
		for _, creationId := range vSet {
			itemID, err := strconv.ParseInt(creationId, 10, 64)
			if err != nil {
				log.Printf("error: ParseInt %v", err)
			}
			userWeight[itemID] = viewWeight
		}
		id, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			log.Printf("error: ParseInt %v", err)
		}
		histories[id] = userWeight
	}

	return histories, nil
}

func GetAllItemUsers(idStrs []string) (map[int64]map[int64]float64, error) {
	const (
		viewWeight = 1
	)
	ctx := context.Background()
	pipe := CacheClient.Pipeline()

	// 用来存储 pipeline 请求的结果
	historyCmds := make([]*redis.StringSliceCmd, len(idStrs))

	// 依次遍历作品 ID，把请求加入 pipeline
	for i, str := range idStrs {
		historyKey := fmt.Sprintf("ZSet_Item_Users_%s", str) // 观看记录 (ZSet)

		// 用 ZRange 取 ZSet，避免用 SMembers 读错数据类型
		historyCmds[i] = pipe.ZRevRange(ctx, historyKey, 0, 199)
	}

	// 统一执行 Pipeline
	_, err := pipe.Exec(ctx)
	if err != nil {
		log.Printf("error: pipeline Exec %v", err)
		return nil, err
	}

	// 先初始化 map，避免 nil map 导致 panic
	Item_Users := make(map[int64]map[int64]float64)
	// 解析 pipeline 结果
	for i, str := range idStrs {
		vSet, err := historyCmds[i].Result()
		if err != nil {
			log.Printf("error: ZSet_Item_Users_ %v", err)
			return nil, err
		}

		creationWeight := make(map[int64]float64)
		// 计算观看的权
		for _, userIdStr := range vSet {
			userId, err := strconv.ParseInt(userIdStr, 10, 64)
			if err != nil {
				log.Printf("error: ParseInt %v", err)
			}
			creationWeight[userId] = viewWeight
		}
		creationId, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			log.Printf("error: ParseInt %v", err)
		}
		Item_Users[creationId] = creationWeight
	}

	return Item_Users, nil
}

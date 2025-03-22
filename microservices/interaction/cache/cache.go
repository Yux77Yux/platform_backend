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

type CacheMethodStruct struct {
	CacheClient CacheInterface
}

// GET

func (c *CacheMethodStruct) ToBaseInteraction(results []redis.Z) ([]*generated.Interaction, error) {
	count := len(results)
	res := make([]*generated.Interaction, count)
	for i, val := range results {
		idStr, ok := val.Member.(string)
		if !ok {
			err := fmt.Errorf("val.Member is not a string, actual type: %T, value: %v", val.Member, val.Member)
			log.Printf("err %v", err)
			return nil, err
		}
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			return nil, err
		}
		times := int64(math.Round(val.Score))
		timestamp := timestamppb.New(time.Unix(times, 0))
		res[i] = &generated.Interaction{}
		res[i].Base = &generated.BaseInteraction{
			CreationId: id,
		}
		res[i].SaveAt = timestamp
		res[i].UpdatedAt = timestamp
	}
	return res, nil
}

// 历史记录
func (c *CacheMethodStruct) GetHistories(ctx context.Context, userId int64, page int32) ([]*generated.Interaction, error) {
	const scope = 30
	start := int64((page - 1) * scope)
	stop := start + scope

	userIdStr := strconv.FormatInt(userId, 10)
	results, err := c.CacheClient.RevRangeZSetWithScore(ctx, "User_Histories", userIdStr, start, stop)
	if err != nil {
		return nil, err
	}
	res, err := c.ToBaseInteraction(results)
	if err != nil {
		return nil, err
	}

	for i := range res {
		res[i].Base.UserId = userId
	}
	return res, nil
}

// 收藏夹
func (c *CacheMethodStruct) GetCollections(ctx context.Context, userId int64, page int32) ([]*generated.Interaction, error) {
	const scope = 25
	start := int64((page - 1) * scope)
	stop := start + scope - 1
	userIdStr := strconv.FormatInt(userId, 10)

	results, err := c.CacheClient.RevRangeZSetWithScore(ctx, "User_Collections", userIdStr, start, stop)
	if err != nil {
		return nil, err
	}
	res, err := c.ToBaseInteraction(results)
	if err != nil {
		return nil, err
	}

	for i := range res {
		res[i].Base.UserId = userId
	}
	return res, nil
}

// Like
func (c *CacheMethodStruct) GetLikes(ctx context.Context, userId int64, page int32) ([]*generated.Interaction, error) {
	const scope = 25
	start := int64((page - 1) * scope)
	stop := start + scope - 1
	userIdStr := strconv.FormatInt(userId, 10)

	result, err := c.CacheClient.RevRangeZSetWithScore(ctx, "User_Likes", userIdStr, start, stop)
	if err != nil {
		return nil, fmt.Errorf("error: %w", err)
	}

	res, err := c.ToBaseInteraction(result)
	if err != nil {
		return nil, fmt.Errorf("error: %w", err)
	}

	for i := range res {
		res[i].Base.UserId = userId
	}
	return res, nil
}

// 观看作品的用户
func (c *CacheMethodStruct) GetUsers(ctx context.Context, creationId int64) ([]int64, error) {
	creationIdStr := strconv.FormatInt(creationId, 10)
	results, err := c.CacheClient.RevRangeZSet(ctx, "Item_Histories", creationIdStr, 0, 199)

	if err != nil {
		return nil, err
	}
	strs := results
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

// 展示页·点赞收藏情况
func (c *CacheMethodStruct) GetInteraction(ctx context.Context, interaction *generated.BaseInteraction) (*generated.Interaction, error) {
	userId := interaction.GetUserId()
	creationId := interaction.GetCreationId()

	pipe := c.CacheClient.Pipeline()
	isLike := fmt.Sprintf("ZSet_User_Likes_%d", userId)             // 自己是否点赞
	isCollection := fmt.Sprintf("ZSet_User_Collections_%d", userId) // 自己是否收藏

	creationIdStr := strconv.FormatInt(creationId, 10)

	// 使用 pipeline 执行多个查询
	setLikeCmd := pipe.ZScore(ctx, isLike, creationIdStr)
	zScoreCollectionCmd := pipe.ZScore(ctx, isCollection, creationIdStr)

	// 执行 pipeline
	_, err := pipe.Exec(ctx)
	if err != nil && err != redis.Nil {
		return nil, fmt.Errorf("redis pipeline execution failed: %w", err)
	}

	// 解析返回结果
	likeScore, err := setLikeCmd.Result()
	if err == redis.Nil {
		likeScore = -1 // 代表用户没有点赞
	} else if err != nil {
		return nil, fmt.Errorf("failed to get like score: %w", err)
	}

	collectionScore, err := zScoreCollectionCmd.Result()
	if err == redis.Nil {
		collectionScore = -1 // 代表用户没有收藏
	} else if err != nil {
		return nil, fmt.Errorf("failed to get collection score: %w", err)
	}

	action_tag := int32(1)
	if likeScore != -1 {
		action_tag |= 2
	}
	if collectionScore != -1 {
		action_tag |= 4
	}

	return &generated.Interaction{
		Base:      interaction,
		ActionTag: action_tag,
	}, nil
}

// RemSet
// CountSet
// GetMembersSet

// 数量不够随机拿已经发布的作品
func (c *CacheMethodStruct) GetPublicCreations(ctx context.Context, count int) ([]int64, error) {
	idStrs, err := c.CacheClient.GetRandZSetMember(ctx, "Public", "Creations", count)
	if err != nil {
		return nil, err
	}
	ids := make([]int64, 0, len(idStrs))
	for _, idStr := range idStrs {
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	return ids, nil
}

func (c *CacheMethodStruct) GetRecommendBaseUser(ctx context.Context, id int64) ([]int64, int64, error) {
	const popCount = 6
	pipe := c.CacheClient.Pipeline()

	sliceCmd := pipe.SPopN(ctx, fmt.Sprintf("Set_RecommendBaseUser_%d", id), popCount)
	intCmd := pipe.SCard(ctx, fmt.Sprintf("Set_RecommendBaseUser_%d", id))

	_, err := pipe.Exec(ctx)
	if err != nil && err != redis.Nil {
		return nil, -1, err
	}

	strs, err := sliceCmd.Result()
	if err != nil {
		return nil, -1, err
	}
	count, err := intCmd.Result()
	if err != nil {
		return nil, -1, err
	}

	ids := make([]int64, len(strs))
	for i, str := range strs {
		creationId, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			return nil, -1, err
		}
		ids[i] = creationId
	}

	return ids, count, nil
}

func (c *CacheMethodStruct) GetRecommendBaseItem(ctx context.Context, id int64) ([]int64, bool, error) {
	const popCount = 10
	pipe := c.CacheClient.Pipeline()

	sliceCmd := pipe.SRandMemberN(ctx, fmt.Sprintf("Set_RecommendBaseItem_%d", id), popCount)
	floatCmd := pipe.ZScore(ctx, "ZSet_RecommendBaseItem_Creation", strconv.FormatInt(id, 10))

	_, err := pipe.Exec(ctx)
	if err != nil && err != redis.Nil {
		return nil, false, err
	}

	strs, err := sliceCmd.Result()
	if err != nil {
		if err == redis.Nil {
			return nil, true, nil
		}
		return nil, false, err
	}

	score, err := floatCmd.Result()
	if err != nil {
		if err == redis.Nil {
			return nil, true, nil
		}
		return nil, false, err
	}

	// 获取当前时间戳（秒）
	now := time.Now().Unix()
	reset := now-int64(score) >= 86400

	ids := make([]int64, len(strs))
	for i, str := range strs {
		creationId, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			return nil, false, err
		}
		ids[i] = creationId
	}

	return ids, reset, nil
}

// 查看是否过期，是否重新计算

// POST
func (c *CacheMethodStruct) SetRecommendBaseUser(ctx context.Context, id int64, ids []int64) error {
	count := len(ids)
	values := make([]any, count)
	for i, val := range ids {
		values[i] = val
	}

	err := c.CacheClient.AddToSet(ctx, "RecommendBaseUser", strconv.FormatInt(id, 10), values...)
	if err != nil {
		return err
	}
	return nil
}

func (c *CacheMethodStruct) SetRecommendBaseItem(ctx context.Context, id int64, ids []int64) error {
	count := len(ids)
	values := make([]any, count)
	for i, val := range ids {
		values[i] = val
	}

	pipe := c.CacheClient.Pipeline()
	// 作品的相似视频集合
	pipe.SAdd(ctx, fmt.Sprintf("Set_RecommendBaseItem_%d", id), values...)
	// 作品的相似视频集合的过期时间
	pipe.ZAdd(ctx, "ZSet_RecommendBaseItem_Creation", &redis.Z{
		Member: id,
		Score:  float64(time.Now().Unix()),
	})

	_, err := pipe.Exec(ctx)
	if err != nil && err != redis.Nil {
		return err
	}

	return nil
}

// POST & UPDATE
// 历史记录 更新时间戳
func (c *CacheMethodStruct) UpdateHistories(ctx context.Context, data []*generated.OperateInteraction) error {
	pipe := c.CacheClient.Pipeline()
	for _, option := range data {
		base := option.GetBase()

		userId := base.GetUserId()
		creationId := base.GetCreationId()
		timestampScore := option.GetUpdatedAt().GetSeconds()

		// 用户的历史记录 基于用户协同过滤
		keyUser := fmt.Sprintf("ZSet_User_Histories_%d", userId)
		pipe.ZAdd(ctx, keyUser, &redis.Z{
			Score:  float64(timestampScore),
			Member: creationId,
		})

		// 记录观看作品的用户 基于物品协同过滤
		keyItem := fmt.Sprintf("ZSet_Item_Histories_%d", creationId)
		pipe.ZAdd(ctx, keyItem, &redis.Z{
			Score:  float64(timestampScore),
			Member: userId,
		})

	}
	_, err := pipe.Exec(ctx)
	if err != nil && err != redis.Nil {
		return fmt.Errorf("pipeline execution failed: %w", err)
	}
	return nil
}

// 收藏夹
func (c *CacheMethodStruct) ModifyCollections(ctx context.Context, data []*generated.OperateInteraction) error {
	pipe := c.CacheClient.Pipeline()
	for _, option := range data {
		base := option.GetBase()

		userId := base.GetUserId()
		creationId := base.GetCreationId()
		timestampScore := option.GetUpdatedAt().GetSeconds()

		// 用户的收藏记录 基于用户协同过滤
		keyUser := fmt.Sprintf("ZSet_User_Collections_%d", userId)
		pipe.ZAdd(ctx, keyUser, &redis.Z{
			Score:  float64(timestampScore),
			Member: creationId,
		})

		// 记录收藏作品的用户 基于物品协同过滤
		keyItem := fmt.Sprintf("ZSet_Item_Collections_%d", creationId)
		pipe.ZAdd(ctx, keyItem, &redis.Z{
			Score:  float64(timestampScore),
			Member: userId,
		})
	}
	_, err := pipe.Exec(ctx)
	if err != nil && err != redis.Nil {
		return fmt.Errorf("pipeline execution failed: %w", err)
	}

	return nil
}

// 点赞
func (c *CacheMethodStruct) ModifyLike(ctx context.Context, data []*generated.OperateInteraction) error {
	pipe := c.CacheClient.Pipeline()
	for _, option := range data {
		base := option.GetBase()

		userId := base.GetUserId()
		creationId := base.GetCreationId()
		timestampScore := option.GetUpdatedAt().GetSeconds()

		// 用户的点赞记录 基于用户协同过滤
		keyUser := fmt.Sprintf("ZSet_User_Likes_%d", userId)
		pipe.ZAdd(ctx, keyUser, &redis.Z{
			Score:  float64(timestampScore),
			Member: creationId,
		})

		// 记录点赞作品的用户 基于物品协同过滤
		keyItem := fmt.Sprintf("ZSet_Item_Likes_%d", creationId)
		pipe.ZAdd(ctx, keyItem, &redis.Z{
			Score:  float64(timestampScore),
			Member: userId,
		})
	}
	_, err := pipe.Exec(ctx)
	if err != nil && err != redis.Nil {
		return fmt.Errorf("pipeline execution failed: %w", err)
	}

	return nil
}

// DELETE

func (c *CacheMethodStruct) DelHistories(ctx context.Context, data []*generated.BaseInteraction) error {
	pipe := c.CacheClient.Pipeline()
	for _, base := range data {
		userId := base.GetUserId()
		creationId := base.GetCreationId()

		key := fmt.Sprintf("ZSet_User_Histories_%d", userId)
		pipe.ZRem(ctx, key, creationId)
	}

	_, err := pipe.Exec(ctx)
	if err != nil && err != redis.Nil {
		return err
	}
	return nil
}

func (c *CacheMethodStruct) DelCollections(ctx context.Context, data []*generated.BaseInteraction) error {
	pipe := c.CacheClient.Pipeline()
	for _, base := range data {
		userId := base.GetUserId()
		creationId := base.GetCreationId()

		key := fmt.Sprintf("ZSet_User_Collections_%d", userId)
		pipe.ZRem(ctx, key, creationId)
	}

	_, err := pipe.Exec(ctx)
	if err != nil && err != redis.Nil {
		return err
	}
	return nil
}

func (c *CacheMethodStruct) DelLike(ctx context.Context, data []*generated.BaseInteraction) error {
	pipe := c.CacheClient.Pipeline()
	for _, base := range data {
		userId := base.GetUserId()
		creationId := base.GetCreationId()

		key := fmt.Sprintf("ZSet_User_Likes_%d", userId)
		pipe.ZRem(ctx, key, creationId)
	}

	_, err := pipe.Exec(ctx)
	if err != nil && err != redis.Nil {
		return err
	}
	return nil
}

// Scan
// 拿到别人的历史记录
func (c *CacheMethodStruct) ScanZSetsByHistories(ctx context.Context) ([]string, error) {
	results, _, err := c.CacheClient.ScanZSet(ctx, "User_Histories", "*", 0, 2500)
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

func (c *CacheMethodStruct) ScanZSetsByCreationId(ctx context.Context) ([]string, error) {
	results, _, err := c.CacheClient.ScanZSet(ctx, "Item_Histories", "*", 0, 2500)
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

func (c *CacheMethodStruct) GetAllInteractions(ctx context.Context, idStrs []string) (map[int64]map[int64]float64, error) {
	const (
		viewWeight = 1
	)
	pipe := c.CacheClient.Pipeline()

	// 用来存储 pipeline 请求的结果
	historyCmds := make([]*redis.StringSliceCmd, len(idStrs))

	// 依次遍历用户 ID，把请求加入 pipeline
	for i, str := range idStrs {
		historyKey := fmt.Sprintf("ZSet_User_Histories_%s", str) // 观看记录 (ZSet)

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
			log.Printf("error: ZSet_User_Histories %v", err)
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

func (c *CacheMethodStruct) GetAllItemUsers(ctx context.Context, idStrs []string) (map[int64]map[int64]float64, error) {
	const (
		viewWeight = 1
	)
	pipe := c.CacheClient.Pipeline()

	// 用来存储 pipeline 请求的结果
	historyCmds := make([]*redis.StringSliceCmd, len(idStrs))

	// 依次遍历作品 ID，把请求加入 pipeline
	for i, str := range idStrs {
		historyKey := fmt.Sprintf("ZSet_Item_Histories_%s", str) // 观看记录 (ZSet)

		// 用 ZRange 取 ZSet，避免用 SMembers 读错数据类型
		historyCmds[i] = pipe.ZRevRange(ctx, historyKey, 0, 199)
	}

	// 统一执行 Pipeline
	_, err := pipe.Exec(ctx)
	if err != nil && err != redis.Nil {
		log.Printf("error: pipeline Exec %v", err)
		return nil, err
	}

	// 先初始化 map，避免 nil map 导致 panic
	Item_Users := make(map[int64]map[int64]float64)
	// 解析 pipeline 结果
	for i, str := range idStrs {
		vSet, err := historyCmds[i].Result()
		if err != nil {
			log.Printf("error: ZSet_Item_Histories_ %v", err)
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

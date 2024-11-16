package redis_cache

import (
	"context"
	"crypto/sha256"
	"fmt"
	"github.com/go-redis/redis/v8"
	"hash/fnv"
	"log"
	"math"
	"strings"
)

type RedisMethods interface {
	Close()

	DelRelatedKeys(ctx context.Context, kindPrefix string, kindType string) error
	DelKey(ctx context.Context, kind string, unique string) error

	SetString(ctx context.Context, kind string, unique string, value interface{}) error
	ModifyString(ctx context.Context, kind string, unique string, value interface{}) error
	GetString(ctx context.Context, kind string, unique string) (string, error)
	ExistsString(ctx context.Context, kind string, unique string) (bool, error)

	ScanHash(ctx context.Context, kind string, unique string, fliter string, cursor uint64, count int64) ([]string, uint64, error)
	SetFieldsHash(ctx context.Context, kind string, unique string, fieldValues ...interface{}) error
	SetFieldHash(ctx context.Context, kind string, unique string, field string, value interface{}) error
	ModifyFieldHash(ctx context.Context, kind string, unique string, field string, value interface{}) error
	GetHash(ctx context.Context, kind string, unique string, field string) (string, error)
	GetAnyHash(ctx context.Context, kind string, unique string, fields ...string) ([]interface{}, error)
	GetAllHash(ctx context.Context, kind string, unique string) (map[string]string, error)
	GetFieldsHash(ctx context.Context, kind string, unique string) ([]string, error)
	GetValuesHash(ctx context.Context, kind string, unique string) ([]string, error)
	ExistsHash(ctx context.Context, kind string, unique string, field string) (bool, error)
	GetLenHash(ctx context.Context, kind string, unique string, field string) (int64, error)
	DelHash(ctx context.Context, kind string, unique string, fields ...string) (int64, error)

	LPushList(ctx context.Context, direction string, kind string, unique string, value ...interface{}) error
	RPushList(ctx context.Context, direction string, kind string, unique string, value ...interface{}) error
	LPopList(ctx context.Context, kind string, unique string) (string, error)
	RPopList(ctx context.Context, kind string, unique string) (string, error)
	InsertBeforeList(ctx context.Context, kind string, unique string, pivot, value interface{}) error
	InsertAfterList(ctx context.Context, kind string, unique string, pivot, value interface{}) error
	IndexList(ctx context.Context, kind string, unique string, index int64) (string, error)
	FindElementList(ctx context.Context, kind string, unique string, value string) (int64, error)
	FindElementsList(ctx context.Context, kind string, unique string, value string, count int64) ([]int64, error)
	TrimList(ctx context.Context, kind string, unique string, start, stop int64) error
	GetLenList(ctx context.Context, kind string, unique string) (int64, error)
	GetElementsList(ctx context.Context, kind string, unique string, start, stop int64) ([]string, error)

	ScanSet(ctx context.Context, kind string, unique string, fliter string, cursor uint64, count int64) ([]string, uint64, error)
	AddToSet(ctx context.Context, kind string, unique string, value interface{}) error
	RemSet(ctx context.Context, kind string, unique string, value interface{}) error
	ExistsInSet(ctx context.Context, kind string, unique string, value interface{}) (bool, error)
	CountSet(ctx context.Context, kind string, unique string) (int64, error)
	GetMembersSet(ctx context.Context, kind string, unique string) ([]string, error)

	ScanZSet(ctx context.Context, kind string, unique string, fliter string, cursor uint64, count int64) ([]string, uint64, error)
	AddZSet(ctx context.Context, kind string, unique string, member string, score float64) error
	ModifyScoreZSet(ctx context.Context, kind string, unique string, member string, score float64) error
	ZRemMemberZSet(ctx context.Context, kind string, unique string, members ...interface{}) error
	GetRankZSet(ctx context.Context, kind string, unique string, member string) (int64, error)
	GetScoreZSet(ctx context.Context, kind string, unique string, member string) (float64, error)
	RangeZSet(ctx context.Context, kind string, unique string, start, stop int64) ([]string, error)
	RevRangeZSet(ctx context.Context, kind string, unique string, start, stop int64) ([]string, error)
	RangeByScoreZSet(ctx context.Context, kind, unique, min, max string, offset, count int64) ([]string, error)
	RevRangeByScoreZSet(ctx context.Context, kind, unique, min, max string, offset, count int64) ([]string, error)
	GetCountZSet(ctx context.Context, kind, unique string) (int64, error)
	GetScoreCountZSet(ctx context.Context, kind, unique, min, max string) (int64, error)

	SetBit(ctx context.Context, kind string, unique string, offset int64, value bool) error
	GetBit(ctx context.Context, kind string, unique string, offset int64) (bool, error)
	ClearBit(ctx context.Context, kind string, unique string, offset int64) error
	ClearRangeBits(ctx context.Context, kind string, unique string, startOffset, endOffset int64) error
	CountBit(ctx context.Context, kind string, unique string, startOffset, endOffset int64) (int64, error)
	OpertorBit(ctx context.Context, operation string, destKind string, unique string, srcKinds []string) error
	ModifyBit(ctx context.Context, kind string, unique string, offset int64, newValue bool) error
	ResetBitmap(ctx context.Context, kind string, unique string, length int64) error
	FindPositionBit(ctx context.Context, kind string, unique string, bit bool) (int64, error)
	CombineBitmaps(ctx context.Context, destKind string, unique string, srcKinds []string) error
	AddToBloomFilter(ctx context.Context, kind string, unique string) error
	CheckBloomFilter(ctx context.Context, kind string, unique string) (bool, error)

	BitField(ctx context.Context, kind string, unique string, command ...interface{}) ([]int64, error)
}

func OpenRedis(connStr string, password string) (*RedisClient, error) {
	var (
		redisClient *redis.Client
		err         error
	)

	redisClient = redis.NewClient(&redis.Options{
		Addr:     connStr,
		Password: password,
	})

	_, err = redisClient.Ping(context.Background()).Result()
	if err != nil {
		err = fmt.Errorf("failed to connect redis client: %w", err)
		return nil, err
	}

	return &RedisClient{
		redisClient: redisClient,
	}, nil
}

type RedisClient struct {
	redisClient *redis.Client
}

func (r *RedisClient) Close() {
	if err := r.redisClient.Close(); err != nil {
		wiredErr := fmt.Errorf("failed to close redis client: %w", err)
		log.Printf("error: %v", wiredErr)
	}
}

// 类型的删除

// 删除指定类型符合要求的键
func (r *RedisClient) DelRelatedKeys(ctx context.Context, kindPrefix string, kindType string) error {
	// 确认有效的键类型
	switch kindType {
	case "list", "string", "zset", "set", "hash", "stream", "hyperloglog", "geospatial":
	case "bitmap":
		// 将 bitmap 类型映射到 string，因为 Redis 中位图是字符串类型
		kindType = "string"
	default:
		return fmt.Errorf("this type is not found")
	}

	keys, err := r.redisClient.Keys(ctx, fmt.Sprintf("%s_*", kindPrefix)).Result()
	if err != nil {
		return err
	}
	if len(keys) == 0 {
		return nil
	}
	for _, key := range keys {
		//获取键的类型
		_type, err := r.redisClient.Type(ctx, key).Result()
		if err != nil {
			return err
		}

		if _type == kindType {
			// 删除键
			if err := r.redisClient.Del(ctx, key).Err(); err != nil {
				return err
			}
		}
	}
	return nil
}

// 删除Key 适用于删除整个Key
func (r *RedisClient) DelKey(ctx context.Context, kind string, unique string) error {
	key := fmt.Sprintf("%s_%s", kind, unique)
	return r.redisClient.Del(ctx, key).Err()
}

// String

func (r *RedisClient) SetString(ctx context.Context, kind string, unique string, value interface{}) error {
	key := fmt.Sprintf("%s_%s", kind, unique)
	return r.redisClient.SetNX(ctx, key, value, 0).Err()
}

func (r *RedisClient) ModifyString(ctx context.Context, kind string, unique string, value interface{}) error {
	key := fmt.Sprintf("%s_%s", kind, unique)
	return r.redisClient.SetXX(ctx, key, value, 0).Err()
}

func (r *RedisClient) GetString(ctx context.Context, kind string, unique string) (string, error) {
	key := fmt.Sprintf("%s_%s", kind, unique)
	value, err := r.redisClient.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}

	return value, nil
}

func (r *RedisClient) ExistsString(ctx context.Context, kind string, unique string) (bool, error) {
	key := fmt.Sprintf("%s_%s", kind, unique)
	val, err := r.redisClient.Exists(ctx, key).Result()
	if err != nil {
		return false, fmt.Errorf("redis_utils checking key is not exists: %v", err)
	}

	return val > 0, nil
}

// Hash

// HScan 适用于字段多的大哈希表
func (r *RedisClient) ScanHash(ctx context.Context, kind string, unique string, fliter string, cursor uint64, count int64) ([]string, uint64, error) {
	key := fmt.Sprintf("%s_%s", kind, unique)
	result, newCursor, err := r.redisClient.HScan(ctx, key, cursor, fliter, count).Result()
	if err != nil {
		return nil, 0, err
	}

	return result, newCursor, nil
}

// 用于批量 设置/更新
// fieldValues : name "Mike" age 20
func (r *RedisClient) SetFieldsHash(ctx context.Context, kind string, unique string, fieldValues ...interface{}) error {
	key := fmt.Sprintf("%s_%s", kind, unique)
	return r.redisClient.HSet(ctx, key, fieldValues...).Err()
}

// 一次一个字段，不存在时才加入
func (r *RedisClient) SetFieldHash(ctx context.Context, kind string, unique string, field string, value interface{}) error {
	key := fmt.Sprintf("%s_%s", kind, unique)
	return r.redisClient.HSetNX(ctx, key, field, value).Err()
}

// 更新一个字段，若存在则更新
func (r *RedisClient) ModifyFieldHash(ctx context.Context, kind string, unique string, field string, value interface{}) error {
	key := fmt.Sprintf("%s_%s", kind, unique)

	exists, err := r.redisClient.HExists(ctx, key, field).Result()
	if err != nil {
		return err
	}

	if !exists {
		return nil
	}

	return r.redisClient.HSet(ctx, key, field, value).Err()
}

// 仅获取一个字段值
func (r *RedisClient) GetHash(ctx context.Context, kind string, unique string, field string) (string, error) {
	key := fmt.Sprintf("%s_%s", kind, unique)
	value, err := r.redisClient.HGet(ctx, key, field).Result()
	if err != nil {
		return "", err
	}
	return value, nil
}

// 获取复数字段名和值
func (r *RedisClient) GetAnyHash(ctx context.Context, kind string, unique string, fields ...string) ([]interface{}, error) {
	key := fmt.Sprintf("%s_%s", kind, unique)
	result, err := r.redisClient.HMGet(ctx, key, fields...).Result()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// 获取全部字段名和值
func (r *RedisClient) GetAllHash(ctx context.Context, kind string, unique string) (map[string]string, error) {
	key := fmt.Sprintf("%s_%s", kind, unique)
	result, err := r.redisClient.HGetAll(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// 获取哈希表中所有字段名
func (r *RedisClient) GetFieldsHash(ctx context.Context, kind string, unique string) ([]string, error) {
	key := fmt.Sprintf("%s_%s", kind, unique)
	result, err := r.redisClient.HKeys(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// 获取哈希表中所有字段值
func (r *RedisClient) GetValuesHash(ctx context.Context, kind string, unique string) ([]string, error) {
	key := fmt.Sprintf("%s_%s", kind, unique)
	result, err := r.redisClient.HVals(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// 获取哈希表中是否有此字段
func (r *RedisClient) ExistsHash(ctx context.Context, kind string, unique string, field string) (bool, error) {
	key := fmt.Sprintf("%s_%s", kind, unique)
	exists, err := r.redisClient.HExists(ctx, key, field).Result()
	if err != nil {
		return false, err
	}
	return exists, nil
}

// 获取哈希表中字段数量
func (r *RedisClient) GetLenHash(ctx context.Context, kind string, unique string, field string) (int64, error) {
	key := fmt.Sprintf("%s_%s", kind, unique)
	exists, err := r.redisClient.HLen(ctx, key).Result()
	if err != nil {
		return 0, err
	}
	return exists, nil
}

// 删除哈希表中的某些字段
func (r *RedisClient) DelHash(ctx context.Context, kind string, unique string, fields ...string) (int64, error) {
	key := fmt.Sprintf("%s_%s", kind, unique)
	num, err := r.redisClient.HDel(ctx, key, fields...).Result()
	if err != nil {
		return 0, err
	}
	return num, nil
}

// List

// 向 List 头部推送元素
func (r *RedisClient) LPushList(ctx context.Context, direction string, kind string, unique string, value ...interface{}) error {
	key := fmt.Sprintf("%s_%s", kind, unique)
	return r.redisClient.LPush(ctx, key, value...).Err()
}

// 向 List 尾部推送元素
func (r *RedisClient) RPushList(ctx context.Context, direction string, kind string, unique string, value ...interface{}) error {
	key := fmt.Sprintf("%s_%s", kind, unique)
	return r.redisClient.RPush(ctx, key, value...).Err()
}

// 从 List 头部弹出元素
func (r *RedisClient) LPopList(ctx context.Context, kind string, unique string) (string, error) {
	key := fmt.Sprintf("%s_%s", kind, unique)
	value, err := r.redisClient.LPop(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return value, nil
}

// 从 List 尾部弹出元素
func (r *RedisClient) RPopList(ctx context.Context, kind string, unique string) (string, error) {
	key := fmt.Sprintf("%s_%s", kind, unique)
	value, err := r.redisClient.RPop(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return value, nil
}

// 从 List 指定元素之前插入
func (r *RedisClient) InsertBeforeList(ctx context.Context, kind string, unique string, pivot, value interface{}) error {
	key := fmt.Sprintf("%s_%s", kind, unique)
	err := r.redisClient.LInsertBefore(ctx, key, pivot, value).Err()
	if err != nil {
		return err
	}
	return nil
}

// 从 List 指定元素之后插入
func (r *RedisClient) InsertAfterList(ctx context.Context, kind string, unique string, pivot, value interface{}) error {
	key := fmt.Sprintf("%s_%s", kind, unique)
	err := r.redisClient.LInsertAfter(ctx, key, pivot, value).Err()
	if err != nil {
		return err
	}
	return nil
}

// 从 List 获取指定位置元素
func (r *RedisClient) IndexList(ctx context.Context, kind string, unique string, index int64) (string, error) {
	key := fmt.Sprintf("%s_%s", kind, unique)
	value, err := r.redisClient.LIndex(ctx, key, index).Result()
	if err != nil {
		return "", err
	}
	return value, nil
}

// 返回 List 中某个元素的第一个索引，MaxLen 查找最大长度，Rank 查找方向
func (r *RedisClient) FindElementList(ctx context.Context, kind string, unique string, value string) (int64, error) {
	key := fmt.Sprintf("%s_%s", kind, unique)
	result, err := r.redisClient.LPos(ctx, key, value, redis.LPosArgs{Rank: -1, MaxLen: 0}).Result()
	if err != nil {
		return 0, err
	}
	return result, nil
}

// 返回 List 中元素的索引，count 索引个数，MaxLen 查找最大长度，Rank 查找方向
func (r *RedisClient) FindElementsList(ctx context.Context, kind string, unique string, value string, count int64) ([]int64, error) {
	key := fmt.Sprintf("%s_%s", kind, unique)
	result, err := r.redisClient.LPosCount(ctx, key, value, count, redis.LPosArgs{Rank: -1, MaxLen: 0}).Result()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// 从 List 截取列表的一部分 并删除不在该范围内的元素
func (r *RedisClient) TrimList(ctx context.Context, kind string, unique string, start, stop int64) error {
	if start == 0 && stop == 0 {
		stop = -1
	}

	key := fmt.Sprintf("%s_%s", kind, unique)
	err := r.redisClient.LTrim(ctx, key, start, stop).Err()
	if err != nil {
		return err
	}
	return nil
}

// 获取 List 长度
func (r *RedisClient) GetLenList(ctx context.Context, kind string, unique string) (int64, error) {
	key := fmt.Sprintf("%s_%s", kind, unique)
	length, err := r.redisClient.LLen(ctx, key).Result()
	if err != nil {
		return 0, err
	}
	return length, nil
}

// 按范围获取 List 中的元素，默认全部
func (r *RedisClient) GetElementsList(ctx context.Context, kind string, unique string, start, stop int64) ([]string, error) {
	if start == 0 && stop == 0 {
		stop = -1
	}

	key := fmt.Sprintf("%s_%s", kind, unique)
	elements, err := r.redisClient.LRange(ctx, key, start, stop).Result()
	if err != nil {
		return nil, err
	}
	return elements, nil
}

// Set

// SScan
func (r *RedisClient) ScanSet(ctx context.Context, kind string, unique string, fliter string, cursor uint64, count int64) ([]string, uint64, error) {
	key := fmt.Sprintf("%s_%s", kind, unique)
	result, newCursor, err := r.redisClient.SScan(ctx, key, cursor, fliter, count).Result()
	if err != nil {
		return nil, 0, err
	}

	return result, newCursor, nil
}

// 向 Set 中添加/更新元素
func (r *RedisClient) AddToSet(ctx context.Context, kind string, unique string, value interface{}) error {
	key := fmt.Sprintf("%s_%s", kind, unique)
	return r.redisClient.SAdd(ctx, key, value).Err()
}

// 在 Set 中删除元素
func (r *RedisClient) RemSet(ctx context.Context, kind string, unique string, value interface{}) error {
	key := fmt.Sprintf("%s_%s", kind, unique)
	return r.redisClient.SRem(ctx, key, value).Err()
}

// 检查 Set 中是否存在某个元素
func (r *RedisClient) ExistsInSet(ctx context.Context, kind string, unique string, value interface{}) (bool, error) {
	key := fmt.Sprintf("%s_%s", kind, unique)
	exists, err := r.redisClient.SIsMember(ctx, key, value).Result()
	if err != nil {
		return false, err
	}
	return exists, nil
}

// 在 Set 中统计元素个数
func (r *RedisClient) CountSet(ctx context.Context, kind string, unique string) (int64, error) {
	key := fmt.Sprintf("%s_%s", kind, unique)
	num, err := r.redisClient.SCard(ctx, key).Result()
	if err != nil {
		return 0, err
	}

	return num, nil
}

// 获取 Set 中的所有成员
func (r *RedisClient) GetMembersSet(ctx context.Context, kind string, unique string) ([]string, error) {
	key := fmt.Sprintf("%s_%s", kind, unique)
	members, err := r.redisClient.SMembers(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	return members, nil
}

// ZSet

// 没做 ZInter交集 ZDiff差集 ZPopMax弹出最高分数成员 ZPopMin弹出最低分数成员 ZCard获取成员总数

// ZScan
func (r *RedisClient) ScanZSet(ctx context.Context, kind string, unique string, fliter string, cursor uint64, count int64) ([]string, uint64, error) {
	key := fmt.Sprintf("%s_%s", kind, unique)
	result, newCursor, err := r.redisClient.ZScan(ctx, key, cursor, fliter, count).Result()
	if err != nil {
		return nil, 0, err
	}

	return result, newCursor, nil
}

// 正常添加, 若重复添加则不会添加
func (r *RedisClient) AddZSet(ctx context.Context, kind string, unique string, member string, score float64) error {
	key := fmt.Sprintf("%s_%s", kind, unique)
	return r.redisClient.ZAddNX(ctx, key, &redis.Z{Score: score, Member: member}).Err()
}

// 修改，当没有该成员则不会更新
func (r *RedisClient) ModifyScoreZSet(ctx context.Context, kind string, unique string, member string, score float64) error {
	key := fmt.Sprintf("%s_%s", kind, unique)
	return r.redisClient.ZAddXX(ctx, key, &redis.Z{Score: score, Member: member}).Err()
}

// 删除，当没有该成员则无事发生
func (r *RedisClient) ZRemMemberZSet(ctx context.Context, kind string, unique string, members ...interface{}) error {
	key := fmt.Sprintf("%s_%s", kind, unique)
	return r.redisClient.ZRem(ctx, key, members...).Err()
}

// 返回排名
func (r *RedisClient) GetRankZSet(ctx context.Context, kind string, unique string, member string) (int64, error) {
	key := fmt.Sprintf("%s_%s", kind, unique)
	rank, err := r.redisClient.ZRank(ctx, key, member).Result()
	if err != nil {
		return 0, err
	}
	return rank, nil
}

// 返回分数
func (r *RedisClient) GetScoreZSet(ctx context.Context, kind string, unique string, member string) (float64, error) {
	key := fmt.Sprintf("%s_%s", kind, unique)

	rank, err := r.redisClient.ZScore(ctx, key, member).Result()
	if err == redis.Nil {
		return 0, fmt.Errorf("member does not exist")
	} else if err != nil {
		return 0, err
	} else {
		return rank, nil
	}
}

// 返回基于分数递增的排名(索引下标)范围的成员
func (r *RedisClient) RangeZSet(ctx context.Context, kind string, unique string, start, stop int64) ([]string, error) {
	if start == 0 && stop == 0 {
		stop = -1
	}
	key := fmt.Sprintf("%s_%s", kind, unique)
	members, err := r.redisClient.ZRange(ctx, key, start, stop).Result()
	if err == redis.Nil {
		return nil, fmt.Errorf("member does not exist")
	} else if err != nil {
		return nil, err
	} else {
		return members, nil
	}
}

// 返回基于分数递减的排名(索引下标)范围的成员
func (r *RedisClient) RevRangeZSet(ctx context.Context, kind string, unique string, start, stop int64) ([]string, error) {
	if start == 0 && stop == 0 {
		stop = -1
	}
	key := fmt.Sprintf("%s_%s", kind, unique)
	members, err := r.redisClient.ZRevRange(ctx, key, start, stop).Result()
	if err == redis.Nil {
		return nil, fmt.Errorf("member does not exist")
	} else if err != nil {
		return nil, err
	} else {
		return members, nil
	}
}

// 返回基于分数递增的分数范围（min，max）的成员
func (r *RedisClient) RangeByScoreZSet(ctx context.Context, kind, unique, min, max string, offset, count int64) ([]string, error) {
	opt := &redis.ZRangeBy{
		Min:    min,
		Max:    max,
		Offset: offset,
	}

	if count != 0 {
		opt.Count = count
	}

	key := fmt.Sprintf("%s_%s", kind, unique)
	members, err := r.redisClient.ZRangeByScore(ctx, key, opt).Result()
	if err == redis.Nil {
		return nil, fmt.Errorf("member does not exist")
	} else if err != nil {
		return nil, err
	} else {
		return members, nil
	}
}

// 返回基于分数递减的分数范围（min，max）的成员
func (r *RedisClient) RevRangeByScoreZSet(ctx context.Context, kind, unique, min, max string, offset, count int64) ([]string, error) {
	opt := &redis.ZRangeBy{
		Min:    min,
		Max:    max,
		Offset: offset,
	}

	if count != 0 {
		opt.Count = count
	}

	key := fmt.Sprintf("%s_%s", kind, unique)
	members, err := r.redisClient.ZRevRangeByScore(ctx, key, opt).Result()
	if err == redis.Nil {
		return nil, fmt.Errorf("member does not exist")
	} else if err != nil {
		return nil, err
	} else {
		return members, nil
	}
}

// 返回有序集合的成员总数
func (r *RedisClient) GetCountZSet(ctx context.Context, kind, unique string) (int64, error) {
	key := fmt.Sprintf("%s_%s", kind, unique)
	rank, err := r.redisClient.ZCard(ctx, key).Result()
	if err != nil {
		return 0, err
	}

	return rank, nil
}

// 返回指定分数范围内的 成员数量
// min 和 max: 定义分数范围的字符串,
// -inf负无穷大 ，+inf正无穷大
// "(2", "4" 大于2小于等于4的 成员数量
func (r *RedisClient) GetScoreCountZSet(ctx context.Context, kind, unique, min, max string) (int64, error) {
	key := fmt.Sprintf("%s_%s", kind, unique)
	rank, err := r.redisClient.ZCount(ctx, key, min, max).Result()
	if err != nil {
		return 0, err
	}
	return rank, nil
}

// Bit
// 设置位
func (r *RedisClient) SetBit(ctx context.Context, kind string, unique string, offset int64, value bool) error {
	if offset < 0 {
		return fmt.Errorf("invalid offset: %d, must be non-negative", offset)
	}

	bit := 0
	if value {
		bit = 1
	}
	key := fmt.Sprintf("%s_%s", kind, unique)
	return r.redisClient.SetBit(ctx, key, offset, bit).Err()
}

// 获取位
func (r *RedisClient) GetBit(ctx context.Context, kind string, unique string, offset int64) (bool, error) {
	if offset < 0 {
		return false, fmt.Errorf("invalid offset: %d, must be non-negative", offset)
	}

	key := fmt.Sprintf("%s_%s", kind, unique)
	value, err := r.redisClient.GetBit(ctx, key, offset).Result()
	if err != nil {
		return false, err
	}
	return value == 1, nil
}

// 清除特定位（设置为0）
func (r *RedisClient) ClearBit(ctx context.Context, kind string, unique string, offset int64) error {
	key := fmt.Sprintf("%s_%s", kind, unique)
	return r.redisClient.SetBit(ctx, key, offset, 0).Err()
}

// 清除指定范围的位值
func (r *RedisClient) ClearRangeBits(ctx context.Context, kind string, unique string, startOffset, endOffset int64) error {
	key := fmt.Sprintf("%s_%s", kind, unique)
	pipeline := r.redisClient.Pipeline()

	for offset := startOffset; offset <= endOffset; offset++ {
		pipeline.SetBit(ctx, key, offset, 0)
	}

	_, err := pipeline.Exec(ctx)
	return err
}

// 计算位图某个范围中 1 的个数
func (r *RedisClient) CountBit(ctx context.Context, kind string, unique string, startOffset, endOffset int64) (int64, error) {
	key := fmt.Sprintf("%s_%s", kind, unique)
	count, err := r.redisClient.BitCount(ctx, key, &redis.BitCount{Start: startOffset, End: endOffset}).Result() // 计算整个键的位图
	if err != nil {
		return 0, err
	}
	return count, nil
}

// 位操作（AND、OR、XOR、NOT）
// 假设有如下的 Redis 键：
// A_user123，B_user123，C_user123，
// 这些键存储的是二进制位图数据。
// 例子1：
// 如果你要执行 A_user123、B_user123 和 C_user123
// 的按位与（AND）操作并将结果存储到 result_user123 中，你可以调用：
// redisClient.BitOp(ctx, "AND", "result", "user123", []string{"A", "B", "C"})
// 例子2：
// 如果你要对 A_user123 执行按位取反（NOT）操作并将结果存储到 result_user123 中：
// redisClient.BitOp(ctx, "NOT", "result", "user123", []string{"A"})
// NOT 仅支持对单个键执行 NOT 操作,忽略第二个srcKinds
func (r *RedisClient) OpertorBit(ctx context.Context, operation string, destKind string, unique string, srcKinds []string) error {
	destKey := fmt.Sprintf("%s_%s", destKind, unique)
	keys := make([]string, len(srcKinds))
	for i, srcKind := range srcKinds {
		keys[i] = fmt.Sprintf("%s_%s", srcKind, unique)
	}
	operation = strings.ToUpper(operation)
	// 根据操作类型执行不同的位操作
	switch operation {
	case "AND":
		return r.redisClient.BitOpAnd(ctx, destKey, keys...).Err()
	case "OR":
		return r.redisClient.BitOpOr(ctx, destKey, keys...).Err()
	case "XOR":
		return r.redisClient.BitOpXor(ctx, destKey, keys...).Err()
	case "NOT":
		return r.redisClient.BitOpNot(ctx, destKey, keys[0]).Err() // 仅支持对单个键执行 NOT 操作
	default:
		return fmt.Errorf("unsupported bit operation. please choose one of AND OR XOR NOT")
	}
}

// 修改特定位的值
func (r *RedisClient) ModifyBit(ctx context.Context, kind string, unique string, offset int64, newValue bool) error {
	key := fmt.Sprintf("%s_%s", kind, unique)
	value := 0
	if newValue {
		value = 1
	}
	return r.redisClient.SetBit(ctx, key, offset, value).Err()
}

// 重置整个位图
func (r *RedisClient) ResetBitmap(ctx context.Context, kind string, unique string, length int64) error {
	key := fmt.Sprintf("%s_%s", kind, unique)
	// 创建一个全 0 的字符串，用作新值
	// 为什么使用 (length+7)/8？
	// 这是因为 Redis 中的位图是按字节（每个字节有 8 位）存储的。
	// 如果位图的长度不是 8 的倍数，最后一个字节可能会有未使用的部分，
	// 因此我们加 7 以确保能分配足够的字节存储所有位。
	zeroData := make([]byte, (length+7)/8) // 字节数 = (位数+7)/8
	return r.redisClient.Set(ctx, key, zeroData, 0).Err()
}

// 找到第一个设置为 1 或 0 的位置
func (r *RedisClient) FindPositionBit(ctx context.Context, kind string, unique string, bit bool) (int64, error) {
	key := fmt.Sprintf("%s_%s", kind, unique)
	bitValue := int64(0)
	if bit {
		bitValue = 1
	}
	position, err := r.redisClient.BitPos(ctx, key, bitValue).Result()
	if err != nil {
		return -1, fmt.Errorf("failed to find bit position: %w", err)
	}
	return position, nil
}

// 按位统计多个位图
func (r *RedisClient) CombineBitmaps(ctx context.Context, destKind string, unique string, srcKinds []string) error {
	destKey := fmt.Sprintf("%s_%s", destKind, unique)
	srcKeys := make([]string, len(srcKinds))
	for i, kind := range srcKinds {
		srcKeys[i] = fmt.Sprintf("%s_%s", kind, unique)
	}

	return r.redisClient.BitOpOr(ctx, destKey, srcKeys...).Err() // 按位 OR 操作
}

// 这个方法是用于将多个哈希值添加到 Redis 的 布隆过滤器（Bloom Filter） 中的一个实现。
// 布隆过滤器是一种概率型数据结构，用于检测一个元素是否在集合中。它有两个特点：
// 可能会误判：可能会判断一个元素存在（即假阳性），但不可能错判元素不存在。
// 空间效率高：它使用固定的内存大小来处理非常大的集合。
func (r *RedisClient) AddToBloomFilter(ctx context.Context, kind string, unique string) error {
	key := fmt.Sprintf("%s_%s", kind, unique)
	pipeline := r.redisClient.Pipeline()

	hashValues := generateMultipleHashes(key, 8)

	for _, hash := range hashValues {
		pipeline.SetBit(ctx, key, hash, 1)
	}

	_, err := pipeline.Exec(ctx)
	return err
}

// 检查元素是否可能存在
func (r *RedisClient) CheckBloomFilter(ctx context.Context, kind string, unique string) (bool, error) {
	key := fmt.Sprintf("%s_%s", kind, unique)

	hashValues := generateMultipleHashes(key, 8)

	for _, hash := range hashValues {
		bit, err := r.redisClient.GetBit(ctx, key, hash).Result()
		if err != nil || bit == 0 {
			return false, err // 如果某个位为 0，则元素一定不存在
		}
	}
	return true, nil // 所有位都为 1，则元素可能存在
}

// BloomFilter 的哈希值生成
func generateMultipleHashes(input string, numHashes int) []int64 {
	hashValues := make([]int64, numHashes)
	// 使用 SHA256 和 FNV 哈希
	sha256Hash := sha256.New()
	sha256Hash.Write([]byte(input))
	sha256Sum := sha256Hash.Sum(nil)

	fnvHash := fnv.New64a()
	fnvHash.Write([]byte(input))
	fnvSum := fnvHash.Sum64()

	for i := 0; i < numHashes; i++ {
		// 使用 SHA256 和 FNV 的组合来生成多个哈希值
		hashValues[i] = int64(math.Abs(float64(fnvSum + uint64(i) + uint64(sha256Sum[i%len(sha256Sum)]))))
	}

	return hashValues
}

// Bitfield
// 使用 BITFIELD 操作批量设置/获取位图中的多个位
// command 格式 "SET/GET/INCRBY" "i8/i16/i32/u8/u16/u32" "#0/#1/..." GET不需要"value"
func (r *RedisClient) BitField(ctx context.Context, kind string, unique string, command ...interface{}) ([]int64, error) {
	key := fmt.Sprintf("%s_%s", kind, unique)
	result, err := r.redisClient.BitField(ctx, key, command...).Result()
	if err != nil {
		return nil, err
	}
	return result, nil
}

package cache

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	common "github.com/Yux77Yux/platform_backend/generated/common"
	generated "github.com/Yux77Yux/platform_backend/generated/user"
	tools "github.com/Yux77Yux/platform_backend/microservices/user/tools"
	"github.com/go-redis/redis/v8"
)

type CacheMethodStruct struct {
	CacheClient CacheInterface
}

func (c *CacheMethodStruct) ExistsUsername(ctx context.Context, username string) (bool, error) {
	exist, err := c.CacheClient.ExistsHashField(ctx, "User", "Credentials", username)
	if err != nil {
		log.Printf("error: failed to execute cache method: ExistsUsername %s", err.Error())
		return false, err
	}

	// 正常返回结果
	return exist, nil
}

func (c *CacheMethodStruct) ExistsEmail(ctx context.Context, email string) (bool, error) {
	exist, err := c.CacheClient.ExistsHashField(ctx, "User", "Credentials", email)

	if err != nil {
		return false, err
	}
	return exist, nil
}

func (c *CacheMethodStruct) ExistsUserInfo(ctx context.Context, user_id int64) (bool, error) {
	exist, err := c.CacheClient.ExistsHash(ctx, "UserInfo", strconv.FormatInt(user_id, 10))
	if err != nil {
		return false, err
	}
	return exist, nil
}

// POST

// 触发的可能有，过期，登录返回，设置邮箱字段
func (c *CacheMethodStruct) StoreEmail(ctx context.Context, credentials []*generated.UserCredentials) error {
	count := len(credentials)
	fieldValues := make([]interface{}, 0, count*2)

	for _, credential := range credentials {
		email := credential.GetUserEmail()
		if email == "" {
			continue
		}
		newCredential := &generated.UserCredentials{
			Password:  credential.GetPassword(),
			UserId:    credential.GetUserId(),
			UserEmail: email,
			UserRole:  credential.GetUserRole(),
		}
		data, err := proto.Marshal(newCredential)
		if err != nil {
			return fmt.Errorf("failed to marshal credentials: %w", err)
		}
		fieldValues = append(fieldValues, email, data)
	}

	if len(fieldValues) == 0 {
		return nil
	}
	result := c.CacheClient.SetFieldsHash(ctx, "User", "Credentials",
		fieldValues...,
	)
	if result != nil {
		return result
	}
	return nil
}

func (c *CacheMethodStruct) StoreUsername(ctx context.Context, credentials []*generated.UserCredentials) error {
	count := len(credentials)
	fieldValues := make([]interface{}, count*2)

	for i, credential := range credentials {
		username := credential.GetUsername()
		newCredential := &generated.UserCredentials{
			Password:  credential.GetPassword(),
			UserId:    credential.GetUserId(),
			UserEmail: credential.GetUserEmail(),
			UserRole:  credential.GetUserRole(),
		}
		data, err := proto.Marshal(newCredential)
		if err != nil {
			return fmt.Errorf("failed to marshal credentials: %w", err)
		}
		fieldValues[i*2] = username
		fieldValues[i*2+1] = data
	}

	return c.CacheClient.SetFieldsHash(ctx, "User", "Credentials",
		fieldValues...,
	)
}

func (c *CacheMethodStruct) StoreUserInfo(ctx context.Context, users []*generated.User) error {
	pipe := c.CacheClient.Pipeline()
	for _, user := range users {

		var userBday interface{}
		// 判断是否为空
		if user.GetUserBday() != nil {
			// 将 Timestamp 转换为 time.Time 类型
			userBday = user.GetUserBday().AsTime()
		} else {
			userBday = "none"
		}

		pipe.HSet(ctx, fmt.Sprintf("Hash_UserInfo_%d", user.GetUserDefault().GetUserId()),
			"user_name", user.GetUserDefault().GetUserName(),
			"user_avatar", user.GetUserAvatar(),
			"user_bio", user.GetUserBio(),
			"user_status", user.GetUserStatus().String(),
			"user_gender", user.GetUserGender().String(),
			"user_bday", userBday,
			"user_created_at", user.GetUserCreatedAt().AsTime(),
			"user_updated_at", user.GetUserUpdatedAt().AsTime(),

			"followers", 0,
			"followees", 0,
		)
	}
	_, err := pipe.Exec(ctx)
	if err != nil && err != redis.Nil {
		return fmt.Errorf("pipeline execution failed: %w", err)
	}

	return nil
}

func (c *CacheMethodStruct) Follow(ctx context.Context, subs []*generated.Follow) error {
	now := float64(timestamppb.Now().Seconds)
	for _, follow := range subs {
		followerId := follow.GetFollowerId()
		followeeId := follow.GetFolloweeId()

		pipe := c.CacheClient.TxPipeline()

		pipe.HIncrBy(ctx, fmt.Sprintf("Hash_UserInfo_%d", followerId), "followees", 1)
		pipe.HIncrBy(ctx, fmt.Sprintf("Hash_UserInfo_%d", followeeId), "followers", 1)

		pipe.ZAdd(ctx, fmt.Sprintf("ZSet_Time_Followees_%d", followerId), &redis.Z{
			Score:  now,
			Member: follow.FolloweeId,
		})

		pipe.ZAdd(ctx, fmt.Sprintf("ZSet_View_Followees_%d", followerId), &redis.Z{
			Score:  0,
			Member: follow.FolloweeId,
		})

		pipe.ZAdd(ctx, fmt.Sprintf("ZSet_Followers_%d", followeeId), &redis.Z{
			Score:  now,
			Member: follow.FollowerId,
		})
		_, err := pipe.Exec(ctx)
		if err != nil {
			return fmt.Errorf("error: StoreFollowee %v : %w", follow, err)
		}
	}
	return nil
}

// UPDATE
func (c *CacheMethodStruct) UpdateUserSpace(ctx context.Context, users []*generated.UserUpdateSpace) error {
	pipe := c.CacheClient.Pipeline()
	for _, user := range users {
		var userBday interface{}
		// 判断是否为空
		if user.GetUserBday() != nil {
			// 将 Timestamp 转换为 time.Time 类型
			userBday = user.GetUserBday().AsTime()
		} else {
			userBday = "none"
		}

		pipe.HSet(ctx, fmt.Sprintf("Hash_UserInfo_%d", user.GetUserDefault().GetUserId()),
			"user_name", user.GetUserDefault().GetUserName(),
			"user_bio", user.GetUserBio(),
			"user_gender", user.GetUserGender().String(),
			"user_bday", userBday,
			"user_updated_at", time.Now(),
		)
	}
	_, err := pipe.Exec(ctx)
	if err != nil && err != redis.Nil {
		return fmt.Errorf("pipeline execution failed: %w", err)
	}
	return nil

}

func (c *CacheMethodStruct) UpdateUserAvatar(ctx context.Context, users []*generated.UserUpdateAvatar) error {
	pipe := c.CacheClient.Pipeline()
	for _, user := range users {
		pipe.HSet(ctx, fmt.Sprintf("Hash_UserInfo_%d", user.GetUserId()),
			"user_avatar", user.GetUserAvatar(),
			"user_updated_at", time.Now(),
		)
	}
	_, err := pipe.Exec(ctx)
	if err != nil && err != redis.Nil {
		return fmt.Errorf("pipeline execution failed: %w", err)
	}
	return nil
}

func (c *CacheMethodStruct) UpdateUserBio(ctx context.Context, users []*generated.UserUpdateBio) error {
	pipe := c.CacheClient.Pipeline()
	for _, user := range users {
		pipe.HSet(ctx, fmt.Sprintf("Hash_UserInfo_%d", user.GetUserId()),
			"user_bio", user.GetUserBio(),
			"user_updated_at", time.Now(),
		)
	}
	_, err := pipe.Exec(ctx)
	if err != nil {
		return fmt.Errorf("pipeline execution failed: %w", err)
	}

	return nil
}

func (c *CacheMethodStruct) UpdateUserStatus(ctx context.Context, users []*generated.UserUpdateStatus) error {
	pipe := c.CacheClient.Pipeline()
	for _, user := range users {
		pipe.HSet(ctx, fmt.Sprintf("Hash_UserInfo_%d", user.GetUserId()),
			"user_status", user.GetUserStatus().String(),
			"user_updated_at", time.Now(),
		)
	}
	_, err := pipe.Exec(ctx)
	if err != nil && err != redis.Nil {
		return fmt.Errorf("pipeline execution failed: %w", err)
	}
	return nil
}

// GET
func (c *CacheMethodStruct) GetUserInfo(ctx context.Context, user_id int64, fields []string) (map[string]string, error) {
	var (
		result map[string]string
		err    error
		values []interface{}
	)
	if len(fields) == 0 {
		result, err = c.CacheClient.GetAllHash(ctx, "UserInfo", strconv.FormatInt(user_id, 10))
	} else {
		values, err = c.CacheClient.GetAnyHash(ctx, "UserInfo", strconv.FormatInt(user_id, 10), fields...)
		// 构造结果 map
		result = make(map[string]string, len(fields))
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
	}

	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *CacheMethodStruct) GetUserCredentials(ctx context.Context, userCrdentials *generated.UserCredentials) (*generated.UserCredentials, error) {
	email := userCrdentials.GetUserEmail()
	field := userCrdentials.GetUsername()
	if email != "" {
		field = email
	}

	credentials, err := c.CacheClient.GetHash(ctx, "User", "Credentials", field)
	if err != nil {
		if err != redis.Nil {
			return nil, err
		}
	}

	if credentials == "" {
		return nil, nil
	}

	var credentialsInStore generated.UserCredentials
	err = proto.Unmarshal([]byte(credentials), &credentialsInStore)
	if err != nil {
		return nil, err
	}

	// 验证密码
	match, err := tools.VerifyPassword(credentialsInStore.GetPassword(), userCrdentials.GetPassword())
	if err != nil {
		return nil, fmt.Errorf("failed to verify password: %w", err)
	}
	if !match {
		return nil, nil
	}

	return &credentialsInStore, nil
}

// Follow methods

func (c *CacheMethodStruct) GetUserCards(ctx context.Context, userIds []int64) ([]*common.UserCreationComment, error) {
	length := len(userIds)
	users := make([]*common.UserCreationComment, length)

	pipe := c.CacheClient.Pipeline()
	// 用来存储 pipeline 请求的结果
	cmds := make([]*redis.SliceCmd, length)
	for i, userId := range userIds {
		cmds[i] = pipe.HMGet(ctx, fmt.Sprintf("Hash_UserInfo_%d", userId), "user_name", "user_bio", "user_avatar")
	}
	_, err := pipe.Exec(ctx)
	if err != nil && err != redis.Nil {
		return nil, err
	}

	for i, userId := range userIds {
		results, err := cmds[i].Result()
		if err != nil {
			return nil, err
		}
		users[i] = &common.UserCreationComment{
			UserDefault: &common.UserDefault{
				UserId:   userId,
				UserName: results[0].(string),
			},
			UserBio:    results[1].(string),
			UserAvatar: results[2].(string),
		}
	}
	return users, nil
}

func (c *CacheMethodStruct) ExistFollowee(ctx context.Context, followeeId, followerId int64) (bool, error) {
	return false, nil
}

func (c *CacheMethodStruct) GetFolloweesByTime(ctx context.Context, userId int64, page int32) ([]int64, error) {
	const LIMIT = 20
	start := int64((page - 1) * LIMIT)
	end := start + 19
	strs, err := c.CacheClient.RevRangeZSet(ctx, "Time_Followees", strconv.FormatInt(userId, 10), start, end)
	if err != nil {
		return nil, err
	}
	length := len(strs)
	ids := make([]int64, length)
	for i, str := range strs {
		id, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			return nil, err
		}
		ids[i] = id
	}

	return ids, nil
}

func (c *CacheMethodStruct) GetFolloweesByView(ctx context.Context, userId int64, page int32) ([]int64, error) {
	const LIMIT = 20
	start := int64((page - 1) * LIMIT)
	end := start + 19
	strs, err := c.CacheClient.RevRangeZSet(ctx, "View_Followees", strconv.FormatInt(userId, 10), start, end)
	if err != nil {
		return nil, err
	}
	length := len(strs)
	ids := make([]int64, length)
	for i, str := range strs {
		id, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			return nil, err
		}
		ids[i] = id
	}

	return ids, nil
}

func (c *CacheMethodStruct) GetFollowers(ctx context.Context, userId int64, page int32) ([]int64, error) {
	const LIMIT = 20
	start := int64((page - 1) * LIMIT)
	end := start + 19
	strs, err := c.CacheClient.RevRangeZSet(ctx, "Followers", strconv.FormatInt(userId, 10), start, end)
	if err != nil {
		return nil, err
	}
	length := len(strs)
	ids := make([]int64, length)
	for i, str := range strs {
		id, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			return nil, err
		}
		ids[i] = id
	}

	return ids, nil
}

// Del
func (c *CacheMethodStruct) CancelFollow(ctx context.Context, follow *generated.Follow) error {
	pipe := c.CacheClient.TxPipeline()
	followerId := follow.GetFollowerId()
	followeeId := follow.GetFolloweeId()
	pipe.HIncrBy(ctx, fmt.Sprintf("Hash_UserInfo_%d", followerId), "followees", -1)
	pipe.HIncrBy(ctx, fmt.Sprintf("Hash_UserInfo_%d", followeeId), "followers", -1)
	pipe.ZRem(ctx, fmt.Sprintf("ZSet_Time_Followees_%d", followerId), followeeId)
	pipe.ZRem(ctx, fmt.Sprintf("ZSet_View_Followees_%d", followerId), followeeId)
	pipe.ZRem(ctx, fmt.Sprintf("ZSet_Followers_%d", followeeId), followerId)

	_, err := pipe.Exec(ctx)
	if err != nil && err != redis.Nil {
		return fmt.Errorf("error: StoreFollowee %w", err)
	}

	return nil
}

func (c *CacheMethodStruct) DelCredentials(ctx context.Context, username string) error {
	_, err := c.CacheClient.DelHashField(ctx, "User", "Credentials", username)
	return err
}

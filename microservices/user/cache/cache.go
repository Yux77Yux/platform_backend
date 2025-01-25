package cache

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"google.golang.org/protobuf/proto"

	generated "github.com/Yux77Yux/platform_backend/generated/user"
	tools "github.com/Yux77Yux/platform_backend/microservices/user/tools"
)

func ExistsUsername(username string) (bool, error) {
	ctx := context.Background()

	resultCh := make(chan struct {
		exist bool
		err   error
	}, 1)

	// 将闭包发至通道
	cacheRequestChannel <- func(CacheClient CacheInterface) {
		exist, err := CacheClient.ExistsHashField(ctx, "User", "Credentials", username)

		select {
		case resultCh <- struct {
			exist bool
			err   error
		}{exist, err}:
			log.Printf("info: completely execute for cache method: ExistsUsername")
		case <-ctx.Done():
			log.Printf("warning: context canceled for cache method: ExistsUsername")
		}
	}

	// 使用 select 来监听超时和结果
	select {
	case <-ctx.Done():
		// 超时
		return false, fmt.Errorf("timeout: %w", ctx.Err())
	case result := <-resultCh:
		if result.err != nil {
			log.Printf("error: failed to execute cache method: ExistsUsername")
			return false, result.err
		}

		// 正常返回结果
		return result.exist, nil
	}
}

func ExistsEmail(email string) (bool, error) {
	ctx := context.Background()

	resultCh := make(chan struct {
		exist bool
		err   error
	}, 1)

	cacheRequestChannel <- func(CacheClient CacheInterface) {
		exist, err := CacheClient.ExistsHashField(ctx, "User", "Credentials", email)

		select {
		case resultCh <- struct {
			exist bool
			err   error
		}{exist, err}:
			log.Printf("info: completely execute for cache method: ExistsEmail")
		case <-ctx.Done():
			log.Printf("warning: context canceled for cache method: ExistsEmail")
		}
	}

	// 使用 select 来监听超时和结果
	select {
	case <-ctx.Done():
		return false, fmt.Errorf("timeout: %w", ctx.Err())
	case result := <-resultCh:
		if result.err != nil {
			return false, result.err
		}
		return result.exist, nil
	}
}

func ExistsUserInfo(user_id int64) (bool, error) {
	ctx := context.Background()
	resultCh := make(chan struct {
		exist bool
		err   error
	}, 1)

	cacheRequestChannel <- func(CacheClient CacheInterface) {
		exist, err := CacheClient.ExistsHash(ctx, "UserInfo", strconv.FormatInt(user_id, 10))

		select {
		case resultCh <- struct {
			exist bool
			err   error
		}{exist, err}:
			log.Printf("info: completely execute for cache method: ExistsUserInfo")
		case <-ctx.Done():
			log.Printf("warning: context canceled for cache method: ExistsUserInfo")
		}
	}

	// 使用 select 来监听超时和结果
	select {
	case <-ctx.Done():
		return false, fmt.Errorf("timeout: %w", ctx.Err())
	case result := <-resultCh:
		if result.err != nil {
			return false, result.err
		}
		return result.exist, nil
	}
}

// POST

// 触发的可能有，过期，登录返回，设置邮箱字段
func StoreEmail(credentials []*generated.UserCredentials) error {
	ctx := context.Background()

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
	resultCh := make(chan error, 1)

	cacheRequestChannel <- func(CacheClient CacheInterface) {
		err := CacheClient.SetFieldsHash(ctx, "User", "Credentials",
			fieldValues...,
		)
		resultCh <- err
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

func StoreUsername(credentials []*generated.UserCredentials) error {
	ctx := context.Background()

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
	resultCh := make(chan error, 1)

	cacheRequestChannel <- func(CacheClient CacheInterface) {
		err := CacheClient.SetFieldsHash(ctx, "User", "Credentials",
			fieldValues...,
		)
		resultCh <- err
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

func StoreUserInfo(users []*generated.User) error {
	ctx := context.Background()

	count := len(users)

	resultCh := make(chan error, 1)

	func(count int) {
		cacheRequestChannel <- func(CacheClient CacheInterface) {
			pipe := CacheClient.Pipeline()
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
				)
			}
			_, err := pipe.Exec(ctx)
			if err != nil {
				resultCh <- fmt.Errorf("pipeline execution failed: %w", err)
			}
		}
	}(count)

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

// UPDATE
func UpdateUser(users []*generated.UserUpdateSpace) error {
	ctx := context.Background()

	count := len(users)

	resultCh := make(chan error, 1)

	func(count int) {
		cacheRequestChannel <- func(CacheClient CacheInterface) {
			pipe := CacheClient.Pipeline()
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
			if err != nil {
				resultCh <- fmt.Errorf("pipeline execution failed: %w", err)
			}
		}
	}(count)

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

func UpdateUserAvatar(users []*generated.UserUpdateAvatar) error {
	ctx := context.Background()

	count := len(users)

	resultCh := make(chan error, 1)

	func(count int) {
		cacheRequestChannel <- func(CacheClient CacheInterface) {
			pipe := CacheClient.Pipeline()
			for _, user := range users {
				pipe.HSet(ctx, fmt.Sprintf("Hash_UserInfo_%d", user.GetUserId()),
					"user_avatar", user.GetUserAvatar(),
					"user_updated_at", time.Now(),
				)
			}
			_, err := pipe.Exec(ctx)
			if err != nil {
				resultCh <- fmt.Errorf("pipeline execution failed: %w", err)
			}
		}
	}(count)

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

func UpdateUserBio(users []*generated.UserUpdateBio) error {
	ctx := context.Background()

	count := len(users)

	resultCh := make(chan error, 1)

	func(count int) {
		cacheRequestChannel <- func(CacheClient CacheInterface) {
			pipe := CacheClient.Pipeline()
			for _, user := range users {
				pipe.HSet(ctx, fmt.Sprintf("Hash_UserInfo_%d", user.GetUserId()),
					"user_bio", user.GetUserBio(),
					"user_updated_at", time.Now(),
				)
			}
			_, err := pipe.Exec(ctx)
			if err != nil {
				resultCh <- fmt.Errorf("pipeline execution failed: %w", err)
			}
		}
	}(count)

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

func UpdateUserStatus(users []*generated.UserUpdateStatus) error {
	ctx := context.Background()

	count := len(users)

	resultCh := make(chan error, 1)

	func(count int) {
		cacheRequestChannel <- func(CacheClient CacheInterface) {
			pipe := CacheClient.Pipeline()
			for _, user := range users {
				pipe.HSet(ctx, fmt.Sprintf("Hash_UserInfo_%d", user.GetUserId()),
					"user_status", user.GetUserStatus(),
					"user_updated_at", time.Now(),
				)
			}
			_, err := pipe.Exec(ctx)
			if err != nil {
				resultCh <- fmt.Errorf("pipeline execution failed: %w", err)
			}
		}
	}(count)

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
func GetUserInfo(user_id int64, fields []string) (map[string]string, error) {
	ctx := context.Background()

	resultCh := make(chan struct {
		user map[string]string
		err  error
	}, 1)

	cacheRequestChannel <- func(CacheClient CacheInterface) {
		if len(fields) == 0 {
			result, err := CacheClient.GetAllHash(ctx, "UserInfo", strconv.FormatInt(user_id, 10))
			resultCh <- struct {
				user map[string]string
				err  error
			}{
				user: result,
				err:  err,
			}
		} else {
			values, err := CacheClient.GetAnyHash(ctx, "UserInfo", strconv.FormatInt(user_id, 10), fields...)
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
				user map[string]string
				err  error
			}{
				user: result,
				err:  err,
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

		return result.user, nil
	}
}

func GetUserCredentials(userCrdentials *generated.UserCredentials) (*generated.UserCredentials, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	email := userCrdentials.GetUserEmail()
	field := userCrdentials.GetUsername()
	if email != "" {
		field = email
	}

	resultCh := make(chan struct {
		credentials string
		err         error
	}, 1)

	cacheRequestChannel <- func(CacheClient CacheInterface) {
		result, err := CacheClient.GetHash(ctx, "User", "Credentials", field)
		resultCh <- struct {
			credentials string
			err         error
		}{
			credentials: result,
			err:         err,
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

		if result.credentials == "" {
			return nil, nil
		}

		credentials := new(generated.UserCredentials)
		err := proto.Unmarshal([]byte(result.credentials), credentials)
		if err != nil {
			return nil, err
		}

		// 验证密码
		match, err := tools.VerifyPassword(credentials.GetPassword(), userCrdentials.GetPassword())
		if err != nil {
			return nil, fmt.Errorf("failed to verify password: %w", err)
		}
		if !match {
			return nil, nil
		}

		return credentials, nil
	}
}

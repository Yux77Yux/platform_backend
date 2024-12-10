package cache

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	generated "github.com/Yux77Yux/platform_backend/generated/user"
)

func ExistsUsername(username string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	resultCh := make(chan struct {
		exist bool
		err   error
	}, 1)

	// 将闭包发至通道
	cacheRequestChannel <- func(CacheClient CacheInterface) {
		exist, err := CacheClient.ExistsInSet(ctx, "User", "Username", username)

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
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	resultCh := make(chan struct {
		exist bool
		err   error
	}, 1)

	cacheRequestChannel <- func(CacheClient CacheInterface) {
		exist, err := CacheClient.ExistsInSet(ctx, "User", "Email", email)

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

func StoreUsername(username string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	resultCh := make(chan error, 1)

	cacheRequestChannel <- func(CacheClient CacheInterface) {
		err := CacheClient.AddToSet(ctx, "User", "Username", username)
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

func StoreEmail(email string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	resultCh := make(chan error, 1)

	cacheRequestChannel <- func(CacheClient CacheInterface) {
		err := CacheClient.AddToSet(ctx, "User", "Email", email)
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

func StoreUserInfo(user *generated.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	resultCh := make(chan error, 1)

	id := strconv.FormatInt(user.GetUserDefault().GetUserId(), 10)

	var userBday interface{}
	// 判断是否为空
	if user.GetUserBday() != nil {
		// 将 Timestamp 转换为 time.Time 类型
		userBday = user.GetUserBday().AsTime()
	} else {
		userBday = "none"
	}

	cacheRequestChannel <- func(CacheClient CacheInterface) {
		err := CacheClient.SetFieldsHash(ctx, "UserInfo", id,
			"user_name", user.GetUserDefault().GetUserName(),
			"user_avator", user.GetUserAvator(),
			"user_bio", user.GetUserBio(),
			"user_status", user.GetUserStatus().String(),
			"user_gender", user.GetUserGender().String(),
			"user_email", user.GetUserEmail(),
			"user_bday", userBday,
			"user_created_at", user.GetUserCreatedAt().AsTime(),
			"user_updated_at", user.GetUserUpdatedAt().AsTime(),
			"user_role", user.GetUserRole().String(),
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

func ExistsUserInfo(user_id int64) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
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

func GetUserInfo(user_id int64, fields []string) (map[string]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

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

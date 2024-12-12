package repository

import (
	"context"
	"fmt"
	"strings"
	"time"

	generated "github.com/Yux77Yux/platform_backend/generated/user"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func UserAddInfoInTransaction(user_info *generated.User) error {
	query := `insert into db_user_1.User 
	(user_id,
	user_name,
	user_avator,
	user_bio,
	user_status,
	user_gender,
	user_email,
	user_bday,
	user_created_at,
	user_updated_at,
	user_role)
	values(?,?,?,?,?,?,?,?,?,?,?)`

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

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
		UserId        int64     = user_info.GetUserDefault().GetUserId()
		UserName      string    = user_info.GetUserDefault().GetUserName()
		UserAvator    string    = user_info.GetUserAvator()
		UserBio       string    = user_info.GetUserBio()
		UserStatus    string    = user_info.GetUserStatus().String()
		UserGender    string    = user_info.GetUserGender().String()
		UserEmail     *string   = &user_info.UserEmail
		UserCreatedAt time.Time = user_info.GetUserCreatedAt().AsTime()
		UserUpdatedAt time.Time = user_info.GetUserUpdatedAt().AsTime()
		UserRole      string    = user_info.GetUserRole().String()
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
			query,
			UserId,
			UserName,
			UserAvator,
			UserBio,
			UserStatus,
			UserGender,
			UserEmail,
			nil,
			UserCreatedAt,
			UserUpdatedAt,
			UserRole,
		)
		if err != nil {
			err = fmt.Errorf("transaction exec failed because %v", err)
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

func UserGetInfoInTransaction(user_id int64, fields []string) (map[string]interface{}, error) {
	var query string
	if len(fields) > 0 {
		// 查询指定字段
		query = fmt.Sprintf("SELECT %s FROM db_user_1.User WHERE user_id = ?", strings.Join(fields, ", "))
	} else {
		// 查询全部字段
		query = "SELECT * FROM db_user_1.User WHERE user_id = ?"
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

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

	var result map[string]interface{}
	select {
	case <-ctx.Done():
		err = fmt.Errorf("exec timeout :%w", ctx.Err())
		if errSecond := db.RollbackTransaction(tx); errSecond != nil {
			err = fmt.Errorf("%w and %w", err, errSecond)
		}

		return nil, err
	default:
		rows, err := tx.Query(query, user_id)
		if err != nil {
			err = fmt.Errorf("transaction exec failed because %v", err)
			if errSecond := db.RollbackTransaction(tx); errSecond != nil {
				err = fmt.Errorf("%w and %w", err, errSecond)
			}

			return nil, err
		}
		defer rows.Close()

		// 获取列名
		cols, err := rows.Columns()
		if err != nil {
			return nil, fmt.Errorf("failed to get columns: %w", err)
		}

		// 确保有结果
		if !rows.Next() {
			return nil, nil
		}

		// 创建一个存储列值的切片
		values := make([]interface{}, len(cols))
		pointers := make([]interface{}, len(cols))
		for i := range values {
			pointers[i] = &values[i]
		}

		// 扫描结果到指针数组
		if err := rows.Scan(pointers...); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		// 将结果填充到 map 中
		result = make(map[string]interface{})
		for i, colName := range cols {
			switch colName {
			case "user_id":
				result[colName] = user_id
			case "user_status":
				if value, ok := values[i].(int); ok {
					result[colName] = generated.UserStatus(value)
				} else {
					return nil, fmt.Errorf("assert user_status type failed ")
				}
			case "user_gender":
				if value, ok := values[i].(int); ok {
					result[colName] = generated.UserGender(value)
				} else {
					return nil, fmt.Errorf("assert user_gender type failed ")
				}
			case "user_bday", "user_created_at", "user_updated_at":
				if value, ok := values[i].(time.Time); ok {
					result[colName] = timestamppb.New(value)
				} else {
					return nil, fmt.Errorf("assert %v type failed ", values[i])
				}
			default:
				if value, ok := values[i].(string); ok {
					result[colName] = value
				} else {
					return nil, fmt.Errorf("assert %v type failed ", values[i])
				}
			}
		}

		if err = db.CommitTransaction(tx); err != nil {
			return nil, err
		}
	}

	return result, nil
}

func UserVerifyInTranscation(user_credential *generated.UserCredentials) (int64, error) {
	query := `select 
	password,
	user_id
	from db_user_credentials_1.UserCredentials 
	where username = ?`

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	tx, err := db.BeginTransaction()
	if err != nil {
		return 0, err
	}

	// 切换到 db_user_credentials_1
	_, err = tx.Exec("USE db_user_credentials_1")
	if err != nil {
		return 0, fmt.Errorf("change the database error: %v", err)
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
		passwordHash string
		user_id      int64
	)
	select {
	case <-ctx.Done():
		err = fmt.Errorf("exec timeout :%w", ctx.Err())
		if errSecond := db.RollbackTransaction(tx); errSecond != nil {
			err = fmt.Errorf("%w and %w", err, errSecond)
		}

		return 0, err
	default:
		err := tx.QueryRow(query, user_credential.GetUsername()).Scan(&passwordHash, &user_id)
		if err != nil {
			err = fmt.Errorf("transaction exec failed because %v", err)
			if errSecond := db.RollbackTransaction(tx); errSecond != nil {
				err = fmt.Errorf("%w and %w", err, errSecond)
			}

			return 0, err
		}

		if err = db.CommitTransaction(tx); err != nil {
			return 0, err
		}
	}

	match, err := verifyPassword(passwordHash, user_credential.GetPassword())
	if err != nil {
		return 0, fmt.Errorf("failed to verify password: %w", err)
	}
	if !match {
		return 0, nil
	}

	if err != nil {
		return 0, fmt.Errorf("not the database error but the others occurred :%w", err)
	}
	return user_id, nil
}

func UserRegisterInTransaction(user_credential *generated.UserCredentials, id int64) error {
	pwd, err := hashPassword(user_credential.GetPassword())

	// 进行复杂加密
	if err != nil {
		return fmt.Errorf("decrypt hash password failed because %w", err)
	}

	query := `insert into db_user_credentials_1.UserCredentials(
	username,
	password,
	user_email,
	user_id,
	user_role)values
	(?,?,?,?,?) 
	`

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

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
		var email interface{}
		if user_credential.GetUserEmail() == "" {
			email = nil
		} else {
			email = user_credential.GetUserEmail()
		}

		_, err = tx.Exec(query, user_credential.GetUsername(), pwd, email, id, user_credential.GetUserRole().String())

		if err != nil {
			err = fmt.Errorf("transaction exec failed because %v", err)
			if errSecond := db.RollbackTransaction(tx); errSecond != nil {
				err = fmt.Errorf("%w and %w", err, errSecond)
			}

			return err
		}

		if err = db.CommitTransaction(tx); err != nil {
			return err
		}
	}

	if err != nil {
		return fmt.Errorf("not the database error but the others occurred :%w", err)
	}
	return nil
}

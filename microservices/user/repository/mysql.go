package repository

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	generated "github.com/Yux77Yux/platform_backend/generated/user"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// SET
func UserAddInfoInTransaction(user_info *generated.User) error {
	query := `insert into db_user_1.User 
	(user_id,
	user_name,
	user_avatar,
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
		UserAvatar    string    = user_info.GetUserAvatar()
		UserBio       string    = user_info.GetUserBio()
		UserStatus    string    = user_info.GetUserStatus().String()
		UserGender    string    = user_info.GetUserGender().String()
		UserEmail     *string   = nil
		UserCreatedAt time.Time = user_info.GetUserCreatedAt().AsTime()
		UserUpdatedAt time.Time = user_info.GetUserUpdatedAt().AsTime()
		UserRole      string    = user_info.GetUserRole().String()
	)

	if user_info.GetUserEmail() != "" {
		UserEmail = &user_info.UserEmail
	}

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
			UserAvatar,
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
		var email *string = nil
		if user_credential.GetUserEmail() != "" {
			email = &user_credential.UserEmail
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

// GET
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
	log.Println("url: ", query)
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
			case "user_bday":
				if values[i] == nil {
					result[colName] = "none"
				} else {
					if value, ok := values[i].([]byte); ok {
						// 将字符串解析为 time.Time（假设格式是 "YYYY-MM-DD"）
						parsedTime, err := time.Parse("2006-01-02", string(value))
						if err != nil {
							log.Printf("colName %s value %v with error: %v", colName, values[i], err)
							return nil, fmt.Errorf("failed to parse time %v: %v", value, err)
						}

						result[colName] = timestamppb.New(parsedTime)
					} else {
						log.Printf("colName %s value %v with error: %v", colName, values[i], err)
						return nil, fmt.Errorf("assert %v timeType failed ", values[i])
					}
				}
			case "user_created_at", "user_updated_at":
				if value, ok := values[i].([]byte); ok {
					// 将字符串解析为 time.Time（假设格式是 "YYYY-MM-DD HH:MM:SS"）
					parsedTime, err := time.Parse("2006-01-02 15:04:05", string(value))
					if err != nil {
						log.Printf("colName %s value %v with error: %v", colName, values[i], err)
						return nil, fmt.Errorf("failed to parse time %v: %v", string(value), err)
					}

					result[colName] = timestamppb.New(parsedTime)
				} else {
					log.Printf("colName %s value %v with error: %v", colName, values[i], err)
					return nil, fmt.Errorf("assert %v timeType failed ", values[i])
				}
			case "user_email":
				if values[i] == nil {
					continue
				} else {
					if value, ok := values[i].([]byte); ok {
						result[colName] = string(value)
						log.Printf("%s :%s", colName, result[colName])
					} else {
						log.Printf("colName %s value %v with error: %v", colName, values[i], err)
						return nil, fmt.Errorf("assert %v type failed ", values[i])
					}
				}
			default:
				if value, ok := values[i].([]byte); ok {
					result[colName] = string(value)
					log.Printf("%s :%s", colName, result[colName])
				} else {
					log.Printf("colName %s value %v with error: %v", colName, values[i], err)
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

// Verify
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

// UPDATE
func UserRegisterUpdateInTransaction(user_info *generated.UserUpdateSpace) error {
	query := `UPDATE db_user_credentials_1.UserCredentials
		SET 
    		user_email = ?
		WHERE user_id = ? `

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
		UserId    int64   = user_info.GetUserDefault().GetUserId()
		UserEmail *string = nil
	)
	if user_info.GetUserEmail() != "" {
		UserEmail = &user_info.UserEmail
	}

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
			UserEmail,
			UserId,
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

func UserUpdateInTransaction(user_info *generated.UserUpdateSpace) error {
	query := `UPDATE db_user_1.User 
		SET 
    		user_name = ?,
    		user_bio = ?,
    		user_gender = ?,
    		user_email = ?,
    		user_bday = ?,
    		user_updated_at = ?
		WHERE user_id = ? `
	email := user_info.GetUserEmail()
	if email == "" {
		query = `UPDATE db_user_1.User 
		SET 
    		user_name = ?,
    		user_bio = ?,
    		user_gender = ?,
    		user_bday = ?,
    		user_updated_at = ?
		WHERE user_id = ? `
	}

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
		UserBio       string    = user_info.GetUserBio()
		UserGender    string    = user_info.GetUserGender().String()
		UserEmail     string    = email
		UserUpdatedAt time.Time = time.Now()
		UserBday      time.Time = user_info.GetUserBday().AsTime()
	)

	select {
	case <-ctx.Done():
		err = fmt.Errorf("exec timeout :%w", ctx.Err())
		if errSecond := db.RollbackTransaction(tx); errSecond != nil {
			err = fmt.Errorf("%w and %w", err, errSecond)
		}

		return err
	default:
		var err error
		if email == "" {
			_, err = tx.Exec(
				query,
				UserName,
				UserBio,
				UserGender,
				UserBday,
				UserUpdatedAt,
				UserId,
			)
		} else {
			_, err = tx.Exec(
				query,
				UserName,
				UserBio,
				UserGender,
				UserEmail,
				UserBday,
				UserUpdatedAt,
				UserId,
			)
		}

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

func UserUpdateAvatarInTransaction(user_info *generated.UserUpdateAvatar) error {
	query := `UPDATE db_user_1.User 
		SET 
    		user_avatar = ?
		WHERE user_id = ? `

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
		UserId     int64  = user_info.GetUserId()
		UserAvatar string = user_info.GetUserAvatar()
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
			UserAvatar,
			UserId,
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

func UserUpdateStatusInTransaction(user_info *generated.UserUpdateStatus) error {
	query := `UPDATE db_user_1.User 
		SET 
    		user_status = ?
		WHERE user_id = ? `

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
		UserId     int64                = user_info.GetUserId()
		UserStatus generated.UserStatus = user_info.GetUserStatus()
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
			UserStatus,
			UserId,
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

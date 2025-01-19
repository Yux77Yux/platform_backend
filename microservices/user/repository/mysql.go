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
	(id,
	name,
	avatar,
	bio,
	status,
	gender,
	bday,
	created_at,
	updated_at
	)
	values(?,?,?,?,?,?,?,?,?)`

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
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
		UserCreatedAt time.Time = user_info.GetUserCreatedAt().AsTime()
		UserUpdatedAt time.Time = user_info.GetUserUpdatedAt().AsTime()
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
			UserAvatar,
			UserBio,
			UserStatus,
			UserGender,
			nil,
			UserCreatedAt,
			UserUpdatedAt,
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

func UserRegisterInTransaction(user_credential *generated.UserCredentials) error {
	// 进行复杂加密
	pwd, err := hashPassword(user_credential.GetPassword())
	if err != nil {
		return fmt.Errorf("decrypt hash password failed because %w", err)
	}

	query := `INSERT INTO db_user_credentials_1.UserCredentials(
			user_id,
			username,
			password,
			email,
			role)
		VALUES
			(?,?,?,?,?)`

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
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

		_, err = tx.Exec(query,
			user_credential.GetUserId(),
			user_credential.GetUsername(),
			pwd,
			email,
			user_credential.GetUserRole().String())

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
		query = fmt.Sprintf("SELECT %s FROM db_user_1.User WHERE id = ?", strings.Join(fields, ", "))
	} else {
		// 查询全部字段
		query = "SELECT * FROM db_user_1.User WHERE id = ?"
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
			case "id":
				result[colName] = user_id
			case "bday":
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
			case "created_at", "updated_at":
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

		// 再查UserCredentials拿身份和邮箱
		var (
			email interface{}
			role  string
		)
		query = "SELECT email,role FROM db_user_credentials_1.UserCredentials WHERE user_id = ?"
		if err := tx.QueryRow(query, user_id).Scan(&email, &role); err != nil {
			return nil, err
		}

		log.Printf("email %v", email)
		if email == nil {
			result["email"] = ""
		} else {
			if value, ok := email.([]byte); ok {
				result["email"] = string(value)
			} else {

				return nil, fmt.Errorf("assert email type failed ")
			}
		}
		result["role"] = role

		if err = db.CommitTransaction(tx); err != nil {
			return nil, err
		}
	}
	return result, nil
}

// Verify
func UserVerifyInTranscation(user_credential *generated.UserCredentials) (*generated.UserCredentials, error) {
	query := `select 
	password,
	user_id,
	role,
	email 
	from db_user_credentials_1.UserCredentials 
	where username = ?`

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	tx, err := db.BeginTransaction()
	if err != nil {
		return nil, err
	}

	// 切换到 db_user_credentials_1
	_, err = tx.Exec("USE db_user_credentials_1")
	if err != nil {
		return nil, fmt.Errorf("change the database error: %v", err)
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
		email        interface{}
		role         string
	)

	select {
	case <-ctx.Done():
		err = fmt.Errorf("exec timeout :%w", ctx.Err())
		if errSecond := db.RollbackTransaction(tx); errSecond != nil {
			err = fmt.Errorf("%w and %w", err, errSecond)
		}

		return nil, err
	default:
		err := tx.QueryRow(query, user_credential.GetUsername()).Scan(&passwordHash, &user_id, &role, &email)
		if err != nil {
			err = fmt.Errorf("transaction exec failed because %v", err)
			if errSecond := db.RollbackTransaction(tx); errSecond != nil {
				err = fmt.Errorf("%w and %w", err, errSecond)
			}

			return nil, err
		}

		if err = db.CommitTransaction(tx); err != nil {
			return nil, err
		}
	}

	match, err := verifyPassword(passwordHash, user_credential.GetPassword())
	if err != nil {
		return nil, fmt.Errorf("failed to verify password: %w", err)
	}
	if !match {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("not the database error but the others occurred :%w", err)
	}

	user_email := ""
	if value, ok := email.([]byte); ok {
		user_email = string(value)
	}

	return &generated.UserCredentials{
		UserEmail: user_email,
		UserRole:  generated.UserRole(generated.UserRole_value[role]),
		UserId:    user_id,
	}, nil
}

// UPDATE
func UserRegisterUpdateInTransaction(user_info *generated.UserCredentials) error {
	query := `UPDATE db_user_credentials_1.UserCredentials
		SET 
    		email = ?
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
		UserId    int64   = user_info.GetUserId()
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
    		name = ?,
    		bio = ?,
    		gender = ?,
    		bday = ?,
    		updated_at = ?
		WHERE id = ? `

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
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
		_, err = tx.Exec(
			query,
			UserName,
			UserBio,
			UserGender,
			UserBday,
			UserUpdatedAt,
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

func UserUpdateAvatarInTransaction(user_info *generated.UserUpdateAvatar) error {
	query := `UPDATE db_user_1.User 
		SET 
    		avatar = ?
		WHERE id = ? `

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
    		status = ?
		WHERE id = ? `

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

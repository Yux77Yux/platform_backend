package repository

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	common "github.com/Yux77Yux/platform_backend/generated/common"
	generatedUser "github.com/Yux77Yux/platform_backend/generated/user"
)

func UserAddInfoInTransaction(user_info *generatedUser.User) error {
	query := `insert into user 
	(user_id,
	user_name,
	user_avator,
	user_bio,
	user_status,
	user_gender,
	user_bday,
	user_created_at,
	user_updated_at,
	)
	values(?,?,?,?,?,?,?,?,?)`

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
		UserId        int64                     = user_info.GetUserDefault().GetUserId()
		UserName      string                    = user_info.GetUserDefault().GetUserName()
		UserAvator    string                    = user_info.GetUserAvator()
		UserBio       string                    = user_info.GetUserBio()
		UserStatus    generatedUser.User_Status = user_info.GetUserStatus()
		UserGender    generatedUser.User_Gender = user_info.GetUserGender()
		UserBday      *timestamppb.Timestamp
		UserCreatedAt *timestamppb.Timestamp
		UserUpdatedAt *timestamppb.Timestamp
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
			UserBday,
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

func UserGetInfoInTransaction(user_id int64) (*generatedUser.User, error) {
	query := `select 
	user_name 
	user_avator 
	user_bio 
	user_status 
	user_gender 
	user_bday 
	user_created_at 
	user_updated_at 
	where user_id = ?`

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

	var (
		UserName      string
		UserAvator    string
		UserBio       string
		UserStatus    generatedUser.User_Status
		UserGender    generatedUser.User_Gender
		UserBday      *timestamppb.Timestamp
		UserCreatedAt *timestamppb.Timestamp
		UserUpdatedAt *timestamppb.Timestamp
	)
	select {
	case <-ctx.Done():
		err = fmt.Errorf("exec timeout :%w", ctx.Err())
		if errSecond := db.RollbackTransaction(tx); errSecond != nil {
			err = fmt.Errorf("%w and %w", err, errSecond)
		}

		return nil, err
	default:
		err := tx.QueryRow(query, user_id).Scan(&UserName,
			&UserAvator,
			&UserBio,
			&UserStatus,
			&UserGender,
			&UserBday,
			&UserCreatedAt,
			&UserUpdatedAt,
		)
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
	user_info := &generatedUser.User{
		UserDefault: &common.UserDefault{
			UserId:   user_id,
			UserName: UserName,
		},
		UserAvator:    UserAvator,
		UserBio:       UserBio,
		UserStatus:    UserStatus,
		UserGender:    UserGender,
		UserBday:      UserBday,
		UserCreatedAt: UserCreatedAt,
		UserUpdatedAt: UserUpdatedAt,
	}

	return user_info, nil
}

func UserVerifyInTranscation(user_credential *generatedUser.UserCredentials) (int64, error) {
	query := `select password user_id where username = ?`

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	tx, err := db.BeginTransaction()
	if err != nil {
		return 0, err
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

func UserRegisterInTransaction(user_credential *generatedUser.UserCredentials, id int64) error {
	pwd, err := hashPassword(user_credential.GetPassword())

	// 进行复杂加密
	if err != nil {
		return fmt.Errorf("decrypt hash password failed because %w", err)
	}

	query := `insert into UserCredentials(username,password,email,user_id) 
	values(?,?,?,?) 
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
		_, err = tx.Exec(query, user_credential.GetUsername(), pwd, user_credential.GetEmail(), id)
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

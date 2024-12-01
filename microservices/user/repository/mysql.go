package repository

import (
	"context"
	"fmt"
	"time"

	model "github.com/Yux77Yux/platform_backend/generated/user"
)

func UserRegisterInTransaction(user_credential *model.UserCredentials) error {
	pwd, err := hashPassword(user_credential.GetPassword())

	// 进行复杂加密
	if err != nil {
		return fmt.Errorf("decrypt hash password failed because %w", err)
	}

	query := `insert into UserCredentials(username,password,email) 
	values(?,?,?) 
	`

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
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
		_, err = tx.Exec(query, user_credential.GetUsername(), pwd, user_credential.GetEmail())
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

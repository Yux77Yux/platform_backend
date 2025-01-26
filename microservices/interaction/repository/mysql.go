package repository

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	generated "github.com/Yux77Yux/platform_backend/generated/interaction"
	tools "github.com/Yux77Yux/platform_backend/microservices/interaction/tools"
)

// GET

func GetActionTag(req *generated.BaseInteraction) (*generated.Interaction, error) {
	ctx := context.Background()

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

	select {
	case <-ctx.Done():
		err = fmt.Errorf("exec timeout :%w", ctx.Err())
		if errSecond := db.RollbackTransaction(tx); errSecond != nil {
			err = fmt.Errorf("%w and %w", err, errSecond)
		}

		return nil, err
	default:
	}

	return nil, nil
}

func GetCollections(interactions []*generated.AnyInteraction) ([]*generated.Interaction, error) {
	ctx := context.Background()

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

	select {
	case <-ctx.Done():
		err = fmt.Errorf("exec timeout :%w", ctx.Err())
		if errSecond := db.RollbackTransaction(tx); errSecond != nil {
			err = fmt.Errorf("%w and %w", err, errSecond)
		}

		return nil, err
	default:
	}

	return nil, nil
}

func GetHistories(userId int64) ([]*generated.Interaction, error) {
	ctx := context.Background()

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

	select {
	case <-ctx.Done():
		err = fmt.Errorf("exec timeout :%w", ctx.Err())
		if errSecond := db.RollbackTransaction(tx); errSecond != nil {
			err = fmt.Errorf("%w and %w", err, errSecond)
		}

		return nil, err
	default:
	}

	return nil, nil
}

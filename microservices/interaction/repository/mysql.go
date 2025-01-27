package repository

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	generated "github.com/Yux77Yux/platform_backend/generated/interaction"
)

// GET

func GetActionTag(req *generated.BaseInteraction) (*generated.Interaction, error) {
	query := `
		SELECT 
			action_tag
		FROM db_interaction_1.Interaction
		WHERE user_id = ?
		AND creation_id = ?`

	userId := req.GetUserId()
	creationId := req.GetCreationId()

	var actionTag int32

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
		err := tx.QueryRowContext(ctx,
			query,
			userId,
			creationId).Scan(&actionTag)
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

	return &generated.Interaction{
		Base: &generated.BaseInteraction{
			UserId:     userId,
			CreationId: creationId,
		},
		ActionTag: actionTag,
	}, nil
}

func GetCollections(userId int64) ([]*generated.Interaction, error) {
	query := `
		SELECT 
			creation_id,
			save_at
		FROM db_interaction_1.Interaction
		WHERE user_id = ?
		AND action_tag & 4 = 4
		ORDER BY save_at DESC`

	var interactions []*generated.Interaction

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
		rows, err := tx.QueryContext(ctx, query, userId)
		if err != nil {
			err = fmt.Errorf("transaction exec failed because %v", err)
			if errSecond := db.RollbackTransaction(tx); errSecond != nil {
				err = fmt.Errorf("%w and %w", err, errSecond)
			}

			return nil, err
		}
		defer rows.Close()

		for rows.Next() {
			var (
				creation_id int64
				save_at     time.Time
			)
			err := rows.Scan(&creation_id, &save_at)
			if err != nil {
				err = fmt.Errorf("error: GetCollections rows.Scan error %v", err)
				if errSecond := db.RollbackTransaction(tx); errSecond != nil {
					err = fmt.Errorf("%w and %w", err, errSecond)
				}
				return nil, err
			}
			interactions = append(interactions, &generated.Interaction{
				Base:   &generated.BaseInteraction{CreationId: creation_id},
				SaveAt: timestamppb.New(save_at),
			})
		}

		if err = db.CommitTransaction(tx); err != nil {
			return nil, err
		}
	}

	return interactions, nil
}

func GetHistories(userId int64) ([]*generated.Interaction, error) {
	query := `
		SELECT 
			creation_id,
			updated_at
		FROM db_interaction_1.Interaction
		WHERE user_id = ?
		AND action_tag & 1 = 1
		ORDER BY updated_at DESC`

	var interactions []*generated.Interaction

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
		rows, err := tx.QueryContext(ctx, query, userId)
		if err != nil {
			err = fmt.Errorf("transaction exec failed because %v", err)
			if errSecond := db.RollbackTransaction(tx); errSecond != nil {
				err = fmt.Errorf("%w and %w", err, errSecond)
			}

			return nil, err
		}
		defer rows.Close()

		for rows.Next() {
			var (
				creation_id int64
				updated_at  time.Time
			)
			err := rows.Scan(&creation_id, &updated_at)
			if err != nil {
				err = fmt.Errorf("error: GetHistories rows.Scan error %v", err)
				if errSecond := db.RollbackTransaction(tx); errSecond != nil {
					err = fmt.Errorf("%w and %w", err, errSecond)
				}
				return nil, err
			}
			interactions = append(interactions, &generated.Interaction{
				Base:      &generated.BaseInteraction{CreationId: creation_id},
				UpdatedAt: timestamppb.New(updated_at),
			})
		}

		if err = db.CommitTransaction(tx); err != nil {
			return nil, err
		}
	}

	return interactions, nil
}

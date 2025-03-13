package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	generated "github.com/Yux77Yux/platform_backend/generated/interaction"
)

// GET

func GetActionTag(ctx context.Context, req *generated.BaseInteraction) (*generated.Interaction, error) {
	query := `
		SELECT 
			action_tag
		FROM db_interaction_1.Interaction
		WHERE user_id = ?
		AND creation_id = ?`

	userId := req.GetUserId()
	creationId := req.GetCreationId()

	var actionTag int32

	select {
	case <-ctx.Done():
		err := fmt.Errorf("exec timeout :%w", ctx.Err())

		return nil, err
	default:
		err := db.QueryRowContext(
			ctx,
			query,
			userId,
			creationId).Scan(&actionTag)
		if err != nil {
			if err != sql.ErrNoRows {
				return nil, err
			}
		}
	}

	return &generated.Interaction{
		Base:      req,
		ActionTag: actionTag,
	}, nil
}

func GetCollections(ctx context.Context, userId int64, page int32) ([]*generated.Interaction, error) {
	offset := (page - 1) * 30
	query := `
		SELECT 
			creation_id,
			save_at
		FROM db_interaction_1.Interaction
		WHERE user_id = ?
		AND action_tag & 4 = 4
		ORDER BY save_at DESC
		LIMIT 30 OFFSET ?`

	var interactions []*generated.Interaction

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		rows, err := db.QueryContext(ctx, query, userId, offset)
		if err != nil {
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
				return nil, err
			}
			interactions = append(interactions, &generated.Interaction{
				Base:   &generated.BaseInteraction{CreationId: creation_id},
				SaveAt: timestamppb.New(save_at),
			})
		}
	}

	return interactions, nil
}

func GetHistories(ctx context.Context, userId int64, page int32) ([]*generated.Interaction, error) {
	offset := (page - 1) * 30
	query := `
		SELECT 
			creation_id,
			updated_at
		FROM db_interaction_1.Interaction
		WHERE user_id = ?
		AND action_tag & 1 = 1
		ORDER BY updated_at DESC
		LIMIT 30 OFFSET ?`

	var interactions []*generated.Interaction

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		rows, err := db.QueryContext(ctx, query, userId, offset)
		if err != nil {
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
				return nil, err
			}
			interactions = append(interactions, &generated.Interaction{
				Base:      &generated.BaseInteraction{CreationId: creation_id},
				UpdatedAt: timestamppb.New(updated_at),
			})
		}
	}

	return interactions, nil
}

// 用于推荐系统返回
func GetOtherUserHistories(ctx context.Context, userId int64, page int32) ([]*generated.Interaction, error) {
	const limit = 5000
	offset := (page - 1) * limit
	query := `
		SELECT 
			creation_id,
			action_tag,
			updated_at
		FROM db_interaction_1.Interaction
		WHERE user_id < ? 
		OR user_id > ?
		ORDER BY updated_at DESC
		LIMIT 5000 offset ?`

	var interactions []*generated.Interaction

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		rows, err := db.QueryContext(ctx, query, userId, userId, offset)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {
			var (
				creation_id int64
				action_tag  int32
				updated_at  time.Time
			)
			err := rows.Scan(&creation_id, &action_tag, &updated_at)
			if err != nil {
				err = fmt.Errorf("error: GetHistories rows.Scan error %v", err)
				return nil, err
			}
			interactions = append(interactions, &generated.Interaction{
				Base:      &generated.BaseInteraction{CreationId: creation_id},
				ActionTag: action_tag,
				UpdatedAt: timestamppb.New(updated_at),
			})
		}
	}

	return interactions, nil
}

// UPDATE
func UpdateInteractions(req []*generated.OperateInteraction) error {
	const (
		QM = "(?,?,?,?,?)"
	)
	count := len(req)
	sqlStr := make([]string, count)
	values := make([]interface{}, count*5)
	for i, val := range req {
		sqlStr[i] = QM
		values[i*5] = val.GetBase().GetUserId()
		values[i*5+1] = val.GetBase().GetCreationId()
		values[i*5+2] = val.GetAction().Number()
		values[i*5+3] = val.GetUpdatedAt().AsTime()
		saveTime := val.GetSaveAt().AsTime()
		var value sql.NullTime
		if saveTime.Unix() != 0 { // 只在非1970-01-01时写入时间
			value = sql.NullTime{Time: saveTime, Valid: true}
		} else {
			value = sql.NullTime{Valid: false}
		}
		values[i*5+4] = value
	}

	query := fmt.Sprintf(`
		INSERT INTO db_interaction_1.Interaction (user_id, creation_id, action_tag,updated_at,save_at)
		VALUES %s
		ON DUPLICATE KEY UPDATE action_tag = CASE 
        	WHEN VALUES(action_tag) = 1 THEN action_tag | 1
        	WHEN VALUES(action_tag) = 2 THEN action_tag | 2
        	WHEN VALUES(action_tag) = 4 THEN action_tag | 4
        	WHEN VALUES(action_tag) = 3 THEN action_tag & 3
        	WHEN VALUES(action_tag) = 5 THEN action_tag & 5
        	WHEN VALUES(action_tag) = 6 THEN action_tag & 6
        	ELSE action_tag
		END,
    		updated_at = CASE
        	WHEN VALUES(updated_at) IS NOT NULL THEN VALUES(updated_at)
        	ELSE updated_at
		END,
    		save_at = VALUES(save_at);`, strings.Join(sqlStr, ","))

	ctx := context.Background()

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		_, err := db.ExecContext(
			ctx,
			query,
			values...,
		)
		if err != nil {
			return err
		}
	}

	return nil
}

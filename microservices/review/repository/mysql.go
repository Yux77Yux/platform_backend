package repository

import (
	"context"
	"database/sql"
	"fmt"
	"math"
	"strings"
	"time"

	generated "github.com/Yux77Yux/platform_backend/generated/review"
	errMap "github.com/Yux77Yux/platform_backend/pkg/error"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// GET
func GetReviews(
	ctx context.Context,
	reviewId int64,
	reviewType generated.TargetType,
	status generated.ReviewStatus,
	page int32,
) ([]*generated.Review, int32, error) {
	const Limit = 10
	var (
		count   int32 = 0
		orderBy       = " ORDER BY created_at, id "
	)
	if status != generated.ReviewStatus_PENDING {
		orderBy = " ORDER BY updated_at DESC, id  "
	}

	offset := (page - 1) * Limit
	query := fmt.Sprintf(`
		SELECT 
			id,
			target_id,
			detail,
			updated_at,
			created_at,
			remark
		FROM db_review_1.Review
		WHERE reviewer_id = ? 
		AND 
			target_type = ?
		AND 
			status = ? 
		%s
		LIMIT ? 
		OFFSET ?`, orderBy)

	if page <= 1 {
		var num int32 = 0
		queryCount := `
		SELECT 
			count(*) AS count
		FROM db_review_1.Review
		WHERE reviewer_id = ? 
		AND 
			target_type = ?
		AND 
			status = ?`
		err := db.QueryRowContext(
			ctx,
			queryCount,
			reviewId,
			reviewType.String(),
			status.String(),
		).Scan(&num)
		if err != nil {
			return nil, -1, errMap.MapMySQLErrorToStatus(err)
		}
		if num <= 0 {
			return nil, 0, nil
		}

		count = int32(math.Ceil(float64(num) / float64(Limit)))
	}

	rows, err := db.QueryContext(
		ctx,
		query,
		reviewId,
		reviewType.String(),
		status.String(),
		Limit,
		offset,
	)
	if err != nil {
		return nil, -1, err
	}
	defer rows.Close()

	reviews := make([]*generated.Review, 0, Limit)
	for rows.Next() {
		var (
			id         int64
			target_id  int64
			detail     string
			updated_at time.Time
			created_at time.Time
			remark     sql.NullString
		)

		err := rows.Scan(&id, &target_id, &detail, &updated_at, &created_at, &remark)
		if err != nil {
			return nil, -1, err
		}

		review := &generated.Review{
			New: &generated.NewReview{
				Id:         id,
				TargetId:   target_id,
				TargetType: reviewType,
				CreatedAt:  timestamppb.New(created_at),
				Msg:        detail,
			},
			ReviewerId: reviewId,
			Status:     status,
			Remark:     remark.String,
			UpdatedAt:  timestamppb.New(updated_at),
		}
		reviews = append(reviews, review)
	}

	return reviews, count, nil
}

func GetTarget(ctx context.Context, id int64) (int64, *generated.TargetType, error) {
	query := `
		SELECT target_id,target_type
		FROM db_review_1.Review
		WHERE id = ? `

	var (
		targetId   int64
		targetType string
	)

	err := db.QueryRow(query, id).Scan(&targetId, &targetType)
	if err != nil {
		return -1, nil, errMap.MapMySQLErrorToStatus(err)
	}
	target_Type := generated.TargetType(generated.TargetType_value[targetType])
	return targetId, &target_Type, nil
}

// POST
func PostReviews(ctx context.Context, reviews []*generated.NewReview) error {
	const (
		QM          = "(?,?,?,?)"
		FieldsCount = 4
	)
	length := len(reviews)
	if length <= 0 {
		return nil
	}
	sqlStr := make([]string, length)
	values := make([]any, length*FieldsCount)
	for i, review := range reviews {
		sqlStr[i] = QM

		values[FieldsCount*i] = review.GetId()
		values[FieldsCount*i+1] = review.GetTargetId()
		values[FieldsCount*i+2] = review.GetTargetType().String()
		values[FieldsCount*i+3] = nil
		msg := review.GetMsg()
		if msg != "" {
			values[FieldsCount*i+3] = msg
		}
	}

	query := fmt.Sprintf(`
		INSERT INTO db_review_1.Review(id,target_id,target_type,detail)
		VALUES %s
		ON DUPLICATE KEY UPDATE id = id;`, strings.Join(sqlStr, ","))

	_, err := db.ExecContext(
		ctx,
		query,
		values...,
	)

	return err
}

// UPDATE
func UpdateReviews(ctx context.Context, reviews []*generated.Review) error {
	const (
		QM          = "?"
		QQM         = "WHEN id = ? THEN ?"
		FieldsCount = 2 * 3
	)
	length := len(reviews)
	if length <= 0 {
		return nil
	}

	QMS := make([]string, length)
	sqlStr := make([]string, length)
	values := make([]any, length*7)
	for i, review := range reviews {
		QMS[i] = QM
		sqlStr[i] = QQM
		id := review.GetNew().GetId()

		values[i*2+0] = id
		values[i*2+1] = review.GetStatus().String()

		values[length*2+i*2+0] = id
		values[length*2+i*2+1] = nil
		msg := review.GetRemark()
		if msg != "" {
			values[length+i*2+1] = msg
		}

		values[length*4+i*2+0] = id
		values[length*4+i*2+1] = review.GetReviewerId()

		values[length*6+i] = id
	}

	join := strings.Join(sqlStr, " ")
	query := fmt.Sprintf(`
		UPDATE db_review_1.Review 
		SET 
			status = CASE 
				%s
			END,
			remark = CASE 
				%s
			END,
			reviewer_id = CASE 
				%s
			END
		WHERE id IN (%s)`,
		join,
		join,
		join,
		strings.Join(QMS, ","))

	_, err := db.ExecContext(
		ctx,
		query,
		values...,
	)

	return err
}

func UpdateReview(ctx context.Context, review *generated.Review) error {
	query := `
		UPDATE db_review_1.Review 
		SET 
			status = ?,
			remark =?,
			reviewer_id = ?
		WHERE id = ?`

	var (
		status      = review.GetStatus().String()
		remark      = review.GetRemark()
		reviewer_id = review.GetReviewerId()
		id          = review.GetNew().GetId()
	)

	_, err := db.ExecContext(
		ctx,
		query,
		status,
		remark,
		reviewer_id,
		id,
	)

	return err
}

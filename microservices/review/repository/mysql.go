package repository

import (
	"context"
	"fmt"
	"strings"
	"time"

	generated "github.com/Yux77Yux/platform_backend/generated/review"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// GET
func GetReviews(
	ctx context.Context,
	reviewId int64,
	reviewType generated.TargetType,
	status generated.ReviewStatus,
	page int32,
) ([]*generated.Review, error) {
	const Limit = 50
	offset := (page - 1) * Limit
	query := `
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
		ORDER BY created_at
		LIMIT ? 
		OFFSET ?`

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
		return nil, err
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
			remark     string
		)

		err := rows.Scan(&id, &target_id, &detail, &updated_at, &created_at, &remark)
		if err != nil {
			return nil, err
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
			Remark:     remark,
			UpdatedAt:  timestamppb.New(updated_at),
		}
		reviews = append(reviews, review)
	}

	return reviews, nil
}

// POST
func PostReviews(reviews []*generated.NewReview) error {
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

	ctx := context.Background()
	_, err := db.ExecContext(
		ctx,
		query,
		values...,
	)

	return err
}

// UPDATE
func UpdateReviews(reviews []*generated.Review) error {
	const (
		QM          = "?"
		QQM         = "WHEN id = ? THEN ?"
		FieldsCount = 2 * 3
	)
	length := len(reviews)
	if length <= 0 {
		return nil
	}
	idStart := FieldsCount * length

	QMS := make([]string, length)
	sqlStr := make([]string, length)
	values := make([]any, length+idStart)
	for i, review := range reviews {
		QMS[i] = QM
		sqlStr[i] = QQM
		id := review.GetNew().GetId()

		values[FieldsCount*i] = id
		values[FieldsCount*i+1] = review.GetStatus().String()

		values[FieldsCount*i+2] = id
		values[FieldsCount*i+3] = nil
		msg := review.GetRemark()
		if msg != "" {
			values[FieldsCount*i+3] = msg
		}

		values[FieldsCount*i+4] = id
		values[FieldsCount*i+5] = review.GetReviewerId()

		values[idStart+i] = id
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

	ctx := context.Background()
	_, err := db.ExecContext(
		ctx,
		query,
		values...,
	)

	return err
}

func UpdateReview(review *generated.Review) error {
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

	ctx := context.Background()
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

package repository

import (
	"context"
	"fmt"
	"strings"

	generated "github.com/Yux77Yux/platform_backend/generated/review"
	snow "github.com/Yux77Yux/platform_backend/pkg/snow"
)

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
		id := snow.GetId()

		values[FieldsCount*i] = id
		values[FieldsCount*i+1] = review.GetTargetId()
		values[FieldsCount*i+2] = review.GetTargetType()
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

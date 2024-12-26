package repository

import (
	"context"
	"fmt"
	"time"

	generated "github.com/Yux77Yux/platform_backend/generated/creation"
)

// POST
func CreationAddInTransaction(creation *generated.Creation) error {
	queryCreation := `insert into db_creation_1.Creation 
	(id,
	author_id,
	src,
	thumbnail,
	title,
	bio,
	status,
	duration,
	category_id,
	upload_time
	)
	values(?,?,?,?,?,?,?,?,?,?)`

	// queryCreationEngagement := `insert into db_creation_engagment_1.CreationEngagement
	// (creation_id
	// )
	// values(?)`

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
		id          = creation.GetCreationId()
		author_id   = creation.GetBaseInfo().GetAuthorId()
		arc         = creation.GetBaseInfo().GetSrc()
		thumbnail   = creation.GetBaseInfo().GetThumbnail()
		title       = creation.GetBaseInfo().GetTitle()
		bio         = creation.GetBaseInfo().GetBio()
		status      = creation.GetBaseInfo().GetStatus().String()
		duration    = creation.GetBaseInfo().GetDuration()
		category_id = creation.GetBaseInfo().GetCategoryId()
		upload_time = creation.GetUploadTime().AsTime()
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
			queryCreation,
			id,
			author_id,
			arc,
			thumbnail,
			title,
			bio,
			status,
			duration,
			category_id,
			upload_time,
		)
		if err != nil {
			err = fmt.Errorf("queryCreation transaction exec failed because %v", err)
			if errSecond := db.RollbackTransaction(tx); errSecond != nil {
				err = fmt.Errorf("%w and %w", err, errSecond)
			}

			return err
		}

		// _, err = tx.Exec(
		// 	queryCreationEngagement,
		// 	id,
		// )
		// if err != nil {
		// 	err = fmt.Errorf("queryCreationEngagement transaction exec failed because %v", err)
		// 	if errSecond := db.RollbackTransaction(tx); errSecond != nil {
		// 		err = fmt.Errorf("%w and %w", err, errSecond)
		// 	}

		// 	return err
		// }

		if err = db.CommitTransaction(tx); err != nil {
			return err
		}
	}

	return nil
}

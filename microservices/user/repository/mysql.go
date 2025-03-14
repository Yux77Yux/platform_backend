package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	common "github.com/Yux77Yux/platform_backend/generated/common"
	generated "github.com/Yux77Yux/platform_backend/generated/user"
	tools "github.com/Yux77Yux/platform_backend/microservices/user/tools"
)

// SET
func UserAddInfoInTransaction(users []*generated.User) error {
	const QM = "(?,?,?,?,?,?,?,?,?)"
	const fieldsCount = 9
	count := len(users)
	if count == 0 {
		return nil // 没有需要更新的内容
	}

	sqlStr := make([]string, count)
	for i := range sqlStr {
		sqlStr[i] = QM
	}

	query := fmt.Sprintf(`insert into db_user_1.User 
	(id,
	name,
	avatar,
	bio,
	status,
	gender,
	bday,
	created_at,
	updated_at)
	VALUES %s`, strings.Join(sqlStr, ","))

	values := make([]interface{}, count*fieldsCount)
	for i, user_info := range users {
		var (
			UserId        int64       = user_info.GetUserDefault().GetUserId()
			UserName      string      = user_info.GetUserDefault().GetUserName()
			UserAvatar    string      = user_info.GetUserAvatar()
			UserBio       string      = user_info.GetUserBio()
			UserStatus    string      = user_info.GetUserStatus().String()
			UserGender    string      = user_info.GetUserGender().String()
			UserBday      interface{} = nil
			UserCreatedAt time.Time   = user_info.GetUserCreatedAt().AsTime()
			UserUpdatedAt time.Time   = user_info.GetUserUpdatedAt().AsTime()
		)
		values[i*9] = UserId
		values[i*9+1] = UserName
		values[i*9+2] = UserAvatar
		values[i*9+3] = UserBio
		values[i*9+4] = UserStatus
		values[i*9+5] = UserGender
		values[i*9+6] = UserBday
		values[i*9+7] = UserCreatedAt
		values[i*9+8] = UserUpdatedAt
		// values = append(values,
		// 	UserId,
		// 	UserName,
		// 	UserAvatar,
		// 	UserBio,
		// 	UserStatus,
		// 	UserGender,
		// 	UserBday,
		// 	UserCreatedAt,
		// 	UserUpdatedAt,
		// )
	}

	ctx := context.Background()

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
		_, err := tx.ExecContext(
			ctx,
			query,
			values...,
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

func UserRegisterInTransaction(user_credentials []*generated.UserCredentials) error {
	const QM = "(?,?,?,?,?)"
	const fieldsCount = 5
	count := len(user_credentials)
	if count == 0 {
		return nil // 没有需要更新的内容
	}

	sqlStr := make([]string, count)
	for i := range sqlStr {
		sqlStr[i] = QM
	}

	query := fmt.Sprintf(`INSERT INTO db_user_credentials_1.UserCredentials(
			id,
			username,
			password,
			email,
			role)
		VALUES%s`, strings.Join(sqlStr, ","))
	values := make([]interface{}, count*fieldsCount)
	for i, user_credential := range user_credentials {
		var email interface{} = nil
		if user_credential.GetUserEmail() != "" {
			email = user_credential.GetUserEmail()
		}

		var (
			UserId       = user_credential.GetUserId()
			Username     = user_credential.GetUsername()
			UserPassword = user_credential.GetPassword()
			UserEmail    = email
			UserRole     = user_credential.GetUserRole().String()
		)
		values[i*5] = UserId
		values[i*5+1] = Username
		values[i*5+2] = UserPassword
		values[i*5+3] = UserEmail
		values[i*5+4] = UserRole
	}

	ctx := context.Background()

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
		_, err = tx.ExecContext(
			ctx,
			query,
			values...,
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

	if err != nil {
		return fmt.Errorf("not the database error but the others occurred :%w", err)
	}
	return nil
}

func Follow(subs []*generated.Follow) error {
	const (
		QM = "(?,?)"
	)
	count := len(subs)
	if count <= 0 {
		return nil
	}
	sqlStr := make([]string, count)
	values := make([]interface{}, count*2)
	for i, val := range subs {
		sqlStr[i] = QM
		values[i*2] = val.FollowerId
		values[i*2+1] = val.FolloweeId
	}

	query := fmt.Sprintf(`
		INSERT INTO db_user_1.Follow (follower_id, followee_id)
		VALUES %s 
		ON DUPLICATE KEY UPDATE
		follower_id = follower_id;`, strings.Join(sqlStr, ","))

	_, err := db.Exec(
		query,
		values...,
	)
	if err != nil {
		err = fmt.Errorf("transaction exec failed because %v", err)

		return err
	}
	return nil
}

// GET
func UserGetInfoInTransaction(ctx context.Context, id int64) (*generated.User, error) {
	query := `
    	SELECT 
    		u.name,
    		u.avatar,
    		u.bio,
    		u.status,
    		u.gender,
    		u.bday,
    		u.created_at,
    		u.updated_at,
    		(SELECT COUNT(*) FROM db_user_1.Follow WHERE followee_id = u.id) AS followers,
    		(SELECT COUNT(*) FROM db_user_1.Follow WHERE follower_id = u.id) AS followees
		FROM db_user_1.User u
		WHERE u.id = ?;`

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		var (
			name       string
			avatar     string
			bio        string
			statusStr  string
			genderStr  string
			bdayOrNull sql.NullTime
			created_at time.Time
			updated_at time.Time
			followers  int32
			followees  int32
		)
		err := db.QueryRowContext(ctx, query, id).Scan(
			&name, &avatar, &bio, &statusStr, &genderStr,
			&bdayOrNull, &created_at, &updated_at, &followers,
			&followees,
		)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}
			return nil, err
		}

		status, ok := generated.UserStatus_value[statusStr]
		if !ok {
			status = int32(generated.UserStatus_INACTIVE) // 设默认值
		}

		gender, ok := generated.UserGender_value[genderStr]
		if !ok {
			gender = int32(generated.UserGender_UNDEFINED)
		}
		bday := bdayOrNull.Time

		return &generated.User{
			UserDefault: &common.UserDefault{
				UserId:   id,
				UserName: name,
			},
			UserAvatar:    avatar,
			UserBio:       bio,
			UserStatus:    generated.UserStatus(status),
			UserGender:    generated.UserGender(gender),
			UserBday:      timestamppb.New(bday),
			UserCreatedAt: timestamppb.New(created_at),
			UserUpdatedAt: timestamppb.New(updated_at),
		}, nil
	}
}

func GetUsers(ctx context.Context, userIds []int64) ([]*common.UserCreationComment, error) {
	count := len(userIds)
	sqlStr := make([]string, count)
	values := make([]any, count)
	users := make([]*common.UserCreationComment, 0, count)
	for i := 0; i < count; i++ {
		sqlStr[i] = "?"
		values[i] = userIds[i]
	}

	query := fmt.Sprintf(`
		SELECT 
			id,
			name,
			avatar,
			bio
		FROM 
			db_user_1.User
		WHERE id IN (%s)`, strings.Join(sqlStr, ","))

	rows, err := db.QueryContext(ctx, query, values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int64
		var name, avatar, bio string

		if err := rows.Scan(&id, &name, &avatar, &bio); err != nil {
			log.Printf("error: row Scan %v", err) // 处理扫描错误
			return nil, err
		}
		users = append(users, &common.UserCreationComment{
			UserDefault: &common.UserDefault{
				UserId:   id,
				UserName: name,
			},
			UserAvatar: avatar,
			UserBio:    bio,
		})
	}

	// 检查是否有错误发生在遍历过程中
	if err := rows.Err(); err != nil {
		log.Printf("error: rows iteration %v", err)
		return nil, err
	}

	return users, nil
}

func GetFolloweers(ctx context.Context, userId int64, page int32) ([]*common.UserCreationComment, error) {
	const LIMIT = 20
	start := int64((page - 1) * LIMIT)

	query := `
		SELECT 
			id,
			name,
			avatar,
			bio
		FROM 
			db_user_1.User u
		JOIN
		(
			SELECT follower_id
			FROM db_user_1.Follow
			WHERE followee_id = ?
			ORDER BY created_at DESC
			LIMIT ? 
			OFFSET ?
		) f
		ON u.id = f.follower_id;`

	results := make([]*common.UserCreationComment, 0, 20)
	rows, err := db.QueryContext(ctx, query, userId, LIMIT, start)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int64
		var name, avatar, bio string
		if err := rows.Scan(&id, &name, &avatar, &bio); err != nil {
			log.Printf("error: row Scan %v", err) // 处理扫描错误
			return nil, err
		}
		results = append(results, &common.UserCreationComment{
			UserDefault: &common.UserDefault{
				UserId:   id,
				UserName: name,
			},
			UserAvatar: avatar,
			UserBio:    bio,
		})
	}

	// 检查是否有错误发生在遍历过程中
	if err := rows.Err(); err != nil {
		log.Printf("error: rows iteration %v", err)
		return nil, err
	}

	return results, nil
}

func GetFolloweesByTime(ctx context.Context, userId int64, page int32) ([]*common.UserCreationComment, error) {
	const LIMIT = 20
	start := int64((page - 1) * LIMIT)

	query := `
		SELECT 
			id,
			name,
			avatar,
			bio
		FROM 
			db_user_1.User u
		JOIN
		(
			SELECT followee_id
			FROM db_user_1.Follow
			WHERE follower_id = ?
			ORDER BY created_at DESC
			LIMIT ? 
			OFFSET ?
		) f
		ON u.id = f.followee_id;`

	results := make([]*common.UserCreationComment, 0, 20)
	rows, err := db.QueryContext(ctx, query, userId, LIMIT, start)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int64
		var name, avatar, bio string
		if err := rows.Scan(&id, &name, &avatar, &bio); err != nil {
			log.Printf("error: row Scan %v", err) // 处理扫描错误
			return nil, err
		}
		results = append(results, &common.UserCreationComment{
			UserDefault: &common.UserDefault{
				UserId:   id,
				UserName: name,
			},
			UserAvatar: avatar,
			UserBio:    bio,
		})
	}

	// 检查是否有错误发生在遍历过程中
	if err := rows.Err(); err != nil {
		log.Printf("error: rows iteration %v", err)
		return nil, err
	}

	return results, nil
}

func GetFolloweesByViews(ctx context.Context, userId int64, page int32) ([]*common.UserCreationComment, error) {
	const LIMIT = 20
	start := int64((page - 1) * LIMIT)

	query := `
		SELECT 
			id,
			name,
			avatar,
			bio
		FROM 
			db_user_1.User u
		JOIN
		(
			SELECT followee_id
			FROM db_user_1.Follow
			WHERE follower_id = ?
			ORDER BY views DESC
			LIMIT ? 
			OFFSET ?
		) f
		ON u.id = f.followee_id;`

	results := make([]*common.UserCreationComment, 0, 20)
	rows, err := db.QueryContext(ctx, query, userId, LIMIT, start)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int64
		var name, avatar, bio string
		if err := rows.Scan(&id, &name, &avatar, &bio); err != nil {
			log.Printf("error: row Scan %v", err) // 处理扫描错误
			return nil, err
		}
		results = append(results, &common.UserCreationComment{
			UserDefault: &common.UserDefault{
				UserId:   id,
				UserName: name,
			},
			UserAvatar: avatar,
			UserBio:    bio,
		})
	}

	// 检查是否有错误发生在遍历过程中
	if err := rows.Err(); err != nil {
		log.Printf("error: rows iteration %v", err)
		return nil, err
	}

	return results, nil
}

// Verify
func UserVerifyInTranscation(ctx context.Context, user_credential *generated.UserCredentials) (*generated.UserCredentials, error) {
	identifier := "username = ?"
	value := user_credential.GetUsername()
	if user_credential.GetUserEmail() != "" {
		identifier = "email = ?"
		value = user_credential.GetUserEmail()
	}

	query := fmt.Sprintf(`SELECT 
			id,
			password,
			email,
			role
		FROM db_user_credentials_1.UserCredentials 
		WHERE %s`, identifier)

	var (
		passwordHash string
		id           int64
		email        interface{}
		role         string
	)

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		err := db.QueryRow(query, value).Scan(&id, &passwordHash, &email, &role)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}
			return nil, err
		}
	}

	match, err := tools.VerifyPassword(passwordHash, user_credential.GetPassword())
	if err != nil {
		return nil, fmt.Errorf("failed to verify password: %w", err)
	}
	if !match {
		return nil, nil
	}

	user_email := ""
	if value, ok := email.([]byte); ok {
		user_email = string(value)
	}

	return &generated.UserCredentials{
		Username:  user_credential.GetUsername(),
		UserEmail: user_email,
		UserRole:  generated.UserRole(generated.UserRole_value[role]),
		UserId:    id,
		Password:  passwordHash,
	}, nil
}

// UPDATE
func UserEmailUpdateInTransaction(user_credentials []*generated.UserCredentials) error {
	const QM = "?"
	const Conf = "WHEN id = ? THEN ?"
	const fieldsCount = 1*2 + 1 // 一行2+1个问号
	count := len(user_credentials)
	if count == 0 {
		return nil // 没有需要更新的内容
	}

	// 构建 CASE 表达式
	Cases := make([]string, count)
	sqlStr := make([]string, count)

	capacity := count * fieldsCount
	length := capacity - count

	values := make([]any, capacity)
	for i, cred := range user_credentials {
		sqlStr[i] = QM
		Cases[i] = Conf
		id := cred.GetUserId()

		values[i*2] = id
		values[i*2+1] = cred.GetUserEmail()
		values[length+i] = id
	}

	query := fmt.Sprintf(`
		UPDATE db_user_credentials_1.UserCredentials
		SET 
			email = CASE 
				%s 
			END
		WHERE id IN (%s)`, strings.Join(Cases, " "), strings.Join(sqlStr, ","))

	ctx := context.Background()
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
		_, err := tx.ExecContext(
			ctx,
			query,
			values...,
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

func UserUpdateSpaceInTransaction(users []*generated.UserUpdateSpace) error {
	const QM = "?"
	const Conf = "WHEN id = ? THEN ?"
	const fieldsCount = 4*2 + 1 // 一个用户需要4*2+1个问号
	count := len(users)
	if count == 0 {
		return nil // 没有需要更新的内容
	}

	// 构建 CASE 表达式
	Cases := make([]string, count)

	casesFillCount := count * 2
	nameCount := casesFillCount * 0
	bioCount := casesFillCount * 1
	genderCount := casesFillCount * 2
	bdayCount := casesFillCount * 3

	sqlStr := make([]string, count)

	caseLength := casesFillCount * 4
	capacity := fieldsCount * count

	values := make([]any, capacity)
	for i, user := range users {
		sqlStr[i] = QM
		Cases[i] = Conf

		userDefault := user.GetUserDefault()
		id := userDefault.GetUserId()

		values[nameCount+i*2] = id
		values[nameCount+i*2+1] = userDefault.GetUserName()

		values[bioCount+i*2] = id
		values[bioCount+i*2+1] = user.GetUserBio()

		values[genderCount+i*2] = id
		values[genderCount+i*2+1] = user.GetUserGender().String()

		values[bdayCount+i*2] = id
		values[bdayCount+i*2+1] = user.GetUserBday().AsTime()

		values[caseLength+i] = id
	}

	// 拼接最终的 SQL
	query := fmt.Sprintf(`
		UPDATE db_user_1.User
		SET 
			name = CASE 
				%s
			END,
			bio = CASE 
				%s
			END,
			gender = CASE 
				%s
			END,
			bday = CASE 
				%s
			END
		WHERE id 
		IN (%s)`,
		strings.Join(Cases, " "),
		strings.Join(Cases, " "),
		strings.Join(Cases, " "),
		strings.Join(Cases, " "),
		strings.Join(sqlStr, ","),
	)

	ctx := context.Background()

	_, err := db.ExecContext(
		ctx,
		query,
		values...,
	)

	if err != nil {
		return err
	}
	return nil
}

func UserUpdateAvatarInTransaction(users []*generated.UserUpdateAvatar) error {
	const QM = "?"
	const Conf = "WHEN id = ? THEN ?"
	const fieldsCount = 1*2 + 1 // 一行2+1个问号
	count := len(users)
	if count == 0 {
		return nil // 没有需要更新的内容
	}

	// 构建 CASE 表达式
	Cases := make([]string, count)
	sqlStr := make([]string, count)

	capacity := count * fieldsCount
	length := capacity - count

	values := make([]any, capacity)
	for i, user := range users {
		sqlStr[i] = QM
		Cases[i] = Conf
		id := user.GetUserId()

		values[i*2] = id
		values[i*2+1] = user.GetUserAvatar()
		values[length+i] = id
	}

	query := fmt.Sprintf(`UPDATE db_user_1.User 
		SET 
    		avatar = CASE
				%s
			END
		WHERE id IN (%s)`, strings.Join(Cases, " "), strings.Join(sqlStr, ","))

	ctx := context.Background()

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
		_, err := tx.ExecContext(
			ctx,
			query,
			values...,
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

func UserUpdateStatusInTransaction(users []*generated.UserUpdateStatus) error {
	const QM = "?"
	const Conf = "WHEN id = ? THEN ?"
	const fieldsCount = 1*2 + 1 // 一行2+1个问号
	count := len(users)
	if count == 0 {
		return nil // 没有需要更新的内容
	}

	// 构建 CASE 表达式
	Cases := make([]string, count)
	sqlStr := make([]string, count)

	capacity := count * fieldsCount
	length := capacity - count

	values := make([]any, capacity)
	for i, user := range users {
		sqlStr[i] = QM
		Cases[i] = Conf
		id := user.GetUserId()

		values[i*2] = id
		values[i*2+1] = user.GetUserStatus().String()
		values[length+i] = id
	}

	query := fmt.Sprintf(`UPDATE db_user_1.User 
		SET 
    		status = CASE
				%s
			END
		WHERE id IN (%s)`, strings.Join(Cases, " "), strings.Join(sqlStr, ","))

	ctx := context.Background()

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
		_, err := tx.ExecContext(
			ctx,
			query,
			values...,
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

func UserUpdateBioInTransaction(users []*generated.UserUpdateBio) error {
	const QM = "?"
	const Conf = "WHEN id = ? THEN ?"
	const fieldsCount = 1*2 + 1 // 一行2+1个问号
	count := len(users)
	if count == 0 {
		return nil // 没有需要更新的内容
	}

	// 构建 CASE 表达式
	Cases := make([]string, count)
	sqlStr := make([]string, count)

	capacity := count * fieldsCount
	length := capacity - count

	values := make([]any, capacity)
	for i, user := range users {
		sqlStr[i] = QM
		Cases[i] = Conf
		id := user.GetUserId()

		values[i*2] = id
		values[i*2+1] = user.GetUserBio()
		values[length+i] = id
	}

	query := fmt.Sprintf(`UPDATE db_user_1.User 
		SET 
    		bio = CASE
				%s
			END
		WHERE id IN (%s)`, strings.Join(Cases, " "), strings.Join(sqlStr, ","))

	ctx := context.Background()

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
		_, err := tx.ExecContext(
			ctx,
			query,
			values...,
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

func DelReviewer(reviewerId int64) (string, string, error) {
	querySELECT := `
		SELECT 
			username,
			email
		FROM db_user_credentials_1.UserCredentials 
		WHERE id = ?
		FOR UPDATE`
	queryUpdate := `
		UPDATE db_user_credentials_1.UserCredentials 
		SET role = USER 
		WHERE id = ?`

	var (
		username string
		email    sql.NullString
	)

	ctx := context.Background()

	// 开始事务
	tx, err := db.BeginTransaction()
	if err != nil {
		return "", "", err
	}

	// 确保在错误时回滚事务
	defer func() {
		if err != nil {
			_ = db.RollbackTransaction(tx) // 确保事务回滚
		}
	}()

	// 查询 username 并加行锁
	err = tx.QueryRowContext(ctx, querySELECT, reviewerId).Scan(&username, &email)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", "", nil
		}
		return "", "", fmt.Errorf("failed to query username: %w", err)
	}

	// 更新角色
	_, err = tx.ExecContext(ctx, queryUpdate, reviewerId)
	if err != nil {
		return "", "", fmt.Errorf("failed to update role: %w", err)
	}

	// 提交事务
	err = db.CommitTransaction(tx)
	if err != nil {
		return "", "", fmt.Errorf("failed to commit transaction: %w", err)
	}

	return username, email.String, nil
}

func ViewFollowee(subs []*generated.Follow) error {
	const (
		QM = "(?,?)"
	)
	count := len(subs)
	if count <= 0 {
		return nil
	}
	sqlStr := make([]string, count)
	values := make([]interface{}, count*2)
	for i, val := range subs {
		sqlStr[i] = QM
		values[i*2] = val.FollowerId
		values[i*2+1] = val.FolloweeId
	}

	query := fmt.Sprintf(`
		UPDATE db_user_1.Follow 
		SET views = views + 1
		WHERE (follower_id,followee_id) 
		IN 
			(%s);`, strings.Join(sqlStr, ","))

	_, err := db.Exec(
		query,
		values...,
	)
	if err != nil {
		err = fmt.Errorf("transaction exec failed because %v", err)

		return err
	}
	return nil
}

// Del
func CancelFollow(f *generated.Follow) error {
	query := `
		DELETE FROM db_user_1.Follow 
		WHERE follower_id = ?
	 	AND followee_id = ?`
	_, err := db.Exec(
		query,
		f.FollowerId,
		f.FolloweeId,
	)
	if err != nil {
		err = fmt.Errorf("transaction exec failed because %v", err)

		return err
	}

	return nil
}

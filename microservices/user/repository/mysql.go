package repository

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/Yux77Yux/platform_backend/generated/common"
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
		// values = append(values,
		// 	UserId,
		// 	Username,
		// 	UserPassword,
		// 	UserEmail,
		// 	UserRole,
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
		VALUES %s ;`, strings.Join(sqlStr, ","))

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
func UserGetInfoInTransaction(ctx context.Context, id int64, fields []string) (map[string]interface{}, error) {
	var query string
	if len(fields) > 0 {
		// 查询指定字段
		query = fmt.Sprintf("SELECT %s FROM db_user_1.User WHERE id = ?", strings.Join(fields, ", "))
	} else {
		// 查询全部字段
		query = "SELECT * FROM db_user_1.User WHERE id = ?"
	}

	var result map[string]interface{}
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		rows, err := db.QueryContext(ctx, query, id)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		// 获取列名
		cols, err := rows.Columns()
		if err != nil {
			err = fmt.Errorf("failed to get columns: %v", err)
			return nil, err
		}

		// 确保有结果,无结果直接返回
		if !rows.Next() {
			return nil, nil
		}

		// 创建一个存储列值的切片
		values := make([]interface{}, len(cols))
		pointers := make([]interface{}, len(cols))
		for i := range values {
			pointers[i] = &values[i]
		}

		// 扫描结果到指针数组
		if err := rows.Scan(pointers...); err != nil {
			return nil, err
		}

		// 将结果填充到 map 中
		result = make(map[string]interface{})
		for i, colName := range cols {
			switch colName {
			case "id":
				result[colName] = id
			case "bday":
				if values[i] == nil {
					result[colName] = "none"
				} else {
					if value, ok := values[i].([]byte); ok {
						// 将字符串解析为 time.Time（假设格式是 "YYYY-MM-DD"）
						parsedTime, err := time.Parse("2006-01-02", string(value))
						if err != nil {
							return nil, err
						}

						result[colName] = timestamppb.New(parsedTime)
					} else {
						return nil, err
					}
				}
			case "created_at", "updated_at":
				if value, ok := values[i].([]byte); ok {
					// 将字符串解析为 time.Time（假设格式是 "YYYY-MM-DD HH:MM:SS"）
					parsedTime, err := time.Parse("2006-01-02 15:04:05", string(value))
					if err != nil {
						return nil, err
					}

					result[colName] = timestamppb.New(parsedTime)
				} else {
					return nil, err
				}
			default:
				if value, ok := values[i].([]byte); ok {
					result[colName] = string(value)
				} else {
					return nil, err
				}
			}
		}

		// 再查UserCredentials拿身份和邮箱
		var (
			email interface{}
			role  string
		)
		query = "SELECT email,role FROM db_user_credentials_1.UserCredentials WHERE id = ?"
		if err := db.QueryRow(query, id).Scan(&email, &role); err != nil {
			return nil, err
		}

		log.Printf("email %v", email)
		if email == nil {
			result["email"] = ""
		} else {
			if value, ok := email.([]byte); ok {
				result["email"] = string(value)
			} else {
				err = fmt.Errorf("assert email type failed ")
				return nil, err
			}
		}
		result["role"] = role
	}
	return result, nil
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
		WHERE IN (%s)`, strings.Join(sqlStr, ","))

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
		UserEmail: user_email,
		UserRole:  generated.UserRole(generated.UserRole_value[role]),
		UserId:    id,
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
	const fieldsCount = 4*2 + 1 // 一行4*2+1个问号
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
		id := user.GetUserDefault().GetUserId()

		values[i*8] = id
		values[i*8+1] = user.GetUserDefault().GetUserName()
		values[i*8+2] = id
		values[i*8+3] = user.GetUserBio()
		values[i*8+4] = id
		values[i*8+5] = user.GetUserGender().String()
		values[i*8+6] = id
		values[i*8+7] = user.GetUserBday().AsTime()
		values[length+i] = id
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

func DelReviewer(reviewerId int64) (string, error) {
	querySELECT := `
		SELECT username 
		FROM db_user_credentials_1.UserCredentials 
		WHERE id = ?
		FOR UPDATE`
	queryUpdate := `
		UPDATE db_user_credentials_1.UserCredentials 
		SET role = USER 
		WHERE id = ?`

	var username string
	ctx := context.Background()

	// 开始事务
	tx, err := db.BeginTransaction()
	if err != nil {
		return "", err
	}

	// 确保在错误时回滚事务
	defer func() {
		if err != nil {
			_ = db.RollbackTransaction(tx) // 确保事务回滚
		}
	}()

	// 查询 username 并加行锁
	err = tx.QueryRowContext(ctx, querySELECT, reviewerId).Scan(&username)
	if err != nil {
		return "", fmt.Errorf("failed to query username: %w", err)
	}

	// 更新角色
	_, err = tx.ExecContext(ctx, queryUpdate, reviewerId)
	if err != nil {
		return "", fmt.Errorf("failed to update role: %w", err)
	}

	// 提交事务
	err = db.CommitTransaction(tx)
	if err != nil {
		return "", fmt.Errorf("failed to commit transaction: %w", err)
	}

	return username, nil
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
	// const (
	// 	QM = "(?,?)"
	// )
	// count := len(subs)
	// if count <= 0 {
	// 	return nil
	// }
	// sqlStr := make([]string, count)
	// values := make([]interface{}, count*2)
	// for i, val := range subs {
	// 	sqlStr[i] = QM
	// 	values[i*2] = val.FollowerId
	// 	values[i*2+1] = val.FolloweeId
	// }

	// query := fmt.Sprintf(`
	// 	DELETE FROM db_user_1.Follow
	// 	WHERE (follower_id,followee_id)
	// 	IN
	// 		(%s);`, strings.Join(sqlStr, ","))

	// _, err := db.Exec(
	// 	query,
	// 	values...,
	// )
	// if err != nil {
	// 	err = fmt.Errorf("transaction exec failed because %v", err)

	// 	return err
	// }
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

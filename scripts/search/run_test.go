package search

import (
	"database/sql"
	"log"
	"strconv"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func gets() []map[string]any {
	dsn := "yuxyuxx:yuxyuxx@tcp(127.0.0.1:13306)/db_creation_1?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=True"

	// 打开数据库连接
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 检查数据库连接是否有效
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	// 执行查询
	query := `
		SELECT id, title, bio
		FROM db_creation_1.Creation`

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	values := make([]map[string]any, 0, 400)
	for rows.Next() {
		var creationId int64
		var title, bio string

		err := rows.Scan(&creationId, &title, &bio)
		if err != nil {
			log.Fatal(err)
		}

		value := make(map[string]any)
		value["id"] = strconv.FormatInt(creationId, 10)
		value["title"] = title
		value["bio"] = bio
		values = append(values, value)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return values
}

func TestRun(t *testing.T) {
	searchService := NewSearchService()

	// 新增健康检查
	if _, err := searchService.Client.Health(); err != nil {
		t.Fatal("MeiliSearch health check failed:", err)
	}

	values := gets()
	log.Printf("第一条数据例: %+v", values[0])
	if err := searchService.AddDocuments("creations", values); err != nil {
		t.Fatal("AddDocuments error:", err)
	}
}

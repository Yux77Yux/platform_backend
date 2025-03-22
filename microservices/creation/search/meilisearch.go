package search

import (
	"fmt"
	"math"
	"strconv"

	"github.com/meilisearch/meilisearch-go"
)

type SearchService struct {
	Client *meilisearch.Client
}

func NewSearchService() *SearchService {
	return &SearchService{
		Client: meilisearch.NewClient(meilisearch.ClientConfig{
			Host:   host,
			APIKey: apiKey,
		}),
	}
}

func (s *SearchService) SearchWithPagination(index, query string, page, pageSize int) ([]int64, int32, error) {
	offset := (page - 1) * pageSize

	// 执行搜索
	searchRes, err := s.Client.Index(index).Search(query, &meilisearch.SearchRequest{
		Limit:  int64(pageSize),
		Offset: int64(offset),
	})
	if err != nil {
		return nil, -1, err
	}

	var ids []int64
	for _, hit := range searchRes.Hits {
		data := hit.(map[string]interface{})

		var id int64
		switch v := data["id"].(type) {
		case string:
			idVal, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				return nil, -1, fmt.Errorf("invalid id format: %v", v)
			}
			id = idVal
		case float64:
			id = int64(v)
		case int64:
			id = v
		default:
			return nil, -1, fmt.Errorf("unexpected id type: %T", v)
		}
		ids = append(ids, id)
	}

	count := searchRes.EstimatedTotalHits
	totalPages := int32(math.Ceil(float64(count) / float64(pageSize)))
	return ids, totalPages, nil
}

func (s *SearchService) AddDocuments(indexName string, documents []map[string]interface{}) error {
	// 1. 创建或获取索引时显式指定主键（兼容旧版SDK）
	index := s.Client.Index(indexName)

	// 2. 确保索引存在并设置主键（需先判断索引是否存在）
	exists, err := s.Client.GetIndex(indexName)
	if err != nil {
		// 如果索引不存在则创建（带主键）
		_, err = s.Client.CreateIndex(&meilisearch.IndexConfig{
			Uid:        indexName,
			PrimaryKey: "id", // 关键设置
		})
		if err != nil {
			return fmt.Errorf("CreateIndex failed: %v", err)
		}
	} else {
		// 如果索引已存在但主键未设置，需先删除重建（旧版本不支持直接修改主键）
		if exists.PrimaryKey != "id" {
			if _, err := s.Client.DeleteIndex(indexName); err != nil {
				return fmt.Errorf("DeleteIndex failed: %v", err)
			}
			if _, err := s.Client.CreateIndex(&meilisearch.IndexConfig{
				Uid:        indexName,
				PrimaryKey: "id",
			}); err != nil {
				return fmt.Errorf("RecreateIndex failed: %v", err)
			}
		}
	}

	// 3. 添加文档（需通过参数指定主键）
	task, err := index.AddDocuments(documents, "id") // 第二个参数指定主键字段名
	if err != nil {
		return err
	}

	// 4. 等待任务完成
	if _, err := s.Client.WaitForTask(task.TaskUID); err != nil { // 旧版 WaitForTask 参数为任务ID
		return err
	}
	return nil
}

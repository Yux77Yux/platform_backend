package search

import (
	"math"

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

		id := int64(data["id"].(float64))
		ids = append(ids, id)
	}

	count := searchRes.EstimatedTotalHits
	totalPages := int32(math.Ceil(float64(count) / float64(pageSize)))
	return ids, totalPages, nil
}

func (s *SearchService) AddDocuments(index string, documents []map[string]interface{}) error {
	// 获取异步任务ID
	task, err := s.Client.Index(index).AddDocuments(documents)
	if err != nil {
		return err
	}

	// 等待任务完成
	_, err = s.Client.WaitForTask(task.TaskUID)
	return err
}

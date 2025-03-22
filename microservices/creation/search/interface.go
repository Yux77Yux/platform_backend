package search

type SearchServiceInterface interface {
	SearchWithPagination(index, query string, page, pageSize int) ([]int64, int32, error)
	AddDocuments(index string, documents []map[string]interface{}) error
}

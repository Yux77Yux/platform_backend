package search

import (
	"github.com/Yux77Yux/platform_backend/microservices/creation/internal"
	"github.com/Yux77Yux/platform_backend/microservices/creation/messaging/receiver"
)

func Run() {
	searchService := NewSearchService()

	internal.InitSearch(searchService)
	receiver.InitSearch(searchService)
}

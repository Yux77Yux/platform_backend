package receiver

import (
	"context"

	messaging "github.com/Yux77Yux/platform_backend/microservices/interaction/messaging"
)

const (
	ComputeSimilarCreation = messaging.ComputeSimilarCreation
	ComputeUser            = messaging.ComputeUser

	BatchUpdateDb = messaging.BatchUpdateDb
	UpdateDb      = messaging.UpdateDb
	AddCollection = messaging.AddCollection
	AddLike       = messaging.AddLike
	AddView       = messaging.AddView
	CancelLike    = messaging.CancelLike

	// Creation
	UPDATE_CREATION_ACTION_COUNT = messaging.UPDATE_CREATION_ACTION_COUNT
)

var (
	ExchangesConfig = messaging.ExchangesConfig
)

// 非RPC类型的消息队列的交换机声明
func Run(ctx context.Context) {
	messaging.Init()
	for exchange := range ExchangesConfig {
		switch exchange {
		// 不同的exchange使用不同函数
		case ComputeSimilarCreation:
			go messaging.ListenToQueue(exchange, ComputeSimilarCreation, ComputeSimilarCreation, computeSimilarProcessor)
		case ComputeUser:
			go messaging.ListenToQueue(exchange, ComputeUser, ComputeUser, computeUserProcessor)

		case UpdateDb:
			go messaging.ListenToQueue(exchange, UpdateDb, UpdateDb, updateDbInteraction)
		case AddView:
			go messaging.ListenToQueue(exchange, AddView, AddView, addViewProcessor)
		case AddCollection:
			go messaging.ListenToQueue(exchange, AddCollection, AddCollection, addCollectionProcessor)
		case AddLike:
			go messaging.ListenToQueue(exchange, AddLike, AddLike, addLikeProcessor)
		case CancelLike:
			go messaging.ListenToQueue(exchange, CancelLike, CancelLike, cancelLikeProcessor)
		case BatchUpdateDb:
			go messaging.ListenToQueue(exchange, BatchUpdateDb, BatchUpdateDb, batchUpdateDbProcessor)
		}
	}

	<-ctx.Done()
	messaging.Close(ctx)
}

package receiver

import (
	"fmt"
	"log"

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
	ExchangesConfig = map[string]string{
		ComputeSimilarCreation: "direct",
		ComputeUser:            "direct",

		UpdateDb:      "direct",
		AddCollection: "direct",
		AddLike:       "direct",
		AddView:       "direct",
		CancelLike:    "direct",
		BatchUpdateDb: "direct",
		// Add more exchanges here
	}
	ListenRPCs = []string{
		"agg_user",
	}
)

// 非RPC类型的消息队列的交换机声明
func Init() {
	rabbitMQ := messaging.GetRabbitMQ()
	defer rabbitMQ.Close()

	if rabbitMQ == nil {
		log.Printf("error: message queue open failed")
		return
	}
	for exchange, kind := range ExchangesConfig {
		if err := rabbitMQ.ExchangeDeclare(exchange, kind, true, false, false, false, nil); err != nil {
			wiredErr := fmt.Errorf("failed to declare exchange %s : %w", exchange, err)
			log.Printf("error: %v", wiredErr)
		}

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
}

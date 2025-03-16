package receiver

import (
	"context"

	messaging "github.com/Yux77Yux/platform_backend/microservices/review/messaging"
)

const (
	New_review      = messaging.New_review
	Comment_review  = messaging.Comment_review
	User_review     = messaging.User_review
	Creation_review = messaging.Creation_review
	PendingCreation = messaging.PendingCreation

	Update      = messaging.Update
	BatchUpdate = messaging.BatchUpdate

	// USER
	USER_APPROVE  = messaging.USER_APPROVE
	USER_REJECTED = messaging.USER_REJECTED

	// CREATION
	CREATION_APPROVE  = messaging.CREATION_APPROVE
	CREATION_REJECTED = messaging.CREATION_REJECTED
	CREATION_DELETED  = messaging.CREATION_DELETED

	// COMMENT
	COMMENT_REJECTED = messaging.COMMENT_REJECTED
	COMMENT_DELETED  = messaging.COMMENT_DELETED
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
		case New_review:
			go messaging.ListenToQueue(exchange, New_review, New_review, NewReviewProcessor)
		case PendingCreation:
			go messaging.ListenToQueue(exchange, PendingCreation, PendingCreation, PendingCreationProcessor)
		case BatchUpdate:
			go messaging.ListenToQueue(exchange, BatchUpdate, BatchUpdate, BatchUpdateProcessor)
		case Update:
			go messaging.ListenToQueue(exchange, Update, Update, UpdateProcessor)
		}
	}

	<-ctx.Done()
	messaging.Close(ctx)
}

var (
	db        SqlMethod
	messaging MessageQueueMethod
	cache     CacheMethod
)

func InitDb(_db SqlMethod) {
	db = _db
}

func InitMQ(_messaging MessageQueueMethod) {
	messaging = _messaging
}

func InitCache(_cache CacheMethod) {
	cache = _cache
}

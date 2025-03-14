package receiver

import (
	"context"

	messaging "github.com/Yux77Yux/platform_backend/microservices/comment/messaging"
)

const (
	PublishComment = messaging.PublishComment
	DeleteComment  = messaging.DeleteComment
)

var (
	ExchangesConfig = messaging.ExchangesConfig
)

func Run(ctx context.Context) {
	messaging.Init()
	for exchange := range ExchangesConfig {
		switch exchange {
		// 不同的exchange使用不同函数
		case PublishComment:
			go messaging.ListenToQueue(exchange, PublishComment, PublishComment, JoinCommentProcessor)
		case DeleteComment:
			go messaging.ListenToQueue(exchange, DeleteComment, DeleteComment, DeleteCommentProcessor)
		}
	}
	<-ctx.Done()
	messaging.Close(ctx)
}

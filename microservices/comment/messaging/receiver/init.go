package receiver

import (
	messaging "github.com/Yux77Yux/platform_backend/microservices/comment/messaging"
)

const (
	PublishComment = messaging.PublishComment
	DeleteComment  = messaging.DeleteComment
)

var (
	ExchangesConfig = messaging.ExchangesConfig
)

func Init(addr string) {
	messaging.InitStr(addr)
	for exchange := range ExchangesConfig {
		switch exchange {
		// 不同的exchange使用不同函数
		case PublishComment:
			go messaging.ListenToQueue(exchange, PublishComment, PublishComment, JoinCommentProcessor)
		case DeleteComment:
			go messaging.ListenToQueue(exchange, DeleteComment, DeleteComment, DeleteCommentProcessor)
		}
	}
}

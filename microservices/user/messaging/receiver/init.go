package messaging

import (
	"context"

	messaging "github.com/Yux77Yux/platform_backend/microservices/user/messaging"
)

const (
	Register         = messaging.Register
	StoreUser        = messaging.StoreUser
	StoreCredentials = messaging.StoreCredentials
	UpdateUserSpace  = messaging.UpdateUserSpace
	UpdateUserAvatar = messaging.UpdateUserAvatar
	UpdateUserBio    = messaging.UpdateUserBio
	Follow           = messaging.Follow

	// review
	UpdateUserStatus = messaging.UpdateUserStatus
	DelReviewer      = messaging.DelReviewer
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
		case Register:
			go messaging.ListenToQueue(exchange, Register, Register, registerProcessor)
		case StoreUser:
			go messaging.ListenToQueue(exchange, StoreUser, StoreUser, storeUserProcessor)
		case StoreCredentials:
			go messaging.ListenToQueue(exchange, StoreCredentials, StoreCredentials, storeCredentialsProcessor)
		case UpdateUserSpace:
			go messaging.ListenToQueue(exchange, UpdateUserSpace, UpdateUserSpace, updateUserSpaceProcessor)
		case UpdateUserAvatar:
			go messaging.ListenToQueue(exchange, UpdateUserAvatar, UpdateUserAvatar, updateUserAvatarProcessor)
		case UpdateUserBio:
			go messaging.ListenToQueue(exchange, UpdateUserBio, UpdateUserBio, updateUserBioProcessor)
		case UpdateUserStatus:
			go messaging.ListenToQueue(exchange, UpdateUserStatus, UpdateUserStatus, updateUserStatusProcessor)
		case DelReviewer:
			go messaging.ListenToQueue(exchange, DelReviewer, DelReviewer, delReviewerProcessor)
		case Follow:
			go messaging.ListenToQueue(exchange, Follow, Follow, followProcessor)
		}
	}

	<-ctx.Done()
	messaging.Close(ctx)
}

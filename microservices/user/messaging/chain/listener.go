package chain

import (
	"log"
	"time"

	"google.golang.org/protobuf/proto"

	generated "github.com/Yux77Yux/platform_backend/generated/user"
	// cache "github.com/Yux77Yux/platform_backend/microservices/user/cache"
	// db "github.com/Yux77Yux/platform_backend/microservices/user/repository"
)

// 监听者结构体
type CredentialsListener struct {
	credentialsChain *CredentialsChain
	channel          chan *generated.UserCredentials // 用于接收评论的通道
	count            int
	timeoutDuration  time.Duration        // 超时持续时间（触发销毁）
	timeoutTimer     *time.Timer          // 用于刷新存活时间
	next             *CredentialsListener // 下一个监听者
}

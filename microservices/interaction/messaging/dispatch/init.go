package dispatch

import (
	"sync"

	"google.golang.org/protobuf/reflect/protoreflect"

	generated "github.com/Yux77Yux/platform_backend/generated/interaction"
)

const (
	LISTENER_CHANNEL_COUNT = 80
	MAX_BATCH_SIZE         = 50
	EXE_CHANNEL_COUNT      = 5

	// sql的插入更新删除都用一套
	DbInteraction = "DbInteraction"

	ViewCache             = "ViewCache"
	LikeCache             = "LikeCache"
	CollectionCache       = "CollectionCache"
	CancelLikeCache       = "CancelLikeCache"
	CancelViewCache       = "CancelViewCache"
	CancelCollectionCache = "CancelCollectionCache"
)

var (
	// 持久化数据库的插入与更新一致
	dbInteractionsChain  *DbInteractionChain
	viewCacheChain       *ViewCacheChain
	likeCacheChain       *LikeCacheChain
	collectionCacheChain *CollectionCacheChain
	interactionsPool     = sync.Pool{
		New: func() any {
			slice := make([]*generated.Interaction, 0, MAX_BATCH_SIZE)
			return &slice
		},
	}

	cancelCollectionCacheChain *CancelCollectionCacheChain
	cancelLikeCacheChain       *CancelLikeCacheChain
	cancelViewCacheChain       *CancelViewCacheChain
	baseInteractionsPool       = sync.Pool{
		New: func() any {
			slice := make([]*generated.BaseInteraction, 0, MAX_BATCH_SIZE)
			return &slice
		},
	}
)

func init() {
	// 初始化责任链
	dbInteractionsChain = InitialDbChain()

	viewCacheChain = InitialViewCacheChain()
	likeCacheChain = InitialLikeCacheChain()
	collectionCacheChain = InitialCollectionCacheChain()

	cancelLikeCacheChain = InitialCancelLikeCacheChain()
	cancelCollectionCacheChain = InitialCancelCollectionCacheChain()
	cancelViewCacheChain = InitialCancelViewCacheChain()
}

func HandleRequest(msg protoreflect.ProtoMessage, typeName string) {
	switch typeName {
	case DbInteraction:
		dbInteractionsChain.HandleRequest(msg)

	case ViewCache:
		viewCacheChain.HandleRequest(msg)
	case LikeCache:
		likeCacheChain.HandleRequest(msg)
	case CollectionCache:
		collectionCacheChain.HandleRequest(msg)
	case CancelLikeCache:
		cancelLikeCacheChain.HandleRequest(msg)
	case CancelCollectionCache:
		cancelCollectionCacheChain.HandleRequest(msg)
	case CancelViewCache:
		cancelViewCacheChain.HandleRequest(msg)
	}
}

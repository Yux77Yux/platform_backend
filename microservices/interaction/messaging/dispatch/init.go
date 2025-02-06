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
	CancelViewCache       = "CancelViewCache"
	CancelLikeCache       = "CancelLikeCache"
	CancelCollectionCache = "CancelCollectionCache"
)

var (
	// 持久化数据库的插入与更新一致
	dbInteractionsChain  *DbInteractionChain
	viewCacheChain       *ViewCacheChain
	collectionCacheChain *CollectionCacheChain
	interactionsPool     = sync.Pool{
		New: func() any {
			slice := make([]*generated.Interaction, 0, MAX_BATCH_SIZE)
			return &slice
		},
	}

	likeCacheChain             *LikeCacheChain
	cancelViewCacheChain       *CancelViewCacheChain
	cancelLikeCacheChain       *CancelLikeCacheChain
	cancelCollectionCacheChain *CancelCollectionCacheChain
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
	cancelViewCacheChain = InitialCancelViewCacheChain()
	cancelLikeCacheChain = InitialCancelLikeCacheChain()
	cancelCollectionCacheChain = InitialCancelCollectionCacheChain()
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
	case CancelViewCache:
		cancelViewCacheChain.HandleRequest(msg)
	case CancelLikeCache:
		cancelLikeCacheChain.HandleRequest(msg)
	case CancelCollectionCache:
		cancelCollectionCacheChain.HandleRequest(msg)
	}
}

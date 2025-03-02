package dispatch

import (
	"log"
	"math"
	"sync"

	"google.golang.org/protobuf/reflect/protoreflect"

	generated "github.com/Yux77Yux/platform_backend/generated/interaction"
)

const (
	LISTENER_CHANNEL_COUNT = 80
	MAX_BATCH_SIZE         = 50
	EXE_CHANNEL_COUNT      = 5

	// sql的插入更新删除都用一套
	DbInteraction      = "DbInteraction"
	DbBatchInteraction = "DbBatchInteraction"

	ViewCache             = "ViewCache"
	LikeCache             = "LikeCache"
	CollectionCache       = "CollectionCache"
	DelViewCache          = "DelViewCache"
	CancelLikeCache       = "CancelLikeCache"
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
			slice := make([]*generated.OperateInteraction, 0, MAX_BATCH_SIZE)
			return &slice
		},
	}

	cancelLikeCacheChain *CancelLikeCacheChain
	baseInteractionsPool = sync.Pool{
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
	case DbBatchInteraction:
		req, ok := msg.(*generated.AnyOperateInteraction)
		if !ok {
			log.Printf("error: req not *generated.AnyOperateInteraction")
			return
		}
		operateInteractions := req.GetOperateInteractions()

		// 计算操作批次的大小
		batchSize := len(operateInteractions)
		batchCount := int(math.Ceil(float64(batchSize) / float64(MAX_BATCH_SIZE))) // 计算需要的批次数

		// 分批处理
		for i := 0; i < batchCount; i++ {
			start := i * MAX_BATCH_SIZE
			end := (i + 1) * MAX_BATCH_SIZE
			if end > batchSize {
				end = batchSize
			}

			// 将分批后的部分赋值给池中的切片
			poolObj := interactionsPool.Get().(*[]*generated.OperateInteraction)
			*poolObj = operateInteractions[start:end]

			// 发往 exeChannel 处理
			dbInteractionsChain.exeChannel <- poolObj
		}
	}
}

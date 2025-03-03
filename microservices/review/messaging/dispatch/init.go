package dispatch

import (
	"log"
	"math"
	"sync"

	"google.golang.org/protobuf/reflect/protoreflect"

	generated "github.com/Yux77Yux/platform_backend/generated/review"
)

const (
	Insert      = "Insert"
	Update      = "Update"
	BatchUpdate = "BatchUpdate"

	LISTENER_CHANNEL_COUNT = 120
	MAX_BATCH_SIZE         = 50
	EXE_CHANNEL_COUNT      = 5
)

var (
	insertChain *InsertChain
	insertPool  = sync.Pool{
		New: func() any {
			slice := make([]*generated.NewReview, 0, MAX_BATCH_SIZE)
			return &slice
		},
	}

	updateChain *UpdateChain
	updatePool  = sync.Pool{
		New: func() any {
			slice := make([]*generated.Review, 0, MAX_BATCH_SIZE)
			return &slice
		},
	}
)

func init() {
	// 初始化责任链
	insertChain = InitialInsertChain()
	updateChain = InitialUpdateChain()
}

func HandleRequest(msg protoreflect.ProtoMessage, typeName string) {
	switch typeName {
	case Insert:
		insertChain.HandleRequest(msg)
	case Update:
		updateChain.HandleRequest(msg)
	case BatchUpdate:
		req, ok := msg.(*generated.AnyReview)
		if !ok {
			log.Printf("error: req not *generated.AnyReview")
			return
		}
		reviews := req.GetReviews()

		// 计算操作批次的大小
		batchSize := len(reviews)
		batchCount := int(math.Ceil(float64(batchSize) / float64(MAX_BATCH_SIZE))) // 计算需要的批次数

		// 分批处理
		for i := 0; i < batchCount; i++ {
			go func() {
				start := i * MAX_BATCH_SIZE
				end := (i + 1) * MAX_BATCH_SIZE
				if end > batchSize {
					end = batchSize
				}

				// 将分批后的部分赋值给池中的切片
				poolObj := updatePool.Get().(*[]*generated.Review)
				*poolObj = reviews[start:end]

				// 发往 exeChannel 处理
				updateChain.exeChannel <- poolObj
			}()
		}
	}
}

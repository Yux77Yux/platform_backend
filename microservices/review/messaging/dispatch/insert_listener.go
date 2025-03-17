package dispatch

import (
	"log"
	"sync/atomic"
	"time"

	"google.golang.org/protobuf/reflect/protoreflect"

	generated "github.com/Yux77Yux/platform_backend/generated/review"
)

// 监听者结构体
type InsertListener struct {
	chain        ChainInterface
	exeChannel   chan *[]*generated.NewReview // 批量发送评论的通道
	datasChannel chan *generated.NewReview    // 用于接收评论的通道
	count        uint32

	timeoutDuration     time.Duration   // 超时持续时间（触发销毁）
	timeoutTimer        *time.Timer     // 用于刷新存活时间
	updateInterval      time.Duration   // 批量插入的间隔时间
	updateIntervalTimer *time.Timer     // 用于周期性执行批量更新
	next                *InsertListener // 下一个监听者
	prev                *InsertListener // 上一个监听者
}

func (listener *InsertListener) GetId() int64 {
	return 0
}

// 启动监听者
func (listener *InsertListener) StartListening() {
	listener.datasChannel = make(chan *generated.NewReview, LISTENER_CHANNEL_COUNT)
	go listener.RestartUpdateIntervalTimer()
	go listener.RestartTimeoutTimer()
}

// 分发评论至通道
func (listener *InsertListener) Dispatch(data protoreflect.ProtoMessage) {
	// 长度加1
	count := atomic.AddUint32(&listener.count, 1)

	_data := data.(*generated.NewReview)
	// 处理评论的逻辑
	listener.datasChannel <- _data

	if count%MAX_BATCH_SIZE == 0 {
		go listener.SendBatch()
		listener.RestartUpdateIntervalTimer()
	}
}

// 执行批量更新
func (listener *InsertListener) SendBatch() {
	const BatchSize = MAX_BATCH_SIZE

	count := atomic.LoadUint32(&listener.count)
	count = calculateBatchSize(count, BatchSize)
	if count == 0 {
		return
	}

	obj := listener.chain.GetPoolObj()
	datasPtr, ok := obj.(*[]*generated.NewReview)
	if !ok {
		log.Fatalf("unexpected type from sync.Pool: %T", obj)
	}
	*datasPtr = (*datasPtr)[:count]
	insertUsers := *datasPtr
	for i := 0; uint32(i) < count; i++ {
		insertUsers[i] = <-listener.datasChannel
	}
	atomic.AddUint32(&listener.count, ^uint32(count-1)) //再减去

	listener.exeChannel <- datasPtr // 送去批量执行,可能被阻塞
}

// 启动周期执行批量更新的定时器
func (listener *InsertListener) RestartUpdateIntervalTimer() {
	if listener.updateIntervalTimer != nil {
		if !listener.updateIntervalTimer.Stop() {
			select {
			case <-listener.updateIntervalTimer.C:
				break
			default:
				break
			}
		}
	}

	// 再执行
	listener.updateIntervalTimer = time.AfterFunc(listener.updateInterval, func() {
		count := atomic.LoadUint32(&listener.count)

		if count > 0 {
			go listener.SendBatch() // 执行批量更新
			listener.RestartTimeoutTimer()
		}
		listener.RestartUpdateIntervalTimer() // 重启定时器
	})
}

// 启动存活时间的定时器
func (listener *InsertListener) RestartTimeoutTimer() {
	// 先重置
	if listener.timeoutTimer != nil {
		// 如果 timer 已存在，确保安全地重置
		if !listener.timeoutTimer.Stop() {
			select {
			case <-listener.timeoutTimer.C:
				break
			default:
				break
			}
		}
	}

	listener.timeoutTimer = time.AfterFunc(listener.timeoutDuration, func() {
		count := atomic.LoadUint32(&listener.count)

		if count == 0 {
			// 超时后销毁监听者
			listener.Cleanup()
			listener.chain.DestroyListener(listener)
		} else {
			listener.RestartTimeoutTimer() // 重启定时器
		}
	})
}

// 清理监听者资源
func (listener *InsertListener) Cleanup() {
	// 关闭评论通道
	close(listener.datasChannel)

	// 清理其他资源（例如定时器、缓存等）
	if listener.timeoutTimer != nil {
		listener.timeoutTimer.Stop() // 停止定时器
	}

	if listener.updateIntervalTimer != nil {
		listener.updateIntervalTimer.Stop() // 停止定时器
	}
}

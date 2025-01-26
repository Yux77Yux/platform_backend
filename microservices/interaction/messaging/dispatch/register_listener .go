package dispatch

import (
	"sync/atomic"
	"time"

	"google.golang.org/protobuf/reflect/protoreflect"

	generated "github.com/Yux77Yux/platform_backend/generated/user"
)

// 监听者结构体
type RegisterListener struct {
	exeChannel             chan *[]*generated.UserCredentials // 批量发送的通道
	userCredentialsChannel chan *generated.UserCredentials    // 用于接收的通道
	count                  uint32

	timeoutDuration     time.Duration     // 超时持续时间（触发销毁）
	timeoutTimer        *time.Timer       // 用于刷新存活时间
	updateInterval      time.Duration     // 批量插入的间隔时间
	updateIntervalTimer *time.Timer       // 用于周期性执行批量更新
	next                *RegisterListener // 下一个监听者
	prev                *RegisterListener // 上一个监听者
}

func (listener *RegisterListener) GetId() int64 {
	return 0
}

// 启动监听者
func (listener *RegisterListener) StartListening() {
	listener.RestartUpdateIntervalTimer()
	listener.RestartTimeoutTimer()
}

// 分发至通道
func (listener *RegisterListener) Dispatch(data protoreflect.ProtoMessage) {
	// 长度加1
	count := atomic.AddUint32(&listener.count, 1)

	userCredentials := data.(*generated.UserCredentials)
	listener.userCredentialsChannel <- userCredentials

	if count%MAX_BATCH_SIZE == 0 {
		go listener.SendBatch()
		listener.RestartUpdateIntervalTimer()
	}
}

// 做成·批量数据·然后发到别的通道
func (listener *RegisterListener) SendBatch() {
	const BatchSize = MAX_BATCH_SIZE

	count := atomic.LoadUint32(&listener.count)
	count = calculateBatchSize(count, BatchSize)
	if count == 0 {
		return
	}
	insertUserCredentialsPtr := insertUserCredentialsPool.Get().(*[]*generated.UserCredentials)
	*insertUserCredentialsPtr = (*insertUserCredentialsPtr)[:count]
	insertUserCredentials := *insertUserCredentialsPtr
	for i := uint32(0); i < count; i++ {
		insertUserCredentials[i] = <-listener.userCredentialsChannel
	}
	atomic.AddUint32(&listener.count, ^uint32(count-1)) //再减去

	listener.exeChannel <- insertUserCredentialsPtr // 送去批量执行,可能被阻塞
}

// 启动周期执行批量更新的定时器
func (listener *RegisterListener) RestartUpdateIntervalTimer() {
	if listener.updateIntervalTimer != nil {
		if !listener.updateIntervalTimer.Stop() {
			<-listener.updateIntervalTimer.C // 清理可能遗留的信号
		}
	}

	// 再执行
	listener.updateIntervalTimer = time.AfterFunc(listener.updateInterval, func() {
		count := atomic.LoadUint32(&listener.count)

		if count > 0 {
			go listener.SendBatch()        // 执行批量更新
			listener.RestartTimeoutTimer() // 重启定时器
		}
		listener.RestartUpdateIntervalTimer() // 重启定时器
	})
}

// 启动存活时间的定时器
func (listener *RegisterListener) RestartTimeoutTimer() {
	// 先重置
	if listener.timeoutTimer != nil {
		if !listener.timeoutTimer.Stop() {
			<-listener.timeoutTimer.C // 清理可能遗留的信号
		}
	}

	listener.timeoutTimer = time.AfterFunc(listener.timeoutDuration, func() {
		count := atomic.LoadUint32(&listener.count)

		if count == 0 {
			listener.Cleanup()
			// 超时后销毁监听者
			registerChain.DestroyListener(listener)
		} else {
			listener.RestartTimeoutTimer() // 重启定时器
		}
	})
}

// 清理监听者资源
func (listener *RegisterListener) Cleanup() {
	// 关闭通道
	close(listener.userCredentialsChannel)

	// 清理其他资源（例如定时器、缓存等）
	if listener.timeoutTimer != nil {
		listener.timeoutTimer.Stop() // 停止定时器
	}

	if listener.updateIntervalTimer != nil {
		listener.updateIntervalTimer.Stop() // 停止定时器
	}
}

package dispatch

import (
	"sync/atomic"
	"time"

	"google.golang.org/protobuf/reflect/protoreflect"

	generated "github.com/Yux77Yux/platform_backend/generated/user"
)

// 监听者结构体
type UserSpaceCacheListener struct {
	chain                  ChainInterface
	exeChannel             chan *[]*generated.UserUpdateSpace // 批量发送的通道
	userUpdateSpaceChannel chan *generated.UserUpdateSpace    // 用于接收的通道
	count                  uint32

	timeoutDuration     time.Duration           // 超时持续时间（触发销毁）
	timeoutTimer        *time.Timer             // 用于刷新存活时间
	updateInterval      time.Duration           // 批量插入的间隔时间
	updateIntervalTimer *time.Timer             // 用于周期性执行批量更新
	next                *UserSpaceCacheListener // 下一个监听者
	prev                *UserSpaceCacheListener // 上一个监听者
}

func (listener *UserSpaceCacheListener) GetId() int64 {
	return 0
}

// 启动监听者
func (listener *UserSpaceCacheListener) StartListening() {
	listener.userUpdateSpaceChannel = make(chan *generated.UserUpdateSpace, LISTENER_CHANNEL_COUNT)
	go listener.RestartUpdateIntervalTimer()
	go listener.RestartTimeoutTimer()
}

// 分发至通道
func (listener *UserSpaceCacheListener) Dispatch(data protoreflect.ProtoMessage) {
	// 长度加1
	count := atomic.AddUint32(&listener.count, 1)

	UserUpdateSpace := data.(*generated.UserUpdateSpace)
	listener.userUpdateSpaceChannel <- UserUpdateSpace

	if count%MAX_BATCH_SIZE == 0 {
		go listener.SendBatch()
		listener.RestartUpdateIntervalTimer()
	}
}

// 执行批量更新
func (listener *UserSpaceCacheListener) SendBatch() {
	const BatchSize = MAX_BATCH_SIZE

	count := atomic.LoadUint32(&listener.count)
	count = calculateBatchSize(count, BatchSize)
	if count == 0 {
		return
	}

	userUpdateSpacePtr := listener.chain.GetPoolObj().(*[]*generated.UserUpdateSpace)
	*userUpdateSpacePtr = (*userUpdateSpacePtr)[:count]
	userUpdateSpace := *userUpdateSpacePtr
	for i := uint32(0); i < count; i++ {
		userUpdateSpace[i] = <-listener.userUpdateSpaceChannel
	}
	atomic.AddUint32(&listener.count, ^uint32(count-1)) //再减去

	listener.exeChannel <- userUpdateSpacePtr // 送去批量执行,可能被阻塞
}

// 启动周期执行批量更新的定时器
func (listener *UserSpaceCacheListener) RestartUpdateIntervalTimer() {
	// 先重置
	if listener.updateIntervalTimer != nil {
		// 如果 timer 已存在，确保安全地重置
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
			go listener.SendBatch()        // 执行批量更新
			listener.RestartTimeoutTimer() // 重启定时器
		}
		listener.RestartUpdateIntervalTimer() // 重启定时器
	})
}

// 启动存活时间的定时器
func (listener *UserSpaceCacheListener) RestartTimeoutTimer() {
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
			listener.Cleanup()
			// 超时后销毁监听者
			listener.chain.DestroyListener(listener)
		} else {
			listener.RestartTimeoutTimer() // 重启定时器
		}
	})
}

// 清理监听者资源
func (listener *UserSpaceCacheListener) Cleanup() {
	// 关闭通道
	close(listener.userUpdateSpaceChannel)

	// 清理其他资源（例如定时器、缓存等）
	if listener.timeoutTimer != nil {
		listener.timeoutTimer.Stop() // 停止定时器
	}

	if listener.updateIntervalTimer != nil {
		listener.updateIntervalTimer.Stop() // 停止定时器
	}
}

package dispatch

import (
	"log"
	"sync/atomic"
	"time"

	"google.golang.org/protobuf/reflect/protoreflect"

	common "github.com/Yux77Yux/platform_backend/generated/common"
)

// 监听者结构体
type UpdateCountListener struct {
	exeChannel chan *ExeBody

	id int64

	viewCount int32
	likeCount int32
	saveCount int32

	timeoutDuration     time.Duration // 超时持续时间（触发销毁）
	timeoutTimer        *time.Timer   // 用于刷新存活时间
	updateInterval      time.Duration // 批量插入的间隔时间
	updateIntervalTimer *time.Timer   // 用于周期性执行批量更新
}

func (listener *UpdateCountListener) GetId() int64 {
	return atomic.LoadInt64(&listener.id)
}

// 启动监听者
func (listener *UpdateCountListener) StartListening() {
	go listener.RestartUpdateIntervalTimer()
	go listener.RestartTimeoutTimer()
}

func (listener *UpdateCountListener) Dispatch(data protoreflect.ProtoMessage) {
	_data := data.(*common.UserAction)

	operate := _data.GetOperate()

	switch operate {
	case common.Operate_VIEW:
		atomic.AddInt32(&listener.viewCount, 1)
	case common.Operate_LIKE:
		atomic.AddInt32(&listener.likeCount, 1)
	case common.Operate_COLLECT:
		atomic.AddInt32(&listener.saveCount, 1)

	case common.Operate_CANCEL_COLLECT:
		atomic.AddInt32(&listener.saveCount, -1)
	case common.Operate_CANCEL_LIKE:
		atomic.AddInt32(&listener.viewCount, -1)
	}
}

// 执行批量更新
func (listener *UpdateCountListener) SendBatch() {
	datasPtr := updatePool.Get().(*ExeBody)

	id := atomic.LoadInt64(&listener.id)
	saveCount := atomic.SwapInt32(&listener.saveCount, 0)
	likeCount := atomic.SwapInt32(&listener.likeCount, 0)
	viewCount := atomic.SwapInt32(&listener.viewCount, 0)

	log.Println("info:id")
	datasPtr.id = id
	log.Printf("info:saveCount %v", saveCount)
	datasPtr.newSaveCount = saveCount
	log.Printf("info:likeCount %v", likeCount)
	datasPtr.newLikeCount = likeCount
	log.Printf("info:viewCount %v", viewCount)
	datasPtr.newViewCount = viewCount

	listener.exeChannel <- datasPtr // 送去批量执行,可能被阻塞
}

// 启动周期执行批量更新的定时器
func (listener *UpdateCountListener) RestartUpdateIntervalTimer() {
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
		likeCount := atomic.LoadInt32(&listener.likeCount)
		saveCount := atomic.LoadInt32(&listener.saveCount)
		viewCount := atomic.LoadInt32(&listener.viewCount)

		if likeCount > 0 || saveCount > 0 || viewCount > 0 {
			go listener.SendBatch() // 执行批量更新
			listener.RestartTimeoutTimer()
		}
		listener.RestartUpdateIntervalTimer() // 重启定时器
	})
}

// 启动存活时间的定时器
func (listener *UpdateCountListener) RestartTimeoutTimer() {
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
		likeCount := atomic.LoadInt32(&listener.likeCount)
		saveCount := atomic.LoadInt32(&listener.saveCount)
		viewCount := atomic.LoadInt32(&listener.viewCount)

		if likeCount == 0 && saveCount == 0 && viewCount == 0 {
			// 超时后销毁监听者
			listener.Cleanup()
			updateCountChain.DestroyListener(listener)
		} else {
			listener.RestartTimeoutTimer() // 重启定时器
		}
	})
}

// 清理监听者资源
func (listener *UpdateCountListener) Cleanup() {
	// 清理其他资源（例如定时器、缓存等）
	if listener.timeoutTimer != nil {
		listener.timeoutTimer.Stop() // 停止定时器
	}

	if listener.updateIntervalTimer != nil {
		listener.updateIntervalTimer.Stop() // 停止定时器
	}
}

package messaging

import (
	"log"
	"time"

	generated "github.com/Yux77Yux/platform_backend/generated/comment"
	cache "github.com/Yux77Yux/platform_backend/microservices/comment/cache"
	db "github.com/Yux77Yux/platform_backend/microservices/comment/repository"
)

// 监听者结构体
type CommentListener struct {
	creationID          int64
	commentChannel      chan *generated.Comment // 用于接收评论的通道
	timeoutDuration     time.Duration           // 超时持续时间（触发销毁）
	timeoutTimer        *time.Timer             // 用于刷新存活时间
	updateInterval      time.Duration           // 批量插入的间隔时间
	updateIntervalTimer *time.Timer             // 用于周期性执行批量更新
	next                *CommentListener        // 下一个监听者
}

// 启动监听者
func (listener *CommentListener) StartListening() {
	listener.startProcessing()
	listener.startUpdateIntervalTimer()
	listener.startTimeoutTimer()
}

// 分发评论至通道
func (listener *CommentListener) DispatchComment(comment *generated.Comment) {
	// 处理评论的逻辑
	listener.commentChannel <- comment
	listener.refreshTimeoutTimer() // 刷新存活时间
}

// 抽取处理逻辑
func (listener *CommentListener) handleComment(comment *generated.Comment) {
	err := cache.SetTemporaryComments(comment)
	if err != nil {
		log.Printf("handleComment error %v", err)
	}
}

// 执行批量更新
func (listener *CommentListener) executeBatchInsert() {
	values, err := cache.GetTemporaryComments(listener.creationID)
	if err != nil {
		log.Printf("executeBatchInsert error :%v", err)
	}
	err = db.BatchInsert(values)
	if err != nil {
		log.Printf("batchInsert error :%v", err)
	}

	// 重置相关redis字段
	cache.RefreshTemporaryComment(listener.creationID)

	// 结束重置更新时间
	listener.updateIntervalTimer.Reset(listener.updateInterval)
}

// 具体处理逻辑
func (listener *CommentListener) startProcessing() {
	go func() {
		for comment := range listener.commentChannel {
			listener.handleComment(comment)
		}
	}()
}

// 启动周期执行批量更新的定时器
func (listener *CommentListener) startUpdateIntervalTimer() {
	listener.updateIntervalTimer = time.AfterFunc(listener.updateInterval, func() {
		listener.executeBatchInsert()       // 执行批量更新
		listener.startUpdateIntervalTimer() // 重启定时器
	})
}

// 启动存活时间的定时器
func (listener *CommentListener) startTimeoutTimer() {
	listener.timeoutTimer = time.AfterFunc(listener.timeoutDuration, func() {

		// 超时后销毁监听者
		chain.DestroyListener(listener)
	})
}

// 刷新存活时间的定时器
func (listener *CommentListener) refreshTimeoutTimer() {
	listener.timeoutTimer.Reset(listener.timeoutDuration)
}

// 清理监听者资源
func (listener *CommentListener) cleanup() {
	// 关闭评论通道
	close(listener.commentChannel)

	// 清理其他资源（例如定时器、缓存等）
	if listener.timeoutTimer != nil {
		listener.timeoutTimer.Stop() // 停止定时器
	}

	if listener.updateIntervalTimer != nil {
		listener.updateIntervalTimer.Stop() // 停止定时器
	}

	// 清空redis相关监听键
	err := cache.DelTemporaryComments(listener.creationID)
	if err != nil {
		log.Printf("delTemporaryComments error :%v", err)
	}
}

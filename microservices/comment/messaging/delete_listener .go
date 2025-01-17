package messaging

import (
	"log"
	"time"

	"google.golang.org/protobuf/proto"

	generated "github.com/Yux77Yux/platform_backend/generated/comment"
	cache "github.com/Yux77Yux/platform_backend/microservices/comment/cache"
	db "github.com/Yux77Yux/platform_backend/microservices/comment/repository"
)

// 监听者结构体
type DeleteListener struct {
	creationId          int64
	commentChannel      chan *generated.AfterAuth // 用于接收评论的通道
	count               int
	timeoutDuration     time.Duration   // 超时持续时间（触发销毁）
	timeoutTimer        *time.Timer     // 用于刷新存活时间
	updateInterval      time.Duration   // 批量插入的间隔时间
	updateIntervalTimer *time.Timer     // 用于周期性执行批量更新
	next                *DeleteListener // 下一个监听者
}

// 启动监听者
func (listener *DeleteListener) StartListening() {
	listener.startProcessing()
	listener.startUpdateIntervalTimer()
	listener.startTimeoutTimer()
}

func (listener *DeleteListener) Next() ListenerInterface {
	return listener.next
}

// 分发评论至通道
func (listener *DeleteListener) Dispatch(data []byte) {
	comment := new(generated.AfterAuth)
	err := proto.Unmarshal(data, comment)
	if err != nil {
		log.Printf("error: DispatchComment Unmarshal :%v", err)
	}
	// 处理评论的逻辑
	listener.commentChannel <- comment
}

// 抽取处理逻辑
func (listener *DeleteListener) handle(data []byte) {
	comment := new(generated.AfterAuth)
	err := proto.Unmarshal(data, comment)
	if err != nil {
		log.Printf("error: handle Unmarshal :%v", err)
	}

	err = cache.PushDeleteStatusComment(comment)
	if err != nil {
		log.Printf("handleComment error %v", err)
	}
	// 长度加1
	listener.count = listener.count + 1
}

// 执行批量更新
func (listener *DeleteListener) executeBatch() {
	values, err := cache.GetDeleteStatusComments(listener.creationId)
	if err != nil {
		log.Printf("executeBatchDelete GetDeleteStatusComments error :%v", err)
	}

	// 存进待永久删除部分
	err = cache.PushPreDeleteComments(values)
	if err != nil {
		log.Printf("executeBatchDelete PushPreDeleteComments error :%v", err)
	}

	count := len(values)
	comments := make([]*generated.AfterAuth, 0, count)
	for _, value := range values {
		comment := new(generated.AfterAuth)
		err := proto.Unmarshal([]byte(value), comment)
		if err != nil {
			log.Printf("error: executeBatchDelete Unmarshal :%v", err)
		}
		comments = append(comments, comment)
	}

	// 更新数据库
	err = db.BatchUpdate(comments)
	if err != nil {
		log.Printf("error: batchDelete error :%v", err)
	}

	// 丢弃已更新部分
	err = cache.RefreshDeleteStatusComments(listener.creationId, int64(count))
	if err != nil {
		log.Printf("executeBatchDelete RefreshTemporaryComment error %v", err)
	}

	// 重置时间
	listener.updateIntervalTimer.Reset(listener.updateInterval)

	// 去掉已完成部分
	listener.count = listener.count - count
}

// 具体处理逻辑
func (listener *DeleteListener) startProcessing() {
	go func() {
		for comment := range listener.commentChannel {
			msg, err := proto.Marshal(comment)
			if err != nil {
				log.Printf("error: startProcessing comment proto.Marshal error %v", err)
			}
			listener.handle(msg)
		}
	}()
}

// 启动周期执行批量更新的定时器
func (listener *DeleteListener) startUpdateIntervalTimer() {
	listener.updateIntervalTimer = time.AfterFunc(listener.updateInterval, func() {
		listener.executeBatch()             // 执行批量更新
		listener.startUpdateIntervalTimer() // 重启定时器
	})
}

// 启动存活时间的定时器
func (listener *DeleteListener) startTimeoutTimer() {
	listener.timeoutTimer = time.AfterFunc(listener.timeoutDuration, func() {
		if listener.count <= 0 {
			// 超时后销毁监听者
			chain.DestroyListener(listener)
		} else {
			listener.timeoutTimer.Reset(listener.timeoutDuration)
		}
	})
}

// 清理监听者资源
func (listener *DeleteListener) cleanup() {
	// 关闭评论通道
	close(listener.commentChannel)

	// 清理其他资源（例如定时器、缓存等）
	if listener.timeoutTimer != nil {
		listener.timeoutTimer.Stop() // 停止定时器
	}

	if listener.updateIntervalTimer != nil {
		listener.updateIntervalTimer.Stop() // 停止定时器
	}
}

// 删除信息的第二次过滤的监听者
type DispatchListener struct {
	id                  int16
	channel             chan *generated.AfterAuth // 用于接收评论的通道
	count               int
	timeoutDuration     time.Duration   // 超时持续时间（触发销毁）
	timeoutTimer        *time.Timer     // 用于刷新存活时间
	updateInterval      time.Duration   // 批量插入的间隔时间
	updateIntervalTimer *time.Timer     // 用于周期性执行批量更新
	next                *DeleteListener // 下一个监听者
}

package messaging

import (
	"log"
	"time"

	"google.golang.org/protobuf/proto"

	generated "github.com/Yux77Yux/platform_backend/generated/comment"
	cache "github.com/Yux77Yux/platform_backend/microservices/comment/cache"
	db "github.com/Yux77Yux/platform_backend/microservices/comment/repository"
)

type CommentInterface interface {
	StartListening()
	DispatchComment(comment *generated.Comment)
	handleComment(comment *generated.Comment)
	executeBatchInsert()
	startProcessing()
	startUpdateIntervalTimer()
	startTimeoutTimer()
	cleanup()
}

// 监听者结构体
type CommentListener struct {
	creationID          int64
	commentChannel      chan *generated.Comment // 用于接收评论的通道
	count               int
	timeoutDuration     time.Duration    // 超时持续时间（触发销毁）
	timeoutTimer        *time.Timer      // 用于刷新存活时间
	updateInterval      time.Duration    // 批量插入的间隔时间
	updateIntervalTimer *time.Timer      // 用于周期性执行批量更新
	next                *CommentListener // 下一个监听者
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
}

// 抽取处理逻辑
func (listener *CommentListener) handleComment(comment *generated.Comment) {
	err := cache.PushTemporaryComments(comment)
	if err != nil {
		log.Printf("handleComment error %v", err)
	}
	// 长度加1
	listener.count = listener.count + 1
}

// 执行批量插入
func (listener *CommentListener) executeBatchInsert() {
	values, err := cache.GetTemporaryComments(listener.creationID)
	if err != nil {
		log.Printf("executeBatchInsert GetTemporaryComments error :%v", err)
	}

	// 存进最后一次交互部分
	err = cache.ChangingTemporaryComments(listener.creationID, values)
	if err != nil {
		log.Printf("executeBatchInsert ChangingTemporaryComments error :%v", err)
	}

	count := len(values)
	comments := make([]*generated.Comment, 0, count)

	for _, commentStr := range values {
		comment := &generated.Comment{}
		err = proto.Unmarshal([]byte(commentStr), comment)
		if err != nil {
			log.Printf("error: comment proto.Unmarshal error %v", err)
		}

		comments = append(comments, comment)
	}

	err = db.BatchInsert(comments)
	if err != nil {
		log.Printf("error: batchInsert error :%v", err)
	}

	err = cache.RefreshTemporaryComment(listener.creationID, int64(count))
	if err != nil {
		log.Printf("executeBatchInsert RefreshTemporaryComment error %v", err)
	}

	err = cache.DelChangingTemporaryComments(listener.creationID)
	if err != nil {
		log.Printf("executeBatchInsert DelChangingTemporaryComments error %v", err)
	}

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
		if listener.count <= 0 {
			// 超时后销毁监听者
			chain.DestroyListener(listener)
		} else {
			listener.timeoutTimer.Reset(listener.timeoutDuration)
		}
	})
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
}

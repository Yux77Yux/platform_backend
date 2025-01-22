package messaging

import (
	"log"
	"sync"
	"time"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	generated "github.com/Yux77Yux/platform_backend/generated/comment"
	cache "github.com/Yux77Yux/platform_backend/microservices/comment/cache"
	db "github.com/Yux77Yux/platform_backend/microservices/comment/repository"
)

var insertChannel = make(chan []*generated.Comment, 3)

// 监听者结构体
type InsertListener struct {
	creationID          int64
	commentChannel      chan *generated.Comment // 用于接收评论的通道
	count               int
	mutex               sync.Mutex
	timeoutDuration     time.Duration   // 超时持续时间（触发销毁）
	timeoutTimer        *time.Timer     // 用于刷新存活时间
	updateInterval      time.Duration   // 批量插入的间隔时间
	updateIntervalTimer *time.Timer     // 用于周期性执行批量更新
	next                *InsertListener // 下一个监听者
}

// 启动监听者
func (listener *InsertListener) StartListening() {
	listener.startProcessing()
	listener.startUpdateIntervalTimer()
	listener.startTimeoutTimer()
}

// 分发评论至通道
func (listener *InsertListener) Dispatch(data protoreflect.ProtoMessage) {
	comment := data.(*generated.Comment)
	// 处理评论的逻辑
	listener.commentChannel <- comment

	// 获取锁
	listener.mutex.Lock()
	listener.count++
	if listener.count == 50 {
		go listener.executeBatch()
	}
	listener.mutex.Unlock()
}

// 抽取处理逻辑
func (listener *InsertListener) handle(data protoreflect.ProtoMessage) {
	comment := data.(*generated.Comment)

	// 长度加1
	listener.count = listener.count + 1
}

// 执行批量插入
func (listener *InsertListener) executeBatch() {
	values, err := cache.GetTemporaryComments(listener.creationID)
	if err != nil {
		log.Printf("executeBatchInsert GetTemporaryComments error :%v", err)
	}

	// 存进最后一次交互部分
	err = cache.PushChangingTemporaryComments(listener.creationID, values)
	if err != nil {
		log.Printf("executeBatchInsert ChangingTemporaryComments error :%v", err)
	}

	count := len(values)
	comments := make([]*generated.Comment, 0, count)

	for _, commentStr := range values {
		comment := &generated.Comment{}
		err = proto.Unmarshal(protoreflect.ProtoMessage(commentStr), comment)
		if err != nil {
			log.Printf("error: comment proto.Unmarshal error %v", err)
		}

		parent := comment.GetParent()
		if parent == 0 {
			// 是层主则通知作者消息
		} else {
			// 按parent，统一数量，通知对象
		}

		comments = append(comments, comment)
	}

	err = db.BatchInsert(comments)
	if err != nil {
		log.Printf("error: batchInsert error :%v", err)
	}

	err = cache.RefreshTemporaryComments(listener.creationID, int64(count))
	if err != nil {
		log.Printf("executeBatchInsert RefreshTemporaryComments error %v", err)
	}

	err = cache.RefreshDeleteStatusComments(listener.creationID, int64(count))
	if err != nil {
		log.Printf("executeBatchInsert DelChangingTemporaryComments error %v", err)
	}

	// 重置时间
	listener.updateIntervalTimer.Reset(listener.updateInterval)

	// 去掉已完成部分
	listener.count = listener.count - count
}

// 具体处理逻辑
func (listener *InsertListener) startProcessing() {
	go func() {
		for comment := range listener.commentChannel {
			msg, err := proto.Marshal(comment)
			if err != nil {
				log.Printf("error: comment proto.Marshal error %v", err)
			}
			listener.handle(msg)
		}
	}()
}

// 启动周期执行批量更新的定时器
func (listener *InsertListener) startUpdateIntervalTimer() {
	listener.updateIntervalTimer = time.AfterFunc(listener.updateInterval, func() {
		if listener.count > 0 {
			listener.executeBatch() // 执行批量更新
		}
		listener.startUpdateIntervalTimer() // 重启定时器
	})
}

// 启动存活时间的定时器
func (listener *InsertListener) startTimeoutTimer() {
	listener.timeoutTimer = time.AfterFunc(listener.timeoutDuration, func() {
		if listener.count <= 0 && len(listener.commentChannel) <= 0 {
			// 超时后销毁监听者
			insertChain.DestroyListener(listener)
		} else {
			listener.timeoutTimer.Reset(listener.timeoutDuration)
		}
	})
}

// 清理监听者资源
func (listener *InsertListener) cleanup() {
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

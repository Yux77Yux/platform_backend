package messaging

import (
	"log"
	"time"

	"google.golang.org/protobuf/proto"

	generated "github.com/Yux77Yux/platform_backend/generated/comment"
	cache "github.com/Yux77Yux/platform_backend/microservices/comment/cache"
	db "github.com/Yux77Yux/platform_backend/microservices/comment/repository"
)

var selectListener *SelectListener

func GetSelectListener() *SelectListener {
	if selectListener != nil {
		return selectListener
	}
	selectListener = &SelectListener{
		timeoutDuration: 15 * time.Second,
		updateInterval:  10 * time.Second,
		commentChannel:  make(chan *generated.AfterAuth, 400),
	}
	return selectListener
}

// 监听者结构体
type SelectListener struct {
	commentChannel      chan *generated.AfterAuth // 用于接收评论的通道
	count               int
	timeoutDuration     time.Duration // 超时持续时间（触发销毁）
	timeoutTimer        *time.Timer   // 用于刷新存活时间
	updateInterval      time.Duration // 批量插入的间隔时间
	updateIntervalTimer *time.Timer   // 用于周期性执行批量更新
}

// 启动监听者
func (listener *SelectListener) StartListening() {
	listener.startProcessing()
	listener.startUpdateIntervalTimer()
	listener.startTimeoutTimer()
}

func (listener *SelectListener) Next() ListenerInterface {
	return listener
}

// 分发评论至通道
func (listener *SelectListener) Dispatch(data []byte) {
	comment := new(generated.AfterAuth)
	err := proto.Unmarshal(data, comment)
	if err != nil {
		log.Printf("error: DispatchComment Unmarshal :%v", err)
	}
	// 处理评论的逻辑
	listener.commentChannel <- comment
}

// 抽取处理逻辑
func (listener *SelectListener) handle(data []byte) {
	comment := new(generated.AfterAuth)
	err := proto.Unmarshal(data, comment)
	if err != nil {
		log.Printf("error: handle Unmarshal :%v", err)
	}

	err = cache.PushSelectComment(comment)
	if err != nil {
		log.Printf("handleComment error %v", err)
	}
	// 长度加1
	listener.count = listener.count + 1
}

// 执行批量查询
func (listener *SelectListener) executeBatch() {
	values, err := cache.GetSelectComments()
	if err != nil {
		log.Printf("executeBatch GetSelectComments error :%v", err)
	}

	// 存进待永久删除部分
	err = cache.PushChangingSelectComments(values)
	if err != nil {
		log.Printf("executeBatch PushChangingSelectComments error :%v", err)
	}

	count := len(values)
	comments := make([]*generated.AfterAuth, 0, count)
	for _, value := range values {
		comment := new(generated.AfterAuth)
		err := proto.Unmarshal([]byte(value), comment)
		if err != nil {
			log.Printf("error: executeBatch Unmarshal :%v", err)
		}
		comments = append(comments, comment)
	}

	// 查询数据库
	result, err := db.GetCommentInfo(comments)
	if err != nil {
		log.Printf("error: executeBatch error :%v", err)
	}

	for _, comment := range result {
		data, err := proto.Marshal(comment)
		if err != nil {
			log.Printf("error: executeBatch comment proto.Marshal error %v", err)
		}
		deleteChain.HandleRequest(data)
	}

	// 丢弃已查询部分
	err = cache.RefreshSelectComments(int64(count))
	if err != nil {
		log.Printf("error: executeBatch RefreshTemporaryComment error %v", err)
	}
	err = cache.RefreshChangingSelectComments(int64(count))
	if err != nil {
		log.Printf("error: executeBatch RefreshTemporaryComment error %v", err)
	}

	// 重置时间
	listener.updateIntervalTimer.Reset(listener.updateInterval)

	// 去掉已完成部分
	listener.count = listener.count - count
}

// 具体处理逻辑
func (listener *SelectListener) startProcessing() {
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
func (listener *SelectListener) startUpdateIntervalTimer() {
	listener.updateIntervalTimer = time.AfterFunc(listener.updateInterval, func() {
		if listener.count > 0 {
			listener.executeBatch() // 执行批量更新
		}
		listener.startUpdateIntervalTimer() // 重启定时器
	})
}

// 启动存活时间的定时器
func (listener *SelectListener) startTimeoutTimer() {
	listener.timeoutTimer = time.AfterFunc(listener.timeoutDuration, func() {
		if listener.count <= 0 && len(listener.commentChannel) <= 0 {
			// 超时后销毁监听者
			listener.cleanup()
		} else {
			listener.timeoutTimer.Reset(listener.timeoutDuration)
		}
	})
}

// 清理监听者资源
func (listener *SelectListener) cleanup() {
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

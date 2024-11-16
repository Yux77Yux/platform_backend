package internal

import (
	"context"
	"log"

	"github.com/google/uuid"
)

type RequestHandlerFunc func(string) error

var UserRequestChan chan RequestHandlerFunc

func DispatchUserRequest(ctx context.Context) {
	// 创建带缓冲区的通道
	UserRequestChan = make(chan RequestHandlerFunc, 20) // 设置合适的缓冲区大小以适应并发请求

	// 启动一个 goroutine 来处理请求
	go handleRequests(ctx)

	// 等待上下文的取消信号，收到信号后关闭通道
	go func() {
		<-ctx.Done() // 等待 ctx 的 Done 信号
		log.Printf("info: received signal: %v. Shutting down gracefully...", ctx.Err())

		// 关闭通道
		close(UserRequestChan)
		log.Println("info: DispatchUserRequest stopped gracefully.")
	}()
}

// 处理请求的函数，使用 ctx 以支持优雅关闭
func handleRequests(ctx context.Context) {
	for {
		select {
		case handler, ok := <-UserRequestChan:
			// 如果通道已关闭且所有请求已处理完，则退出
			if !ok {
				log.Println("info: no more requests, exiting handleRequests.")
				return
			}

			reqId := uuid.New().String()
			if err := handler(reqId); err != nil {
				log.Printf("error: id:%s failed to handle : %v", reqId, err)
			}

		case <-ctx.Done():
			// 如果收到取消信号，退出处理循环
			log.Printf("info: context cancelled. Exiting request handling.")
			return
		}
	}
}

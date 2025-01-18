package internal

import (
	"log"
	"sync"

	"github.com/google/uuid"
)

type RequestDispatcher struct {
	requestChan chan RequestHandlerFunc
	wg          sync.WaitGroup
}

func InitDispatch() *RequestDispatcher {
	return &RequestDispatcher{
		// 闭包通道容纳个数
		requestChan: make(chan RequestHandlerFunc, 100),
	}
}

func (r *RequestDispatcher) Start() {
	go func() {
		for requestHandler := range r.requestChan {
			reqId := uuid.New().String()
			r.wg.Add(1)

			go func() {
				defer r.wg.Done()
				if err := requestHandler(reqId); err != nil {
					log.Printf("error: request_id: %s failed: %v", reqId, err)
				}
			}()
		}
		log.Println("info: RequestDispatcher stopped.")
	}()
}

// RegisterRequest 注册请求，供 internal 文件夹调用
func (r *RequestDispatcher) GetChannel() chan RequestHandlerFunc {
	return r.requestChan
}

// Shutdown 停止调度器
func (r *RequestDispatcher) Shutdown() {
	close(r.requestChan)
	r.wg.Wait() // 等待所有任务完成
	log.Println("info: RequestDispatcher shutdown gracefully.")
}

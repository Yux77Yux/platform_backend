package cache

import (
	"log"
	"sync"
)

type RequestHandlerFunc = func(CacheInterface)

type RequestProcessor struct {
	requestChan chan RequestHandlerFunc
	wg          sync.WaitGroup
}

func InitDispatch() *RequestProcessor {
	return &RequestProcessor{
		// 闭包通道容纳个数
		requestChan: make(chan RequestHandlerFunc, 100),
	}
}

func (r *RequestProcessor) Start() {
	go func() {
		for handler := range r.requestChan {
			log.Println("cache request input")
			r.wg.Add(1)

			go func() {
				defer r.wg.Done()
				handler(CacheClient)
			}()
		}
		log.Println("info: RequestProcessor stopped.")
	}()
}

// Shutdown 停止调度器
func (r *RequestProcessor) Shutdown() {
	close(r.requestChan)
	r.wg.Wait() // 等待所有任务完成
	log.Println("info: RequestProcessor shutdown gracefully.")
}

func (r *RequestProcessor) GetChannel() chan RequestHandlerFunc {
	return r.requestChan
}

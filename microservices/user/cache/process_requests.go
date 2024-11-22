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
		requestChan: make(chan RequestHandlerFunc, 20),
	}
}

func (r *RequestProcessor) Start() {
	go func() {
		for handler := range r.requestChan {
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

package cache

import (
	"log"

	pkgCache "github.com/Yux77Yux/platform_backend/pkg/cache"
)

var (
	addr                string
	password            string
	cacheRequestChannel chan RequestHandlerFunc
	CacheClient         CacheInterface
)

func InitStr(Addr, Password string) {
	addr, password = Addr, Password
}

func GetCacheClient() CacheInterface {
	var cache CacheInterface = &pkgCache.RedisClient{}
	err := cache.Open(addr, password)
	if err != nil {
		log.Printf("error: failed to connect the cache client: %v", err)
		return nil
	}

	return cache
}

func CloseClient() {
	log.Println("info: cache client gracefully.")
	if err := CacheClient.Close(); err != nil {
		log.Println("error: cache client close error.")
		return
	}
	log.Println("info: cache client close ok.")
}

func InitWorker(master *RequestProcessor) {
	cacheRequestChannel = master.GetChannel()
}

func Init() {
	CacheClient = GetCacheClient()
}

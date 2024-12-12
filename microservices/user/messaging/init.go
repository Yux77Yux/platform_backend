package messaging

import (
	"fmt"
	"log"

	pkgMQ "github.com/Yux77Yux/platform_backend/pkg/messagequeue"
)

var (
	connStr         string
	ExchangesConfig = map[string]string{
		"register":         "direct",
		"storeUserInCache": "direct",
		// Add more exchanges here
	}
	ListenRPCs = []string{
		"agg_user",
	}
)

func InitStr(_str string) {
	connStr = _str
}

func GetRabbitMQ() MessageQueueInterface {
	var messageQueue MessageQueueInterface = &pkgMQ.RabbitMQClass{}
	if err := messageQueue.Open(connStr); err != nil {
		wiredErr := fmt.Errorf("failed to connect the rabbit client: %w", err)
		log.Printf("error: %v", wiredErr)
		return nil
	}

	return messageQueue
}

// 非RPC类型的消息队列的交换机声明
func Init() {
	rabbitMQ := GetRabbitMQ()
	defer rabbitMQ.Close()

	if rabbitMQ == nil {
		log.Printf("error: message queue open failed")
		return
	}
	for exchange, kind := range ExchangesConfig {
		if err := rabbitMQ.ExchangeDeclare(exchange, kind, true, false, false, false, nil); err != nil {
			wiredErr := fmt.Errorf("failed to declare exchange %s : %w", exchange, err)
			log.Printf("error: %v", wiredErr)
		}

		switch exchange {
		// 不同的exchange使用不同函数
		case "register":
			go ListenToQueue(exchange, "register", "register", registerProcessor)
		case "storeUserInCache":
			go ListenToQueue(exchange, "storeUserInCache", "storeUserInCache", storeUserInCache)
		case "agg_user":
			go ListenToQueue(exchange, "getUser", "getUser", storeUserInCache)
		}
	}

	for _, exchange := range ListenRPCs {
		switch exchange {
		// 不同的exchange使用不同函数
		case "agg_user":
			go ListenRPC(exchange, "getUser", "getUser", getUserProcessor)
		}
	}
}

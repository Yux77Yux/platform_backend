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
		"storeUser":        "direct",
		"storeCredentials": "direct",
		"updateUserSpace":  "direct",
		"delReviewer":      "direct",

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
		case "storeUser":
			go ListenToQueue(exchange, "storeUser", "storeUser", storeUserProcessor)
		case "storeCredentials":
			go ListenToQueue(exchange, "storeCredentials", "storeCredentials", storeCredentialsProcessor)
		case "updateUserSpace":
			go ListenToQueue(exchange, "updateUserSpace", "updateUserSpace", updateUserSpaceProcessor)
		case "delReviewer":
			go ListenToQueue(exchange, "delReviewer", "delReviewer", delReviewerProcessor)
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

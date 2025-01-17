package messaging

// 由于不同的exchange，需要不同的接收者，事实上需要被调度，统一开关

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/protobuf/proto"

	generated "github.com/Yux77Yux/platform_backend/generated/comment"
	db "github.com/Yux77Yux/platform_backend/microservices/comment/repository"
)

func JoinCommentProcessor(msg amqp.Delivery) error {
	// 传递至责任链
	insertChain.HandleRequest(msg.Body)
	return nil
}

func DeleteCommentProcessor(msg amqp.Delivery) error {
	req := new(generated.AfterAuth)
	// 反序列化
	err := proto.Unmarshal(msg.Body, req)
	if err != nil {
		log.Printf("error: DeleteCommentProcessor unmarshaling message: %v", err)
		return fmt.Errorf("deleteCommentProcessor processor error: %w", err)
	}

	// 开始第二次过滤，验证评论发布者与token的是否匹配
	// 这里做一个监听者，将收到的请求拉出来发到redis，
	// 一段时间从redis取出，然后批量查询返回信息

	// 发送集中处理
	deleteChain.HandleRequest(msg.Body)

	return nil
}

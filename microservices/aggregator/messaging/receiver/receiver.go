package receiver

// 由于不同的exchange，需要不同的接收者，事实上需要被调度，统一开关

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/protobuf/proto"

	common "github.com/Yux77Yux/platform_backend/generated/common"
	cache "github.com/Yux77Yux/platform_backend/microservices/aggregator/cache"
	dispatch "github.com/Yux77Yux/platform_backend/microservices/aggregator/messaging/dispatch"
)

func addViewProcessor(msg amqp.Delivery) error {
	req := new(common.ViewCreation)
	err := proto.Unmarshal(msg.Body, req)
	if err != nil {
		log.Printf("error: proto.Unmarshal %v", err)
		return err
	}

	id := req.GetId()
	ip := req.GetIpv4()
	if id <= 0 {
		return fmt.Errorf("error: creationId not exist")
	}
	if ip == "" {
		return fmt.Errorf("error: ip not exist")
	}

	exist, err := cache.ExistIpInSet(req)
	if err != nil {
		err = fmt.Errorf("error: ExistIpInSet %w", err)
		log.Printf("%v", err)
		return err
	}
	if exist {
		return nil
	} else {
		err = cache.AddIpInSet(req)
		if err != nil {
			err = fmt.Errorf("error: AddIpInSet %w", err)
			log.Printf("%v", err)
			return err
		}
	}

	action := &common.UserAction{
		Id: &common.CreationId{
			Id: id,
		},
		ActionTag: 1,
	}
	go dispatch.HandleRequest(action, dispatch.AddView)

	return nil
}

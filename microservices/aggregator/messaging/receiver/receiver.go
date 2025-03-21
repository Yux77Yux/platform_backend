package receiver

// 由于不同的exchange，需要不同的接收者，事实上需要被调度，统一开关

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/protobuf/types/known/anypb"

	common "github.com/Yux77Yux/platform_backend/generated/common"
	"github.com/Yux77Yux/platform_backend/microservices/aggregator/messaging/dispatch"
)

func increaseViewProcessor(ctx context.Context, msg *anypb.Any) error {
	req := new(common.ViewCreation)
	err := msg.UnmarshalTo(req)
	if err != nil {
		log.Printf("error: anypb.Unmarshal %v", err)
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

	exist, err := cache.ExistIpInSet(ctx, req)
	if err != nil {
		err = fmt.Errorf("error: ExistIpInSet %w", err)
		log.Printf("%v", err)
		return err
	}
	if exist {
		return nil
	} else {
		err = cache.AddIpInSet(ctx, req)
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
		Operate: common.Operate_VIEW,
	}
	go dispatcher.HandleRequest(action, dispatch.IncreaseView)

	return nil
}

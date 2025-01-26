package internal

import (
	"context"
	"fmt"
	"log"

	common "github.com/Yux77Yux/platform_backend/generated/common"
	generated "github.com/Yux77Yux/platform_backend/generated/interaction"
	cache "github.com/Yux77Yux/platform_backend/microservices/interaction/cache"
	mq "github.com/Yux77Yux/platform_backend/microservices/interaction/messaging"
	db "github.com/Yux77Yux/platform_backend/microservices/interaction/repository"
	tools "github.com/Yux77Yux/platform_backend/microservices/interaction/tools"
)

func GetActionTag(ctx context.Context, req *generated.GetCreationInteractionRequest) (*generated.GetCreationInteractionResponse, error) {
	response := &generated.GetCreationInteractionResponse{
		Msg: &common.ApiResponse{
			Status: common.ApiResponse_SUCCESS,
			Code:   "202",
		},
	}
	return response, nil
}

func GetCollections(ctx context.Context, req *generated.GetCollectionsRequest) (*generated.GetInteractionsResponse, error) {
	response := &generated.GetInteractionsResponse{
		Msg: &common.ApiResponse{
			Status: common.ApiResponse_SUCCESS,
			Code:   "202",
		},
	}
	return response, nil
}

func GetHistories(ctx context.Context, req *generated.GetHistoriesRequest) (*generated.GetInteractionsResponse, error) {
	response := &generated.GetInteractionsResponse{
		Msg: &common.ApiResponse{
			Status: common.ApiResponse_SUCCESS,
			Code:   "202",
		},
	}
	return response, nil
}

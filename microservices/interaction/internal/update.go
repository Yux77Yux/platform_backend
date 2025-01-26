package internal

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/protobuf/types/known/timestamppb"

	common "github.com/Yux77Yux/platform_backend/generated/common"
	generated "github.com/Yux77Yux/platform_backend/generated/interaction"
	mq "github.com/Yux77Yux/platform_backend/microservices/interaction/messaging"
	dispatch "github.com/Yux77Yux/platform_backend/microservices/interaction/messaging/dispatch"
	oss "github.com/Yux77Yux/platform_backend/microservices/interaction/oss"
	jwt "github.com/Yux77Yux/platform_backend/pkg/jwt"
)

func CancelCollections(ctx context.Context, req *generated.UpdateInteractionsRequest) (*generated.UpdateInteractionResponse, error) {
	response := &generated.UpdateInteractionResponse{
		Msg: &common.ApiResponse{
			Status: common.ApiResponse_SUCCESS,
			Code:   "202",
		},
	}
	return response, nil
}

func CancelLike(ctx context.Context, req *generated.UpdateInteractionRequest) (*generated.UpdateInteractionResponse, error) {
	response := &generated.UpdateInteractionResponse{
		Msg: &common.ApiResponse{
			Status: common.ApiResponse_SUCCESS,
			Code:   "202",
		},
	}
	return response, nil
}

func ClickCollection(ctx context.Context, req *generated.UpdateInteractionRequest) (*generated.UpdateInteractionResponse, error) {
	response := &generated.UpdateInteractionResponse{
		Msg: &common.ApiResponse{
			Status: common.ApiResponse_SUCCESS,
			Code:   "202",
		},
	}
	return response, nil
}

func ClickLike(ctx context.Context, req *generated.UpdateInteractionRequest) (*generated.UpdateInteractionResponse, error) {
	response := &generated.UpdateInteractionResponse{
		Msg: &common.ApiResponse{
			Status: common.ApiResponse_SUCCESS,
			Code:   "202",
		},
	}
	return response, nil
}

func DelHistories(ctx context.Context, req *generated.UpdateInteractionRequest) (*generated.UpdateInteractionResponse, error) {
	response := &generated.UpdateInteractionResponse{
		Msg: &common.ApiResponse{
			Status: common.ApiResponse_SUCCESS,
			Code:   "202",
		},
	}
	return response, nil
}

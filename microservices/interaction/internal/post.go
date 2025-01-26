package internal

import (
	"context"
	"log"

	"github.com/Yux77Yux/platform_backend/generated/common"
	generated "github.com/Yux77Yux/platform_backend/generated/interaction"
	internal "github.com/Yux77Yux/platform_backend/microservices/interaction/internal"
)

func PostInteraction(ctx context.Context, req *generated.PostInteractionRequest) (*generated.PostInteractionResponse, error) {
	response := &generated.PostInteractionResponse{
		Msg: &common.ApiResponse{
			Status: common.ApiResponse_SUCCESS,
			Code:   "202",
		},
	}
	return response, nil
}

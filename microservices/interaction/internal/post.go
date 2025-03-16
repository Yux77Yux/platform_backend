package internal

import (
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"

	common "github.com/Yux77Yux/platform_backend/generated/common"
	generated "github.com/Yux77Yux/platform_backend/generated/interaction"
	tools "github.com/Yux77Yux/platform_backend/microservices/interaction/tools"
	auth "github.com/Yux77Yux/platform_backend/pkg/auth"
)

func PostInteraction(ctx context.Context, req *generated.PostInteractionRequest) (*generated.PostInteractionResponse, error) {
	var response = new(generated.PostInteractionResponse)
	token := req.GetAccessToken().GetValue()
	pass, userId, err := auth.Auth("post", "interaction", token)
	if err != nil {
		return &generated.PostInteractionResponse{
			Msg: &common.ApiResponse{
				Status: common.ApiResponse_FAILED,
				Code:   "500",
			},
		}, err
	}
	if !pass {
		return &generated.PostInteractionResponse{
			Msg: &common.ApiResponse{
				Status: common.ApiResponse_ERROR,
				Code:   "403",
			},
		}, nil
	}

	operateInteraction := &generated.OperateInteraction{
		Base: &generated.BaseInteraction{
			UserId:     userId,
			CreationId: req.GetBase().GetCreationId(),
		},
		Action:    common.Operate_VIEW,
		UpdatedAt: timestamppb.Now(),
	}

	go func(operateInteraction *generated.OperateInteraction, ctx context.Context) {
		traceId, fullName := tools.GetMetadataValue(ctx, "trace-id"), tools.GetMetadataValue(ctx, "full-name")
		err = messaging.SendMessage(ctx, messaging.AddView, messaging.AddView, operateInteraction)
		if err != nil {
			tools.LogError(traceId, fullName, err)
		}
	}(operateInteraction, ctx)

	response.Msg = &common.ApiResponse{
		Status: common.ApiResponse_SUCCESS,
		Code:   "202",
	}
	return response, nil
}

package internal

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	common "github.com/Yux77Yux/platform_backend/generated/common"
	generated "github.com/Yux77Yux/platform_backend/generated/interaction"
	dispatch "github.com/Yux77Yux/platform_backend/microservices/interaction/messaging/dispatch"
	auth "github.com/Yux77Yux/platform_backend/pkg/auth"
)

func PostInteraction(req *generated.PostInteractionRequest) (*generated.PostInteractionResponse, error) {
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

	interaciton := &generated.Interaction{
		Base: &generated.BaseInteraction{
			UserId:     userId,
			CreationId: req.GetBase().GetCreationId(),
		},
		ActionTag: int32(generated.Operate_VIEW),
		UpdatedAt: timestamppb.Now(),
	}
	go dispatch.HandleRequest(interaciton, dispatch.DbInteraction)
	go dispatch.HandleRequest(interaciton, dispatch.ViewCache)

	response.Msg = &common.ApiResponse{
		Status: common.ApiResponse_SUCCESS,
		Code:   "202",
	}
	return response, nil
}

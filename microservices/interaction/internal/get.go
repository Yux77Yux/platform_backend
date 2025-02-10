package internal

import (
	"context"

	common "github.com/Yux77Yux/platform_backend/generated/common"
	generated "github.com/Yux77Yux/platform_backend/generated/interaction"
	cache "github.com/Yux77Yux/platform_backend/microservices/interaction/cache"

	// mq "github.com/Yux77Yux/platform_backend/microservices/interaction/messaging"
	recommend "github.com/Yux77Yux/platform_backend/microservices/interaction/recommend"
	db "github.com/Yux77Yux/platform_backend/microservices/interaction/repository"
	auth "github.com/Yux77Yux/platform_backend/pkg/auth"
)

func GetActionTag(ctx context.Context, req *generated.GetCreationInteractionRequest) (*generated.GetCreationInteractionResponse, error) {
	var response = new(generated.GetCreationInteractionResponse)
	base := req.GetBase()
	interaction, err := cache.GetInteraction(ctx, base)
	if err != nil {
		response.Interaction = interaction
		response.Msg = &common.ApiResponse{
			Status: common.ApiResponse_ERROR,
			Code:   "500",
		}
		interaction, err := db.GetActionTag(ctx, base)
		if err != nil {
			return response, nil
		}

		response.Interaction = interaction
		// 补充Redis
		// action:=interaction.GetActionTag()
		// if action & 2 ==2{

		// }
		// if action & 4 == 4{

		// }
	}

	response.Msg = &common.ApiResponse{
		Status: common.ApiResponse_SUCCESS,
		Code:   "200",
	}
	return response, nil
}

func GetCollections(ctx context.Context, req *generated.GetCollectionsRequest) (*generated.GetInteractionsResponse, error) {
	var response = new(generated.GetInteractionsResponse)
	pageNum := req.GetPage()
	token := req.GetAccessToken().GetValue()
	pass, userId, err := auth.Auth("get", "interaction", token)
	if err != nil {
		return &generated.GetInteractionsResponse{
			Msg: &common.ApiResponse{
				Status: common.ApiResponse_FAILED,
				Code:   "500",
			},
		}, err
	}
	if !pass {
		return &generated.GetInteractionsResponse{
			Msg: &common.ApiResponse{
				Status: common.ApiResponse_ERROR,
				Code:   "403",
			},
		}, nil
	}

	interactions, err := cache.GetCollections(ctx, userId, pageNum)
	if err != nil {
		response.Msg = &common.ApiResponse{
			Status: common.ApiResponse_ERROR,
			Code:   "500",
		}
		interactions, err = db.GetCollections(ctx, userId, pageNum)

		if err != nil {
			return response, nil
		}

		// 补充Redis
	}
	response.AnyInteraction = &generated.AnyInteraction{
		AnyInterction: interactions,
	}

	response.Msg = &common.ApiResponse{
		Status: common.ApiResponse_SUCCESS,
		Code:   "200",
	}
	return response, nil
}

func GetHistories(ctx context.Context, req *generated.GetHistoriesRequest) (*generated.GetInteractionsResponse, error) {
	var response = new(generated.GetInteractionsResponse)
	pageNum := req.GetPage()
	token := req.GetAccessToken().GetValue()
	pass, userId, err := auth.Auth("get", "interaction", token)
	if err != nil {
		return &generated.GetInteractionsResponse{
			Msg: &common.ApiResponse{
				Status: common.ApiResponse_FAILED,
				Code:   "500",
			},
		}, err
	}
	if !pass {
		return &generated.GetInteractionsResponse{
			Msg: &common.ApiResponse{
				Status: common.ApiResponse_ERROR,
				Code:   "403",
			},
		}, nil
	}

	interactions, err := cache.GetHistories(ctx, userId, pageNum)
	if err != nil {
		response.Msg = &common.ApiResponse{
			Status: common.ApiResponse_ERROR,
			Code:   "500",
		}
		interactions, err = db.GetHistories(ctx, userId, pageNum)

		if err != nil {
			return response, nil
		}
	}
	response.AnyInteraction = &generated.AnyInteraction{
		AnyInterction: interactions,
	}

	response.Msg = &common.ApiResponse{
		Status: common.ApiResponse_SUCCESS,
		Code:   "200",
	}
	return response, nil
}

func GetRecommend(ctx context.Context, req *generated.GetRecommendRequest) (*generated.GetRecommendResponse, error) {
	var response = new(generated.GetRecommendResponse)
	token := req.GetAccessToken().GetValue()
	pass, userId, err := auth.Auth("get", "interaction", token)
	if err != nil {
		return &generated.GetRecommendResponse{
			Msg: &common.ApiResponse{
				Status: common.ApiResponse_FAILED,
				Code:   "500",
			},
		}, err
	}
	if !pass {
		return &generated.GetRecommendResponse{
			Msg: &common.ApiResponse{
				Status: common.ApiResponse_ERROR,
				Code:   "403",
			},
		}, nil
	}

	interactions, err := recommend.Recommend(userId)
	if err != nil {
		response.Msg = &common.ApiResponse{
			Status: common.ApiResponse_ERROR,
			Code:   "500",
		}
	}
	response.Creations = interactions

	response.Msg = &common.ApiResponse{
		Status: common.ApiResponse_SUCCESS,
		Code:   "200",
	}
	return response, nil
}

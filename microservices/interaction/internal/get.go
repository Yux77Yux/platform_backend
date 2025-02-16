package internal

import (
	"context"
	"log"

	common "github.com/Yux77Yux/platform_backend/generated/common"
	generated "github.com/Yux77Yux/platform_backend/generated/interaction"
	cache "github.com/Yux77Yux/platform_backend/microservices/interaction/cache"
	messaging "github.com/Yux77Yux/platform_backend/microservices/interaction/messaging"
	auth "github.com/Yux77Yux/platform_backend/pkg/auth"

	db "github.com/Yux77Yux/platform_backend/microservices/interaction/repository"
)

func GetActionTag(ctx context.Context, req *generated.GetCreationInteractionRequest) (*generated.GetCreationInteractionResponse, error) {
	var response = new(generated.GetCreationInteractionResponse)
	token := req.GetAccessToken().GetValue()
	pass, userId, err := auth.Auth("get", "interaction", token)
	if err != nil {
		response.Msg = &common.ApiResponse{
			Code:    "500",
			Status:  common.ApiResponse_ERROR,
			Details: err.Error(),
		}
		return response, err
	}
	if !pass {
		response.Msg = &common.ApiResponse{
			Code:   "403",
			Status: common.ApiResponse_ERROR,
		}
		return response, nil
	}

	base := req.GetBase()
	base.UserId = userId
	interaction, err := cache.GetInteraction(ctx, base)
	if err != nil {
		response.Msg = &common.ApiResponse{
			Status: common.ApiResponse_ERROR,
			Code:   "500",
		}
		interaction, err = db.GetActionTag(ctx, base)
		if err != nil {
			return response, nil
		}
	}

	response.Interaction = interaction
	response.Msg = &common.ApiResponse{
		Status: common.ApiResponse_SUCCESS,
		Code:   "200",
	}
	return response, nil
}

func GetCollections(ctx context.Context, req *generated.GetCollectionsRequest) (*generated.GetInteractionsResponse, error) {
	var response = new(generated.GetInteractionsResponse)
	pageNum := req.GetPage()
	userId := req.GetUserId()

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
	userId := req.GetUserId()

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

func GetRecommendBaseUser(ctx context.Context, req *generated.GetRecommendRequest) (*generated.GetRecommendResponse, error) {
	var response = new(generated.GetRecommendResponse)

	userId := req.GetId()
	interactions, count, err := cache.GetRecommendBaseUser(ctx, userId)
	if err != nil {
		response.Msg = &common.ApiResponse{
			Status: common.ApiResponse_ERROR,
			Code:   "500",
		}
		return response, err
	}

	go func() {
		if count <= 17 {
			err = messaging.SendMessage(messaging.ComputeUser, messaging.ComputeUser, &common.UserDefault{
				UserId: userId,
			})
			if err != nil {
				log.Printf("error:GetRecommendBaseUser SendMessage %v", err)
			}
		}
	}()

	response.Creations = interactions
	response.Msg = &common.ApiResponse{
		Status: common.ApiResponse_SUCCESS,
		Code:   "200",
	}
	return response, nil
}

func GetRecommendBaseCreation(ctx context.Context, req *generated.GetRecommendRequest) (*generated.GetRecommendResponse, error) {
	var response = new(generated.GetRecommendResponse)

	id := req.GetId()
	creations, reset, err := cache.GetRecommendBaseItem(ctx, id)
	if err != nil {
		response.Msg = &common.ApiResponse{
			Status: common.ApiResponse_ERROR,
			Code:   "500",
		}
		return response, err
	}

	go func() {
		if reset {
			err = messaging.SendMessage(messaging.ComputeSimilarCreation, messaging.ComputeSimilarCreation, &common.CreationId{
				Id: id,
			})
			if err != nil {
				log.Printf("error:GetRecommendBaseItem SendMessage %v", err)
			}
		}
	}()

	response.Creations = creations
	response.Msg = &common.ApiResponse{
		Status: common.ApiResponse_SUCCESS,
		Code:   "200",
	}
	return response, nil
}

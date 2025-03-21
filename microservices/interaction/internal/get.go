package internal

import (
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"

	common "github.com/Yux77Yux/platform_backend/generated/common"
	generated "github.com/Yux77Yux/platform_backend/generated/interaction"
	tools "github.com/Yux77Yux/platform_backend/microservices/interaction/tools"
	auth "github.com/Yux77Yux/platform_backend/pkg/auth"
	errMap "github.com/Yux77Yux/platform_backend/pkg/error"
)

func GetActionTag(ctx context.Context, req *generated.GetCreationInteractionRequest) (*generated.GetCreationInteractionResponse, error) {
	response := new(generated.GetCreationInteractionResponse)
	token := req.GetAccessToken().GetValue()
	pass, userId, err := auth.Auth("update", "interaction", token)
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
			Code:    "403",
			Status:  common.ApiResponse_ERROR,
			Details: "no pass",
		}
		return response, nil
	}

	base := req.GetBase()
	base.UserId = userId

	newInteraction := &generated.OperateInteraction{
		Base:      base,
		UpdatedAt: timestamppb.Now(),
		Action:    common.Operate_VIEW,
	}

	go func(newInteraction *generated.OperateInteraction, ctx context.Context) {
		traceId, fullName := tools.GetMetadataValue(ctx, "trace-id"), tools.GetMetadataValue(ctx, "full-name")
		err = messaging.SendMessage(ctx, EXCHANGE_ADD_VIEW, KEY_ADD_VIEW, newInteraction)
		if err != nil {
			tools.LogError(traceId, fullName, err)
		}
	}(newInteraction, ctx)

	interaction, err := cache.GetInteraction(ctx, base)
	if err != nil {
		interaction, err = db.GetActionTag(ctx, base)
		if err != nil {
			if errMap.IsServerError(err) {
				response.Msg = &common.ApiResponse{
					Status:  common.ApiResponse_ERROR,
					Code:    errMap.GrpcCodeToHTTPStatusString(err),
					Details: err.Error(),
				}
				return response, err
			}
			response.Msg = &common.ApiResponse{
				Status:  common.ApiResponse_ERROR,
				Code:    errMap.GrpcCodeToHTTPStatusString(err),
				Details: err.Error(),
			}
			return response, nil
		}
	}

	response.ActionTag = interaction.GetActionTag()
	response.Msg = &common.ApiResponse{
		Status: common.ApiResponse_SUCCESS,
		Code:   "200",
	}
	return response, nil
}

func GetCollections(ctx context.Context, req *generated.GetCollectionsRequest) (*generated.GetInteractionsResponse, error) {
	response := new(generated.GetInteractionsResponse)
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
			if errMap.IsServerError(err) {
				response.Msg = &common.ApiResponse{
					Status:  common.ApiResponse_ERROR,
					Code:    errMap.GrpcCodeToHTTPStatusString(err),
					Details: err.Error(),
				}
				return response, err
			}
			response.Msg = &common.ApiResponse{
				Status:  common.ApiResponse_ERROR,
				Code:    errMap.GrpcCodeToHTTPStatusString(err),
				Details: err.Error(),
			}
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
	response := new(generated.GetInteractionsResponse)
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
			if errMap.IsServerError(err) {
				response.Msg = &common.ApiResponse{
					Status:  common.ApiResponse_ERROR,
					Code:    errMap.GrpcCodeToHTTPStatusString(err),
					Details: err.Error(),
				}
				return response, err
			}
			response.Msg = &common.ApiResponse{
				Status:  common.ApiResponse_ERROR,
				Code:    errMap.GrpcCodeToHTTPStatusString(err),
				Details: err.Error(),
			}
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
	response := new(generated.GetRecommendResponse)

	userId := req.GetId()
	interactions, count, err := cache.GetRecommendBaseUser(ctx, userId)
	if err != nil {
		if errMap.IsServerError(err) {
			response.Msg = &common.ApiResponse{
				Status:  common.ApiResponse_ERROR,
				Code:    errMap.GrpcCodeToHTTPStatusString(err),
				Details: err.Error(),
			}
			return response, err
		}
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    errMap.GrpcCodeToHTTPStatusString(err),
			Details: err.Error(),
		}
		return response, nil
	}

	go func(count, userId int64, ctx context.Context) {
		if count <= 17 {
			traceId, fullName := tools.GetMetadataValue(ctx, "trace-id"), tools.GetMetadataValue(ctx, "full-name")
			err = messaging.SendMessage(ctx, EXCHANGE_COMPUTE_USER, KEY_COMPUTE_USER, &common.UserDefault{
				UserId: userId,
			})
			if err != nil {
				tools.LogError(traceId, fullName, err)
			}
		}
	}(count, userId, ctx)

	response.Creations = interactions
	response.Msg = &common.ApiResponse{
		Status: common.ApiResponse_SUCCESS,
		Code:   "200",
	}
	return response, nil
}

func GetRecommendBaseCreation(ctx context.Context, req *generated.GetRecommendRequest) (*generated.GetRecommendResponse, error) {
	response := new(generated.GetRecommendResponse)

	id := req.GetId()
	creations, reset, err := cache.GetRecommendBaseItem(ctx, id)
	if err != nil {
		response.Msg = &common.ApiResponse{
			Status: common.ApiResponse_ERROR,
			Code:   "500",
		}
		return response, err
	}

	go func(reset bool, id int64, ctx context.Context) {
		if !reset {
			return
		}
		traceId, fullName := tools.GetMetadataValue(ctx, "trace-id"), tools.GetMetadataValue(ctx, "full-name")
		err = messaging.SendMessage(ctx, EXCHANGE_COMPUTE_CREATION, KEY_COMPUTE_CREATION, &common.CreationId{
			Id: id,
		})
		if err != nil {
			tools.LogError(traceId, fullName, err)
		}
	}(reset, id, ctx)

	response.Creations = creations
	response.Msg = &common.ApiResponse{
		Status: common.ApiResponse_SUCCESS,
		Code:   "200",
	}
	return response, nil
}

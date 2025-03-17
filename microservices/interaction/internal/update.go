package internal

import (
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"

	common "github.com/Yux77Yux/platform_backend/generated/common"
	generated "github.com/Yux77Yux/platform_backend/generated/interaction"
	tools "github.com/Yux77Yux/platform_backend/microservices/interaction/tools"
	auth "github.com/Yux77Yux/platform_backend/pkg/auth"
)

func ClickCollection(ctx context.Context, req *generated.UpdateInteractionRequest) (*generated.UpdateInteractionResponse, error) {
	token := req.GetAccessToken().GetValue()
	response := new(generated.UpdateInteractionResponse)
	pass, userId, err := auth.Auth("update", "interaction", token)
	if err != nil {
		response.Msg = &common.ApiResponse{
			Status: common.ApiResponse_FAILED,
			Code:   "500",
		}
		return response, err
	}
	if !pass {
		response.Msg = &common.ApiResponse{
			Status: common.ApiResponse_ERROR,
			Code:   "403",
		}
		return response, nil
	}

	base_interaction := req.GetBase()
	base_interaction.UserId = userId

	timest := timestamppb.Now()
	operateInteraction := &generated.OperateInteraction{
		Base: &generated.BaseInteraction{
			UserId:     userId,
			CreationId: req.GetBase().GetCreationId(),
		},
		Action:    common.Operate_COLLECT,
		UpdatedAt: timest,
		SaveAt:    timest,
	}

	go func(operateInteraction *generated.OperateInteraction, ctx context.Context) {
		traceId, fullName := tools.GetMetadataValue(ctx, "trace-id"), tools.GetMetadataValue(ctx, "full-name")
		err = messaging.SendMessage(ctx, EXCHANGE_ADD_COLLECTION, KEY_ADD_COLLECTION, operateInteraction)
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

func ClickLike(ctx context.Context, req *generated.UpdateInteractionRequest) (*generated.UpdateInteractionResponse, error) {
	token := req.GetAccessToken().GetValue()
	response := new(generated.UpdateInteractionResponse)
	pass, userId, err := auth.Auth("update", "interaction", token)
	if err != nil {
		response.Msg = &common.ApiResponse{
			Status: common.ApiResponse_FAILED,
			Code:   "500",
		}
		return response, err
	}
	if !pass {
		response.Msg = &common.ApiResponse{
			Status: common.ApiResponse_ERROR,
			Code:   "403",
		}
		return response, nil
	}

	base_interaction := req.GetBase()
	base_interaction.UserId = userId

	timest := timestamppb.Now()
	operateInteraction := &generated.OperateInteraction{
		Base: &generated.BaseInteraction{
			UserId:     userId,
			CreationId: req.GetBase().GetCreationId(),
		},
		Action:    common.Operate_LIKE,
		UpdatedAt: timest,
	}

	go func(operateInteraction *generated.OperateInteraction, ctx context.Context) {
		traceId, fullName := tools.GetMetadataValue(ctx, "trace-id"), tools.GetMetadataValue(ctx, "full-name")
		err = messaging.SendMessage(ctx, EXCHANGE_ADD_LIKE, KEY_ADD_LIKE, operateInteraction)
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

func CancelCollections(ctx context.Context, req *generated.UpdateInteractionsRequest) (*generated.UpdateInteractionResponse, error) {
	token := req.GetAccessToken().GetValue()
	response := new(generated.UpdateInteractionResponse)
	pass, userId, err := auth.Auth("update", "interaction", token)
	if err != nil {
		response.Msg = &common.ApiResponse{
			Status: common.ApiResponse_FAILED,
			Code:   "500",
		}
		return response, err
	}
	if !pass {
		response.Msg = &common.ApiResponse{
			Status: common.ApiResponse_ERROR,
			Code:   "403",
		}
		return response, nil
	}

	base_interactions := req.GetBases()
	for _, val := range base_interactions {
		val.UserId = userId
	}

	err = cache.DelCollections(ctx, base_interactions)
	if err != nil {
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "500",
			Details: err.Error(),
		}
		return response, nil
	}

	timest := timestamppb.Now()
	length := len(base_interactions)
	operateInteractions := make([]*generated.OperateInteraction, length)
	for i := 0; i < length; i++ {
		operateInteractions[i] = &generated.OperateInteraction{
			Base:      base_interactions[i],
			UpdatedAt: timest,
			Action:    common.Operate_CANCEL_COLLECT,
		}
	}

	go func(operateInteraction []*generated.OperateInteraction, ctx context.Context) {
		traceId, fullName := tools.GetMetadataValue(ctx, "trace-id"), tools.GetMetadataValue(ctx, "full-name")
		anyOperateInteraction := &generated.AnyOperateInteraction{
			OperateInteractions: operateInteractions,
		}
		err = messaging.SendMessage(ctx, EXCHANGE_BATCH_UPDATE_DB, KEY_BATCH_UPDATE_DB, anyOperateInteraction)
		if err != nil {
			tools.LogError(traceId, fullName, err)
		}
	}(operateInteractions, ctx)

	response.Msg = &common.ApiResponse{
		Status: common.ApiResponse_SUCCESS,
		Code:   "202",
	}
	return response, nil
}

func DelHistories(ctx context.Context, req *generated.UpdateInteractionsRequest) (*generated.UpdateInteractionResponse, error) {
	token := req.GetAccessToken().GetValue()
	response := new(generated.UpdateInteractionResponse)
	pass, userId, err := auth.Auth("update", "interaction", token)
	if err != nil {
		response.Msg = &common.ApiResponse{
			Status: common.ApiResponse_FAILED,
			Code:   "500",
		}
		return response, err
	}
	if !pass {
		response.Msg = &common.ApiResponse{
			Status: common.ApiResponse_ERROR,
			Code:   "403",
		}
		return response, nil
	}
	base_interactions := req.GetBases()
	for _, val := range base_interactions {
		val.UserId = userId
	}

	err = cache.DelHistories(ctx, base_interactions)
	if err != nil {
		traceId, fullName := tools.GetMetadataValue(ctx, "trace-id"), tools.GetMetadataValue(ctx, "full-name")
		tools.LogError(traceId, fullName, err)
	}

	timest := timestamppb.Now()
	length := len(base_interactions)
	operateInteractions := make([]*generated.OperateInteraction, length)
	for i := 0; i < length; i++ {
		operateInteractions[i] = &generated.OperateInteraction{
			Base:      base_interactions[i],
			UpdatedAt: timest,
			Action:    common.Operate_DEL_VIEW,
		}
	}

	go func(operateInteraction []*generated.OperateInteraction, ctx context.Context) {
		traceId, fullName := tools.GetMetadataValue(ctx, "trace-id"), tools.GetMetadataValue(ctx, "full-name")
		anyOperateInteraction := &generated.AnyOperateInteraction{
			OperateInteractions: operateInteractions,
		}
		err = messaging.SendMessage(ctx, EXCHANGE_BATCH_UPDATE_DB, KEY_BATCH_UPDATE_DB, anyOperateInteraction)
		if err != nil {
			tools.LogError(traceId, fullName, err)
		}
	}(operateInteractions, ctx)

	response.Msg = &common.ApiResponse{
		Status: common.ApiResponse_SUCCESS,
		Code:   "202",
	}

	return response, nil
}

func CancelLike(ctx context.Context, req *generated.UpdateInteractionRequest) (*generated.UpdateInteractionResponse, error) {
	token := req.GetAccessToken().GetValue()
	response := new(generated.UpdateInteractionResponse)
	pass, userId, err := auth.Auth("update", "interaction", token)
	if err != nil {
		response.Msg = &common.ApiResponse{
			Status: common.ApiResponse_FAILED,
			Code:   "500",
		}
		return response, err
	}
	if !pass {
		response.Msg = &common.ApiResponse{
			Status: common.ApiResponse_ERROR,
			Code:   "403",
		}
		return response, nil
	}

	base_interaction := req.GetBase()
	base_interaction.UserId = userId

	interaction := &generated.OperateInteraction{
		Base:      base_interaction,
		Action:    common.Operate_CANCEL_LIKE,
		UpdatedAt: timestamppb.Now(),
	}
	go func(interaction *generated.OperateInteraction, ctx context.Context) {
		traceId, fullName := tools.GetMetadataValue(ctx, "trace-id"), tools.GetMetadataValue(ctx, "full-name")
		err = messaging.SendMessage(ctx, EXCHANGE_CANCEL_LIKE, KEY_CANCEL_LIKE, interaction)
		if err != nil {
			tools.LogError(traceId, fullName, err)
		}
	}(interaction, ctx)

	response.Msg = &common.ApiResponse{
		Status: common.ApiResponse_SUCCESS,
		Code:   "202",
	}

	return response, nil
}

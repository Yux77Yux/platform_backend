package internal

import (
	"log"

	"google.golang.org/protobuf/types/known/timestamppb"

	common "github.com/Yux77Yux/platform_backend/generated/common"
	generated "github.com/Yux77Yux/platform_backend/generated/interaction"
	cache "github.com/Yux77Yux/platform_backend/microservices/interaction/cache"
	// mq "github.com/Yux77Yux/platform_backend/microservices/interaction/messaging"
	dispatch "github.com/Yux77Yux/platform_backend/microservices/interaction/messaging/dispatch"
	db "github.com/Yux77Yux/platform_backend/microservices/interaction/repository"
	auth "github.com/Yux77Yux/platform_backend/pkg/auth"
)

func ClickCollection(req *generated.UpdateInteractionRequest) (*generated.UpdateInteractionResponse, error) {
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
	OperateInteraction := req.GetOperateInteraction()
	base_interaction := OperateInteraction.GetInterction()
	base_interaction.UserId = userId

	timest := timestamppb.Now()
	interaction := &generated.Interaction{
		Base:      base_interaction,
		SaveAt:    timest,
		UpdatedAt: timest,
		ActionTag: int32(OperateInteraction.GetAction()),
	}
	go dispatch.HandleRequest(interaction, dispatch.DbInteraction)
	go dispatch.HandleRequest(interaction, dispatch.CollectionCache)

	response.Msg = &common.ApiResponse{
		Status: common.ApiResponse_SUCCESS,
		Code:   "202",
	}

	return response, nil
}

func ClickLike(req *generated.UpdateInteractionRequest) (*generated.UpdateInteractionResponse, error) {
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

	OperateInteraction := req.GetOperateInteraction()
	base_interaction := OperateInteraction.GetInterction()
	base_interaction.UserId = userId

	timest := timestamppb.Now()
	interaction := &generated.Interaction{
		Base:      base_interaction,
		SaveAt:    timest,
		UpdatedAt: timest,
		ActionTag: int32(OperateInteraction.GetAction()),
	}
	go dispatch.HandleRequest(base_interaction, dispatch.DbInteraction)
	go dispatch.HandleRequest(interaction, dispatch.LikeCache)

	response.Msg = &common.ApiResponse{
		Status: common.ApiResponse_SUCCESS,
		Code:   "202",
	}

	return response, nil
}

func CancelCollections(req *generated.UpdateInteractionsRequest) (*generated.UpdateInteractionResponse, error) {
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

	OperateInteraction := req.GetOperateInteraction()
	base_interactions := OperateInteraction.GetAnyInterction()
	for _, val := range base_interactions {
		val.UserId = userId
	}

	err = cache.DelCollections(base_interactions)
	if err != nil {
		log.Printf("error: DelCollections %v", err)
	}

	length := len(base_interactions)
	interactions := make([]*generated.Interaction, length)
	for i, val := range base_interactions {
		interactions[i].SaveAt = nil
		interactions[i].Base = val
		interactions[i].UpdatedAt = timestamppb.Now()
		interactions[i].ActionTag = int32(OperateInteraction.GetAction())
	}
	err = db.UpdateInteractions(interactions)
	if err != nil {
		log.Printf("error: UpdateInteractions %v", err)
	}

	response.Msg = &common.ApiResponse{
		Status: common.ApiResponse_SUCCESS,
		Code:   "202",
	}

	return response, nil
}

func DelHistories(req *generated.UpdateInteractionsRequest) (*generated.UpdateInteractionResponse, error) {
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
	OperateInteraction := req.GetOperateInteraction()
	base_interactions := OperateInteraction.GetAnyInterction()
	for _, val := range base_interactions {
		val.UserId = userId
	}

	err = cache.DelHistories(base_interactions)
	if err != nil {
		log.Printf("error: DelHistories %v", err)
	}

	length := len(base_interactions)
	interactions := make([]*generated.Interaction, length)
	for i, val := range base_interactions {
		interactions[i].SaveAt = nil
		interactions[i].Base = val
		interactions[i].UpdatedAt = timestamppb.Now()
		interactions[i].ActionTag = int32(OperateInteraction.GetAction())
	}
	err = db.UpdateInteractions(interactions)
	if err != nil {
		log.Printf("error: UpdateInteractions %v", err)
	}

	response.Msg = &common.ApiResponse{
		Status: common.ApiResponse_SUCCESS,
		Code:   "202",
	}

	return response, nil
}

func CancelLike(req *generated.UpdateInteractionRequest) (*generated.UpdateInteractionResponse, error) {
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

	OperateInteraction := req.GetOperateInteraction()
	base_interaction := OperateInteraction.GetInterction()
	base_interaction.UserId = userId

	interaction := &generated.Interaction{
		Base:      base_interaction,
		ActionTag: int32(OperateInteraction.GetAction()),
	}
	go dispatch.HandleRequest(base_interaction, dispatch.DbInteraction)
	go dispatch.HandleRequest(interaction, dispatch.CancelLikeCache)

	response.Msg = &common.ApiResponse{
		Status: common.ApiResponse_SUCCESS,
		Code:   "202",
	}

	return response, nil
}

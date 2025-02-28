package internal

import (
	"log"

	"google.golang.org/protobuf/types/known/timestamppb"

	common "github.com/Yux77Yux/platform_backend/generated/common"
	generated "github.com/Yux77Yux/platform_backend/generated/interaction"
	cache "github.com/Yux77Yux/platform_backend/microservices/interaction/cache"
	messaging "github.com/Yux77Yux/platform_backend/microservices/interaction/messaging"

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

	base_interaction := req.GetBase()
	base_interaction.UserId = userId

	timest := timestamppb.Now()
	operateInteraction := &generated.OperateInteraction{
		Base: &generated.BaseInteraction{
			UserId:     userId,
			CreationId: req.GetBase().GetCreationId(),
		},
		Action:    common.Operate_VIEW,
		UpdatedAt: timest,
		SaveAt:    timest,
	}

	err = messaging.SendMessage(messaging.AddCollection, messaging.AddCollection, operateInteraction)
	if err != nil {
		log.Printf("error: SendMessage %v", err)
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "500",
			Details: err.Error(),
		}
		return response, nil
	}

	response.Msg = &common.ApiResponse{
		Status: common.ApiResponse_SUCCESS,
		Code:   "202",
	}

	return response, nil
}

func ClickLike(req *generated.UpdateInteractionRequest) (*generated.UpdateInteractionResponse, error) {
	log.Printf("req %v", req)
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
		Action:    common.Operate_VIEW,
		UpdatedAt: timest,
	}

	err = messaging.SendMessage(messaging.AddLike, messaging.AddLike, operateInteraction)
	if err != nil {
		log.Printf("error: SendMessage %v", err)
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "500",
			Details: err.Error(),
		}
		return response, nil
	}

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

	base_interactions := req.GetBases()
	for _, val := range base_interactions {
		val.UserId = userId
	}

	err = cache.DelCollections(base_interactions)
	if err != nil {
		log.Printf("error: DelCollections %v", err)
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

	anyOperateInteraction := &generated.AnyOperateInteraction{
		OperateInteractions: operateInteractions,
	}
	err = messaging.SendMessage(messaging.BatchUpdateDb, messaging.BatchUpdateDb, anyOperateInteraction)
	if err != nil {
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "500",
			Details: err.Error(),
		}
		return response, err
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
	base_interactions := req.GetBases()
	for _, val := range base_interactions {
		val.UserId = userId
	}

	err = cache.DelHistories(base_interactions)
	if err != nil {
		log.Printf("error: DelHistories %v", err)
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

	anyOperateInteraction := &generated.AnyOperateInteraction{
		OperateInteractions: operateInteractions,
	}
	err = messaging.SendMessage(messaging.BatchUpdateDb, messaging.BatchUpdateDb, anyOperateInteraction)
	if err != nil {
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "500",
			Details: err.Error(),
		}
		return response, err
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

	base_interaction := req.GetBase()
	base_interaction.UserId = userId

	interaction := &generated.OperateInteraction{
		Base:      base_interaction,
		Action:    common.Operate_DEL_VIEW,
		UpdatedAt: timestamppb.Now(),
	}
	err = messaging.SendMessage(messaging.CancelLike, messaging.CancelLike, interaction)
	if err != nil {
		log.Printf("error: SendMessage CancelLike %v", err)
	}

	response.Msg = &common.ApiResponse{
		Status: common.ApiResponse_SUCCESS,
		Code:   "202",
	}

	return response, nil
}

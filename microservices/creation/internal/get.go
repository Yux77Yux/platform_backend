package internal

import (
	"context"
	"log"

	common "github.com/Yux77Yux/platform_backend/generated/common"
	generated "github.com/Yux77Yux/platform_backend/generated/creation"
	auth "github.com/Yux77Yux/platform_backend/pkg/auth"

	cache "github.com/Yux77Yux/platform_backend/microservices/creation/cache"
	messaging "github.com/Yux77Yux/platform_backend/microservices/creation/messaging"
	db "github.com/Yux77Yux/platform_backend/microservices/creation/repository"
	tools "github.com/Yux77Yux/platform_backend/microservices/creation/tools"
)

func GetCreation(ctx context.Context, req *generated.GetCreationRequest) (*generated.GetCreationResponse, error) {
	response := new(generated.GetCreationResponse)
	// 取数据
	creationId := req.GetCreationId()
	creation, err := cache.GetCreationInfo(ctx, creationId)
	if err != nil {
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "500",
			Message: "Internal Server Error",
			Details: err.Error(),
		}
		return response, nil
	}

	status := creation.GetCreation().GetBaseInfo().GetStatus()

	if creation == nil {
		creation, err = db.GetDetailInTransaction(ctx, creationId)
		if err != nil {
			response.Msg = &common.ApiResponse{
				Status:  common.ApiResponse_ERROR,
				Code:    "500",
				Message: "Internal Server Error",
				Details: err.Error(),
			}
			return response, err
		}

		status = creation.GetCreation().GetBaseInfo().GetStatus()
		if status == generated.CreationStatus_PUBLISHED {
			// 存作品至redis
			go func(creation *generated.CreationInfo) {
				err := messaging.SendMessage(messaging.StoreCreationInfo, messaging.StoreCreationInfo, creation)
				if err != nil {
					log.Printf("error: GetCreation SendMessage %v", err)
				}
			}(creation)
		}
	}

	if status == generated.CreationStatus_DELETE {
		response.Msg = &common.ApiResponse{
			Status: common.ApiResponse_SUCCESS,
			Code:   "404",
		}
		return response, nil
	}

	response.CreationInfo = creation
	response.Msg = &common.ApiResponse{
		Status: common.ApiResponse_SUCCESS,
		Code:   "200",
	}
	return response, nil
}

func GetCreationPrivate(ctx context.Context, req *generated.GetCreationRequest) (*generated.GetCreationResponse, error) {
	// 取数据
	creationId := req.GetCreationId()

	creation, err := db.GetDetailInTransaction(ctx, creationId)
	if err != nil {
		return &generated.GetCreationResponse{
			Msg: &common.ApiResponse{
				Status:  common.ApiResponse_ERROR,
				Code:    "500",
				Message: "Internal Server Error",
				Details: err.Error(),
			},
		}, nil
	}

	return &generated.GetCreationResponse{
		CreationInfo: creation,
		Msg: &common.ApiResponse{
			Status: common.ApiResponse_SUCCESS,
			Code:   "200",
		},
	}, nil
}

func GetCreationList(ctx context.Context, req *generated.GetCreationListRequest) (*generated.GetCreationListResponse, error) {
	response := new(generated.GetCreationListResponse)

	ids := req.GetIds()
	creations, err := db.GetCreationCardInTransaction(ctx, ids)
	if err != nil {
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "500",
			Details: err.Error(),
		}
		return response, err
	}

	response.CreationInfoGroup = creations
	response.Msg = &common.ApiResponse{
		Status: common.ApiResponse_SUCCESS,
		Code:   "200",
	}
	return response, nil
}

func GetPublicCreationList(ctx context.Context, req *generated.GetCreationListRequest) (*generated.GetCreationListResponse, error) {
	response := new(generated.GetCreationListResponse)

	ids := req.GetIds()
	creations, err := db.GetCreationCardInTransaction(ctx, ids)
	if err != nil {
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "500",
			Details: err.Error(),
		}
		return response, err
	}

	length := len(ids)
	filteredCreations := make([]*generated.CreationInfo, 0, length)
	for _, info := range creations {
		if info.GetCreation().GetBaseInfo().GetStatus() == generated.CreationStatus_PUBLISHED {
			filteredCreations = append(filteredCreations, info)
		}
	}

	response.CreationInfoGroup = filteredCreations
	response.Msg = &common.ApiResponse{
		Status: common.ApiResponse_SUCCESS,
		Code:   "200",
	}
	return response, nil
}

// 拿用户id
func GetSpaceCreations(ctx context.Context, req *generated.GetSpaceCreationsRequest) (*generated.GetCreationListResponse, error) {
	response := new(generated.GetCreationListResponse)
	id := req.GetUserId()
	page := req.GetPage()
	if page == 0 {
		page = 1
	}
	byWhat := req.GetByWhat()
	typeStr := tools.GetSpaceCreationCountType(byWhat)

	ids, count, err := cache.GetSpaceCreationList(ctx, id, page, typeStr)
	if err != nil {
		log.Printf("error: cache GetSpaceCreationList %v", err)
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "500",
			Details: err.Error(),
		}
		return response, err
	}

	infos, err := db.GetCreationCardInTransaction(ctx, ids)
	if err != nil {
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "500",
			Details: err.Error(),
		}
		return response, err
	}

	filter := make([]*generated.CreationInfo, 0, len(infos))
	for _, info := range infos {
		creation := info.GetCreation()
		base := creation.GetBaseInfo()
		if base.GetStatus() != generated.CreationStatus_PUBLISHED {
			continue
		}
		filter = append(filter, info)
	}

	response.CreationInfoGroup = filter
	response.Count = count
	response.Msg = &common.ApiResponse{
		Status: common.ApiResponse_SUCCESS,
		Code:   "200",
	}
	return response, nil
}

func GetUserCreations(ctx context.Context, req *generated.GetUserCreationsRequest) (*generated.GetCreationListResponse, error) {
	response := new(generated.GetCreationListResponse)
	pass, author_id, err := auth.Auth("get", "creation", req.GetAccessToken().GetValue())
	if err != nil {
		return &generated.GetCreationListResponse{
			Msg: &common.ApiResponse{
				Status: common.ApiResponse_FAILED,
				Code:   "500",
			},
		}, err
	}
	if !pass {
		return &generated.GetCreationListResponse{
			Msg: &common.ApiResponse{
				Status: common.ApiResponse_ERROR,
				Code:   "403",
			},
		}, nil
	}

	req.UserId = author_id

	infos, count, err := db.GetUserCreations(ctx, req)
	if err != nil {
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "500",
			Details: err.Error(),
		}
		return response, err
	}

	response.CreationInfoGroup = infos
	response.Count = count
	response.Msg = &common.ApiResponse{
		Status: common.ApiResponse_SUCCESS,
		Code:   "200",
	}
	return response, nil
}

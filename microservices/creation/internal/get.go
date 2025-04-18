package internal

import (
	"context"
	"fmt"

	common "github.com/Yux77Yux/platform_backend/generated/common"
	generated "github.com/Yux77Yux/platform_backend/generated/creation"
	auth "github.com/Yux77Yux/platform_backend/pkg/auth"
	errMap "github.com/Yux77Yux/platform_backend/pkg/error"

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
		return response, err
	}

	if creation == nil {
		creation, err = db.GetDetailInTransaction(ctx, creationId)
		fmt.Printf("cc %v\n", creation)
		if err != nil {
			response.Msg = &common.ApiResponse{
				Status:  common.ApiResponse_ERROR,
				Code:    "500",
				Message: "Internal Server Error",
				Details: err.Error(),
			}
			return response, err
		}
		if creation == nil {
			response.Msg = &common.ApiResponse{
				Status: common.ApiResponse_ERROR,
				Code:   "404",
			}
			return response, nil
		}

		status := creation.GetCreation().GetBaseInfo().GetStatus()
		if status == generated.CreationStatus_PUBLISHED {
			go func(creation *generated.CreationInfo, ctx context.Context) {
				traceId, fullName := tools.GetMetadataValue(ctx, "trace-id"), tools.GetMetadataValue(ctx, "full-name")
				err := messaging.SendMessage(ctx, EXCHANGE_STORE_CREATION, KEY_STORE_CREATION, creation)
				if err != nil {
					tools.LogError(traceId, fullName, err)
				}
			}(creation, ctx)
		}
	}

	status := creation.GetCreation().GetBaseInfo().GetStatus()
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

func GetCreationPrivate(ctx context.Context, req *generated.GetCreationPrivateRequest) (*generated.GetCreationResponse, error) {
	response := new(generated.GetCreationResponse)
	// 取数据
	token := req.GetAccessToken().GetValue()
	pass, authorId, err := auth.Auth("get", "creation", token)
	if err != nil {
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "500",
			Message: "Internal Server Error",
			Details: err.Error(),
		}
		return response, err
	}
	if !pass {
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "403",
			Details: "you are not pass the auth",
		}
		return response, nil
	}

	creationId := req.GetCreationId()
	creationInfo, err := db.GetDetailInTransaction(ctx, creationId)
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

	if creationInfo.Creation.BaseInfo.GetAuthorId() != authorId {
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "403",
			Details: "the req not pass the auth",
		}
		return response, nil
	}

	response.Msg = &common.ApiResponse{
		Status: common.ApiResponse_SUCCESS,
		Code:   "200",
	}
	response.CreationInfo = creationInfo
	return response, nil
}

func GetCreationList(ctx context.Context, req *generated.GetCreationListRequest) (*generated.GetCreationListResponse, error) {
	response := new(generated.GetCreationListResponse)

	ids := req.GetIds()
	creations, err := db.GetCreationCardInTransaction(ctx, ids)
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
	infos, err := db.GetCreationCardInTransaction(ctx, ids)
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
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "500",
			Details: err.Error(),
		}
		return response, err
	}

	infos, err := db.GetCreationCardInTransaction(ctx, ids)
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

	response.CreationInfoGroup = infos
	response.Count = count
	response.Msg = &common.ApiResponse{
		Status: common.ApiResponse_SUCCESS,
		Code:   "200",
	}
	return response, nil
}

func SearchCreation(ctx context.Context, req *generated.SearchCreationRequest) (*generated.GetCreationListResponse, error) {
	const LIMIT = 15
	response := new(generated.GetCreationListResponse)
	title := req.GetTitle()
	page := req.GetPage()
	if page == 0 {
		page = 1
	}

	ids, count, err := search_client.SearchWithPagination("creations", title, int(page), int(LIMIT))
	if err != nil {
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "500",
			Details: err.Error(),
		}
		return response, err
	}

	infos, err := db.GetCreationCardInTransaction(ctx, ids)
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

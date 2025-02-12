package internal

import (
	"context"
	"fmt"

	common "github.com/Yux77Yux/platform_backend/generated/common"
	generated "github.com/Yux77Yux/platform_backend/generated/creation"

	cache "github.com/Yux77Yux/platform_backend/microservices/creation/cache"
	db "github.com/Yux77Yux/platform_backend/microservices/creation/repository"
)

func GetCreation(ctx context.Context, req *generated.GetCreationRequest) (*generated.GetCreationResponse, error) {
	// 取数据
	creationId := req.GetCreationId()
	cache.GetCreationInfo(creationId)

	data, err := db.GetDetailInTransaction(ctx, creationId)
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
		CreationInfo: data,
		Msg: &common.ApiResponse{
			Status: common.ApiResponse_SUCCESS,
			Code:   "200",
		},
	}, nil
}

func GetCreations(ctx context.Context, ids []int64) ([]*generated.CreationInfo, error) {
	data, err := db.GetCardInTransaction(ctx, ids)
	if err != nil {
		return nil, fmt.Errorf("error: get creation in db error :%w", err)
	}
	return data, nil
}

// 此处先使用分区进行相似推荐，
// 拿到此作品的分区，然后到ms或es查相同标签个数的作品id，
// 拿到id在redis取热度分值，然后排序，返回前十六个视频作品相似作品
func GetSimilarCreationList(ctx context.Context, req *generated.GetPublicCreationListRequest) (*generated.GetCreationListResponse, error) {
	response := &generated.GetCreationListResponse{}
	// id := req.GetId()

	// 此处拿到id然后向redis或Meilisearch查相似列表

	var err error
	if err != nil {
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "500",
			Details: err.Error(),
		}
		return response, nil
	}

	response.Msg = &common.ApiResponse{
		Status: common.ApiResponse_SUCCESS,
		Code:   "200",
	}
	return response, nil
}

// 拿用户id，然后author_id = id 作为筛选条件
func GetSpaceCreationList(ctx context.Context, req *generated.GetPublicCreationListRequest) (*generated.GetCreationListResponse, error) {
	response := &generated.GetCreationListResponse{}
	// id := req.GetId()

	// 此处拿到id然后向redis或Meilisearch查用户的作品

	var err error = nil

	if err != nil {
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "500",
			Details: err.Error(),
		}
		return response, nil
	}

	response.Msg = &common.ApiResponse{
		Status: common.ApiResponse_SUCCESS,
		Code:   "200",
	}
	return response, nil
}

// 收藏夹
func GetCollectionCreationList(ctx context.Context, req *generated.GetSpecificCreationListRequest) (*generated.GetCreationListResponse, error) {
	response := &generated.GetCreationListResponse{}
	// token := req.GetAccessToken().GetValue()

	// 此处拿到id然后向redis或Meilisearch查用户的作品

	var err error = nil
	if err != nil {
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "500",
			Details: err.Error(),
		}
		return response, nil
	}

	response.Msg = &common.ApiResponse{
		Status: common.ApiResponse_SUCCESS,
		Code:   "200",
	}
	return response, nil
}

// 主页，打开网页，请求id，取视频推荐的计算结果，拿热度和个性化混合的creation_id数量，请求获取信息
// 此处需要推荐系统配合
func GetHomeCreationList(ctx context.Context, req *generated.GetSpecificCreationListRequest) (*generated.GetCreationListResponse, error) {
	response := &generated.GetCreationListResponse{}

	var err error = nil
	if err != nil {
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "500",
			Details: err.Error(),
		}
		return response, nil
	}

	response.Msg = &common.ApiResponse{
		Status: common.ApiResponse_SUCCESS,
		Code:   "200",
	}
	return response, nil
}

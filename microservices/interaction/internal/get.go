package internal

import (
	"fmt"

	common "github.com/Yux77Yux/platform_backend/generated/common"
	generated "github.com/Yux77Yux/platform_backend/generated/interaction"

	// cache "github.com/Yux77Yux/platform_backend/microservices/interaction/cache"
	// messaging "github.com/Yux77Yux/platform_backend/microservices/interaction/messaging"
	db "github.com/Yux77Yux/platform_backend/microservices/interaction/repository"
)

func Getinteraction(req *generated.GetinteractionRequest) (*generated.GetinteractionResponse, error) {
	// 取数据
	data, err := db.GetDetailInTransaction(req.GetinteractionId())
	if err != nil {
		return &generated.GetinteractionResponse{
			Msg: &common.ApiResponse{
				Status:  common.ApiResponse_ERROR,
				Code:    "500",
				Message: "Internal Server Error",
				Details: err.Error(),
			},
		}, nil
	}

	return &generated.GetinteractionResponse{
		interactionInfo: data,
		Msg: &common.ApiResponse{
			Status: common.ApiResponse_SUCCESS,
			Code:   "200",
		},
	}, nil
}

func getinteractions(ids []int64) ([]*generated.interactionInfo, error) {
	data, err := db.GetCardInTransaction(ids)
	if err != nil {
		return nil, fmt.Errorf("error: get interaction in db error :%w", err)
	}
	return data, nil
}

// 此处先使用分区进行相似推荐，
// 拿到此作品的分区，然后到ms或es查相同标签个数的作品id，
// 拿到id在redis取热度分值，然后排序，返回前十六个视频作品相似作品
func GetSimilarinteractionList(req *generated.GetSpecificinteractionListRequest) (*generated.GetinteractionListResponse, error) {
	response := &generated.GetinteractionListResponse{}
	id := req.GetId()

	if err != nil {
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "500",
			Details: err.Error(),
		}
		return response, nil
	}

	response.interactionInfoGroup = result
	response.Msg = &common.ApiResponse{
		Status: common.ApiResponse_SUCCESS,
		Code:   "200",
	}
	return response, nil
}

// 拿用户id，然后author_id = id 作为筛选条件
func GetSpaceinteractionList(req *generated.GetSpecificinteractionListRequest) (*generated.GetinteractionListResponse, error) {
	response := &generated.GetinteractionListResponse{}
	id := req.GetId()

	if err != nil {
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "500",
			Details: err.Error(),
		}
		return response, nil
	}

	response.interactionInfoGroup = result
	response.Msg = &common.ApiResponse{
		Status: common.ApiResponse_SUCCESS,
		Code:   "200",
	}
	return response, nil
}

// 收藏夹
func GetCollectioninteractionList(req *generated.GetSpecificinteractionListRequest) (*generated.GetinteractionListResponse, error) {
	response := &generated.GetinteractionListResponse{}
	id := req.GetId()

	if err != nil {
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "500",
			Details: err.Error(),
		}
		return response, nil
	}

	response.interactionInfoGroup = result
	response.Msg = &common.ApiResponse{
		Status: common.ApiResponse_SUCCESS,
		Code:   "200",
	}
	return response, nil
}

// 主页，打开网页，请求id，取视频推荐的计算结果，拿热度和个性化混合的interaction_id数量，请求获取信息
// 此处需要推荐系统配合
func GetHomeinteractionList(req *generated.GetSpecificinteractionListRequest) (*generated.GetinteractionListResponse, error) {
	response := &generated.GetinteractionListResponse{}

	id := req.GetId()

	if err != nil {
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "500",
			Details: err.Error(),
		}
		return response, nil
	}

	response.interactionInfoGroup = result
	response.Msg = &common.ApiResponse{
		Status: common.ApiResponse_SUCCESS,
		Code:   "200",
	}
	return response, nil
}

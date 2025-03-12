package internal

import (
	"context"
	"log"

	generated "github.com/Yux77Yux/platform_backend/generated/comment"
	common "github.com/Yux77Yux/platform_backend/generated/common"
	db "github.com/Yux77Yux/platform_backend/microservices/comment/repository"
	auth "github.com/Yux77Yux/platform_backend/pkg/auth"
)

func GetComments(ctx context.Context, req *generated.GetCommentsRequest) (*generated.GetCommentsResponse, error) {
	response := new(generated.GetCommentsResponse)
	ids := req.GetIds()

	comments, err := db.GetComments(ctx, ids)
	if err != nil {
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "500",
			Details: err.Error(),
		}
		return response, err
	}

	return &generated.GetCommentsResponse{
		Comments: comments,
		Msg: &common.ApiResponse{
			Status: common.ApiResponse_SUCCESS,
			Code:   "200",
		},
	}, nil
}

// 第一次请求评论
func InitialComments(ctx context.Context, req *generated.InitialCommentsRequest) (*generated.InitialCommentsResponse, error) {
	creationId := req.GetCreationId()

	area, comments, count, err := db.GetInitialTopCommentsInTransaction(ctx, creationId)
	if err != nil {
		log.Printf("error: creationId %d %v", creationId, err)
		return &generated.InitialCommentsResponse{
			Msg: &common.ApiResponse{
				Status:  common.ApiResponse_ERROR,
				Code:    "404",
				Message: err.Error(),
				Details: err.Error(),
			},
		}, err
	}

	return &generated.InitialCommentsResponse{
		Comments:    comments,
		CommentArea: area,
		PageCount:   count,
		Msg: &common.ApiResponse{
			Status: common.ApiResponse_SUCCESS,
			Code:   "200",
		},
	}, nil
}

// 一级评论
func GetTopComments(ctx context.Context, req *generated.GetTopCommentsRequest) (*generated.GetTopCommentsResponse, error) {
	creationId := req.GetCreationId()
	page := req.GetPage()

	comments, err := db.GetTopCommentsInTransaction(ctx, creationId, page)
	if err != nil {
		return &generated.GetTopCommentsResponse{
			Msg: &common.ApiResponse{
				Status:  common.ApiResponse_ERROR,
				Code:    "500",
				Message: err.Error(),
				Details: err.Error(),
			},
		}, nil
	}

	return &generated.GetTopCommentsResponse{
		Comments: comments,
		Msg: &common.ApiResponse{
			Status: common.ApiResponse_SUCCESS,
			Code:   "200",
		},
	}, nil
}

// 查看二级评论
func GetSecondComments(ctx context.Context, req *generated.GetSecondCommentsRequest) (*generated.GetSecondCommentsResponse, error) {
	creationId := req.GetCreationId()
	root := req.GetRoot()
	page := req.GetPage()

	comments, err := db.GetSecondCommentsInTransaction(ctx, creationId, root, page)
	if err != nil {
		return &generated.GetSecondCommentsResponse{
			Msg: &common.ApiResponse{
				Status:  common.ApiResponse_ERROR,
				Code:    "500",
				Message: err.Error(),
				Details: err.Error(),
			},
		}, nil
	}

	return &generated.GetSecondCommentsResponse{
		Comments: comments,
		Msg: &common.ApiResponse{
			Status: common.ApiResponse_SUCCESS,
			Code:   "200",
		},
	}, nil
}

// 登录用户查看回复自己的评论
func GetReplyComments(ctx context.Context, req *generated.GetReplyCommentsRequest) (*generated.GetCommentsResponse, error) {
	response := &generated.GetCommentsResponse{}
	pass, user_id, err := auth.Auth("get", "comment", req.GetAccessToken().GetValue())
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

	comments, err := db.GetReplyCommentsInTransaction(ctx, user_id, req.GetPage())
	if err != nil {
		return &generated.GetCommentsResponse{
			Msg: &common.ApiResponse{
				Status:  common.ApiResponse_ERROR,
				Code:    "500",
				Message: err.Error(),
				Details: err.Error(),
			},
		}, nil
	}

	// 以上为鉴权
	return &generated.GetCommentsResponse{
		Comments: comments,
		Msg: &common.ApiResponse{
			Status: common.ApiResponse_SUCCESS,
			Code:   "200",
		},
	}, nil
}

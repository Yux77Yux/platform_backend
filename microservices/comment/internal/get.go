package internal

import (
	generated "github.com/Yux77Yux/platform_backend/generated/comment"
	common "github.com/Yux77Yux/platform_backend/generated/common"
	db "github.com/Yux77Yux/platform_backend/microservices/comment/repository"
	auth "github.com/Yux77Yux/platform_backend/pkg/auth"
)

// 一级评论
func InitalComments(req *generated.InitalCommentsRequest) (*generated.InitalCommentsResponse, error) {
	creationId := req.GetCreationId()

	area, comments, err := db.GetFirstCommentsInTransaction(creationId)
	if err != nil {
		return &generated.InitalCommentsResponse{
			Msg: &common.ApiResponse{
				Status:  common.ApiResponse_ERROR,
				Code:    "404",
				Message: err.Error(),
				Details: err.Error(),
			},
		}, nil
	}

	return &generated.InitalCommentsResponse{
		Comments:    comments,
		CommentArea: area,
		Msg: &common.ApiResponse{
			Status: common.ApiResponse_SUCCESS,
			Code:   "200",
		},
	}, nil
}

// 一级评论
func GetTopComments(req *generated.GetTopCommentsRequest) (*generated.GetCommentsResponse, error) {
	creationId := req.GetCreationId()
	page := req.GetPage()

	comments, err := db.GetTopCommentsInTransaction(creationId, page)
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

	return &generated.GetCommentsResponse{
		Comments: comments,
		Msg: &common.ApiResponse{
			Status: common.ApiResponse_SUCCESS,
			Code:   "200",
		},
	}, nil
}

// 查看二级评论
func GetSecondComments(req *generated.GetSecondCommentsRequest) (*generated.GetCommentsResponse, error) {
	creationId := req.GetCreationId()
	root := req.GetRoot()
	page := req.GetPage()

	comments, err := db.GetSecondCommentsInTransaction(creationId, root, page)
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

	return &generated.GetCommentsResponse{
		Comments: comments,
		Msg: &common.ApiResponse{
			Status: common.ApiResponse_SUCCESS,
			Code:   "200",
		},
	}, nil
}

// 登录用户查看回复自己的评论
func GetReplyComments(req *generated.GetReplyCommentsRequest) (*generated.GetCommentsResponse, error) {
	response := &generated.GetCommentsResponse{}
	pass, user_id, err := auth.Auth("post", "comment", req.GetAccessToken().GetValue())
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

	page := req.GetPage()

	comments, err := db.GetReplyCommentsInTransaction(user_id, page)
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

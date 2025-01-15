package internal

import (
	"fmt"

	generated "github.com/Yux77Yux/platform_backend/generated/comment"
	common "github.com/Yux77Yux/platform_backend/generated/common"
	db "github.com/Yux77Yux/platform_backend/microservices/comment/repository"
)

// 一级评论
func GetTopComment(req *generated.GetTopCommentRequest) (*generated.GetCommentsResponse, error) {
	id := req.GetCreationId()

	area, comments, err := db.GetTopCommentInTransaction(id)

	return &generated.GetCommentsResponse{
		Msg: &common.ApiResponse{
			Status: common.ApiResponse_SUCCESS,
			Code:   "200",
		},
	}, nil
}

// 查看二级评论
func GetSecondComment(req *generated.GetSecondCommentRequest) (*generated.GetCommentsResponse, error) {

	return &generated.GetCommentsResponse{
		Msg: &common.ApiResponse{
			Status: common.ApiResponse_SUCCESS,
			Code:   "200",
		},
	}, nil
}

// 登录用户查看回复自己的评论
func GetReplyComment(req *generated.GetReplyCommentRequest) (*generated.GetCommentsResponse, error) {
	return &generated.GetCommentsResponse{
		Msg: &common.ApiResponse{
			Status: common.ApiResponse_SUCCESS,
			Code:   "200",
		},
	}, nil
}

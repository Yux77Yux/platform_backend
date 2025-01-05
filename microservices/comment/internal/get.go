package internal

import (
	// "fmt"

	generated "github.com/Yux77Yux/platform_backend/generated/comment"
	common "github.com/Yux77Yux/platform_backend/generated/common"
	// cache "github.com/Yux77Yux/platform_backend/microservices/comment/cache"
	// messaging "github.com/Yux77Yux/platform_backend/microservices/comment/messaging"
	// db "github.com/Yux77Yux/platform_backend/microservices/comment/repository"
)

func GetTopComment(req *generated.GetTopCommentRequest) (*generated.GetCommentsResponse, error) {
	// comment_id := req.GetcommentId()
	// 用于后来的黑名单,尚未开发
	// accessToken := req.GetAccessToken()
	// block := false

	return &generated.GetCommentsResponse{
		Msg: &common.ApiResponse{
			Status: common.ApiResponse_SUCCESS,
			Code:   "200",
		},
	}, nil
}

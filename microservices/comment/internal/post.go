package internal

import (
	"fmt"

	generated "github.com/Yux77Yux/platform_backend/generated/comment"
	common "github.com/Yux77Yux/platform_backend/generated/common"
	messaging "github.com/Yux77Yux/platform_backend/microservices/comment/messaging"
	auth "github.com/Yux77Yux/platform_backend/pkg/auth"
)

func PublishComment(req *generated.PublishCommentRequest) (*generated.PublishCommentResponse, error) {
	response := &generated.PublishCommentResponse{}
	token := req.GetAccessToken().GetValue()
	pass, user_id, err := auth.Auth("post", "comment", token)
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
	// 以上为鉴权

	// 将token中的userId填充到请求体
	comment := req.GetComment()
	comment.UserId = user_id

	// 异步处理
	err = messaging.SendMessage(messaging.PublishComment, messaging.PublishComment, comment)
	if err != nil {
		err = fmt.Errorf("error: SendMessage PublishComment error %w", err)
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "500",
			Details: err.Error(),
		}
		return response, nil
	}

	response.Msg = &common.ApiResponse{
		Status: common.ApiResponse_PENDING,
		Code:   "202",
	}

	return response, nil
}

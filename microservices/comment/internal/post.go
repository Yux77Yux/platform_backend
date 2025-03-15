package internal

import (
	"context"

	generated "github.com/Yux77Yux/platform_backend/generated/comment"
	common "github.com/Yux77Yux/platform_backend/generated/common"
	messaging "github.com/Yux77Yux/platform_backend/microservices/comment/messaging"
	tools "github.com/Yux77Yux/platform_backend/microservices/comment/tools"
	auth "github.com/Yux77Yux/platform_backend/pkg/auth"
)

func PublishComment(ctx context.Context, req *generated.PublishCommentRequest) (*generated.PublishCommentResponse, error) {
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
	content := comment.GetContent()
	if err := tools.CheckStringLength(content, CONTENT_MIN_LENGTH, CONTENT_MAX_LENGTH); err != nil {
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "400",
			Details: err.Error(),
		}
		return response, err
	}

	comment.UserId = user_id

	// 异步处理
	if err != nil {
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "500",
			Details: err.Error(),
		}
		return response, nil
	}

	go func(comment *generated.Comment, ctx context.Context) {
		traceId, fullName := tools.GetMetadataValue(ctx, "trace-id"), tools.GetMetadataValue(ctx, "full-name")
		err = messaging.SendMessage(ctx, messaging.PublishComment, messaging.PublishComment, comment)
		if err != nil {
			tools.LogError(traceId, fullName, err)
		}
	}(comment, ctx)

	response.Msg = &common.ApiResponse{
		Status: common.ApiResponse_PENDING,
		Code:   "202",
	}

	return response, nil
}

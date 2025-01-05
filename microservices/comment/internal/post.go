package internal

import (
	generated "github.com/Yux77Yux/platform_backend/generated/comment"
	common "github.com/Yux77Yux/platform_backend/generated/common"
	// messaging "github.com/Yux77Yux/platform_backend/microservices/comment/messaging"
	auth "github.com/Yux77Yux/platform_backend/pkg/auth"
)

func PublishComment(req *generated.PublishCommentRequest) (*generated.PublishCommentResponse, error) {
	pass, err := auth.Auth(req.GetComment().GetUserId(), "post", "comment", req.GetAccessToken().GetValue())
	if err != nil {
		return &generated.PublishCommentResponse{
			Msg: &common.ApiResponse{
				Status: common.ApiResponse_FAILED,
				Code:   "500",
			},
		}, err
	}
	if !pass {
		return &generated.PublishCommentResponse{
			Msg: &common.ApiResponse{
				Status: common.ApiResponse_ERROR,
				Code:   "403",
			},
		}, nil
	}
	// 以上为鉴权

	// // 异步处理
	// if req.GetBaseInfo().GetStatus() == generated.CommentStatus_DRAFT {
	// 	messaging.SendMessage("draftComment", "draftComment", req.GetBaseInfo())
	// } else if req.GetBaseInfo().GetStatus() == generated.CommentStatus_PENDING {
	// 	messaging.SendMessage("pendingComment", "pendingComment", req.GetBaseInfo())
	// }

	return &generated.PublishCommentResponse{
		Msg: &common.ApiResponse{
			Status: common.ApiResponse_PENDING,
			Code:   "202",
		},
	}, nil
}

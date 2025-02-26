package internal

import (
	"fmt"
	"log"

	"google.golang.org/protobuf/types/known/timestamppb"

	common "github.com/Yux77Yux/platform_backend/generated/common"
	generated "github.com/Yux77Yux/platform_backend/generated/creation"
	messaging "github.com/Yux77Yux/platform_backend/microservices/creation/messaging"
	db "github.com/Yux77Yux/platform_backend/microservices/creation/repository"
	auth "github.com/Yux77Yux/platform_backend/pkg/auth"
	snow "github.com/Yux77Yux/platform_backend/pkg/snow"
)

func UploadCreation(req *generated.UploadCreationRequest) (*generated.UploadCreationResponse, error) {
	pass, author_id, err := auth.Auth("post", "creation", req.GetAccessToken().GetValue())
	if err != nil {
		return &generated.UploadCreationResponse{
			Msg: &common.ApiResponse{
				Status: common.ApiResponse_FAILED,
				Code:   "500",
			},
		}, err
	}
	if !pass {
		return &generated.UploadCreationResponse{
			Msg: &common.ApiResponse{
				Status: common.ApiResponse_ERROR,
				Code:   "403",
			},
		}, nil
	}
	// 以上为鉴权
	baseInfo := req.GetBaseInfo()
	baseInfo.AuthorId = author_id
	status := baseInfo.GetStatus()

	creation := &generated.Creation{
		CreationId: snow.GetId(),
		BaseInfo:   baseInfo,
		UploadTime: timestamppb.Now(),
	}

	err = db.CreationAddInTransaction(creation)
	if err != nil {
		log.Printf("db CreationAddInTransaction occur error: %v", err)
		return &generated.UploadCreationResponse{
			Msg: &common.ApiResponse{
				Status: common.ApiResponse_ERROR,
				Code:   "500",
			},
		}, err
	}

	// 异步处理
	if status == generated.CreationStatus_PENDING {
		err = messaging.SendMessage(messaging.PendingCreation, messaging.PendingCreation, baseInfo)
		if err != nil {
			log.Printf("error: publish failed because %v", err)
			newErr := fmt.Errorf("error: publish error %w and trun back into draft", err)

			updateStatus := &generated.CreationUpdateStatus{
				CreationId: creation.GetCreationId(),
				AuthorId:   baseInfo.GetAuthorId(),
				Status:     generated.CreationStatus_DRAFT,
			}
			errSecond := db.UpdateCreationStatusInTransaction(updateStatus)
			if errSecond != nil {
				log.Printf("error: db UpdateCreationStatusInTransaction occur error: %v", err)
				newErr = fmt.Errorf("%w,but  %w", newErr, errSecond)
			}

			return &generated.UploadCreationResponse{
				Msg: &common.ApiResponse{
					Status:  common.ApiResponse_ERROR,
					Code:    "201",
					Details: newErr.Error(),
				},
			}, newErr
		}
	}

	return &generated.UploadCreationResponse{
		Msg: &common.ApiResponse{
			Status: common.ApiResponse_SUCCESS,
			Code:   "201",
		},
	}, nil
}

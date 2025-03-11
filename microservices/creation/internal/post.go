package internal

import (
	"fmt"
	"log"

	"google.golang.org/protobuf/types/known/timestamppb"

	common "github.com/Yux77Yux/platform_backend/generated/common"
	generated "github.com/Yux77Yux/platform_backend/generated/creation"
	messaging "github.com/Yux77Yux/platform_backend/microservices/creation/messaging"
	db "github.com/Yux77Yux/platform_backend/microservices/creation/repository"
	tools "github.com/Yux77Yux/platform_backend/microservices/creation/tools"
	auth "github.com/Yux77Yux/platform_backend/pkg/auth"
	snow "github.com/Yux77Yux/platform_backend/pkg/snow"
)

func UploadCreation(req *generated.UploadCreationRequest) (*generated.UploadCreationResponse, error) {
	response := new(generated.UploadCreationResponse)
	pass, author_id, err := auth.Auth("post", "creation", req.GetAccessToken().GetValue())
	if err != nil {
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_FAILED,
			Code:    "500",
			Details: err.Error(),
		}
		return response, err
	}
	if !pass {
		response.Msg = &common.ApiResponse{
			Status: common.ApiResponse_ERROR,
			Code:   "403",
		}
		return response, err
	}
	// 以上为鉴权
	baseInfo := req.GetBaseInfo()

	src := baseInfo.GetSrc()
	if tools.IsValidVideoURL(src) {
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "400",
			Details: "Video source URL is invalid",
		}
		return response, err
	}
	thumbnail := baseInfo.GetThumbnail()
	if tools.IsValidImageURL(thumbnail) {
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "400",
			Details: "Image URL is invalid",
		}
		return response, err
	}

	bio := baseInfo.GetBio()
	if err := tools.CheckStringLength(bio, BIO_MIN_LENGTH, BIO_MAX_LENGTH); err != nil {
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "400",
			Details: err.Error(),
		}
		return response, err
	}

	title := baseInfo.GetTitle()
	if err := tools.CheckStringLength(title, TITLE_MIN_LENGTH, TITLE_MAX_LENGTH); err != nil {
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "400",
			Details: err.Error(),
		}
		return response, err
	}

	baseInfo.AuthorId = author_id
	status := baseInfo.GetStatus()

	creationId := snow.GetId()
	creation := &generated.Creation{
		CreationId: creationId,
		BaseInfo:   baseInfo,
		UploadTime: timestamppb.Now(),
	}

	err = db.CreationAddInTransaction(creation)
	if err != nil {
		log.Printf("db CreationAddInTransaction occur error: %v", err)
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "500",
			Details: err.Error(),
		}
		return response, err
	}

	// 异步处理
	if status == generated.CreationStatus_PENDING {
		err = messaging.SendMessage(messaging.PendingCreation, messaging.PendingCreation, &common.CreationId{
			Id: creationId,
		})
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

			response.Msg = &common.ApiResponse{
				Status:  common.ApiResponse_ERROR,
				Code:    "201",
				Details: newErr.Error(),
			}
			return response, newErr
		}
	}

	response.Msg = &common.ApiResponse{
		Status: common.ApiResponse_ERROR,
		Code:   "201",
	}
	return response, nil
}

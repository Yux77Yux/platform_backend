package internal

import (
	"context"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	common "github.com/Yux77Yux/platform_backend/generated/common"
	generated "github.com/Yux77Yux/platform_backend/generated/creation"
	tools "github.com/Yux77Yux/platform_backend/microservices/creation/tools"
	auth "github.com/Yux77Yux/platform_backend/pkg/auth"
	errMap "github.com/Yux77Yux/platform_backend/pkg/error"
)

func UploadCreation(ctx context.Context, req *generated.UploadCreationRequest) (*generated.UploadCreationResponse, error) {
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
	if !tools.IsValidVideoURL(src) {
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "400",
			Details: "Video source URL is invalid",
		}
		return response, err
	}
	thumbnail := baseInfo.GetThumbnail()
	if !tools.IsValidImageURL(thumbnail) {
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

	creation := &generated.Creation{
		CreationId: tools.GetSnowId(),
		BaseInfo:   baseInfo,
		UploadTime: timestamppb.Now(),
	}
	for i := 0; i < 3; i++ {
		err = db.CreationAddInTransaction(ctx, creation)
		if err != nil {
			if isServerError := errMap.IsServerError(err); isServerError {
				response.Msg = &common.ApiResponse{
					Status:  common.ApiResponse_ERROR,
					Code:    errMap.GrpcCodeToHTTPStatusString(err),
					Details: err.Error(),
				}
				return response, err
			}

			response.Msg = &common.ApiResponse{
				Status:  common.ApiResponse_ERROR,
				Code:    errMap.GrpcCodeToHTTPStatusString(err),
				Details: err.Error(),
			}
			if i == 2 {
				return response, nil
			}
			time.Sleep(2 * time.Second)
			creation.CreationId = tools.GetSnowId()
		} else {
			break
		}
	}
	creationId := creation.GetCreationId()

	// 异步处理
	if status == generated.CreationStatus_PENDING {
		go func(creationId int64, ctx context.Context) {
			traceId, fullName := tools.GetMetadataValue(ctx, "trace-id"), tools.GetMetadataValue(ctx, "full-name")
			err = messaging.SendMessage(ctx, messaging.PendingCreation, messaging.PendingCreation, &common.CreationId{
				Id: creationId,
			})
			if err != nil {
				tools.LogError(traceId, fullName, err)
				updateStatus := &generated.CreationUpdateStatus{
					CreationId: creation.GetCreationId(),
					AuthorId:   baseInfo.GetAuthorId(),
					Status:     generated.CreationStatus_DRAFT,
				}
				errSecond := db.UpdateCreationStatusInTransaction(ctx, updateStatus)
				if errSecond != nil {
					tools.LogError(traceId, fullName, errSecond)
				}
			}
		}(creationId, ctx)
	}

	response.Msg = &common.ApiResponse{
		Status: common.ApiResponse_PENDING,
		Code:   "201",
	}
	return response, nil
}

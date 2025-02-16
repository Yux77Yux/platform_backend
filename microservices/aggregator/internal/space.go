package internal

import (
	"context"
	"fmt"

	generated "github.com/Yux77Yux/platform_backend/generated/aggregator"
	common "github.com/Yux77Yux/platform_backend/generated/common"
	creation "github.com/Yux77Yux/platform_backend/generated/creation"
	client "github.com/Yux77Yux/platform_backend/microservices/aggregator/client"
)

func Space(ctx context.Context, req *generated.SpaceRequest) (*generated.SpaceResponse, error) {
	response := new(generated.SpaceResponse)
	userId := req.GetUserId()

	user_client, err := client.NewUserClient()
	if err != nil {
		err = fmt.Errorf("error: user client %w", err)
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "500",
			Details: err.Error(),
		}
		return response, nil
	}
	user, err := user_client.GetUser(ctx, userId)
	if err != nil {
		var msg *common.ApiResponse
		if user != nil {
			msg = user.GetMsg()
		} else {
			msg = &common.ApiResponse{
				Code:    "500",
				Status:  common.ApiResponse_ERROR,
				Details: err.Error(),
			}
		}
		response.Msg = msg
		return response, err
	}

	creation_client, err := client.NewCreationClient()
	if err != nil {
		err = fmt.Errorf("error: creation client %w", err)
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "500",
			Details: err.Error(),
		}
		return response, nil
	}
	creation_list, err := creation_client.GetSpaceCreations(ctx, &creation.GetSpaceCreationsRequest{
		UserId: userId,
	})
	if err != nil {
		var msg *common.ApiResponse
		if creation_list != nil {
			msg = creation_list.GetMsg()
		} else {
			msg = &common.ApiResponse{
				Code:    "500",
				Status:  common.ApiResponse_ERROR,
				Details: err.Error(),
			}
		}
		response.Msg = msg
		return response, err
	}

	// 组装开始

	response.User = user
	response.CreationList = creation_list
	response.Msg = &common.ApiResponse{
		Status:  common.ApiResponse_SUCCESS,
		Code:    "200",
		Message: "Space Request success",
	}
	// 组装完成返回至前端
	return response, nil
}

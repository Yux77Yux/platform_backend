package internal

import (
	"context"
	"fmt"
	"sync"

	generated "github.com/Yux77Yux/platform_backend/generated/aggregator"
	common "github.com/Yux77Yux/platform_backend/generated/common"
	user "github.com/Yux77Yux/platform_backend/generated/user"
	client "github.com/Yux77Yux/platform_backend/microservices/aggregator/client"
	jwt "github.com/Yux77Yux/platform_backend/pkg/jwt"
)

// func aggUserGetUser(msg *amqp.Delivery) (*user.GetUserResponse, error) {
// 	response := new(user.GetUserResponse)
// 	err := proto.Unmarshal(msg.Body, response)
// 	if err != nil {
// 		log.Printf("Error unmarshaling message: %v", err)
// 		return nil, fmt.Errorf("unmarshaling message: %w", err)
// 	}

// 	// 判断返回体有无内部错误信息
// 	status := response.GetMsg().GetStatus()
// 	if status == common.ApiResponse_ERROR || status == common.ApiResponse_FAILED {
// 		return response, fmt.Errorf(response.GetMsg().GetDetails())
// 	}

// 	return response, nil
// }

func Space(ctx context.Context, req *generated.SpaceRequest) (*generated.SpaceResponse, error) {
	// 等待所有异步服务返回
	var wg sync.WaitGroup

	master := false
	userId := req.GetUserId()
	accessToken := req.GetAccessToken()
	if accessToken.Value != "none" {
		accessClaims, err := jwt.ParseJWT(accessToken.GetValue())
		if err != nil {
			return &generated.SpaceResponse{
				Msg: &common.ApiResponse{
					Status:  common.ApiResponse_ERROR,
					Code:    "500",
					Message: "user client failed",
					Details: err.Error(),
				},
			}, fmt.Errorf("error: user client %v", err)
		}

		// 是否为登录用户主页
		master = accessClaims.UserID == userId
	}

	// 从消息队列取用户信息
	userResultCh := make(chan struct {
		user *user.GetUserResponse
		err  error
	}, 1)
	wg.Add(1)
	go func(userId int64, accessToken *common.AccessToken) {
		defer wg.Done()

		user_client, err := client.NewUserClient()
		if err != nil {
			err = fmt.Errorf("error: user client %w", err)
			userResultCh <- struct {
				user *user.GetUserResponse
				err  error
			}{user: nil, err: err}
			return
		}

		result, err := user_client.GetUser(ctx, userId, accessToken)
		if err != nil {
			err = fmt.Errorf("error: user client during GetUser, %w", err)
			userResultCh <- struct {
				user *user.GetUserResponse
				err  error
			}{user: nil, err: err}
			return
		}

		userResultCh <- struct {
			user *user.GetUserResponse
			err  error
		}{user: result, err: nil}

		// userReq := &user.GetUserRequest{
		// 	UserId:      userId,
		// 	AccessToken: accessToken,
		// }
		// reqId := uuid.New().String()
		// msg, err := mq.RPCPattern("agg_user", "getUser", "getUser", reqId, userReq)
		// if err != nil {
		// 	userResultCh <- struct {
		// 		user *user.GetUserResponse
		// 		err  error
		// 	}{err: err}
		// 	return
		// }

		// result, err := aggUserGetUser(msg)
		// userResultCh <- struct {
		// 	user *user.GetUserResponse
		// 	err  error
		// }{user: result, err: err}
		// return
	}(userId, accessToken)

	// 其他异步服务,如获得空间作品
	// 尚未定义
	//

	// 等待完成
	wg.Wait()

	// user service 部分错误检查
	userResult := <-userResultCh
	if userResult.err != nil {
		return &generated.SpaceResponse{
			Msg: userResult.user.GetMsg(),
		}, userResult.err
	}
	// 其他response

	// 组装开始
	user_response := userResult.user

	// 组装完成返回至前端
	return &generated.SpaceResponse{
		User:   user_response.GetUser(),
		Master: master,
		Block:  user_response.GetBlock(),
		Msg: &common.ApiResponse{
			Status:  common.ApiResponse_SUCCESS,
			Code:    "200",
			Message: "Space Request success",
		},
	}, nil
}

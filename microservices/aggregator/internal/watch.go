package internal

import (
	"context"
	"fmt"

	generated "github.com/Yux77Yux/platform_backend/generated/aggregator"
	comment "github.com/Yux77Yux/platform_backend/generated/comment"
	common "github.com/Yux77Yux/platform_backend/generated/common"
	creation "github.com/Yux77Yux/platform_backend/generated/creation"
	client "github.com/Yux77Yux/platform_backend/microservices/aggregator/client"
	tools "github.com/Yux77Yux/platform_backend/microservices/aggregator/tools"
)

func WatchCreation(ctx context.Context, req *generated.WatchCreationRequest) (*generated.WatchCreationResponse, error) {
	response := new(generated.WatchCreationResponse)
	id := req.GetCreationId()

	creation_client, err := client.GetCreationClient()
	if err != nil {
		err = fmt.Errorf("error: creation client %w", err)
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "500",
			Details: err.Error(),
		}
		return response, err
	}
	creationResponse, err := creation_client.GetCreation(ctx, &creation.GetCreationRequest{
		CreationId: id,
	})
	if err != nil {
		var msg *common.ApiResponse
		if creationResponse != nil {
			msg = creationResponse.GetMsg()
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

	msg := creationResponse.GetMsg()
	status := msg.GetStatus()
	code := msg.GetCode()
	if status != common.ApiResponse_SUCCESS && status != common.ApiResponse_PENDING {
		response.Msg = msg
		if code[0] == '5' {
			return response, err
		}
		return response, nil
	}
	creationInfo := creationResponse.GetCreationInfo()
	response.CreationInfo = creationInfo
	userId := creationInfo.GetCreation().GetBaseInfo().GetAuthorId()

	user_client, err := client.GetUserClient()
	if err != nil {
		err = fmt.Errorf("error: user client %w", err)
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "500",
			Details: err.Error(),
		}
		return response, err
	}
	userResponse, err := user_client.GetUser(ctx, userId)
	if err != nil {
		var msg *common.ApiResponse
		if userResponse != nil {
			msg = userResponse.GetMsg()
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

	msg = userResponse.GetMsg()
	status = msg.GetStatus()
	code = msg.GetCode()
	if status != common.ApiResponse_SUCCESS && status != common.ApiResponse_PENDING {
		response.Msg = msg
		if code[0] == '5' {
			return response, err
		}
		return response, nil
	}

	go func(id int64, ctx context.Context) {
		ipv4 := tools.GetMetadataValue(ctx, "x-forwarded-for")
		if ipv4 == "" {
			return
		}
		traceId, fullName := tools.GetMetadataValue(ctx, "trace-id"), tools.GetMetadataValue(ctx, "full-name")
		err := messaging.SendMessage(
			ctx,
			EXCHANGE_INCREASE_VIEW,
			KEY_INCREASE_VIEW,
			&common.ViewCreation{
				Id:   id,
				Ipv4: ipv4,
			},
		)
		if err != nil {
			tools.LogError(traceId, fullName, err)
		}
	}(id, ctx)

	// 组装开始
	user := userResponse.GetUser()
	userDefault := user.GetUserDefault()
	response.CreationUser = &common.UserCreationComment{
		UserDefault: &common.UserDefault{
			UserId:   userDefault.GetUserId(),
			UserName: userDefault.GetUserName(),
		},
		UserAvatar: user.GetUserAvatar(),
		UserBio:    user.GetUserBio(),
		Followers:  user.GetFollowers(),
	}
	response.Msg = &common.ApiResponse{
		Status: common.ApiResponse_SUCCESS,
		Code:   "200",
	}
	// 组装完成返回至前端
	return response, nil
}

func InitialComments(ctx context.Context, req *generated.InitialCommentsRequest) (*generated.InitialCommentsResponse, error) {
	response := new(generated.InitialCommentsResponse)
	request := req.GetRequest()
	creationId := request.GetCreationId()

	comment_client, err := client.GetCommentClient()
	if err != nil {
		err = fmt.Errorf("error: comment client %w", err)
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "500",
			Details: err.Error(),
		}
		return response, err
	}
	commentsResponse, err := comment_client.InitialComments(ctx, &comment.InitialCommentsRequest{
		CreationId: creationId,
	})
	if err != nil {
		var msg *common.ApiResponse
		if commentsResponse != nil {
			msg = commentsResponse.GetMsg()
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

	msg := commentsResponse.GetMsg()
	status := msg.GetStatus()
	code := msg.GetCode()
	if status != common.ApiResponse_SUCCESS && status != common.ApiResponse_PENDING {
		response.Msg = msg
		if code[0] == '5' {
			return response, err
		}
		return response, nil
	}

	comments := commentsResponse.GetComments()
	comments_len := len(comments)
	if comments_len <= 0 {
		response.Msg = &common.ApiResponse{
			Code:    "404",
			Status:  common.ApiResponse_ERROR,
			Details: "no comment",
		}
		return response, nil
	}

	userIds := make([]int64, 0, comments_len)
	for i := 0; i < comments_len; i++ {
		userIds = append(userIds, comments[i].Comment.GetUserId())
	}

	userMap, err := getUsers(ctx, userIds)
	if err != nil {
		err = fmt.Errorf("error: user client %w", err)
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "500",
			Details: err.Error(),
		}
		return response, err
	}
	if len(userMap) <= 0 {
		response.Msg = &common.ApiResponse{
			Code:    "404",
			Status:  common.ApiResponse_ERROR,
			Details: "no user",
		}
		return response, nil
	}

	userMap[0] = &common.UserCreationComment{
		UserDefault: &common.UserDefault{},
	}
	userMap[-1] = &common.UserCreationComment{
		UserDefault: &common.UserDefault{},
	}
	topComments := make([]*generated.TopCommentInfo, comments_len)
	for i := 0; i < comments_len; i++ {
		comment := comments[i]
		topComments[i] = &generated.TopCommentInfo{
			CommentUser: userMap[comment.Comment.GetUserId()],
			TopComment:  comment,
		}
	}

	response.Comments = topComments
	response.Area = commentsResponse.GetCommentArea()
	response.PageCount = commentsResponse.GetPageCount()
	response.Msg = &common.ApiResponse{
		Status: common.ApiResponse_SUCCESS,
		Code:   "200",
	}
	return response, nil
}

func GetTopComments(ctx context.Context, req *generated.GetTopCommentsRequest) (*generated.GetTopCommentsResponse, error) {
	response := new(generated.GetTopCommentsResponse)
	request := req.GetRequest()
	creationId := request.GetCreationId()

	comment_client, err := client.GetCommentClient()
	if err != nil {
		err = fmt.Errorf("error: comment client %w", err)
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "500",
			Details: err.Error(),
		}
		return response, err
	}
	commentsResponse, err := comment_client.GetTopComments(ctx, &comment.GetTopCommentsRequest{
		CreationId: creationId,
		Page:       request.GetPage(),
	})
	if err != nil {
		var msg *common.ApiResponse
		if commentsResponse != nil {
			msg = commentsResponse.GetMsg()
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

	msg := commentsResponse.GetMsg()
	status := msg.GetStatus()
	code := msg.GetCode()
	if status != common.ApiResponse_SUCCESS && status != common.ApiResponse_PENDING {
		response.Msg = msg
		if code[0] == '5' {
			return response, err
		}
		return response, nil
	}

	comments := commentsResponse.GetComments()
	comments_len := len(comments)
	if comments_len <= 0 {
		response.Msg = &common.ApiResponse{
			Code:    "404",
			Status:  common.ApiResponse_ERROR,
			Details: "no comment",
		}
		return response, nil
	}

	userIds := make([]int64, 0, comments_len)
	for i := 0; i < comments_len; i++ {
		userIds = append(userIds, comments[i].Comment.GetUserId())
	}

	userMap, err := getUsers(ctx, userIds)
	if err != nil {
		err = fmt.Errorf("error: user client %w", err)
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "500",
			Details: err.Error(),
		}
		return response, err
	}
	if len(userMap) <= 0 {
		response.Msg = &common.ApiResponse{
			Code:    "404",
			Status:  common.ApiResponse_ERROR,
			Details: "no user",
		}
		return response, nil
	}

	userMap[0] = &common.UserCreationComment{
		UserDefault: &common.UserDefault{},
	}
	userMap[-1] = &common.UserCreationComment{
		UserDefault: &common.UserDefault{},
	}
	topComments := make([]*generated.TopCommentInfo, comments_len)
	for i := 0; i < comments_len; i++ {
		comment := comments[i]
		topComments[i] = &generated.TopCommentInfo{
			CommentUser: userMap[comment.Comment.GetUserId()],
			TopComment:  comment,
		}
	}

	response.Comments = topComments
	response.Msg = &common.ApiResponse{
		Status: common.ApiResponse_SUCCESS,
		Code:   "200",
	}
	return response, nil
}

func GetSecondComments(ctx context.Context, req *generated.GetSecondCommentsRequest) (*generated.GetSecondCommentsResponse, error) {
	response := new(generated.GetSecondCommentsResponse)
	request := req.GetRequest()
	creationId := request.GetCreationId()

	comment_client, err := client.GetCommentClient()
	if err != nil {
		err = fmt.Errorf("error: comment client %w", err)
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "500",
			Details: err.Error(),
		}
		return response, err
	}
	commentsResponse, err := comment_client.GetSecondComments(ctx, &comment.GetSecondCommentsRequest{
		CreationId: creationId,
		Root:       request.GetRoot(),
		Page:       request.GetPage(),
	})
	if err != nil {
		var msg *common.ApiResponse
		if commentsResponse != nil {
			msg = commentsResponse.GetMsg()
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

	msg := commentsResponse.GetMsg()
	status := msg.GetStatus()
	code := msg.GetCode()
	if status != common.ApiResponse_SUCCESS && status != common.ApiResponse_PENDING {
		response.Msg = msg
		if code[0] == '5' {
			return response, err
		}
		return response, nil
	}

	comments := commentsResponse.GetComments()
	comments_len := len(comments)
	if comments_len <= 0 {
		response.Msg = &common.ApiResponse{
			Code:    "404",
			Status:  common.ApiResponse_ERROR,
			Details: "no comment",
		}
		return response, nil
	}

	userIds := make([]int64, 0, comments_len*2)
	for i := 0; i < comments_len; i++ {
		replyUserId := comments[i].GetReplyUserId()
		if replyUserId > 0 {
			userIds = append(userIds, replyUserId)
		}
		userIds = append(userIds, comments[i].Comment.GetUserId())
	}

	userMap, err := getUsers(ctx, userIds)
	if err != nil {
		err = fmt.Errorf("error: user client %w", err)
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "500",
			Details: err.Error(),
		}
		return response, err
	}
	if len(userMap) <= 0 {
		response.Msg = &common.ApiResponse{
			Code:    "404",
			Status:  common.ApiResponse_ERROR,
			Details: "no user",
		}
		return response, nil
	}

	userMap[0] = &common.UserCreationComment{
		UserDefault: &common.UserDefault{},
	}
	userMap[-1] = &common.UserCreationComment{
		UserDefault: &common.UserDefault{},
	}
	secondComments := make([]*generated.SecondCommentInfo, comments_len)
	for i := 0; i < comments_len; i++ {
		comment := comments[i]
		secondComments[i] = &generated.SecondCommentInfo{
			ReplyUser: userMap[comment.GetReplyUserId()],
		}
		secondComments[i].SecondComment = &generated.CommentInfo{
			CommentUser: userMap[comment.Comment.GetUserId()],
			Comment:     comment.GetComment(),
		}
	}

	response.Comments = secondComments
	response.Msg = &common.ApiResponse{
		Status: common.ApiResponse_SUCCESS,
		Code:   "200",
	}
	return response, nil
}

func getUsers(ctx context.Context, userIds []int64) (map[int64]*common.UserCreationComment, error) {
	userMap := make(map[int64]*common.UserCreationComment)

	user_client, err := client.GetUserClient()
	if err != nil {
		return nil, err
	}
	userResponse, err := user_client.GetUsers(ctx, userIds)
	if err != nil {
		return nil, err
	}
	if userResponse.Msg.GetStatus() != common.ApiResponse_SUCCESS {
		return nil, fmt.Errorf("error: %s", userResponse.Msg.GetDetails())
	}

	users := userResponse.GetUsers()
	if len(users) <= 0 {
		return nil, nil
	}

	for _, user := range users {
		userId := user.GetUserDefault().GetUserId()
		userMap[userId] = user
	}

	return userMap, nil
}

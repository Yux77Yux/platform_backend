package internal

import (
	"context"
	"fmt"

	generated "github.com/Yux77Yux/platform_backend/generated/aggregator"
	comment "github.com/Yux77Yux/platform_backend/generated/comment"
	common "github.com/Yux77Yux/platform_backend/generated/common"
	creation "github.com/Yux77Yux/platform_backend/generated/creation"
	interaction "github.com/Yux77Yux/platform_backend/generated/interaction"
	client "github.com/Yux77Yux/platform_backend/microservices/aggregator/client"
)

func WatchCreation(ctx context.Context, req *generated.WatchCreationRequest) (*generated.WatchCreationResponse, error) {
	response := new(generated.WatchCreationResponse)
	id := req.GetCreationId()

	creation_client, err := client.NewCreationClient()
	if err != nil {
		err = fmt.Errorf("error: user client %w", err)
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "500",
			Details: err.Error(),
		}
		return response, nil
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

	creationInfo := creationResponse.GetCreationInfo()
	userId := creationInfo.GetCreation().GetBaseInfo().GetAuthorId()
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

	// 组装开始
	user := userResponse.GetUser()

	response.CreationInfo = creationInfo
	response.CreationUser = &common.UserCreationComment{
		UserDefault: &common.UserDefault{
			UserId:   user.GetUserDefault().GetUserId(),
			UserName: user.GetUserDefault().GetUserName(),
		},
		UserAvatar: user.GetUserAvatar(),
		UserBio:    user.GetUserBio(),
	}
	response.Msg = &common.ApiResponse{
		Status: common.ApiResponse_SUCCESS,
		Code:   "200",
	}
	// 组装完成返回至前端
	return response, nil
}

func SimilarCreations(ctx context.Context, req *generated.SimilarCreationsRequest) (*generated.GetCardsResponse, error) {
	response := new(generated.GetCardsResponse)
	id := req.GetCreationId()

	// 从 用户数据服务 调取相似列表
	interaction_client, err := client.NewInteractionClient()
	if err != nil {
		err = fmt.Errorf("error: interaction client %w", err)
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "500",
			Details: err.Error(),
		}
		return response, nil
	}
	interactionResponse, err := interaction_client.GetRecommendBaseCreation(ctx, &interaction.GetRecommendRequest{
		Id: id,
	})
	if err != nil {
		var msg *common.ApiResponse
		if interactionResponse != nil {
			msg = interactionResponse.GetMsg()
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
	creationIds := interactionResponse.GetCreations()

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
	creationResponse, err := creation_client.GetPublicCreationList(ctx, &creation.GetCreationListRequest{
		Ids: creationIds,
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

	creationInfos := creationResponse.GetCreationInfoGroup()

	length := len(creationInfos)
	userIds := make([]int64, length)
	for i, info := range creationInfos {
		userIds[i] = info.GetCreation().GetBaseInfo().GetAuthorId()
	}

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
	userResponse, err := user_client.GetUsers(ctx, userIds)
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

	// 构建 userId -> 用户信息的映射表
	users := userResponse.GetUsers()
	limit := len(users)
	userMap := make(map[int64]*common.UserDefault, limit)
	for _, user := range users {
		userMap[user.GetUserDefault().GetUserId()] = user.GetUserDefault()
	}

	cards := make([]*generated.CreationCard, 0, length)
	for _, info := range creationInfos {
		creation := info.GetCreation()
		authorId := creation.GetBaseInfo().GetAuthorId()
		card := &generated.CreationCard{
			Creation:           creation,
			CreationEngagement: info.GetCreationEngagement(),
		}
		if user, exists := userMap[authorId]; exists {
			card.User = user
		}
		cards = append(cards, card)
	}

	response.Cards = cards
	response.Msg = &common.ApiResponse{
		Status: common.ApiResponse_SUCCESS,
		Code:   "200",
	}
	return response, nil
}

func InitialComments(ctx context.Context, req *generated.InitialCommentsRequest) (*generated.InitialCommentsResponse, error) {
	response := new(generated.InitialCommentsResponse)
	request := req.GetRequest()
	creationId := request.GetCreationId()

	comment_client, err := client.NewCommentClient()
	if err != nil {
		err = fmt.Errorf("error: comment client %w", err)
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "500",
			Details: err.Error(),
		}
		return response, nil
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

	comments := commentsResponse.GetComments()
	cards, err := getCards(ctx, comments)
	if err != nil {
		err = fmt.Errorf("error: user client %w", err)
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "500",
			Details: err.Error(),
		}
		return response, nil
	}

	response.Area = commentsResponse.GetCommentArea()
	response.Comments = cards
	response.Msg = &common.ApiResponse{
		Status: common.ApiResponse_SUCCESS,
		Code:   "200",
	}
	return response, nil
}

func GetTopComments(ctx context.Context, req *generated.GetTopCommentsRequest) (*generated.GetCommentsResponse, error) {
	response := new(generated.GetCommentsResponse)
	request := req.GetRequest()
	creationId := request.GetCreationId()

	comment_client, err := client.NewCommentClient()
	if err != nil {
		err = fmt.Errorf("error: comment client %w", err)
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "500",
			Details: err.Error(),
		}
		return response, nil
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

	comments := commentsResponse.GetComments()
	cards, err := getCards(ctx, comments)
	if err != nil {
		err = fmt.Errorf("error: user client %w", err)
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "500",
			Details: err.Error(),
		}
		return response, nil
	}

	response.Comments = cards
	response.Msg = &common.ApiResponse{
		Status: common.ApiResponse_SUCCESS,
		Code:   "200",
	}
	return response, nil
}

func GetSecondComments(ctx context.Context, req *generated.GetSecondCommentsRequest) (*generated.GetCommentsResponse, error) {
	response := new(generated.GetCommentsResponse)
	request := req.GetRequest()
	creationId := request.GetCreationId()

	comment_client, err := client.NewCommentClient()
	if err != nil {
		err = fmt.Errorf("error: comment client %w", err)
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "500",
			Details: err.Error(),
		}
		return response, nil
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

	comments := commentsResponse.GetComments()
	cards, err := getCards(ctx, comments)
	if err != nil {
		err = fmt.Errorf("error: user client %w", err)
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "500",
			Details: err.Error(),
		}
		return response, nil
	}

	response.Comments = cards
	response.Msg = &common.ApiResponse{
		Status: common.ApiResponse_SUCCESS,
		Code:   "200",
	}
	return response, nil
}

func getCards(ctx context.Context, comments []*comment.Comment) ([]*generated.CommentInfo, error) {
	length := len(comments)
	userMap := make(map[int64]*common.UserCreationComment)

	for _, comment := range comments {
		userId := comment.GetUserId()
		userMap[userId] = &common.UserCreationComment{
			UserDefault: &common.UserDefault{
				UserId: userId,
			},
		}
	}
	userIds := make([]int64, 0, len(userMap))
	for id := range userMap {
		userIds = append(userIds, id)
	}

	user_client, err := client.NewUserClient()
	if err != nil {
		return nil, err
	}
	userResponse, err := user_client.GetUsers(ctx, userIds)
	if err != nil {
		return nil, err
	}
	users := userResponse.GetUsers()

	for _, user := range users {
		userId := user.GetUserDefault().GetUserId()
		userMap[userId] = user
	}

	cards := make([]*generated.CommentInfo, length)
	for i, comment := range comments {
		userId := comment.GetUserId()

		card := &generated.CommentInfo{
			Comment:     comment,
			CommentUser: userMap[userId],
		}

		cards[i] = card
	}

	return cards, nil
}

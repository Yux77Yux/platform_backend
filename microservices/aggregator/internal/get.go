package internal

import (
	"context"
	"fmt"

	generated "github.com/Yux77Yux/platform_backend/generated/aggregator"
	common "github.com/Yux77Yux/platform_backend/generated/common"
	creation "github.com/Yux77Yux/platform_backend/generated/creation"
	interaction "github.com/Yux77Yux/platform_backend/generated/interaction"
	client "github.com/Yux77Yux/platform_backend/microservices/aggregator/client"
	auth "github.com/Yux77Yux/platform_backend/pkg/auth"
)

func HomePage(ctx context.Context, req *generated.HomeRequest) (*generated.GetCardsResponse, error) {
	response := new(generated.GetCardsResponse)
	token := req.GetAccessToken().GetValue()
	pass, userId, err := auth.Auth("get", "interaction", token)
	if err != nil {
		response.Msg = &common.ApiResponse{
			Code:    "500",
			Status:  common.ApiResponse_ERROR,
			Details: err.Error(),
		}
		return response, err
	}
	if !pass {
		response.Msg = &common.ApiResponse{
			Code:   "403",
			Status: common.ApiResponse_ERROR,
		}
		return response, nil
	}

	// 从 用户数据服务 调取相似列表
	interaction_client, err := client.GetInteractionClient()
	if err != nil {
		err = fmt.Errorf("error: interaction client %w", err)
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "500",
			Details: err.Error(),
		}
		return response, err
	}
	interactionResponse, err := interaction_client.GetRecommendBaseUser(ctx, &interaction.GetRecommendRequest{
		Id: userId,
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

	msg := interactionResponse.GetMsg()
	code := msg.GetCode()
	status := msg.GetStatus()
	if status != common.ApiResponse_SUCCESS {
		response.Msg = interactionResponse.Msg
		if code[0] == '5' {
			return response, err
		}
		return response, nil
	}

	creationIds := interactionResponse.GetCreations()
	userMap, creationInfos, err := getUserMap(ctx, creationIds)
	if err != nil {
		response.Msg = &common.ApiResponse{
			Code:    "500",
			Status:  common.ApiResponse_ERROR,
			Details: err.Error(),
		}
		return response, err
	}
	if len(userMap) <= 0 || len(creationInfos) <= 0 {
		err := fmt.Errorf("error: no member in data")
		response.Msg = &common.ApiResponse{
			Code:    "404",
			Status:  common.ApiResponse_ERROR,
			Details: err.Error(),
		}
		return response, nil
	}

	cards := make([]*generated.CreationCard, 0, len(creationInfos))
	for _, info := range creationInfos {
		creation := info.GetCreation()
		authorId := creation.GetBaseInfo().GetAuthorId()
		engagement := info.GetCreationEngagement()
		card := &generated.CreationCard{
			Creation:           creation,
			CreationEngagement: engagement,
			TimeAt:             engagement.GetPublishTime(),
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

func Collections(ctx context.Context, req *generated.CollectionsRequest) (*generated.GetCardsResponse, error) {
	response := new(generated.GetCardsResponse)
	token := req.GetAccessToken().GetValue()
	pass, userId, err := auth.Auth("get", "interaction", token)
	if err != nil {
		response.Msg = &common.ApiResponse{
			Code:    "500",
			Status:  common.ApiResponse_ERROR,
			Details: err.Error(),
		}
		return response, err
	}
	if !pass {
		response.Msg = &common.ApiResponse{
			Code:   "403",
			Status: common.ApiResponse_ERROR,
		}
		return response, nil
	}

	page := req.GetPage()
	// 从 用户数据服务 调取相似列表
	interaction_client, err := client.GetInteractionClient()
	if err != nil {
		err = fmt.Errorf("error: interaction client %w", err)
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "500",
			Details: err.Error(),
		}
		return response, err
	}
	interactionResponse, err := interaction_client.GetCollections(ctx, &interaction.GetCollectionsRequest{
		UserId: userId,
		Page:   page,
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

	msg := interactionResponse.GetMsg()
	code := msg.GetCode()
	status := msg.GetStatus()
	if status != common.ApiResponse_SUCCESS {
		response.Msg = interactionResponse.Msg
		if code[0] == '5' {
			return response, err
		}
		return response, nil
	}

	interactions := interactionResponse.GetAnyInteraction().GetAnyInterction()
	length := len(interactions)
	if length <= 0 {
		err := fmt.Errorf("error: no interactions")
		response.Msg = &common.ApiResponse{
			Code:    "404",
			Status:  common.ApiResponse_ERROR,
			Details: err.Error(),
		}
		return response, nil
	}
	creationIds := make([]int64, length)
	for i, _interaction := range interactions {
		id := _interaction.Base.GetCreationId()
		creationIds[i] = id
	}

	userMap, creationInfos, err := getUserMap(ctx, creationIds)
	if err != nil {
		response.Msg = &common.ApiResponse{
			Code:    "500",
			Status:  common.ApiResponse_ERROR,
			Details: err.Error(),
		}
		return response, err
	}
	if len(userMap) <= 0 || len(creationInfos) <= 0 {
		err := fmt.Errorf("error: no member in data")
		response.Msg = &common.ApiResponse{
			Code:    "404",
			Status:  common.ApiResponse_ERROR,
			Details: err.Error(),
		}
		return response, nil
	}

	creationMap := make(map[int64]*creation.CreationInfo)
	for _, info := range creationInfos {
		creation := info.GetCreation()
		creationId := creation.GetCreationId()
		creationMap[creationId] = info
	}

	cards := make([]*generated.CreationCard, 0, length)
	for _, interaction := range interactions {
		creationId := interaction.Base.GetCreationId()
		info := creationMap[creationId]

		creation := info.GetCreation()
		authorId := creation.GetBaseInfo().GetAuthorId()

		card := &generated.CreationCard{
			Creation:           creation,
			CreationEngagement: info.GetCreationEngagement(),
			TimeAt:             interaction.GetSaveAt(),
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

func History(ctx context.Context, req *generated.HistoryRequest) (*generated.GetCardsResponse, error) {
	response := new(generated.GetCardsResponse)
	token := req.GetAccessToken().GetValue()
	pass, userId, err := auth.Auth("get", "interaction", token)
	if err != nil {
		response.Msg = &common.ApiResponse{
			Code:    "500",
			Status:  common.ApiResponse_ERROR,
			Details: err.Error(),
		}
		return response, err
	}
	if !pass {
		response.Msg = &common.ApiResponse{
			Code:   "403",
			Status: common.ApiResponse_ERROR,
		}
		return response, nil
	}

	// 从 用户数据服务 调取相似列表
	interaction_client, err := client.GetInteractionClient()
	if err != nil {
		err = fmt.Errorf("error: interaction client %w", err)
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "500",
			Details: err.Error(),
		}
		return response, nil
	}
	interactionResponse, err := interaction_client.GetHistories(ctx, &interaction.GetHistoriesRequest{
		UserId: userId,
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

	msg := interactionResponse.GetMsg()
	code := msg.GetCode()
	status := msg.GetStatus()
	if status != common.ApiResponse_SUCCESS {
		response.Msg = interactionResponse.Msg
		if code[0] == '5' {
			return response, err
		}
		return response, nil
	}

	interactions := interactionResponse.GetAnyInteraction().GetAnyInterction()

	length := len(interactions)
	if length <= 0 {
		err := fmt.Errorf("error: no interactions")
		response.Msg = &common.ApiResponse{
			Code:    "404",
			Status:  common.ApiResponse_ERROR,
			Details: err.Error(),
		}
		return response, nil
	}
	creationIds := make([]int64, length)
	creationMap := make(map[int64]*creation.CreationInfo)
	for i, _interaction := range interactions {
		id := _interaction.Base.GetCreationId()
		creationIds[i] = id
		creationMap[id] = &creation.CreationInfo{
			Creation: &creation.Creation{
				CreationId: id,
			},
			CreationEngagement: &creation.CreationEngagement{
				CreationId: id,
			},
		}
	}

	userMap, creationInfos, err := getUserMap(ctx, creationIds)
	if err != nil {
		response.Msg = &common.ApiResponse{
			Code:    "500",
			Status:  common.ApiResponse_ERROR,
			Details: err.Error(),
		}
		return response, err
	}
	if len(userMap) <= 0 || len(creationInfos) <= 0 {
		err := fmt.Errorf("error: no member in data")
		response.Msg = &common.ApiResponse{
			Code:    "404",
			Status:  common.ApiResponse_ERROR,
			Details: err.Error(),
		}
		return response, nil
	}

	for _, info := range creationInfos {
		creation := info.GetCreation()
		creationId := creation.GetCreationId()

		if _, exists := creationMap[creationId]; exists {
			creationMap[creationId] = info
		}
	}

	cards := make([]*generated.CreationCard, length)
	for i := 0; i < length; i++ {
		info := creationMap[creationIds[i]]
		creation := info.GetCreation()
		authorId := creation.GetBaseInfo().GetAuthorId()

		card := &generated.CreationCard{
			Creation:           creation,
			CreationEngagement: info.GetCreationEngagement(),
			TimeAt:             interactions[i].GetUpdatedAt(),
		}
		if user, exists := userMap[authorId]; exists {
			card.User = user
		}
		cards[i] = card
	}

	response.Cards = cards
	response.Msg = &common.ApiResponse{
		Status: common.ApiResponse_SUCCESS,
		Code:   "200",
	}
	return response, nil
}

func getUserMap(ctx context.Context, creationIds []int64) (map[int64]*common.UserDefault, []*creation.CreationInfo, error) {
	creation_client, err := client.GetCreationClient()
	if err != nil {
		err = fmt.Errorf("error: creation client %w", err)
		return nil, nil, err
	}
	creationResponse, err := creation_client.GetPublicCreationList(ctx, &creation.GetCreationListRequest{
		Ids: creationIds,
	})
	if err != nil {
		return nil, nil, err
	}

	msg := creationResponse.GetMsg()
	code := msg.GetCode()
	status := msg.GetStatus()
	if status != common.ApiResponse_SUCCESS {
		if code[0] == '5' {
			return nil, nil, fmt.Errorf("error: %s", msg.GetDetails())
		}
		return nil, nil, nil
	}

	creationInfos := creationResponse.GetCreationInfoGroup()

	length := len(creationInfos)
	if length <= 0 {
		return nil, nil, nil
	}
	userIdMap := make(map[int64]struct{})
	for _, info := range creationInfos {
		userIdMap[info.GetCreation().GetBaseInfo().GetAuthorId()] = struct{}{}
	}
	userIds := make([]int64, 0, len(userIdMap))
	for id := range userIdMap {
		userIds = append(userIds, id)
	}

	user_client, err := client.GetUserClient()
	if err != nil {
		err = fmt.Errorf("error: user client %w", err)
		return nil, nil, err
	}
	userResponse, err := user_client.GetUsers(ctx, userIds)
	if err != nil {
		return nil, nil, err
	}

	msg = userResponse.GetMsg()
	code = msg.GetCode()
	status = msg.GetStatus()
	if status != common.ApiResponse_SUCCESS {
		if code[0] == '5' {
			return nil, nil, fmt.Errorf("error: %s", msg.GetDetails())
		}
		return nil, nil, nil
	}

	// 构建 userId -> 用户信息的映射表
	users := userResponse.GetUsers()
	if len(users) <= 0 {
		return nil, nil, nil
	}
	limit := len(users)
	userMap := make(map[int64]*common.UserDefault, limit)
	for _, user := range users {
		userDef := user.GetUserDefault()
		userMap[userDef.GetUserId()] = userDef
	}

	return userMap, creationInfos, nil
}

func Search(ctx context.Context, req *generated.SearchCreationsRequest) (*generated.SearchCreationsResponse, error) {
	response := new(generated.SearchCreationsResponse)

	creation_client, err := client.GetCreationClient()
	if err != nil {
		response.Msg = &common.ApiResponse{
			Code:    "500",
			Status:  common.ApiResponse_ERROR,
			Details: err.Error(),
		}
		return response, err
	}
	creationResponse, err := creation_client.SearchCreation(ctx, &creation.SearchCreationRequest{
		Title: req.GetTitle(),
		Page:  req.GetPage(),
	})
	if err != nil {
		response.Msg = &common.ApiResponse{
			Code:    "500",
			Status:  common.ApiResponse_ERROR,
			Details: err.Error(),
		}
		return response, err
	}

	msg := creationResponse.GetMsg()
	status := msg.GetStatus()
	if status != common.ApiResponse_SUCCESS {
		response.Msg = &common.ApiResponse{
			Code:    msg.GetCode(),
			Status:  common.ApiResponse_ERROR,
			Details: msg.GetDetails(),
		}
		return response, nil
	}

	creationInfos := creationResponse.GetCreationInfoGroup()
	userIdMap := make(map[int64]bool)
	for _, info := range creationInfos {
		userIdMap[info.GetCreation().GetBaseInfo().GetAuthorId()] = true
	}
	userIds := make([]int64, 0, len(userIdMap))
	for id := range userIdMap {
		userIds = append(userIds, id)
	}

	user_client, err := client.GetUserClient()
	if err != nil {
		err = fmt.Errorf("error: user client %w", err)
		return nil, err
	}
	userResponse, err := user_client.GetUsers(ctx, userIds)
	if err != nil {
		return nil, err
	}

	msg = userResponse.GetMsg()
	code := msg.GetCode()
	status = msg.GetStatus()
	if status != common.ApiResponse_SUCCESS {
		if code[0] == '5' {
			return nil, fmt.Errorf("error: %s", msg.GetDetails())
		}
		return nil, nil
	}

	users := userResponse.GetUsers()
	if len(users) <= 0 {
		return nil, nil
	}
	limit := len(users)
	userMap := make(map[int64]*common.UserDefault, limit)
	for _, user := range users {
		userMap[user.GetUserDefault().GetUserId()] = user.GetUserDefault()
	}

	length := len(creationInfos)
	cards := make([]*generated.CreationCard, length)
	for i := 0; i < length; i++ {
		info := creationInfos[i]
		creation := info.GetCreation()
		authorId := creation.GetBaseInfo().GetAuthorId()
		CreationEngagement := info.GetCreationEngagement()
		card := &generated.CreationCard{
			Creation:           creation,
			CreationEngagement: CreationEngagement,
			TimeAt:             CreationEngagement.GetPublishTime(),
		}
		if user, exists := userMap[authorId]; exists {
			card.User = user
		}
		cards[i] = card
	}

	response.Cards = cards
	response.Count = creationResponse.GetCount()
	response.Msg = &common.ApiResponse{
		Code:   "200",
		Status: common.ApiResponse_SUCCESS,
	}
	return response, nil
}

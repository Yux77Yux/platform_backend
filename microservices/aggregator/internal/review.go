package internal

import (
	"context"
	"fmt"

	generated "github.com/Yux77Yux/platform_backend/generated/aggregator"
	comment "github.com/Yux77Yux/platform_backend/generated/comment"
	common "github.com/Yux77Yux/platform_backend/generated/common"
	creation "github.com/Yux77Yux/platform_backend/generated/creation"
	review "github.com/Yux77Yux/platform_backend/generated/review"
	client "github.com/Yux77Yux/platform_backend/microservices/aggregator/client"
	auth "github.com/Yux77Yux/platform_backend/pkg/auth"
)

func getUserReviewCards(ctx context.Context, reviews []*review.Review) ([]*generated.GetUserReviewsResponse_UserReview, error) {
	userMap := make(map[int64]*common.UserCreationComment)
	for _, review := range reviews {
		userMap[review.GetNew().GetTargetId()] = &common.UserCreationComment{
			UserDefault: &common.UserDefault{
				UserId: review.GetNew().GetTargetId(),
			},
		}
	}
	length := len(userMap)
	if length <= 0 {
		return nil, nil
	}
	userIds := make([]int64, 0, length)
	for id := range userMap {
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

	msg := userResponse.GetMsg()
	code := msg.GetCode()
	status := msg.GetStatus()
	if status != common.ApiResponse_SUCCESS {
		if code[0] == '5' {
			return nil, fmt.Errorf("error: userResponse %s", msg.GetDetails())
		}
		return nil, nil
	}

	users := userResponse.GetUsers()
	if len(users) <= 0 {
		return nil, nil
	}
	for _, user := range users {
		userId := user.GetUserDefault().GetUserId()
		userMap[userId] = user
	}

	cards := make([]*generated.GetUserReviewsResponse_UserReview, len(reviews))
	for i, review := range reviews {
		userId := review.GetNew().GetTargetId()

		card := &generated.GetUserReviewsResponse_UserReview{
			Review: review,
			User:   userMap[userId],
		}
		cards[i] = card
	}
	return cards, nil
}

func GetUserReviews(ctx context.Context, req *generated.GetReviewsRequest) (*generated.GetUserReviewsResponse, error) {
	response := new(generated.GetUserReviewsResponse)

	token := req.GetAccessToken().GetValue()
	pass, reviewerId, err := auth.Auth("get", "review", token)
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

	review_client, err := client.GetReviewClient()
	if err != nil {
		err = fmt.Errorf("error: review client %w", err)
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "500",
			Details: err.Error(),
		}
		return response, err
	}
	reviewResponse, err := review_client.GetReviews(ctx, &review.GetReviewsRequest{
		Status:     req.GetStatus(),
		Type:       review.TargetType_USER,
		Page:       req.GetPage(),
		ReviewerId: reviewerId,
	})
	if err != nil {
		var msg *common.ApiResponse
		if reviewResponse != nil {
			msg = reviewResponse.GetMsg()
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
	msg := reviewResponse.GetMsg()
	status := msg.GetStatus()
	code := msg.GetCode()
	if status != common.ApiResponse_SUCCESS && status != common.ApiResponse_PENDING {
		response.Msg = msg
		if code[0] == '5' {
			return response, err
		}
		return response, nil
	}

	reviews := reviewResponse.GetReviews()
	cards, err := getUserReviewCards(ctx, reviews)
	if err != nil {
		response.Msg = &common.ApiResponse{
			Status: common.ApiResponse_ERROR,
			Code:   "500",
		}
		return response, err
	}

	response.Reviews = cards
	response.Count = reviewResponse.GetCount()
	response.Msg = &common.ApiResponse{
		Status: common.ApiResponse_SUCCESS,
		Code:   "200",
	}
	return response, nil
}

func GetNewUserReviews(ctx context.Context, req *generated.GetNewReviewsRequest) (*generated.GetUserReviewsResponse, error) {
	response := new(generated.GetUserReviewsResponse)

	token := req.GetAccessToken().GetValue()
	pass, reviewerId, err := auth.Auth("get", "review", token)
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

	review_client, err := client.GetReviewClient()
	if err != nil {
		err = fmt.Errorf("error: review client %w", err)
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "500",
			Details: err.Error(),
		}
		return response, err
	}
	reviewResponse, err := review_client.GetNewReviews(ctx, &review.GetNewReviewsRequest{
		Type:       review.TargetType_USER,
		ReviewerId: reviewerId,
	})
	if err != nil {
		var msg *common.ApiResponse
		if reviewResponse != nil {
			msg = reviewResponse.GetMsg()
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

	msg := reviewResponse.GetMsg()
	status := msg.GetStatus()
	code := msg.GetCode()
	if status != common.ApiResponse_SUCCESS && status != common.ApiResponse_PENDING {
		response.Msg = msg
		if code[0] == '5' {
			return response, err
		}
		return response, nil
	}

	reviews := reviewResponse.GetReviews()
	cards, err := getUserReviewCards(ctx, reviews)
	if err != nil {
		response.Msg = &common.ApiResponse{
			Status: common.ApiResponse_ERROR,
			Code:   "500",
		}
		return response, err
	}

	response.Reviews = cards
	response.Msg = &common.ApiResponse{
		Status: common.ApiResponse_SUCCESS,
		Code:   "200",
	}
	return response, nil
}

func getCreationReviewCards(ctx context.Context, reviews []*review.Review) ([]*generated.GetCreationReviewsResponse_CreationReview, error) {
	creationMap := make(map[int64]*creation.Creation)
	for _, review := range reviews {
		id := review.GetNew().GetTargetId()
		creationMap[id] = &creation.Creation{
			CreationId: id,
		}
	}
	length := len(creationMap)
	if length <= 0 {
		return nil, nil
	}
	ids := make([]int64, 0, length)
	for id := range creationMap {
		ids = append(ids, id)
	}

	creation_client, err := client.GetCreationClient()
	if err != nil {
		err = fmt.Errorf("error: creation client %w", err)
		return nil, err
	}
	creationResponse, err := creation_client.GetCreationList(ctx, &creation.GetCreationListRequest{
		Ids: ids,
	})
	if err != nil {
		return nil, err
	}

	msg := creationResponse.GetMsg()
	code := msg.GetCode()
	status := msg.GetStatus()
	if status != common.ApiResponse_SUCCESS {
		if code[0] == '5' {
			return nil, fmt.Errorf("error: creationResponse %s", msg.GetDetails())
		}
		return nil, nil
	}

	infos := creationResponse.GetCreationInfoGroup()
	if len(infos) <= 0 {
		return nil, nil
	}
	for _, info := range infos {
		id := info.GetCreation().GetCreationId()
		creationMap[id] = info.GetCreation()
	}

	cards := make([]*generated.GetCreationReviewsResponse_CreationReview, len(reviews))
	for i, review := range reviews {
		id := review.GetNew().GetTargetId()

		card := &generated.GetCreationReviewsResponse_CreationReview{
			Review:   review,
			Creation: creationMap[id],
		}
		cards[i] = card
	}
	return cards, nil
}

func GetCreationReviews(ctx context.Context, req *generated.GetReviewsRequest) (*generated.GetCreationReviewsResponse, error) {
	response := new(generated.GetCreationReviewsResponse)

	token := req.GetAccessToken().GetValue()
	pass, reviewerId, err := auth.Auth("get", "review", token)
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

	review_client, err := client.GetReviewClient()
	if err != nil {
		err = fmt.Errorf("error: review client %w", err)
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "500",
			Details: err.Error(),
		}
		return response, err
	}
	reviewResponse, err := review_client.GetReviews(ctx, &review.GetReviewsRequest{
		Status:     req.GetStatus(),
		Type:       review.TargetType_CREATION,
		Page:       req.GetPage(),
		ReviewerId: reviewerId,
	})
	if err != nil {
		var msg *common.ApiResponse
		if reviewResponse != nil {
			msg = reviewResponse.GetMsg()
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

	msg := reviewResponse.GetMsg()
	status := msg.GetStatus()
	code := msg.GetCode()
	if status != common.ApiResponse_SUCCESS && status != common.ApiResponse_PENDING {
		response.Msg = msg
		if code[0] == '5' {
			return response, err
		}
		return response, nil
	}

	reviews := reviewResponse.GetReviews()
	cards, err := getCreationReviewCards(ctx, reviews)
	if err != nil {
		response.Msg = &common.ApiResponse{
			Status: common.ApiResponse_ERROR,
			Code:   "500",
		}
		return response, err
	}

	response.Count = reviewResponse.GetCount()
	response.Reviews = cards
	response.Msg = &common.ApiResponse{
		Status: common.ApiResponse_SUCCESS,
		Code:   "200",
	}
	return response, nil
}

func GetNewCreationReviews(ctx context.Context, req *generated.GetNewReviewsRequest) (*generated.GetCreationReviewsResponse, error) {
	response := new(generated.GetCreationReviewsResponse)

	token := req.GetAccessToken().GetValue()
	pass, reviewerId, err := auth.Auth("get", "review", token)
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

	review_client, err := client.GetReviewClient()
	if err != nil {
		err = fmt.Errorf("error: review client %w", err)
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "500",
			Details: err.Error(),
		}
		return response, err
	}
	reviewResponse, err := review_client.GetNewReviews(ctx, &review.GetNewReviewsRequest{
		Type:       review.TargetType_CREATION,
		ReviewerId: reviewerId,
	})
	if err != nil {
		var msg *common.ApiResponse
		if reviewResponse != nil {
			msg = reviewResponse.GetMsg()
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

	msg := reviewResponse.GetMsg()
	status := msg.GetStatus()
	code := msg.GetCode()
	if status != common.ApiResponse_SUCCESS && status != common.ApiResponse_PENDING {
		response.Msg = msg
		if code[0] == '5' {
			return response, err
		}
		return response, nil
	}

	reviews := reviewResponse.GetReviews()
	cards, err := getCreationReviewCards(ctx, reviews)
	if err != nil {
		response.Msg = &common.ApiResponse{
			Status: common.ApiResponse_ERROR,
			Code:   "500",
		}
		return response, err
	}

	response.Reviews = cards
	response.Msg = &common.ApiResponse{
		Status: common.ApiResponse_SUCCESS,
		Code:   "200",
	}
	return response, nil
}

func getCommentReviewCards(ctx context.Context, reviews []*review.Review) ([]*generated.GetCommentReviewsResponse_CommentReview, error) {
	commentMap := make(map[int32]*comment.Comment)
	for _, review := range reviews {
		id := review.GetNew().GetTargetId()
		commentMap[int32(id)] = &comment.Comment{
			CommentId: int32(id),
		}
	}
	length := len(commentMap)
	if length <= 0 {
		return nil, nil
	}
	ids := make([]int32, 0, length)
	for id := range commentMap {
		ids = append(ids, id)
	}

	comment_client, err := client.GetCommentClient()
	if err != nil {
		err = fmt.Errorf("error: comment client %w", err)
		return nil, err
	}
	commentResponse, err := comment_client.GetComments(ctx, &comment.GetCommentsRequest{
		Ids: ids,
	})
	if err != nil {
		return nil, err
	}

	msg := commentResponse.GetMsg()
	code := msg.GetCode()
	status := msg.GetStatus()
	if status != common.ApiResponse_SUCCESS {
		if code[0] == '5' {
			return nil, fmt.Errorf("error: commentResponse %s", msg.GetDetails())
		}
		return nil, nil
	}

	infos := commentResponse.GetComments()
	if len(infos) <= 0 {
		return nil, nil
	}
	for _, info := range infos {
		id := info.GetCommentId()
		commentMap[id] = info
	}

	cards := make([]*generated.GetCommentReviewsResponse_CommentReview, len(reviews))
	for i, review := range reviews {
		id := review.GetNew().GetTargetId()

		card := &generated.GetCommentReviewsResponse_CommentReview{
			Review:  review,
			Comment: commentMap[int32(id)],
		}
		cards[i] = card
	}
	return cards, nil
}

func GetCommentReviews(ctx context.Context, req *generated.GetReviewsRequest) (*generated.GetCommentReviewsResponse, error) {
	response := new(generated.GetCommentReviewsResponse)

	token := req.GetAccessToken().GetValue()
	pass, reviewerId, err := auth.Auth("get", "review", token)
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

	review_client, err := client.GetReviewClient()
	if err != nil {
		err = fmt.Errorf("error: review client %w", err)
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "500",
			Details: err.Error(),
		}
		return response, err
	}
	reviewResponse, err := review_client.GetReviews(ctx, &review.GetReviewsRequest{
		Status:     req.GetStatus(),
		Type:       review.TargetType_COMMENT,
		Page:       req.GetPage(),
		ReviewerId: reviewerId,
	})
	if err != nil {
		var msg *common.ApiResponse
		if reviewResponse != nil {
			msg = reviewResponse.GetMsg()
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

	msg := reviewResponse.GetMsg()
	status := msg.GetStatus()
	code := msg.GetCode()
	if status != common.ApiResponse_SUCCESS && status != common.ApiResponse_PENDING {
		response.Msg = msg
		if code[0] == '5' {
			return response, err
		}
		return response, nil
	}

	reviews := reviewResponse.GetReviews()
	cards, err := getCommentReviewCards(ctx, reviews)
	if err != nil {
		response.Msg = &common.ApiResponse{
			Status: common.ApiResponse_ERROR,
			Code:   "500",
		}
		return response, err
	}

	response.Reviews = cards
	response.Count = reviewResponse.GetCount()
	response.Msg = &common.ApiResponse{
		Status: common.ApiResponse_SUCCESS,
		Code:   "200",
	}
	return response, nil
}

func GetNewCommentReviews(ctx context.Context, req *generated.GetNewReviewsRequest) (*generated.GetCommentReviewsResponse, error) {
	response := new(generated.GetCommentReviewsResponse)

	token := req.GetAccessToken().GetValue()
	pass, reviewerId, err := auth.Auth("get", "review", token)
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

	review_client, err := client.GetReviewClient()
	if err != nil {
		err = fmt.Errorf("error: review client %w", err)
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "500",
			Details: err.Error(),
		}
		return response, err
	}
	reviewResponse, err := review_client.GetNewReviews(ctx, &review.GetNewReviewsRequest{
		Type:       review.TargetType_COMMENT,
		ReviewerId: reviewerId,
	})
	if err != nil {
		var msg *common.ApiResponse
		if reviewResponse != nil {
			msg = reviewResponse.GetMsg()
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

	msg := reviewResponse.GetMsg()
	status := msg.GetStatus()
	code := msg.GetCode()
	if status != common.ApiResponse_SUCCESS && status != common.ApiResponse_PENDING {
		response.Msg = msg
		if code[0] == '5' {
			return response, err
		}
		return response, nil
	}

	reviews := reviewResponse.GetReviews()
	cards, err := getCommentReviewCards(ctx, reviews)
	if err != nil {
		response.Msg = &common.ApiResponse{
			Status: common.ApiResponse_ERROR,
			Code:   "500",
		}
		return response, err
	}

	response.Reviews = cards
	response.Msg = &common.ApiResponse{
		Status: common.ApiResponse_SUCCESS,
		Code:   "200",
	}
	return response, nil
}

package recommend

import (
	"context"
	"math"

	"github.com/Yux77Yux/platform_backend/microservices/interaction/tools"
)

// 行为数据类型
type Behavior struct {
	Id     int64
	Weight map[int64]float64
	norm   float64
}

// 基于用户的协同过滤
// 获取用户的行为数据
func GetUserBehavior(ctx context.Context, userID int64) *Behavior {
	const (
		viewWeight = 1
	)

	// 取存档，若无存档则取历史记录
	history, err := cache.GetArchiveData(ctx, userID)
	if err != nil {
		tools.LogError("", "recommend GetUserBehavior", err)
	}

	userWeight := make(map[int64]float64)
	for _, item := range history {
		itemID := item.GetBase().GetCreationId()
		userWeight[itemID] = viewWeight
	}

	var normUser float64
	// 计算 模
	for _, weight := range userWeight {
		normUser += weight * weight
	}

	// 将结果转为 Behavior 格式并返回
	Behavior := &Behavior{
		Id:     userID,
		Weight: userWeight,
		norm:   math.Sqrt(normUser),
	}
	return Behavior
}

func GetOtherUsers(ctx context.Context, ids []int64) ([]*Behavior, error) {
	otherMap, err := cache.GetAnyItemUsers(ctx, ids)
	if err != nil {
		tools.LogError("", "recommend GetOtherUsers", err)
		return nil, err
	}

	behaviorSlice := make([]*Behavior, 0, len(ids))
	for id, val := range otherMap {
		var normUser float64
		// 计算 模
		for _, weight := range val {
			normUser += weight * weight
		}
		behavior := &Behavior{
			Id:     id,
			Weight: val,
			norm:   math.Sqrt(normUser),
		}
		behaviorSlice = append(behaviorSlice, behavior)
	}
	return behaviorSlice, nil
}

// 基于物品的协同过滤
// 获取观看过作品的用户
func GetCreationViewer(ctx context.Context, creationId int64) *Behavior {
	const (
		viewWeight = 1
	)

	itemUsers, err := cache.GetUsers(ctx, creationId)
	if err != nil {
		tools.LogError("", "recommend GetCreationViewer", err)
	}

	itemWeight := make(map[int64]float64)
	// 如果存在于历史记录中
	for _, userId := range itemUsers {
		itemWeight[userId] = viewWeight
	}

	var norm float64
	// 计算 模
	for _, weight := range itemWeight {
		norm += weight * weight
	}

	// 将结果转为 Behavior 格式并返回
	Behavior := &Behavior{
		Id:     creationId,
		Weight: itemWeight,
		norm:   math.Sqrt(norm),
	}
	return Behavior
}

func GetOtherCreationViewer(ctx context.Context, userIds []int64) ([]*Behavior, error) {
	otherMap, err := cache.GetAnyUsersHistory(ctx, userIds)
	if err != nil {
		return nil, err
	}

	length := len(userIds)
	behaviorSlice := make([]*Behavior, 0, length)
	for id, val := range otherMap {
		var norm float64
		// 计算 模
		for _, weight := range val {
			norm += weight * weight
		}
		behavior := &Behavior{
			Id:     id,
			Weight: val,
			norm:   math.Sqrt(norm),
		}
		behaviorSlice = append(behaviorSlice, behavior)
	}
	return behaviorSlice, nil
}

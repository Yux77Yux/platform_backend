package recommend

import (
	"context"

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
func GetUserBehavior(userID int64) *Behavior {
	const (
		viewWeight = 1
	)

	ctx := context.Background()
	history, err := cache.GetHistories(ctx, userID, 1)
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
		norm:   normUser,
	}
	return Behavior
}

func GetOtherUsers(ctx context.Context) ([]*Behavior, error) {
	others, err := cache.ScanZSetsByHistories(ctx)
	if err != nil {
		tools.LogError("", "recommend GetOtherUsers", err)
		return nil, err
	}
	length := len(others)

	otherMap, err := cache.GetAllInteractions(ctx, others)
	if err != nil {
		tools.LogError("", "recommend GetOtherUsers", err)
		return nil, err
	}
	behaviorSlice := make([]*Behavior, 0, length)
	for id, val := range otherMap {
		var normUser float64
		// 计算 模
		for _, weight := range val {
			normUser += weight * weight
		}
		behavior := &Behavior{
			Id:     id,
			Weight: val,
			norm:   normUser,
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
		norm:   norm,
	}
	return Behavior
}

func GetOtherCreationViewer(ctx context.Context) ([]*Behavior, error) {
	others, err := cache.ScanZSetsByCreationId(ctx)
	if err != nil {
		return nil, err
	}
	length := len(others)

	otherMap, err := cache.GetAllItemUsers(ctx, others)
	if err != nil {
		return nil, err
	}
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
			norm:   norm,
		}
		behaviorSlice = append(behaviorSlice, behavior)
	}
	return behaviorSlice, nil
}

package recommend

import (
	cache "github.com/Yux77Yux/platform_backend/microservices/interaction/cache"
	"log"
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

	history, err := cache.GetHistories(userID, 1)
	if err != nil {
		log.Printf("GetHistories err %v", err)
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

func GetOtherUsers() ([]*Behavior, error) {
	others, err := cache.ScanZSetsByHistories()
	if err != nil {
		log.Printf("ScanZSetsByHistories err %v", err)
		return nil, err
	}
	length := len(others)

	otherMap, err := cache.GetAllInteractions(others)
	if err != nil {
		log.Printf("GetAllInteractions err %v", err)
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
func GetCreationViewer(creationId int64) *Behavior {
	const (
		viewWeight = 1
	)

	itemUsers, err := cache.GetUsers(creationId)
	if err != nil {
		log.Printf("GetHistories err %v", err)
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

func GetOtherCreationViewer() ([]*Behavior, error) {
	others, err := cache.ScanZSetsByCreationId()
	if err != nil {
		log.Printf("ScanZSetsByCreationId err %v", err)
		return nil, err
	}
	length := len(others)

	otherMap, err := cache.GetAllItemUsers(others)
	if err != nil {
		log.Printf("GetAllItemUsers err %v", err)
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

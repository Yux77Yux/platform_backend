package recommend

import (
	cache "github.com/Yux77Yux/platform_backend/microservices/interaction/cache"
	"log"
)

// 获取用户的行为数据
func GetUserBehavior(userID int64) *UserBehavior {
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

	// 将结果转为 UserBehavior 格式并返回
	userBehavior := &UserBehavior{
		userId:     userID,
		userWeight: userWeight,
		norm:       normUser,
	}
	return userBehavior
}

func GetOtherUsers() ([]*UserBehavior, error) {
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
	behaviorSlice := make([]*UserBehavior, 0, length)
	for id, val := range otherMap {
		var normUser float64
		// 计算 模
		for _, weight := range val {
			normUser += weight * weight
		}
		behavior := &UserBehavior{
			userId:     id,
			userWeight: val,
			norm:       normUser,
		}
		behaviorSlice = append(behaviorSlice, behavior)
	}
	return behaviorSlice, nil
}

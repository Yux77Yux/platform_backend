package recommend

import (
	"context"
	"math"
)

// 计算两个用户之间的余弦相似度
func CosineSimilarity(user1, user2 *Behavior) float64 {
	var (
		dotProduct float64
		normUser1  = math.Sqrt(user1.norm)
		normUser2  = math.Sqrt(user2.norm)
	)

	// 计算点积
	for itemID, weight1 := range user1.Weight {
		if weight2, exist := user2.Weight[itemID]; exist {
			dotProduct += weight1 * weight2
		}
	}

	return dotProduct / (normUser1 * normUser2)
}

// 根据用户的相似度来推荐作品
func Recommend(ctx context.Context, userID int64) ([]int64, error) {
	// 获取目标用户的行为数据
	targetUser := GetUserBehavior(userID)
	otherUsers, err := GetOtherUsers(ctx)
	if err != nil {
		return nil, err
	}

	recommendations := make([]int64, 0, 201)

	// 计算与其他用户的相似度
	for _, otherUser := range otherUsers {
		if otherUser.Id == userID {
			continue
		}
		similarity := CosineSimilarity(targetUser, otherUser)

		// 两个用户的相似度高于阈值，推荐作品,targetUser是推送目标
		if similarity > 0.5 {
			if len(recommendations) == 200 {
				break
			}
			for itemID := range otherUser.Weight {
				if _, exists := targetUser.Weight[itemID]; !exists {
					recommendations = append(recommendations, itemID)
				}
			}
		}
	}

	return recommendations, nil
}

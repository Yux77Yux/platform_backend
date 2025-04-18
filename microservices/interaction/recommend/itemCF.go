package recommend

import (
	"context"
	"log"
	"math"
	"sort"
)

// 基于物品的协同过滤推荐
// 思路：
// 1. 从所有用户（包括目标用户）构造倒排索引：快速查相似度。
// 2. 对于目标用户每个看过的视频，遍历该视频的用户列表，找出他们看过的其他视频，并计算物品相似度（余弦相似度）。
// 3. 累加相似度得分，最后排序返回未看过的热门视频。
func RecommendItemBased(ctx context.Context, id int64) ([]int64, error) {
	// 获取观看过目标视频的用户与视频的模
	targetCreation := GetCreationViewer(ctx, id)
	//  otherUsersId 为 用户ID集合
	userIds := make([]int64, 0, len(targetCreation.Weight))
	for id := range targetCreation.Weight {
		userIds = append(userIds, id)
	}

	// 获取观看过该视频的所有用户看过的视频
	others, err := GetOtherCreationViewer(ctx, userIds)
	if err != nil {
		log.Printf("GetOtherCreationViewer err: %v", err)
		return nil, err
	}

	creationUserMap := make(map[int64]map[int64]float64)
	for _, val := range others {
		creationUserMap[val.Id] = val.Weight
	}

	itemUserIndex := make(map[int64]float64)
	for _, other := range others {
		var dot float64
		for id, weight1 := range other.Weight {
			if weight2, exist := targetCreation.Weight[id]; exist {
				dot += weight1 * weight2
			}
		}
		dot = math.Sqrt(dot)
		itemUserIndex[other.Id] = dot / other.norm * targetCreation.norm
	}

	// 将推荐结果按分数排序，返回前 N 个视频ID
	type pair struct {
		creationId int64
		score      float64
	}

	var recPairs []pair
	for otherCreationId, userWeightMap := range creationUserMap {
		similarity := itemUserIndex[otherCreationId]

		score := float64(0)
		for targetUserId := range targetCreation.Weight {
			if weight, exist := userWeightMap[targetUserId]; exist {
				score += weight * similarity
			}
		}

		recPairs = append(recPairs, pair{creationId: otherCreationId, score: score})
	}

	sort.Slice(recPairs, func(i, j int) bool {
		return recPairs[i].score > recPairs[j].score
	})

	N := len(recPairs)

	recommendations := make([]int64, N)
	for i := 0; i < N; i++ {
		recommendations[i] = recPairs[i].creationId
	}

	return recommendations, nil
}

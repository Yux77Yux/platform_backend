package recommend

import (
	"log"
	"math"
	"sort"
)

// 基于物品的协同过滤推荐
// 思路：
// 1. 从所有用户（包括目标用户）构造倒排索引：每个视频对应一个 map，记录哪些用户看过该视频及其权重。
// 2. 对于目标用户每个看过的视频，遍历该视频的用户列表，找出他们看过的其他视频，并计算物品相似度（余弦相似度）。
// 3. 累加相似度得分，最后排序返回未看过的热门视频。
func RecommendItemBased(userId int64) ([]int64, error) {
	// 获取目标用户的行为数据
	targetUser := GetUserBehavior(userId)

	// 获取其他所有用户的数据
	otherUsers, err := GetOtherUsers()
	if err != nil {
		log.Printf("GetOtherUsers err: %v", err)
		return nil, err
	}

	// 将目标用户加入用户集合，构建全局倒排索引
	allUsers := append(otherUsers, targetUser)

	// 构建倒排索引：map[creationId]map[userId]weight
	itemUserIndex := make(map[int64]map[int64]float64)
	for _, ub := range allUsers {
		for creationId, weight := range ub.Weight {
			if itemUserIndex[creationId] == nil {
				itemUserIndex[creationId] = make(map[int64]float64)
			}
			itemUserIndex[creationId][ub.Id] = weight
		}
	}

	// 计算每个视频的模（基于所有用户的贡献），用于后续相似度计算
	itemNorm := make(map[int64]float64)
	for creationId, userMap := range itemUserIndex {
		var sum float64
		for _, w := range userMap {
			sum += w * w
		}
		itemNorm[creationId] = math.Sqrt(sum)
	}

	// 推荐分数：对于目标用户看过的每个视频 i，
	// 找出和 i 相似的其他视频 j：余弦相似度 = (sum over common users: w_ui * w_uj) / (norm(i)*norm(j))
	// 最终推荐分数累加：score(j) += similarity(i,j)*targetUser.userWeight[i]
	recScores := make(map[int64]float64)
	// user_creationId 为目标用户的历史记录的作品id，即用户id->作品id->weight 中的 用户id
	// creation_userId 为形成的倒排索引中的  作品id->用户id->weight中的 作品id
	for user_creationId, user_weight := range targetUser.Weight {
		// creation_userIdMap为记录 目标用户的观看记录中的作品 ，观看过该作品的其他用户。
		// 我看过的，你们也看过
		creation_userIdMap := itemUserIndex[user_creationId]
		// 遍历倒排索引中记录的视频id，以及观看过该视频的用户id
		for creationId, creationUserMap := range itemUserIndex {
			if user_creationId == creationId {
				continue
			}
			// 两个 视频->用户 的交集
			// 比如
			// 目标用户的历史记录(a1,a2,a3,a4,...)
			// 倒排索引中的(b1,b2,b3,b4,...)
			// a1中有一批用户id，(au1,au2,au3,au4,...)
			// b1中有一批用户id，(bu1,bu2,bu3,bu4,...)
			// 上面的 user_creationId == creationId 即a1=b1的时候，两者的集合，
			// 即creation_userIdMap和creationUserMap元素完全一致，可以跳过
			// 计算的是，它们的交集即(au1,au2,au3,au4,...)中存在部分元素(e1,e2,e3,e4,...)
			// (e1,e2,e3,e4,...)也存在(bu1,bu2,bu3,bu4,...)中，这既是交集
			var dot float64
			// wI就是用户对该视频的热度贡献，交集中的每个元素也可他们的名字一样，但是可能映射不同的值，这个值便是权重
			// 点积就是交集，组成分子
			// 此处计算的是交集
			for user, wI := range creation_userIdMap {
				if wJ, exists := creationUserMap[user]; exists {
					dot += wI * wJ
				}
			}
			// 分子大于0，即存在它们有一定程度都喜欢另外一个视频
			// itemNorm[user_creationId]是自己的历史记录中视频的模
			// itemNorm[creationId]是倒排索引中视频的模
			// 入循环前已经计算出来了
			if dot > 0 && itemNorm[user_creationId] > 0 && itemNorm[creationId] > 0 {
				// similarity就是两个视频之间的余弦相似度
				similarity := dot / (itemNorm[user_creationId] * itemNorm[creationId])
				// 累加推荐分数。目标用户对观看着的视频的权重user_weight 乘以 两个视频之间的余弦相似度similarity
				//  由于点赞和收藏都未成为权重的一部分，所以这里相乘user_weight只是为了未来可能拓展权重分配而设置
				// 使用+=是因为分数的计算是由两两比较得出的，假设我有4个历史记录，而倒排索引记录了8个视频，
				// 假如两个集合都没有同样的视频，我需要两两比较4*8=32次，而此处creationId的similarity的会经历四次计算
				// 4个历史记录中，目标用户的user_weight可能都不同（虽然这里权重都为1，但未来可能改变权重分配）
				recScores[creationId] += user_weight * similarity
			}
		}
	}

	// 移除目标用户已经看过的视频
	// 由于存在 a1和a2之间的两两比较，即此时b1或b2或b3或b4...可能和a2相同，也可能和a1或a3或a4相同
	for creationId := range targetUser.Weight {
		delete(recScores, creationId)
	}

	// 将推荐结果按分数排序，返回前 N 个视频ID
	type pair struct {
		creationId int64
		score      float64
	}
	var recPairs []pair
	for vid, score := range recScores {
		recPairs = append(recPairs, pair{creationId: vid, score: score})
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

package recommend

// import (
// 	"fmt"
// 	"log"
// 	"math/rand/v2"
// 	"strconv"
// )

func UpdateCategoryHotList(categoryID int64) {
	// key := fmt.Sprintf("ZSET_Category_%d", categoryID)
	// videoList, err := db.GetTopVideosByCategory(categoryID, 100) // 从数据库获取前100热门视频
	// if err != nil {
	// 	log.Printf("Error fetching videos for category %d: %v", categoryID, err)
	// 	return
	// }

	// pipe := redisClient.TxPipeline()
	// pipe.Del(ctx, key) // 先删除原来的 ZSET

	// for _, video := range videoList {
	// 	hotScore := video.ViewCount + 2*video.LikeCount + 3*video.CollectionCount
	// 	pipe.ZAdd(ctx, key, &redis.Z{Score: float64(hotScore), Member: video.ID})
	// }

	// _, err = pipe.Exec(ctx)
	// if err != nil {
	// 	log.Printf("Error updating category hot list for %d: %v", categoryID, err)
	// }
}

func GetCategoryRecommendations(categoryID int64, limit int) ([]int64, error) {
	// key := fmt.Sprintf("ZSET_Category_%d", categoryID)
	// videoIDs, err := redisClient.ZRevRange(ctx, key, 0, 49).Result() // 取前50个
	// if err != nil {
	// 	return nil, err
	// }

	// if len(videoIDs) == 0 {
	// 	return nil, nil
	// }

	// // 洗牌算法随机排序
	// rand.Shuffle(len(videoIDs), func(i, j int) { videoIDs[i], videoIDs[j] = videoIDs[j], videoIDs[i] })

	// // 取前 limit 个返回
	// result := make([]int64, 0, limit)
	// for i, idStr := range videoIDs {
	// 	if i >= limit {
	// 		break
	// 	}
	// 	id, _ := strconv.ParseInt(idStr, 10, 64)
	// 	result = append(result, id)
	// }

	// return result, nil
	return nil, nil
}

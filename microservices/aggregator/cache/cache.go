package cache

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/Yux77Yux/platform_backend/generated/common"
)

func AddIpInSet(ctx context.Context, req *common.ViewCreation) error {
	id := req.GetId()
	ip := req.GetIpv4()

	key := fmt.Sprintf("Set_Ip_Creation_%d", id)
	// 检查集合是否存在
	exists, err := CacheClient.Exists(ctx, key)
	if err != nil {
		return err
	}

	// 创建管道
	pipeline := CacheClient.TxPipeline()

	// 先检查集合是否存在，如果不存在则插入数据并设置过期时间
	if !exists {
		// 集合不存在，插入数据并设置过期时间
		pipeline.SAdd(ctx, key, ip)
		pipeline.Expire(ctx, key, 5*time.Minute)
	} else {
		// 集合已存在，只插入数据
		pipeline.SAdd(ctx, key, ip)
	}

	// 执行管道中的所有命令
	_, err = pipeline.Exec(ctx)
	if err != nil {
		return err
	}

	return err
}

func ExistIpInSet(ctx context.Context, req *common.ViewCreation) (bool, error) {
	id := req.GetId()
	ip := req.GetIpv4()

	idStr := strconv.FormatInt(id, 10)

	exist, err := CacheClient.ExistsInSet(ctx, "Ip_Creation", idStr, ip)
	if err != nil {
		return false, err // 如果 Redis 出现错误，返回错误
	}

	if !exist {
		// 键不存在或集合为空
		return false, nil
	}

	// 键存在且 IP 在集合中
	return true, nil
}

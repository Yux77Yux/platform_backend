package snow

import (
	"log"

	"github.com/bwmarrin/snowflake"
)

func GetId() int64 {
	// 生成id
	node, err := snowflake.NewNode(1) // 传入机器ID，这里假设为1
	if err != nil {
		log.Printf("Failed to create snowflake node: %v", err)
	}

	id := node.Generate().Int64()
	// 生成唯一的ID,确保不为0
	for id == 0 {
		id = node.Generate().Int64()
	}

	return id
}

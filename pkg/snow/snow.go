package snow

import (
	"log"
	"sync"

	"github.com/bwmarrin/snowflake"
)

var (
	node *snowflake.Node
	once sync.Once
)

func initNode() {
	var err error
	node, err = snowflake.NewNode(1)
	if err != nil {
		log.Fatalf("Failed to create snowflake node: %v", err)
	}
}

func GetId() int64 {
	once.Do(initNode)
	id := node.Generate().Int64()
	for id == 0 {
		id = node.Generate().Int64()
	}
	return id
}

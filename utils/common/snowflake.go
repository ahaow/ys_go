package common

import (
	"ys_go/global"

	"github.com/bwmarrin/snowflake"
)

var node *snowflake.Node

func InitSnowflake(machineID int64) {
	var err error
	node, err = snowflake.NewNode(machineID)
	if err != nil {
		global.Log.Sugar().Errorf("Snowflake init failed: %v", err)
	}
}

// func GenerateSnowflakeID() int64 {
// 	return node.Generate().Int64()
// }

func GenerateSnowflakeID() string {
	return node.Generate().String()
}

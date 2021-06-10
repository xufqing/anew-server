package redis

import (
	"anew-server/pkg/common"
	"context"
	"github.com/go-redis/redis/v8"
)

// 处理string类型操作
type StringOperation struct {
	redis *redis.Client
	ctx   context.Context
}

// 初始化string类型redis服务
func NewStringOperation() *StringOperation {
	return &StringOperation{
		redis: common.Redis,
		ctx:   context.Background(),
	}
}

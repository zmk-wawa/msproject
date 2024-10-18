package repo

import (
	"context"
	"time"
)

// 缓存cache的接口，使用统一的接口让Redis/Mysql/MangoDb等接入
//各种存储形式实现这个接口

type Cache interface {
	Put(ctx context.Context, key, value string, expire time.Duration) error
	Get(ctx context.Context, key string) (string, error)
}

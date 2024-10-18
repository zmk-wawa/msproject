package dao

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

// dao层是负责各种数据库的连接

// 这个是Redis数据库实现的repo/cache中的接口
//业务层通过调用repo/cache中的接口来调用Redis
var Rc *RedisCache

type RedisCache struct {
	rdb *redis.Client
}

func init() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	Rc = &RedisCache{
		rdb: rdb,
	}
}

// 重新包装Redis 中的Set与Get方法
func (rc *RedisCache) Put(ctx context.Context, key, value string, expire time.Duration) error {
	err := rc.rdb.Set(ctx, key, value, expire).Err()
	return err
}

func (rc *RedisCache) Get(ctx context.Context, key string) (string, error) {
	result, err := rc.rdb.Get(ctx, key).Result()
	return result, err
}

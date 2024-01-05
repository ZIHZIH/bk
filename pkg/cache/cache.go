package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
	"wzh/logger"
)

type Cache struct {
	Prefix  string
	RedisDb *redis.Client
}

// SetCache 设置缓存
func (cache *Cache) SetCache(ctx context.Context, id string, value interface{}, expiration time.Duration) error {
	key := cache.Prefix + id
	err := cache.RedisDb.Set(ctx, key, value, expiration).Err()
	if err != nil {
		logger.Logger.Println(err)
		return err
	}
	return nil
}

// GetCache 获取缓存
func (cache *Cache) GetCache(ctx context.Context, id string) (string, error) {
	key := cache.Prefix + id
	value, err := cache.RedisDb.Get(ctx, key).Result()
	if err != nil && err != redis.Nil {
		logger.Logger.Println(err)
	}
	return value, err
}

// DelCache 删除缓存
func (cache *Cache) DelCache(ctx context.Context, id string) {
	key := cache.Prefix + id
	if err := cache.RedisDb.Del(ctx, key).Err(); err != nil {
		logger.Logger.Println(err)
	}
}
func NewCache(prefix string, client *redis.Client) *Cache {
	return &Cache{
		Prefix:  prefix,
		RedisDb: client,
	}
}

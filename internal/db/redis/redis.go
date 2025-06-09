package db

import (
    "context"
    "time"

    "github.com/redis/go-redis/v9"
)

type RedisCache struct {
    client *redis.ClusterClient
    ctx    context.Context
}

func NewRedisCache(addrs []string, password string) *RedisCache {
    ctx := context.Background()
    client := redis.NewClusterClient(&redis.ClusterOptions{
        Addrs:    addrs,   
        Password: password, 
    })
    return &RedisCache{
        client: client,
        ctx:    ctx,
    }
}

func (rc *RedisCache) Set(key string, value interface{}, expiration time.Duration) error {
    return rc.client.Set(rc.ctx, key, value, expiration).Err()
}

func (rc *RedisCache) Get(key string) (string, error) {
    return rc.client.Get(rc.ctx, key).Result()
}

func (rc *RedisCache) Del(key string) error {
    return rc.client.Del(rc.ctx, key).Err()
}

func (rc *RedisCache) Exists(key string) (bool, error) {
    n, err := rc.client.Exists(rc.ctx, key).Result()
    return n > 0, err
}

func (rc *RedisCache) SAdd(key string, members ...interface{}) error {
    return rc.client.SAdd(rc.ctx, key, members...).Err()
}

func (rc *RedisCache) SIsMember(key string, member interface{}) (bool, error) {
    return rc.client.SIsMember(rc.ctx, key, member).Result()
}

func (rc *RedisCache) Close() error {
    return rc.client.Close()
}
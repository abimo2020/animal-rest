package repositories

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisRepository interface {
	Get(key string) (string, error)
	Set(key string, value string) error
	Delete(key string) error
}

type redisRepository struct {
	redis *redis.Client
}

func NewRedisRepository(redis *redis.Client) RedisRepository {
	return &redisRepository{redis: redis}
}

func (rr redisRepository) Get(key string) (string, error) {
	value, err := rr.redis.Get(context.Background(), key).Result()
	if err != nil {
		return "", err
	}
	return value, nil
}

func (rr redisRepository) Set(key string, value string) error {
	return rr.redis.Set(context.Background(), key, value, time.Hour*24).Err()
}

func (rr redisRepository) Delete(key string) error {
	return rr.redis.Del(context.Background(), key).Err()
}

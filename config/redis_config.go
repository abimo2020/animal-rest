package config

import (
	"context"
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
)

var Redis *redis.Client

type RedisConfig struct {
	Host     string
	Password string
	DB       int
}

func RedisInit() {
	config := RedisConfig{
		Host:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	}
	Redis = redis.NewClient(&redis.Options{
		Addr:     config.Host,
		Password: config.Password,
		DB:       config.DB,
	})
	_, err := Redis.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("failed to connect to redis:", err)
	} else {
		fmt.Println("connected to redis!")
	}
}

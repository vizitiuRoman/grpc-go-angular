package models

import (
	"context"
	"os"

	"github.com/go-redis/redis/v8"
)

var (
	rds *redis.Client
)

func InitRedis() error {
	host, port := os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")
	rds = redis.NewClient(&redis.Options{
		Addr: host + ":" + port,
	})
	_, err := rds.Ping(context.Background()).Result()
	if err != nil {
		return err
	}
	return nil
}

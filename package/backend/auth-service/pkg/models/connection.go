package models

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var (
	rds *redis.Client
)

func InitRedis() error {
	//host, port := settings.Get().RedisHost, settings.Get().RedisPort
	rds = redis.NewClient(&redis.Options{
		Addr: "redis-master" + ":" + "6379",
	})
	_, err := rds.Ping(context.Background()).Result()
	if err != nil {
		return err
	}
	return nil
}

package infra

import (
	"context"
	"fmt"
	"im/internal/config"
	"log"

	"github.com/redis/go-redis/v9"
)

func NewRedis(conf *config.Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", conf.Redis.Host, conf.Redis.Port),
		Password: conf.Redis.Password,
		DB:       conf.Redis.DB,
	})

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("redis启动失败, %s", err.Error())
	}
	return rdb
}

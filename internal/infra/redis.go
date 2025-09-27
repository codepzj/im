package infra

import (
	"context"
	"fmt"
	"im/internal/config"
	"log"
	"log/slog"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctxBackround = context.Background()

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

// 设置具有过期时间的key
func SetKeyEx(rdb *redis.Client, key string, value any, duration time.Duration) {
	if err := rdb.Set(ctxBackround, key, value, duration).Err(); err != nil {
		slog.Error("写缓存失败")
	}
}

// 根据key读缓存
func GetKey(rdb *redis.Client, key string) (string, error) {
	return rdb.Get(ctxBackround, key).Result()
}

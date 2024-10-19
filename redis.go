package goutils

import (
	"github.com/redis/go-redis/v9"
)

var (
	Rd redis.Client
)

func initRedis() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     GetStringEnv("REDIS_HOST", DefaultString("localhost")),
		Password: GetStringEnv("REDIS_PASSWORD", DefaultString("")),
		DB:       0,
	})

	Rd = *rdb
}

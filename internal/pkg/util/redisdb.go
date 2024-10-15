package util

import (
	"github.com/redis/go-redis/v9"
)

func InitRedisDB(addr string) *redis.Client {
	// init redis db
	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	return rdb
}

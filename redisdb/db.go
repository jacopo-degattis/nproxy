package redisdb

import (
	"context"
	"nproxy/config"

	"github.com/redis/go-redis/v9"
)

var client *redis.Client
var ctx = context.Background()

func InitRedisClient() {
	addr := config.GetRedisUrl()
	password := config.GetRedisPassword()

	options := &redis.Options{
		Addr: addr,
		DB:   0,
	}

	if password != nil {
		options.Password = *password
	}

	client = redis.NewClient(options)
}

func Set(key string, value any) error {
	return client.Set(ctx, key, value, 0).Err()
}

func Get(key string) *redis.StringCmd {
	return client.Get(ctx, key)
}

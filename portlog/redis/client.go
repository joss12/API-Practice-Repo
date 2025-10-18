package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

var Ctx = context.Background()
var Client *redis.Client

func Connect(addr, password string) {
	Client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0,
	})

	//ping to test
	_, err := Client.Ping(Ctx).Result()
	if err != nil {
		panic("Redis connection failed: " + err.Error())
	}

	println("Redis connected successfully!")
}

func CacheSet(key string, value string, ttl time.Duration) error {
	return Client.Set(Ctx, key, value, ttl).Err()
}

func CacheGet(key string) (string, error) {
	return Client.Get(Ctx, key).Result()
}

package cache

import (
	"fmt"
	"gopkg.in/redis.v5"
	"time"
)

var redisClient *redis.Client

func Init() {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
		PoolSize: 100,
	})
	if redisClient == nil {
		panic(fmt.Sprintln("InitRedis Error"))
	}
	redisClient.Set("nihao", 1, time.Duration(0))
}

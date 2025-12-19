package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func InitRedis() {
	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		redisAddr = "localhost:6379"
	}
	redisPassword := os.Getenv("REDIS_PASSWORD") // Load password

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword, // Set password
		DB:       0,
	})
	if _, err := RedisClient.Ping(context.Background()).Result(); err != nil {
		log.Println("Warning: Redis not connected. Ensure Redis is running.", err)
	} else {
		fmt.Println("Redis Connected")
	}
}

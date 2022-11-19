package database

import (
	"fmt"
	"os"

	"github.com/go-redis/redis"
)

func CreateConnectionRedis() *redis.Client {
	REDIS_URL := os.Getenv("REDIS_URL")
	client := redis.NewClient(&redis.Options{
		Addr:     REDIS_URL,
		Password: "",
		DB:       0,
	})

	fmt.Println("Connected redis...")
	return client
}

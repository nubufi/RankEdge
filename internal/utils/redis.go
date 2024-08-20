package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

// ConnectToRedis connects to the redis database
func ConnectToRedis() {
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: "",
		DB:       0,
	})
}

func SetCache(key string, value interface{}) {
	ctx := context.Background()
	// Serialize the orders to JSON
	jsonOrders, err := json.Marshal(value)
	if err != nil {
		fmt.Println("Error serializing orders:", err)
		return
	}
	RedisClient.Set(ctx, key, jsonOrders, 0)
}

func ClearCache(key string) {
	ctx := context.Background()
	RedisClient.Del(ctx, key)
}

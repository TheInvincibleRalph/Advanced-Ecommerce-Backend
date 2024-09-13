package utils

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client

// InitRedisClient initializes the Redis client
func InitRedisClient() *redis.Client {
	if redisClient != nil {
		return redisClient
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis server address
		Password: "",               // No password set
		DB:       0,                // Use default DB
	})

	// Create a context with a timeout of 10 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Ping the Redis server to test the connection
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis at %s: %v", rdb.Options().Addr, err)
		return nil
	}

	log.Println("Connected to Redis at", rdb.Options().Addr)
	redisClient = rdb
	return redisClient
}

// GetRedisClient returns the Redis client
func GetRedisClient() *redis.Client {
	return InitRedisClient()
}

package utils

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis"
)

// Intializes a Redis instance
func InitRedisClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis server address
		Password: "",               // No password set
		DB:       0,                // Use default DB
	})

	// Creates a context with a timeout of 10 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resultChan := make(chan error, 1)

	go func() {
		_, err := rdb.Ping().Result()
		if err != nil {
			resultChan <- err
		} else {
			close(resultChan)
		}
	}()

	select {
	case <-ctx.Done():
		log.Fatalf("Context timeout: %v", ctx.Err())
		return nil
	case err := <-resultChan:
		if err != nil {
			log.Fatalf("Failed to connect to Redis: %v", err)
			return nil
		}
		return rdb
	}
}

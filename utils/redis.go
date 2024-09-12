package utils

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

// Initializes a Redis instance
// func InitRedisClient() *redis.Client {
// 	rdb := redis.NewClient(&redis.Options{
// 		Addr:     "localhost:6379", // Redis server address
// 		Password: "",               // No password set
// 		DB:       0,                // Use default DB
// 	})

// 	// Creates a context with a timeout of 10 seconds
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	resultChan := make(chan error, 1)

// 	go func() {
// 		_, err := rdb.Ping(ctx).Result() // Pass the context to Ping method
// 		if err != nil {
// 			resultChan <- err
// 		} else {
// 			close(resultChan)
// 		}
// 	}()

// 	select {
// 	case <-ctx.Done():
// 		log.Fatalf("Context timeout: %v", ctx.Err())
// 		return nil
// 	case err := <-resultChan:
// 		if err != nil {
// 			log.Fatalf("Failed to connect to Redis: %v", err)
// 			return nil
// 		}
// 		return rdb
// 	}
// }

// func InitRedisClient() *redis.Client {
// 	rdb := redis.NewClient(&redis.Options{
// 		Addr:     "localhost:6379", // Redis server address
// 		Password: "",               // No password set
// 		DB:       0,                // Use default DB
// 	})

// 	// Context with a 5-second timeout
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	// Ping the Redis server
// 	_, err := rdb.Ping(ctx).Result()
// 	if err != nil {
// 		log.Fatalf("Error connecting to Redis: %v", err)
// 	} else {
// 		fmt.Println("Connected to Redis successfully!")
// 	}

//		return rdb
//	}

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

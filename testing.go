package main

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"time"

// 	"github.com/go-redis/redis/v8"
// )

// func main() {
// 	// Create a Redis client
// 	rdb := redis.NewClient(&redis.Options{
// 		Addr:     "localhost:6379", // Use the correct Redis port
// 		Password: "",               // No password set
// 		DB:       0,                // Default DB
// 	})

// 	// Context with a 5-second timeout
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	// Ping the Redis server
// 	_, err := rdb.Ping(ctx).Result()
// 	if err != nil {
// 		log.Fatalf("Error connecting to Redis: %v", err)
// 	} else {
// 		fmt.Println("Connected to Redis successfully!")
// 	}
// }

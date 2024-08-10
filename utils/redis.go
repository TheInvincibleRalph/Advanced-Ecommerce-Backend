package utils

import "github.com/go-redis/redis"

//Intializes a Redis instance
func InitRedisClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:3001", // Redis server address
		Password: "",               // No password set
		DB:       0,                // Use default DB
	})

	return rdb
}

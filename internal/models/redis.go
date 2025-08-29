package models

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client
var ctx = context.Background()

func InitRedis() {
	redisURL := os.Getenv("REDIS_URL")
	if redisURL == "" {
		log.Fatal("REDIS_URL environment variable not set")
	}

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     redisURL,
		Password: "", // no password for development
		DB:       0,  // default DB
	})

	// Test Redis connection with retry
	maxRetries := 10
	for i := 0; i < maxRetries; i++ {
		_, err := RedisClient.Ping(ctx).Result()
		if err == nil {
			log.Println("Redis connected successfully")
			return
		}

		log.Printf("Failed to connect to Redis (attempt %d/%d): %v", i+1, maxRetries, err)

		if i < maxRetries-1 {
			time.Sleep(time.Duration(i+1) * 2 * time.Second)
		}
	}

	log.Fatal("Failed to connect to Redis after retries")
}

func StoreOTP(email string, otp string) error {
	err := RedisClient.Set(ctx, email, otp, 15*time.Minute).Err()
	return err
}

func VerifyOTP(email string, otp string) bool {
	storedOTP, err := RedisClient.Get(ctx, email).Result()
	if err != nil {
		return false
	}
	return storedOTP == otp
}

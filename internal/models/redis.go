package models

import (
	"context"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client
var ctx = context.Background()

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_URL"),
	})
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

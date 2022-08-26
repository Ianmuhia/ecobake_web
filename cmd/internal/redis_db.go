package internal

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

func GetRedisClient(redisURL string, redisDB int, logger *log.Logger) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Network:            "",
		Addr:               redisURL,
		Dialer:             nil,
		OnConnect:          nil,
		Username:           "",
		Password:           "",
		DB:                 redisDB,
		MaxRetries:         0,
		MinRetryBackoff:    0,
		MaxRetryBackoff:    0,
		DialTimeout:        0,
		ReadTimeout:        0,
		WriteTimeout:       0,
		PoolFIFO:           false,
		PoolSize:           0,
		MinIdleConns:       0,
		MaxConnAge:         0,
		PoolTimeout:        0,
		IdleTimeout:        0,
		IdleCheckFrequency: 0,
		TLSConfig:          nil,
		Limiter:            nil,
	})

	ctx := context.Background()

	_, err := client.Ping(ctx).Result()
	if err != nil {
		logger.Printf("error connecting to redis %s", err)
	}
	logger.Println("Redis database connected successfully...")
	return client
}

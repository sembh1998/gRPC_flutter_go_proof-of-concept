package rediscon

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisConnection struct {
	Conn *redis.Client
}

var singletonRedis *RedisConnection

func GetRedisConnection() *RedisConnection {
	if singletonRedis == nil {
		db, err := strconv.Atoi(os.Getenv("REDIS_DB"))
		if err != nil {
			panic(err)
		}
		conn := redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
			Password: os.Getenv(os.Getenv("REDIS_PASSWORD")),
			Username: os.Getenv(os.Getenv("REDIS_USER")),
			DB:       db,
		})
		singletonRedis = &RedisConnection{Conn: conn}
		log.Println("Redis connection created")
	}
	ctx := context.Background()
	timeOutCtx, cancel := context.WithTimeout(ctx, 50*time.Millisecond)
	defer cancel()
	_, err := singletonRedis.Conn.Ping(timeOutCtx).Result()
	if err != nil {
		fmt.Printf("Error connecting to redis: %v\n", err)
		singletonRedis = nil
		return nil
	}
	return singletonRedis
}

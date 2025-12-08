package redisClient

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
	"github.com/x-sushant-x/connective/internal/adapter/config"
)

type RedisClient struct {
	*redis.Client
}

func NewRedisClient(ctx context.Context) *RedisClient {
	c := config.Config

	conn := redis.NewClient(&redis.Options{
		Addr:         c.RedisAddr,
		Password:     c.RedisPassword,
		DB:           c.RedisDB,
		PoolSize:     c.RedisPoolSize,
		MinIdleConns: c.RedisMinIdleConns,
		PoolTimeout:  30 * time.Second,
	})

	_, err := conn.Ping(ctx).Result()
	if err != nil {
		log.Fatal().Err(err).Msg("Unable to connect to redis. Ping Failed")
	}

	return &RedisClient{
		conn,
	}
}

func (r *RedisClient) Set(ctx context.Context, key string, val any, expiry time.Duration) error {
	_, err := r.Client.Set(ctx, key, val, expiry).Result()
	return err
}

func (r *RedisClient) Get(ctx context.Context, key string) (string, error) {
	return r.Client.Get(ctx, key).Result()
}

func (r *RedisClient) SetJson(ctx context.Context, key string, val any, expiry time.Duration) error {
	_, err := r.Client.JSONSet(ctx, key, ".", val).Result()
	return err
}

func (r *RedisClient) GetJson(ctx context.Context, key string) (string, error) {
	return r.Client.JSONGet(ctx, key).Result()
}

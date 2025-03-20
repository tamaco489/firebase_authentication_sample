package repository

import (
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/tamaco489/firebase_authentication_sample/api/core/internal/configuration"
)

type redisClient struct {
	client *redis.Client
}

func NewRedis() *redisClient {
	return &redisClient{
		client: redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", configuration.Get().CoreRedis.Host, configuration.Get().CoreRedis.Port),
			Password: configuration.Get().CoreRedis.Password,
			DB:       0,
			PoolSize: configuration.Get().CoreRedis.PoolSize,
		}),
	}
}

func (rc redisClient) GetClient() *redis.Client {
	return rc.client
}

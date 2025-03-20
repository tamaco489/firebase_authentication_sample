package repository

import (
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/tamaco489/firebase_authentication_sample/api/core/internal/configuration"
)

const redisTTLSeconds = 30

type RedisConnector struct {
	Client *redis.Client
}

func NewRedis() *RedisConnector {
	return &RedisConnector{
		Client: redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", configuration.Get().CoreRedis.Host, configuration.Get().CoreRedis.Port),
			Password: "",
			DB:       0,
			PoolSize: 10,
		}),
	}
}

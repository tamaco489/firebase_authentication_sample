package controller

import (
	"database/sql"

	"github.com/redis/go-redis/v9"
	"github.com/tamaco489/firebase_authentication_sample/api/core/internal/configuration"
	"github.com/tamaco489/firebase_authentication_sample/api/core/internal/usecase"

	repository_gen_sqlc "github.com/tamaco489/firebase_authentication_sample/api/core/internal/repository/gen_sqlc"
)

type Controllers struct {
	config      configuration.Config
	userUseCase usecase.IUserUseCase
}

func NewCoreControllers(
	cfg configuration.Config,
	db *sql.DB,
	queries repository_gen_sqlc.Queries,
	redisClient *redis.Client,
) (*Controllers, error) {
	userUseCase := usecase.NewUserUseCase(db, queries, db, redisClient)
	return &Controllers{
		config:      cfg,
		userUseCase: userUseCase,
	}, nil
}

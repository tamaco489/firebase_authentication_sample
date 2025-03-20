package controller

import (
	"database/sql"

	"github.com/tamaco489/firebase_authentication_sample/api/core/internal/configuration"
	"github.com/tamaco489/firebase_authentication_sample/api/core/internal/usecase"
)

type Controllers struct {
	config      configuration.Config
	db          *sql.DB
	userUseCase usecase.IUserUseCase
}

func NewCoreControllers(
	cfg configuration.Config,
	db *sql.DB,
) (*Controllers, error) {
	userUseCase := usecase.NewUserUseCase()
	return &Controllers{
		config:      cfg,
		db:          db,
		userUseCase: userUseCase,
	}, nil
}

package controller

import (
	"github.com/tamaco489/firebase_authentication_sample/api/core/internal/configuration"
	"github.com/tamaco489/firebase_authentication_sample/api/core/internal/usecase"
)

type Controllers struct {
	config      configuration.Config
	userUseCase usecase.IUserUseCase
}

func NewCoreControllers(cfg configuration.Config) (*Controllers, error) {
	userUseCase := usecase.NewUserUseCase()
	return &Controllers{
		config:      cfg,
		userUseCase: userUseCase,
	}, nil
}

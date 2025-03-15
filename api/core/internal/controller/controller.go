package controller

import "github.com/tamaco489/firebase_authentication_sample/api/core/internal/configuration"

type Controllers struct {
	config configuration.Config
}

func NewCoreControllers(cfg configuration.Config) (*Controllers, error) {
	return &Controllers{config: cfg}, nil
}

package controller

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/tamaco489/firebase_authentication_sample/api/core/internal/gen"
)

func (c *Controllers) CreateUser(ctx *gin.Context, request gen.CreateUserRequestObject) (gen.CreateUserResponseObject, error) {

	uid := "123e4567-e89b-12d3-a456-426614174000"
	slog.InfoContext(ctx, "debug log", slog.String("provider_type", string(request.Body.ProviderType)))

	return &gen.CreateUser201JSONResponse{Uid: uid}, nil
}

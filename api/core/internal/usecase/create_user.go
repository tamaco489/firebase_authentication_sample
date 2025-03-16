package usecase

import (
	"fmt"
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/tamaco489/firebase_authentication_sample/api/core/internal/gen"
)

func (u *userUseCase) CreateUser(ctx *gin.Context, request gen.CreateUserRequestObject) (gen.CreateUserResponseObject, error) {

	slog.InfoContext(ctx, "debug log", slog.String("provider_type", string(request.Body.ProviderType)))

	uuid, err := uuid.NewV7()
	if err != nil {
		return gen.CreateUser500Response{}, fmt.Errorf("failed to new uuid: %w", err)
	}

	return gen.CreateUser201JSONResponse{Uid: uuid.String()}, nil
}

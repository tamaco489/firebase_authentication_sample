package usecase

import (
	"github.com/gin-gonic/gin"
	"github.com/tamaco489/firebase_authentication_sample/api/core/internal/gen"
)

type userUseCase struct{}

type IUserUseCase interface {
	CreateUser(ctx *gin.Context, request gen.CreateUserRequestObject) (gen.CreateUserResponseObject, error)
}

var _ IUserUseCase = (*userUseCase)(nil)

func NewUserUseCase() IUserUseCase {
	return &userUseCase{}
}

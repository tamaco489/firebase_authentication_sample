package controller

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/tamaco489/firebase_authentication_sample/api/core/internal/gen"
	"github.com/tamaco489/firebase_authentication_sample/api/core/internal/utils/ctx_utils"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func (c *Controllers) CreateUser(ctx *gin.Context, request gen.CreateUserRequestObject) (gen.CreateUserResponseObject, error) {

	err := validation.ValidateStruct(request.Body,
		validation.Field(
			&request.Body.ProviderType,
			validation.Required,
		),
	)
	if err != nil {
		_ = ctx.Error(err)
		return gen.CreateUser400Response{}, nil
	}

	sub, ok := ctx_utils.GetFirebaseUID(ctx)
	if !ok {
		_ = ctx.Error(errors.New("failed to get sub from context"))
		return gen.GetMe401Response{}, nil
	}

	res, err := c.userUseCase.CreateUser(ctx, sub, request)
	if err != nil {
		return gen.CreateUser500Response{}, err
	}

	return res, nil
}

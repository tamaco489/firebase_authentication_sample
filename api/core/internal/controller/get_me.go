package controller

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/tamaco489/firebase_authentication_sample/api/core/internal/gen"
	"github.com/tamaco489/firebase_authentication_sample/api/core/internal/utils/ctx_utils"
)

func (c *Controllers) GetMe(ctx *gin.Context, request gen.GetMeRequestObject) (gen.GetMeResponseObject, error) {

	uid, ok := ctx_utils.GetCoreUID(ctx)
	if !ok {
		_ = ctx.Error(errors.New("failed to get uid from context"))
		return gen.GetMe401Response{}, nil
	}

	res, err := c.userUseCase.GetMe(ctx, uid, request)
	if err != nil {
		return gen.GetMe500Response{}, err
	}
	return res, nil
}

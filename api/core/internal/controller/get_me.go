package controller

import (
	"errors"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/tamaco489/firebase_authentication_sample/api/core/internal/gen"
	"github.com/tamaco489/firebase_authentication_sample/api/core/internal/utils/ctx_utils"
)

func (c *Controllers) GetMe(ctx *gin.Context, request gen.GetMeRequestObject) (gen.GetMeResponseObject, error) {

	sub, ok := ctx_utils.GetFirebaseUID(ctx)
	if !ok {
		_ = ctx.Error(errors.New("failed to get sub from context"))
		return gen.GetMe401Response{}, nil
	}

	uid, ok := ctx_utils.GetCoreUID(ctx)
	if !ok {
		_ = ctx.Error(errors.New("failed to get uid from context"))
		return gen.GetMe401Response{}, nil
	}

	provider, ok := ctx_utils.GetFirebaseProviderType(ctx)
	if !ok {
		_ = ctx.Error(errors.New("failed to get provider from context"))
		return gen.GetMe401Response{}, nil
	}

	// todo: 検証後削除
	log.Println("[info] sub:", sub, "uid:", uid, "provider:", provider)

	res, err := c.userUseCase.GetMe(ctx, request)
	if err != nil {
		return gen.GetMe500Response{}, err
	}
	return res, nil
}

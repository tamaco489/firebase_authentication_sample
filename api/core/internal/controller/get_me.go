package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tamaco489/firebase_authentication_sample/api/core/internal/gen"
)

func (c *Controllers) GetMe(ctx *gin.Context, request gen.GetMeRequestObject) (gen.GetMeResponseObject, error) {
	res, err := c.userUseCase.GetMe(ctx, request)
	if err != nil {
		return gen.GetMe500Response{}, err
	}
	return res, nil
}

package usecase

import (
	"github.com/gin-gonic/gin"
	"github.com/tamaco489/firebase_authentication_sample/api/core/internal/gen"
)

func (u *userUseCase) GetMe(ctx *gin.Context, request gen.GetMeRequestObject) (gen.GetMeResponseObject, error) {

	// NOTE: MySQL出来上がるまでは一旦固定値返す
	uuid := "01959e82-2a8a-7b55-87ee-a9c1d9582c9d"

	return gen.GetMe200JSONResponse{Uid: uuid}, nil
}

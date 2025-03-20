package usecase

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/tamaco489/firebase_authentication_sample/api/core/internal/gen"
)

func (u *userUseCase) GetMe(ctx *gin.Context, request gen.GetMeRequestObject) (gen.GetMeResponseObject, error) {

	// NOTE: Redis検証のため一時的にPing送信の処理入れてる
	p, err := u.redisClient.Ping(ctx).Result()
	if err != nil {
		return gen.CreateUser500Response{}, fmt.Errorf("failed to redis ping: %w", err)
	}

	log.Println("[info] redis ping success:", p)

	// NOTE: MySQL出来上がるまでは一旦固定値返す
	uuid := "01959e82-2a8a-7b55-87ee-a9c1d9582c9d"

	return gen.GetMe200JSONResponse{Uid: uuid}, nil
}

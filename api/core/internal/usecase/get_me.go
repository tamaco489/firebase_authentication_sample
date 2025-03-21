package usecase

import (
	"database/sql"
	"errors"
	"fmt"
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/tamaco489/firebase_authentication_sample/api/core/internal/domain/auth"
	"github.com/tamaco489/firebase_authentication_sample/api/core/internal/gen"
)

func (u *userUseCase) GetMe(ctx *gin.Context, request gen.GetMeRequestObject) (gen.GetMeResponseObject, error) {

	// NOTE: 本来であればmiddleware上でjwtを解析して得たものをusecase上で使用する
	// sub := "hoge"
	sub := "2iSI3im4bcOFJDoT7E9QLebbU9G2"

	// redisからセッション情報を取得する
	session := auth.NewGetSession(sub)
	if err := session.Get(ctx, u.redisClient); err != nil {
		return gen.GetMe500Response{}, err
	}

	// redisにセッション情報が存在している場合はこの時点で200で返してしまう
	if session.UID != "" {
		return gen.GetMe200JSONResponse{Uid: session.UID}, nil
	}

	// redisにセッション情報が存在しない場合はMySQLにアクセスしuidを取得する
	uid, err := u.queries.GetUIDByFirebaseUID(ctx, u.db, sub)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return gen.GetMe500Response{}, fmt.Errorf("failed to get uid by firebase uid: %w", err)
	}
	if uid == "" {
		slog.ErrorContext(ctx, "not exists firebase user.", slog.String("sub", sub))
		return gen.GetMe404Response{}, nil
	}

	// redisにセッション情報を登録する
	// NOTE: 本来であればmiddleware上でjwtを解析して得たものをusecase上で使用する
	authTime := int64(1742064672)
	expire := int64(1742068272)
	newSession := auth.NewSaveSession(sub, authTime, expire, uid, "firebase") // providerも一旦固定値
	if err := newSession.Save(ctx, u.redisClient); err != nil {
		return gen.GetMe500Response{}, err
	}

	return gen.GetMe200JSONResponse{Uid: uid}, nil
}

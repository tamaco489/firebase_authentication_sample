package usecase

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"

	"github.com/tamaco489/firebase_authentication_sample/api/core/internal/domain/auth"
	"github.com/tamaco489/firebase_authentication_sample/api/core/internal/gen"
	"github.com/tamaco489/firebase_authentication_sample/api/core/internal/utils/ctx_utils"
)

func (u *userUseCase) GetMe(ctx context.Context, uid, sub string, request gen.GetMeRequestObject) (gen.GetMeResponseObject, error) {

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
		slog.ErrorContext(ctx, "not exists firebase user", slog.String("sub", sub))
		return gen.GetMe404Response{}, nil
	}

	// redisにセッション情報を登録する
	newSession := auth.NewSaveSession(sub, uid, ctx_utils.FirebaseProviderKey.String())
	if err := newSession.Save(ctx, u.redisClient); err != nil {
		return gen.GetMe500Response{}, err
	}

	return gen.GetMe200JSONResponse{Uid: uid}, nil
}

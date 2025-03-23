package usecase

import (
	"context"

	"github.com/tamaco489/firebase_authentication_sample/api/core/internal/gen"
)

func (u *userUseCase) GetMe(ctx context.Context, uid string, request gen.GetMeRequestObject) (gen.GetMeResponseObject, error) {
	return gen.GetMe200JSONResponse{Uid: uid}, nil
}

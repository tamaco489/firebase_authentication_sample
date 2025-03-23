package usecase

import (
	"context"
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/tamaco489/firebase_authentication_sample/api/core/internal/gen"

	repository_gen_sqlc "github.com/tamaco489/firebase_authentication_sample/api/core/internal/repository/gen_sqlc"
)

type userUseCase struct {
	db          *sql.DB
	queries     repository_gen_sqlc.Queries
	dbtx        repository_gen_sqlc.DBTX
	redisClient *redis.Client
}

type IUserUseCase interface {
	CreateUser(ctx *gin.Context, request gen.CreateUserRequestObject) (gen.CreateUserResponseObject, error)
	GetMe(ctx context.Context, uid, sub string, request gen.GetMeRequestObject) (gen.GetMeResponseObject, error)
}

var _ IUserUseCase = (*userUseCase)(nil)

func NewUserUseCase(
	db *sql.DB,
	queries repository_gen_sqlc.Queries,
	dbtx repository_gen_sqlc.DBTX,
	redisClient *redis.Client,
) IUserUseCase {
	return &userUseCase{
		db:          db,
		queries:     queries,
		dbtx:        dbtx,
		redisClient: redisClient,
	}
}

package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/tamaco489/firebase_authentication_sample/api/core/internal/configuration"
	"github.com/tamaco489/firebase_authentication_sample/api/core/internal/gen"
	"github.com/tamaco489/firebase_authentication_sample/api/core/internal/library/logger"

	middleware "github.com/tamaco489/firebase_authentication_sample/api/core/internal/middleware/authorization"
	repository_gen_sqlc "github.com/tamaco489/firebase_authentication_sample/api/core/internal/repository/gen_sqlc"
	repository_store "github.com/tamaco489/firebase_authentication_sample/api/core/internal/repository/store"
)

func NewCoreAPIServer(cnf configuration.Config) (*http.Server, error) {
	corsCfg := NewCorsConfig()

	r := gin.New()
	r.Use(gin.LoggerWithFormatter(logger.LogFormatter))
	r.Use(cors.New(corsCfg))
	r.Use(gin.Recovery())

	// new mysql
	db := repository_store.InitDB()
	queries := repository_gen_sqlc.New()

	// new redis
	redis := repository_store.NewRedis()
	redisClient := redis.GetClient()

	// 認可(JWT検証)
	r.Use(middleware.JWTAuthMiddleware(db, *queries, redisClient))

	// new controller
	apiController, err := NewCoreControllers(cnf, db, *queries, redisClient)
	if err != nil {
		return nil, fmt.Errorf("failed to new controllers %v", err)
	}

	strictServer := gen.NewStrictHandler(apiController, nil)

	gen.RegisterHandlersWithOptions(
		r,
		strictServer,
		gen.GinServerOptions{
			BaseURL:     "/core/",
			Middlewares: []gen.MiddlewareFunc{},
			ErrorHandler: func(ctx *gin.Context, err error, i int) {
				_ = ctx.Error(err)
				ctx.JSON(i, gin.H{"msg": err.Error()})
			},
		},
	)

	server := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf(":%s", cnf.API.Port),
	}

	return server, nil
}

func NewCorsConfig() cors.Config {
	return cors.Config{
		// 許可するオリジンを指定（一旦全許可）
		AllowOrigins: []string{"*"},

		// 必要なメソッドのみ許可
		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"DELETE",
			"HEAD",
			"OPTIONS",
		},

		// 許可するヘッダーを限定
		AllowHeaders: []string{
			"Origin",
			"Content-Length",
			"Content-Type",
			"Authorization",
			"Access-Control-Allow-Origin",
		},

		// クライアントがアクセスできるレスポンスヘッダー
		ExposeHeaders: []string{"Content-Length"},

		// 認証情報を送信可能にする
		AllowCredentials: false,

		// プリフライトリクエストのキャッシュ時間（秒）
		MaxAge: 86400,
	}
}

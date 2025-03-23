package middleware

import (
	"database/sql"
	"errors"
	"log/slog"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/tamaco489/firebase_authentication_sample/api/core/internal/configuration"
	"github.com/tamaco489/firebase_authentication_sample/api/core/internal/domain/auth"
	"github.com/tamaco489/firebase_authentication_sample/api/core/internal/library/firebase"
	"github.com/tamaco489/firebase_authentication_sample/api/core/internal/utils/ctx_utils"

	repository_gen_sqlc "github.com/tamaco489/firebase_authentication_sample/api/core/internal/repository/gen_sqlc"
)

const (
	healthCheckEndpoint = "/core/v1/healthcheck"
	createUserEndpoint  = "/core/v1/users"
)

// JWTAuthMiddleware:
func JWTAuthMiddleware(db *sql.DB, queries repository_gen_sqlc.Queries, redisClient *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		// healthcheckの場合は検証をスキップ
		if c.Request.Method == http.MethodGet && c.Request.URL.Path == healthCheckEndpoint {
			c.Next()
			return
		}

		// Authorization ヘッダーの取得
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			slog.WarnContext(c.Request.Context(), "authorization header is required")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": "authorization header is required"})
			return
		}

		// "Bearer {token}" 形式のチェック
		idToken := strings.TrimPrefix(authHeader, "Bearer ")
		if idToken == authHeader {
			slog.WarnContext(c.Request.Context(), "failed to extract authorization header")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": "failed to extract authorization header"})
			return
		}

		// Firebase Admin SDK の初期化、及びインスタンス生成
		firebaseClient, err := firebase.NewFirebaseClient(c.Request.Context(), configuration.Get().Firebase.GoogleServiceAccount)
		if err != nil {
			slog.ErrorContext(c.Request.Context(), "failed to initialize firebase client", slog.String("error", err.Error()))
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": "failed to initialize firebase client"})
			return
		}

		// jwt検証（フォーマットが正しいかどうか、有効期限内かどうか、署名が正しいかどうか）
		token, err := firebaseClient.VerifyIDToken(c.Request.Context(), idToken)
		if err != nil {
			slog.ErrorContext(c.Request.Context(), "invalid or expired token", slog.String("error", err.Error()))
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": "invalid or expired token"})
			return
		}

		// ユーザ新規登録の処理 （検証済みのsubとprovider_typeをctxに設定して返してしまう）
		if c.Request.Method == http.MethodPost && c.Request.URL.Path == createUserEndpoint {
			ctx_utils.SetFirebaseUID(c, token.Subject)
			ctx_utils.SetFirebaseProviderType(c, ctx_utils.FirebaseProviderKey.String())
			c.Next()
			return
		}

		// 取得したsubをkeyにしてredisからセッション情報を取得（ユーザ新規登録API以外のエンドポイント共通処理）
		var uid string
		session := auth.NewGetSession(token.Subject)
		if err := session.Get(c.Request.Context(), redisClient); err != nil {
			slog.ErrorContext(c.Request.Context(), "failed to new get session", slog.String("error", err.Error()))
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": "failed to new get session"})
			return
		}

		// セッションが存在している場合: MySQLにアクセスせずに、セッションからuidを取得
		if session.UID != "" {
			uid = session.UID
		}

		// セッションが存在していない場合: MySQLにアクセスし、subをkeyにしてuidを取得し、再度Redisにセッションを保存
		if session.UID == "" {
			// MySQLにアクセスし、subをkeyにしてuidを取得
			uid, err = queries.GetUIDByFirebaseUID(c.Request.Context(), db, token.Subject)
			if err != nil && !errors.Is(err, sql.ErrNoRows) {
				slog.ErrorContext(c.Request.Context(), "failed to fetch firebase uid by mysql", slog.String("error", err.Error()))
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": "failed to fetch firebase uid by mysql"})
				return
			}

			// MySQLにも存在しなかった場合は404エラーを返す
			if uid == "" {
				slog.ErrorContext(c.Request.Context(), "not exists user", slog.String("error", err.Error()))
				c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"msg": "not exists user"})
				return
			}

			// Redisに再度セッションを保存
			newSession := auth.NewSaveSession(token.Subject, uid, ctx_utils.FirebaseProviderKey.String())
			if err := newSession.Save(c.Request.Context(), redisClient); err != nil {
				slog.ErrorContext(c.Request.Context(), "failed to save session to redis", slog.String("error", err.Error()))
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": "failed to save session to redis"})
				return
			}
		}

		// contextにuid、sub、providerを入れる
		if token.Firebase.SignInProvider != "" {
			ctx_utils.SetFirebaseUID(c, token.Subject)
			ctx_utils.SetFirebaseProviderType(c, ctx_utils.FirebaseProviderKey.String())
			ctx_utils.SetCoreUID(c, uid)
		}

		c.Next()
	}
}

package middleware

import (
	"log/slog"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/tamaco489/firebase_authentication_sample/api/core/internal/configuration"
	"github.com/tamaco489/firebase_authentication_sample/api/core/internal/domain/auth"
	"github.com/tamaco489/firebase_authentication_sample/api/core/internal/library/firebase"
	"github.com/tamaco489/firebase_authentication_sample/api/core/internal/utils/ctx_utils"
)

const healthCheckEndpoint = "/core/v1/healthcheck"

// JWTAuthMiddleware:
func JWTAuthMiddleware(redisClient *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		// *************** [1. Authorization Headerのチェック] ***************
		// healthcheckの場合は検証をスキップ
		if c.Request.Method == http.MethodGet && c.Request.URL.Path == healthCheckEndpoint {
			c.Next()
			return
		}

		// Authorization ヘッダーの取得
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			slog.WarnContext(c.Request.Context(), "authorization header is required")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// "Bearer {token}" 形式のチェック
		idToken := strings.TrimPrefix(authHeader, "Bearer ")
		if idToken == authHeader {
			slog.WarnContext(c.Request.Context(), "failed to extract authorization header")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Firebase Admin SDK の初期化、及びインスタンス生成
		firebaseClient, err := firebase.NewFirebaseClient(c.Request.Context(), configuration.Get().Firebase.GoogleServiceAccount)
		if err != nil {
			slog.ErrorContext(c.Request.Context(), "failed to initialize Firebase client", slog.String("error", err.Error()))
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		// jwt検証（フォーマットが正しいかどうか、有効期限内かどうか、署名が正しいかどうか）
		token, err := firebaseClient.VerifyIDToken(c.Request.Context(), idToken)
		if err != nil {
			// トークン検証失敗
			slog.ErrorContext(c.Request.Context(), "invalid or expired token", slog.String("error", err.Error()))
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// todo: 検証後削除
		// 検証が成功した場合、トークンから情報を取り出して後続の処理に利用
		slog.InfoContext(c.Request.Context(), "ID token verified",
			slog.String("sub", token.Subject),
			slog.String("uid", token.UID),
			slog.Int64("auth_time", token.AuthTime),
			slog.Int64("expire", token.Expires),
			slog.Int64("issued_at", token.IssuedAt),
			slog.Any("firebase_info", token.Firebase),
		)

		// 取得したsubをkeyにしてredisからセッション情報を取得
		var uid string
		session := auth.NewGetSession(token.Subject)
		if err := session.Get(c.Request.Context(), redisClient); err != nil {
			slog.ErrorContext(c.Request.Context(), "failed to new get session", slog.String("error", err.Error()))
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		// セッションが存在している場合: MySQLにアクセスせずに、セッションからuidを取得
		if session.UID != "" {
			uid = session.UID
		}

		// todo: 7. セッションが存在していない場合: Redisに再度セッションを保存した上でMySQLにアクセスし、subをkeyにしてuidを取得
		if session.UID == "" {
			// 0195b45e-a7e3-7572-b2b8-e85247c422b8
		}

		// todo: 8. contextにuid、sub、providerを入れる
		if token.Firebase.SignInProvider != "" {
			ctx_utils.SetFirebaseUID(c, token.Subject)
			ctx_utils.SetFirebaseProviderType(c, ctx_utils.FirebaseProviderKey.String())
			ctx_utils.SetCoreUID(c, uid)
		}

		c.Next()
	}
}

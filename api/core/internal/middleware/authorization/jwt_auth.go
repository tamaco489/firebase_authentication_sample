package middleware

import (
	"log/slog"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tamaco489/firebase_authentication_sample/api/core/internal/configuration"
	"github.com/tamaco489/firebase_authentication_sample/api/core/internal/library/firebase"
)

const healthCheckEndpoint = "/core/v1/healthcheck"

// JWTAuthMiddleware:
func JWTAuthMiddleware() gin.HandlerFunc {
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
			c.JSON(http.StatusUnauthorized, gin.H{"msg": "Authorization header is required"})
			c.Abort()
			return
		}

		// "Bearer {token}" 形式のチェック
		idToken := strings.TrimPrefix(authHeader, "Bearer ")
		if idToken == authHeader {
			c.JSON(http.StatusUnauthorized, gin.H{"msg": "Invalid token format"})
			c.Abort()
			return
		}

		// todo: 1. Firebase Admin SDK の初期化、及びインスタンス生成
		firebaseClient, err := firebase.NewFirebaseClient(c.Request.Context(), configuration.Get().Firebase.GoogleServiceAccount)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": "Failed to initialize Firebase client"})
			c.Abort()
			return
		}

		// todo: 2. jwtフォーマットの検証
		token, err := firebaseClient.VerifyIDToken(c.Request.Context(), idToken)
		if err != nil {
			// トークン検証失敗
			c.JSON(http.StatusUnauthorized, gin.H{"msg": "Invalid or expired token"})
			c.Abort()
			return
		}

		// 検証が成功した場合、トークンから情報を取り出して後続の処理に利用
		slog.InfoContext(c.Request.Context(), "ID token verified",
			slog.String("sub", token.Subject),
			slog.String("uid", token.UID),
			slog.Int64("auth_time", token.AuthTime),
			slog.Int64("expire", token.Expires),
			slog.Int64("issued_at", token.IssuedAt),
			slog.Any("firebase_info", token.Firebase),
		)

		// todo: 3. 検証済みのjwtを解析して、roleを確認

		// todo: 4. 検証済みのjwtを解析して、subを取得する

		// todo: 5. 取得したsubをkeyにしてredisからセッション情報を取得し、有効期限内かの判定を行う。

		// todo: 6. セッションが有効期限内: MySQLにアクセスせずに、セッションからuidを取得

		// todo: 7. セッションが有効期限切れ: Redisに再度セッションを保存した上でMySQLにアクセスし、subをkeyにしてuidを取得

		// todo: 8. contextにuid、sub、roleを入れる

		c.Next()
	}
}

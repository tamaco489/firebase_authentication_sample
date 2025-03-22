package middleware

import (
	"log/slog"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
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

		// todo: 検証後削除
		slog.InfoContext(c.Request.Context(), "success fetch id_token", slog.String("id_token", idToken))

		// todo: 2. jwtの署名を検証（Firebase Admin SDKを使用する）

		// todo: 3. jwtフォーマットの検証

		// todo: 4. 検証済みのjwtを解析して、roleを確認

		// todo: 5. 検証済みのjwtを解析して、subを取得する

		// todo: 6. 取得したsubをkeyにしてredisからセッション情報を取得し、有効期限内かの判定を行う。

		// todo: 7. セッションが有効期限内: MySQLにアクセスせずに、セッションからuidを取得

		// todo: 8. セッションが有効期限切れ: Redisに再度セッションを保存した上でMySQLにアクセスし、subをkeyにしてuidを取得

		// todo: 9. contextにuid、sub、roleを入れる

		c.Next()
	}
}

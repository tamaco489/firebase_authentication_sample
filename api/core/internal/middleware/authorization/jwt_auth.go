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

		c.Next()
	}
}

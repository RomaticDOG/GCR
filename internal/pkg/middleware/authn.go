package middleware

import (
	"github.com/RomaticDOG/GCR/FastGO/internal/pkg/contextx"
	"github.com/RomaticDOG/GCR/FastGO/internal/pkg/core"
	"github.com/RomaticDOG/GCR/FastGO/internal/pkg/errorsx"
	"github.com/RomaticDOG/GCR/FastGO/pkg/token"
	"github.com/gin-gonic/gin"
)

// AuthN 是认证中间件，用来从 gin.Context 中提取 token 并验证 token 是否合法
// 如果合法则将 token 中的 sub 作为用户名存放在 gin.Context 的 XUsernameKey 键中
func AuthN() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 解析 JWT token
		userID, err := token.ParseRequest(c)
		if err != nil {
			core.WriteResponse(c, errorsx.ErrInvalidToken, nil)
			c.Abort()
			return
		}
		ctx := contextx.WithUserID(c.Request.Context(), userID)
		c.Request = c.Request.WithContext(ctx)
		// 继续后续的操作
		c.Next()
	}
}

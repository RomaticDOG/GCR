package middleware

import (
	"github.com/RomaticDOG/GCR/FastGO/internal/pkg/contextx"
	"github.com/RomaticDOG/GCR/FastGO/internal/pkg/known"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// RequestID 是一个 Gin 中间件，用来在每一个 http 请求的 context，response 中注入 x-request-id 键值对
func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.Request.Header.Get(known.XRequestID)
		if requestID == "" {
			requestID = uuid.New().String()
		}
		// 将 requestID 保存到上下文中，方便调用
		ctx := contextx.WithRequestID(c.Request.Context(), requestID)
		c.Request = c.Request.WithContext(ctx)

		// 将 requestID 保存到 HTTP 返回头中，Header 中的键名称为 X-Request-ID
		c.Writer.Header().Set(known.XRequestID, requestID)

		// 继续处理请求
		c.Next()
	}
}

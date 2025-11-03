package core

import (
	"net/http"

	"github.com/RomaticDOG/GCR/FastGO/internal/pkg/errorsx"
	"github.com/gin-gonic/gin"
)

// ErrorResponse 定义了错误响应的结构，用于 API 请求响应中返回统一的错误信息
type ErrorResponse struct {
	Reason  string `json:"reason"`
	Message string `json:"message"`
}

// WriteResponse 是通用的响应函数，它会根据是否发生错误，生成成功响应或通用的失败响应
func WriteResponse(c *gin.Context, err error, data any) {
	if err != nil {
		ex := errorsx.FromError(err)
		c.JSON(ex.Code, ErrorResponse{
			Reason:  ex.Reason,
			Message: ex.Message,
		})
		return
	}
	c.JSON(http.StatusOK, data)
}

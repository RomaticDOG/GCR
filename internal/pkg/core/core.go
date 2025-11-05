package core

import (
	"net/http"

	"github.com/RomaticDOG/GCR/FastGO/internal/pkg/errorsx"
	"github.com/gin-gonic/gin"
)

// Response 用于 API 请求响应中返回统一的信息
type Response struct {
	Code    string      `json:"code"`
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// WriteResponse 是通用的响应函数，它会根据是否发生错误，生成成功响应或通用的失败响应
func WriteResponse(c *gin.Context, err error, data any) {
	if err != nil {
		ex := errorsx.FromError(err)
		c.JSON(ex.Code, Response{
			Code:    ex.Reason,
			Status:  false,
			Message: ex.Message,
			Data:    nil,
		})
		return
	}
	c.JSON(http.StatusOK, Response{
		Code:    "StatusOK",
		Status:  true,
		Message: "",
		Data:    data,
	})
}

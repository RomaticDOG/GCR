package handler

import (
	"log/slog"

	"github.com/RomaticDOG/GCR/FastGO/internal/pkg/core"
	"github.com/RomaticDOG/GCR/FastGO/internal/pkg/errorsx"
	apiV1 "github.com/RomaticDOG/GCR/FastGO/pkg/api/v1"
	"github.com/gin-gonic/gin"
)

// CreateUser 创建新用户
func (h *Handler) CreateUser(c *gin.Context) {
	slog.Info("Create user function called.")
	var req apiV1.CreateUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		core.WriteResponse(c, errorsx.ErrBind, nil)
		return
	}
	if err := h.v.ValidateCreateUserReq(c.Request.Context(), &req); err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	resp, err := h.biz.UserV1().Create(c, &req)
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	core.WriteResponse(c, nil, resp)
}

func (h *Handler) UpdateUser(c *gin.Context) {

}

func (h *Handler) DeleteUser(c *gin.Context) {

}

func (h *Handler) GetUser(c *gin.Context) {

}

func (h *Handler) ListUser(c *gin.Context) {

}

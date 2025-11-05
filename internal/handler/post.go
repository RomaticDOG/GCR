package handler

import (
	"log/slog"

	"github.com/RomaticDOG/GCR/FastGO/internal/pkg/core"
	"github.com/RomaticDOG/GCR/FastGO/internal/pkg/errorsx"
	apiV1 "github.com/RomaticDOG/GCR/FastGO/pkg/api/v1"
	"github.com/gin-gonic/gin"
)

// CreatePost 创建新用户
func (h *Handler) CreatePost(c *gin.Context) {
	slog.Info("Create user function called.")
	var req apiV1.CreatePostReq
	if err := c.ShouldBindJSON(&req); err != nil {
		core.WriteResponse(c, errorsx.ErrBind, nil)
		return
	}
	if err := h.v.ValidateCreatePostReq(c.Request.Context(), &req); err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	resp, err := h.biz.PostV1().Create(c, &req)
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	core.WriteResponse(c, nil, resp)
}

func (h *Handler) UpdatePost(c *gin.Context) {

}

func (h *Handler) DeletePost(c *gin.Context) {

}

func (h *Handler) GetPost(c *gin.Context) {

}

func (h *Handler) ListPost(c *gin.Context) {

}

package handler

import (
	"log/slog"

	"github.com/RomaticDOG/GCR/FastGO/internal/pkg/core"
	"github.com/RomaticDOG/GCR/FastGO/internal/pkg/errorsx"
	apiV1 "github.com/RomaticDOG/GCR/FastGO/pkg/api/v1"
	"github.com/gin-gonic/gin"
)

// CreatePost 创建新博文
func (h *Handler) CreatePost(c *gin.Context) {
	slog.Info("Create post function called.")
	var req apiV1.CreatePostReq
	if err := c.ShouldBindJSON(&req); err != nil {
		core.WriteResponse(c, errorsx.ErrBind, nil)
		return
	}
	if err := h.v.ValidateCreatePostReq(c.Request.Context(), &req); err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	resp, err := h.biz.PostV1().Create(c.Request.Context(), &req)
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	core.WriteResponse(c, nil, resp)
}

// UpdatePost 更新博文
func (h *Handler) UpdatePost(c *gin.Context) {
	slog.Info("Update post function called.")
	var req apiV1.UpdatePostReq
	if err := c.ShouldBindJSON(&req); err != nil {
		core.WriteResponse(c, errorsx.ErrBind, nil)
		return
	}
	req.PostID = c.Param("postID")
	if err := h.v.ValidateUpdatePostReq(c.Request.Context(), &req); err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	resp, err := h.biz.PostV1().Update(c.Request.Context(), &req)
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	core.WriteResponse(c, nil, resp)
}

// DeletePost 删除博文
func (h *Handler) DeletePost(c *gin.Context) {
	slog.Info("Delete post function called.")
	var req apiV1.DeletePostReq
	if err := c.ShouldBindJSON(&req); err != nil {
		core.WriteResponse(c, errorsx.ErrBind, nil)
		return
	}
	if err := h.v.ValidateDeletePostReq(c.Request.Context(), &req); err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	resp, err := h.biz.PostV1().Delete(c.Request.Context(), &req)
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	core.WriteResponse(c, nil, resp)
}

// GetPost 获取博文
func (h *Handler) GetPost(c *gin.Context) {
	slog.Info("Get post function called.")
	req := apiV1.GetPostReq{
		PostID: c.Param("postID"),
	}
	if err := h.v.ValidateGetPostReq(c.Request.Context(), &req); err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	resp, err := h.biz.PostV1().Get(c.Request.Context(), &req)
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	core.WriteResponse(c, nil, resp)
}

// ListPost 获取博文列表
func (h *Handler) ListPost(c *gin.Context) {
	slog.Info("List post function called.")
	var req apiV1.ListPostReq
	if err := c.ShouldBindJSON(&req); err != nil {
		core.WriteResponse(c, errorsx.ErrBind, nil)
		return
	}
	if err := h.v.ValidateListPostReq(c.Request.Context(), &req); err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	resp, err := h.biz.PostV1().List(c.Request.Context(), &req)
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	core.WriteResponse(c, nil, resp)
}

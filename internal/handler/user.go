package handler

import (
	"log/slog"

	"github.com/RomaticDOG/GCR/FastGO/internal/pkg/core"
	"github.com/RomaticDOG/GCR/FastGO/internal/pkg/errorsx"
	apiV1 "github.com/RomaticDOG/GCR/FastGO/pkg/api/v1"
	"github.com/gin-gonic/gin"
)

// Login 登陆
func (h *Handler) Login(c *gin.Context) {
	slog.Info("Login function called.")
	var req apiV1.LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		core.WriteResponse(c, errorsx.ErrBind, nil)
		return
	}
	if err := h.v.ValidateLoginReq(c.Request.Context(), &req); err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	resp, err := h.biz.UserV1().Login(c.Request.Context(), &req)
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	core.WriteResponse(c, nil, resp)
}

// RefreshToken 刷新 Token
func (h *Handler) RefreshToken(c *gin.Context) {
	slog.Info("RefreshToken function called.")
	var req apiV1.RefreshTokenReq
	if err := c.ShouldBindJSON(&req); err != nil {
		core.WriteResponse(c, errorsx.ErrBind, nil)
		return
	}
	if err := h.v.ValidateRefreshTokenReq(c.Request.Context(), &req); err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	resp, err := h.biz.UserV1().RefreshToken(c.Request.Context(), &req)
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	core.WriteResponse(c, nil, resp)
}

// ChangePassword 变更密码
func (h *Handler) ChangePassword(c *gin.Context) {
	slog.Info("ChangePassword function called.")
	var req apiV1.ChangePasswordReq
	if err := c.ShouldBindJSON(&req); err != nil {
		core.WriteResponse(c, errorsx.ErrBind, nil)
		return
	}
	if err := h.v.ValidateChangePasswordReq(c.Request.Context(), &req); err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	resp, err := h.biz.UserV1().ChangePassword(c.Request.Context(), &req)
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	core.WriteResponse(c, nil, resp)
}

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
	resp, err := h.biz.UserV1().Create(c.Request.Context(), &req)
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	core.WriteResponse(c, nil, resp)
}

// UpdateUser 更新用户
func (h *Handler) UpdateUser(c *gin.Context) {
	slog.Info("Update user function called.")
	var req apiV1.UpdateUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		core.WriteResponse(c, errorsx.ErrBind, nil)
		return
	}
	if err := h.v.ValidateUpdateUserReq(c.Request.Context(), &req); err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	resp, err := h.biz.UserV1().Update(c.Request.Context(), &req)
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	core.WriteResponse(c, nil, resp)
}

// DeleteUser 删除用户
// TODO 传入的 userID 如何使用
func (h *Handler) DeleteUser(c *gin.Context) {
	slog.Info("Delete user function called.")
	var req apiV1.DeleteUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		core.WriteResponse(c, errorsx.ErrBind, nil)
		return
	}
	if err := h.v.ValidateDeleteUserReq(c.Request.Context(), &req); err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	resp, err := h.biz.UserV1().Delete(c.Request.Context(), &req)
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	core.WriteResponse(c, nil, resp)
}

// GetUser 获取用户
func (h *Handler) GetUser(c *gin.Context) {
	slog.Info("Get user function called.")
	var req apiV1.GetUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		core.WriteResponse(c, errorsx.ErrBind, nil)
		return
	}
	if err := h.v.ValidateGetUserReq(c.Request.Context(), &req); err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	resp, err := h.biz.UserV1().Get(c.Request.Context(), &req)
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	core.WriteResponse(c, nil, resp)
}

// ListUser 获取用户列表
func (h *Handler) ListUser(c *gin.Context) {
	slog.Info("Create user function called.")
	var req apiV1.ListUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		core.WriteResponse(c, errorsx.ErrBind, nil)
		return
	}
	if err := h.v.ValidateListUserReq(c.Request.Context(), &req); err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	resp, err := h.biz.UserV1().List(c.Request.Context(), &req)
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	core.WriteResponse(c, nil, resp)
}

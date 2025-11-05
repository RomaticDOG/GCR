package handler

import (
	"github.com/RomaticDOG/GCR/FastGO/internal/biz"
	"github.com/RomaticDOG/GCR/FastGO/internal/pkg/validation"
)

// Handler 处理博客模块的请求
type Handler struct {
	biz biz.IBiz
	v   *validation.Validator
}

func NewHandler(biz biz.IBiz, v *validation.Validator) *Handler {
	return &Handler{biz, v}
}

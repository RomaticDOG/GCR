package errorsx

import (
	"errors"
	"fmt"
)

// ErrorX 定义了本项目中使用的错误类型，用于描述错误的具体信息
type ErrorX struct {
	// Code 表示 HTTP 状态码，用于与客户端交互时标识错误的类型
	Code int `json:"code,omitempty"`
	// Reason 表示错误发生的原因，通常为业务错误码，用于精准反应错误
	Reason string `json:"reason,omitempty"`
	// Message 表示简短的错误信息，通常可以直接暴露给用户查看
	Message string `json:"message,omitempty"`
}

// New 创建一个新的错误
func New(code int, reason, format string, args ...any) *ErrorX {
	return &ErrorX{
		Code:    code,
		Reason:  reason,
		Message: fmt.Sprintf(format, args...),
	}
}

func (e *ErrorX) Error() string {
	return fmt.Sprintf("error: code [%d], reason [%s], message [%s]", e.Code, e.Reason, e.Message)
}

// WithMessage 设置当前 error 的 message 字段
func (e *ErrorX) WithMessage(format string, args ...any) *ErrorX {
	e.Message = fmt.Sprintf(format, args...)
	return e
}

// FromError 尝试将一个通用的 error 类型转换为自定义的 error 类型
func FromError(e error) *ErrorX {
	if e == nil {
		return nil
	}
	if ex := new(ErrorX); errors.As(e, &ex) {
		return ex
	}
	return New(ErrInternal.Code, ErrInternal.Reason, ErrInternal.Message)
}

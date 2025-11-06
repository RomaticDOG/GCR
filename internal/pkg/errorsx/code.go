package errorsx

import "net/http"

var (
	ErrInternal = &ErrorX{
		Code:    http.StatusInternalServerError,
		Reason:  "InternalError",
		Message: "Internal error",
	}
	ErrNotFound = &ErrorX{
		Code:    http.StatusNotFound,
		Reason:  "NotFound",
		Message: "Resource not found",
	}
	ErrDBRead = &ErrorX{
		Code:    http.StatusInternalServerError,
		Reason:  "InternalError.DBRead",
		Message: "DB read error",
	}
	ErrDBWrite = &ErrorX{
		Code:    http.StatusInternalServerError,
		Reason:  "InternalError.DBWrite",
		Message: "DB write error",
	}
	// ErrBind 表示请求体绑定错误.
	ErrBind = &ErrorX{
		Code:    http.StatusBadRequest,
		Reason:  "BindError",
		Message: "Error occurred while binding the request body to the struct.",
	}
	// ErrSignToken 表示签发 Token 失败.
	ErrSignToken = &ErrorX{
		Code:    http.StatusInternalServerError,
		Reason:  "InternalServerError.SignToken",
		Message: "Error occurred while signing the token.",
	}
	// ErrInvalidToken 表示 Token 错误.
	ErrInvalidToken = &ErrorX{
		Code:    http.StatusBadRequest,
		Reason:  "BadRequest.InvalidToken",
		Message: "Invalid token",
	}
)

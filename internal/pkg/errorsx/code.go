package errorsx

import "net/http"

var (
	OK          = &ErrorX{Code: http.StatusOK, Message: ""}
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
)

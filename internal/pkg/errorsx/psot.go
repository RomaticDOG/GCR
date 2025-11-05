package errorsx

import "net/http"

var (
	ErrTitleEmpty = &ErrorX{
		Code:    http.StatusBadRequest,
		Reason:  "BadRequest.Post.TitleEmpty",
		Message: "Title cannot be empty.",
	}
	ErrContentEmpty = &ErrorX{
		Code:    http.StatusBadRequest,
		Reason:  "BadRequest.Post.ContentEmpty",
		Message: "Content cannot be empty.",
	}
)

package errorsx

import "net/http"

var (
	ErrUsernameEmpty = &ErrorX{
		Code:    http.StatusBadRequest,
		Reason:  "BadRequest.UsernameEmpty",
		Message: "Username cannot be empty.",
	}
	ErrInvalidUsernameLength = &ErrorX{
		Code:    http.StatusBadRequest,
		Reason:  "BadRequest.InvalidUsernameLength",
		Message: "Username length must be between 4 and 32 characters.",
	}
	ErrPasswordEmpty = &ErrorX{
		Code:    http.StatusBadRequest,
		Reason:  "BadRequest.PasswordEmpty",
		Message: "Password cannot be empty.",
	}
	ErrInvalidPasswordLength = &ErrorX{
		Code:    http.StatusBadRequest,
		Reason:  "BadRequest.InvalidPasswordLength",
		Message: "Password length must be between 8 and 32 characters.",
	}
	ErrInvalidNicknameLength = &ErrorX{
		Code:    http.StatusBadRequest,
		Reason:  "BadRequest.InvalidNicknameLength",
		Message: "Nickname length must be less than 32 characters.",
	}
	ErrEmailEmpty = &ErrorX{
		Code:    http.StatusBadRequest,
		Reason:  "BadRequest.EmailEmpty",
		Message: "Email cannot be empty.",
	}
	ErrPhoneEmpty = &ErrorX{
		Code:    http.StatusBadRequest,
		Reason:  "BadRequest.PhoneEmpty",
		Message: "Phone cannot be empty.",
	}
)

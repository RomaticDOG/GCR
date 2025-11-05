package validation

import (
	"context"

	"github.com/RomaticDOG/GCR/FastGO/internal/pkg/errorsx"
	apiV1 "github.com/RomaticDOG/GCR/FastGO/pkg/api/v1"
)

func (v *Validator) ValidateCreateUserReq(ctx context.Context, req *apiV1.CreateUserReq) error {
	if req.Username == "" {
		return errorsx.ErrUsernameEmpty
	}
	if len(req.Username) < 4 || len(req.Username) > 32 {
		return errorsx.ErrInvalidUsernameLength
	}
	if req.Password == "" {
		return errorsx.ErrPasswordEmpty
	}
	if len(req.Password) < 8 || len(req.Password) > 32 {
		return errorsx.ErrInvalidPasswordLength
	}
	if req.Nickname != nil && *req.Nickname != "" {
		if len(*req.Nickname) > 32 {
			return errorsx.ErrInvalidNicknameLength
		}
	}
	if req.Email == "" {
		return errorsx.ErrEmailEmpty
	}
	if req.Phone == "" {
		return errorsx.ErrPhoneEmpty
	}
	return nil
}

func (v *Validator) ValidateUpdateUserReq(ctx context.Context, req *apiV1.UpdateUserReq) error {
	return nil
}

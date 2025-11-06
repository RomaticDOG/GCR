package validation

import (
	"context"

	"github.com/RomaticDOG/GCR/FastGO/internal/pkg/errorsx"
	apiV1 "github.com/RomaticDOG/GCR/FastGO/pkg/api/v1"
)

// ValidateLoginReq 登陆请求体校验
func (v *Validator) ValidateLoginReq(ctx context.Context, req *apiV1.LoginReq) error {
	return nil
}

// ValidateRefreshTokenReq Token 刷新请求体校验
func (v *Validator) ValidateRefreshTokenReq(ctx context.Context, req *apiV1.RefreshTokenReq) error {
	return nil
}

// ValidateChangePasswordReq 密码修改请求体校验
func (v *Validator) ValidateChangePasswordReq(ctx context.Context, req *apiV1.ChangePasswordReq) error {
	return nil
}

// ValidateCreateUserReq 创建用户请求体校验
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

// ValidateUpdateUserReq 更新用户请求体校验
func (v *Validator) ValidateUpdateUserReq(ctx context.Context, req *apiV1.UpdateUserReq) error {
	return nil
}

// ValidateDeleteUserReq 删除用户请求体校验
func (v *Validator) ValidateDeleteUserReq(ctx context.Context, req *apiV1.DeleteUserReq) error {
	return nil
}

// ValidateGetUserReq 获取用户请求体校验
func (v *Validator) ValidateGetUserReq(ctx context.Context, req *apiV1.GetUserReq) error {
	return nil
}

// ValidateListUserReq 获取用户列表请求体校验
func (v *Validator) ValidateListUserReq(ctx context.Context, req *apiV1.ListUserReq) error {
	return nil
}

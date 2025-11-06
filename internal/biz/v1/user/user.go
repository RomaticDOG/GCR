package user

import (
	"context"
	"log/slog"
	"sync"

	"github.com/RomaticDOG/GCR/FastGO/internal/model"
	"github.com/RomaticDOG/GCR/FastGO/internal/pkg/contextx"
	"github.com/RomaticDOG/GCR/FastGO/internal/pkg/conversion"
	"github.com/RomaticDOG/GCR/FastGO/internal/pkg/errorsx"
	"github.com/RomaticDOG/GCR/FastGO/internal/pkg/known"
	"github.com/RomaticDOG/GCR/FastGO/internal/store"
	apiV1 "github.com/RomaticDOG/GCR/FastGO/pkg/api/v1"
	"github.com/RomaticDOG/GCR/FastGO/pkg/token"
	"github.com/jinzhu/copier"
	"github.com/onexstack/onexstack/pkg/authn"
	"github.com/onexstack/onexstack/pkg/store/where"
	"golang.org/x/sync/errgroup"
)

// UserBiz 定义处理用户请求的业务层接口
type UserBiz interface {
	Create(ctx context.Context, req *apiV1.CreateUserReq) (*apiV1.CreateUserResp, error)
	Update(ctx context.Context, req *apiV1.UpdateUserReq) (*apiV1.UpdateUserResp, error)
	Delete(ctx context.Context, req *apiV1.DeleteUserReq) (*apiV1.DeleteUserResp, error)
	Get(ctx context.Context, req *apiV1.GetUserReq) (*apiV1.GetUserResp, error)
	List(ctx context.Context, req *apiV1.ListUserReq) (*apiV1.ListUserResp, error)

	Login(ctx context.Context, req *apiV1.LoginReq) (*apiV1.LoginResp, error)
	RefreshToken(ctx context.Context, req *apiV1.RefreshTokenReq) (*apiV1.RefreshTokenResp, error)
	ChangePassword(ctx context.Context, req *apiV1.ChangePasswordReq) (*apiV1.ChangePasswordResp, error)

	UserExpansion
}

// UserExpansion 用户业务层扩展接口
type UserExpansion interface {
}

// userBiz 用户业务层接口的具体实现
type userBiz struct {
	store store.IStore
}

var _ UserBiz = (*userBiz)(nil)

// New 创建 userBiz 实例
func New(store store.IStore) UserBiz {
	return &userBiz{store: store}
}

// Login 用户登陆
func (ub *userBiz) Login(ctx context.Context, req *apiV1.LoginReq) (*apiV1.LoginResp, error) {
	whr := where.F("username", req.Username)
	user, err := ub.store.User().Get(ctx, whr)
	if err != nil {
		return nil, err
	}
	if err = authn.Compare(user.Password, req.Password); err != nil {
		slog.ErrorContext(ctx, "Failed to compare password", "err", err)
		return nil, errorsx.ErrInvalidPassword
	}
	// 匹配成功，签发 token
	tokenStr, expireAt, err := token.Sign(user.UserID)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to sign token", "err", err)
		return nil, errorsx.ErrSignToken.WithMessage(err.Error())
	}
	return &apiV1.LoginResp{
		Token:    tokenStr,
		ExpireAt: expireAt,
	}, nil
}

// RefreshToken 用于刷新用户的身份认证令牌
func (ub *userBiz) RefreshToken(ctx context.Context, req *apiV1.RefreshTokenReq) (*apiV1.RefreshTokenResp, error) {
	tokenStr, expireAt, err := token.Sign(contextx.UserID(ctx))
	if err != nil {
		slog.ErrorContext(ctx, "Failed to sign token", "err", err)
		return nil, errorsx.ErrSignToken.WithMessage(err.Error())
	}
	return &apiV1.RefreshTokenResp{
		Token:    tokenStr,
		ExpireAt: expireAt,
	}, nil
}

// ChangePassword 修改密码
func (ub *userBiz) ChangePassword(ctx context.Context, req *apiV1.ChangePasswordReq) (*apiV1.ChangePasswordResp, error) {
	user, err := ub.store.User().Get(ctx, where.T(ctx))
	if err != nil {
		return nil, err
	}
	if err = authn.Compare(user.Password, req.OldPassword); err != nil {
		slog.ErrorContext(ctx, "Failed to compare old password", "err", err)
		return nil, errorsx.ErrInvalidPassword
	}
	user.Password, _ = authn.Encrypt(req.NewPassword)
	if err = ub.store.User().Update(ctx, user); err != nil {
		return nil, err
	}
	return &apiV1.ChangePasswordResp{}, nil
}

// Create 创建用户请求
func (ub *userBiz) Create(ctx context.Context, req *apiV1.CreateUserReq) (*apiV1.CreateUserResp, error) {
	var user model.User
	_ = copier.Copy(&user, req)
	if err := ub.store.User().Create(ctx, &user); err != nil {
		return nil, err
	}
	return &apiV1.CreateUserResp{
		UserID: user.UserID,
	}, nil
}

// Update 更新用户请求
func (ub *userBiz) Update(ctx context.Context, req *apiV1.UpdateUserReq) (*apiV1.UpdateUserResp, error) {
	user, err := ub.store.User().Get(ctx, where.T(ctx))
	if err != nil {
		return nil, err
	}
	if req.Phone != nil {
		user.Phone = *req.Phone
	}
	if req.Nickname != nil {
		user.Nickname = *req.Nickname
	}
	if req.Email != nil {
		user.Email = *req.Email
	}
	if req.Username != nil {
		user.Username = *req.Username
	}
	if err := ub.store.User().Update(ctx, user); err != nil {
		return nil, err
	}
	return &apiV1.UpdateUserResp{}, nil
}

// Delete 删除用户请求
func (ub *userBiz) Delete(ctx context.Context, req *apiV1.DeleteUserReq) (*apiV1.DeleteUserResp, error) {
	if err := ub.store.User().Delete(ctx, where.F("userID", contextx.UserID(ctx))); err != nil {
		return nil, err
	}
	return &apiV1.DeleteUserResp{}, nil
}

// Get 获取用户请求
func (ub *userBiz) Get(ctx context.Context, req *apiV1.GetUserReq) (*apiV1.GetUserResp, error) {
	user, err := ub.store.User().Get(ctx, where.F("userID", contextx.UserID(ctx)))
	if err != nil {
		return nil, err
	}
	return &apiV1.GetUserResp{
		User: conversion.UserModelToUserV1(user),
	}, nil
}

// List 获取用户列表请求
func (ub *userBiz) List(ctx context.Context, req *apiV1.ListUserReq) (*apiV1.ListUserResp, error) {
	whr := where.P(int(req.Offset), int(req.Limit))
	cnt, users, err := ub.store.User().List(ctx, whr)
	if err != nil {
		return nil, err
	}
	var m sync.Map
	eg, ctx := errgroup.WithContext(ctx)
	eg.SetLimit(known.MaxErrGroupConcurrency)
	for _, user := range users {
		eg.Go(func() error {
			select {
			case <-ctx.Done():
				return nil
			default:
				pCnt, _, err := ub.store.User().List(ctx, where.T(ctx))
				if err != nil {
					return err
				}
				converted := conversion.UserModelToUserV1(user)
				converted.PostCount = pCnt
				m.Store(user.UserID, converted)
				return nil
			}
		})
	}
	if err := eg.Wait(); err != nil {
		slog.ErrorContext(ctx, "Failed to wait all function calls returned", "err", err)
		return nil, err
	}
	ret := make([]*apiV1.User, 0, len(users))
	for _, user := range users {
		each, _ := m.Load(user.UserID)
		ret = append(ret, each.(*apiV1.User))
	}
	slog.DebugContext(ctx, "Get users from backend storage", "count", len(ret))
	return &apiV1.ListUserResp{
		TotalCount: cnt,
		Users:      ret,
	}, nil
}

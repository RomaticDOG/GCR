package v1

import (
	"time"
)

// User 表示用户信息
type User struct {
	// UserID 表示用户 ID
	UserID string `json:"userID"`
	// Username 表示用户名称
	Username string `json:"username"`
	// Nickname 表示用户昵称
	Nickname string `json:"nickname"`
	// Email 表示用户电子邮箱
	Email string `json:"email"`
	// Phone 表示用户手机号
	Phone string `json:"phone"`
	// PostCount 表示用户拥有的博客数量
	PostCount int64 `json:"postCount"`
	// CreatedAt 表示用户注册时间
	CreatedAt time.Time `json:"createdAt"`
	// UpdatedAt 表示用户最后更新时间
	UpdatedAt time.Time `json:"updatedAt"`
}

// CreateUserReq 表示创建用户请求
type CreateUserReq struct {
	// Username 表示用户名称
	Username string `json:"username"`
	// Password 表示用户密码
	Password string `json:"password"`
	// Nickname 表示用户昵称
	Nickname *string `json:"nickname"`
	// Email 表示用户电子邮箱
	Email string `json:"email"`
	// Phone 表示用户手机号
	Phone string `json:"phone"`
}

// LoginReq 表示登录请求
type LoginReq struct {
	// Username 表示用户名称
	Username string `json:"username"`
	// Password 表示用户密码
	Password string `json:"password"`
}

// LoginResp 表示登录响应
type LoginResp struct {
	// Token 表示返回的身份验证令牌
	Token string `json:"token"`
	// ExpireAt 表示该 token 的过期时间
	ExpireAt time.Time `json:"expireAt"`
}

// RefreshTokenRequest 表示刷新令牌的请求
type RefreshTokenReq struct {
}

// RefreshTokenResp 表示刷新令牌的响应
type RefreshTokenResp struct {
	// Token 表示返回的身份验证令牌
	Token string `json:"token"`
	// ExpireAt 表示该 token 的过期时间
	ExpireAt time.Time `json:"expireAt"`
}

// ChangePasswordReq 表示修改密码请求
type ChangePasswordReq struct {
	// OldPassword 表示当前密码
	OldPassword string `json:"oldPassword"`
	// NewPassword 表示准备修改的新密码
	NewPassword string `json:"newPassword"`
}

// ChangePasswordResp 表示修改密码响应
type ChangePasswordResp struct {
}

// CreateUserResp 表示创建用户响应
type CreateUserResp struct {
	// UserID 表示新创建的用户 ID
	UserID string `json:"userID"`
}

// UpdateUserReq 表示更新用户请求
type UpdateUserReq struct {
	// Username 表示可选的用户名称
	Username *string `json:"username"`
	// Nickname 表示可选的用户昵称
	Nickname *string `json:"nickname"`
	// Email 表示可选的用户电子邮箱
	Email *string `json:"email"`
	// Phone 表示可选的用户手机号
	Phone *string `json:"phone"`
}

// UpdateUserResp 表示更新用户响应
type UpdateUserResp struct {
}

// DeleteUserReq 表示删除用户请求
type DeleteUserReq struct {
}

// DeleteUserResp 表示删除用户响应
type DeleteUserResp struct {
}

// GetUserReq 表示获取用户请求
type GetUserReq struct {
}

// GetUserResp 表示获取用户响应
type GetUserResp struct {
	// User 表示返回的用户信息
	User *User `json:"user"`
}

// ListUserReq 表示用户列表请求
type ListUserReq struct {
	// Offset 表示偏移量
	Offset int64 `json:"offset"`
	// Limit 表示每页数量
	Limit int64 `json:"limit"`
}

// ListUserResp 表示用户列表响应
type ListUserResp struct {
	// TotalCount 表示总用户数
	TotalCount int64 `json:"totalCount"`
	// Users 表示用户列表
	Users []*User `json:"users"`
}

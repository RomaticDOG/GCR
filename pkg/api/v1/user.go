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

package v1

import (
	"time"
)

// Post 表示博客文章
type Post struct {
	// PostID 表示博文 ID
	PostID string `json:"postID"`
	// UserID 表示用户 ID
	UserID string `json:"userID"`
	// Title 表示博客标题
	Title string `json:"title"`
	// Content 表示博客内容
	Content string `json:"content"`
	// CreatedAt 表示博客创建时间
	CreatedAt time.Time `json:"createdAt"`
	// UpdatedAt 表示博客最后更新时间
	UpdatedAt time.Time `json:"updatedAt"`
}

// CreatePostReq 表示创建文章请求
type CreatePostReq struct {
	// Title 表示博客标题
	Title string `json:"title"`
	// Content 表示博客内容
	Content string `json:"content"`
}

// CreatePostResp 表示创建文章响应
type CreatePostResp struct {
	// PostID 表示创建的文章 ID
	PostID string `json:"postID"`
}

// UpdatePostReq 表示更新文章请求
type UpdatePostReq struct {
	// PostID 表示要更新的文章 ID，对应 {postID}
	PostID string `json:"postID" uri:"postID"`
	// Title 表示更新后的博客标题
	Title *string `json:"title"`
	// Content 表示更新后的博客内容
	Content *string `json:"content"`
}

// UpdatePostResp 表示更新文章响应
type UpdatePostResp struct {
}

// DeletePostReq 表示删除文章请求
type DeletePostReq struct {
	// PostIDs 表示要删除的文章 ID 列表
	PostIDs []string `json:"postIDs"`
}

// DeletePostResp 表示删除文章响应
type DeletePostResp struct {
}

// GetPostReq 表示获取文章请求
type GetPostReq struct {
	// PostID 表示要获取的文章 ID
	PostID string `json:"postID" uri:"postID"`
}

// GetPostResp 表示获取文章响应
type GetPostResp struct {
	// Post 表示返回的文章信息
	Post *Post `json:"post"`
}

// ListPostReq 表示获取文章列表请求
type ListPostReq struct {
	// Offset 表示偏移量
	Offset int64 `json:"offset"`
	// Limit 表示每页数量
	Limit int64 `json:"limit"`
	// Title 表示可选的标题过滤
	Title *string `json:"title"`
}

// ListPostResp 表示获取文章列表响应
type ListPostResp struct {
	// TotalCount 表示总文章数
	TotalCount int64 `json:"totalCount"`
	// Posts 表示文章列表
	Posts []*Post `json:"posts"`
}

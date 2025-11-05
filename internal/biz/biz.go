package biz

import (
	postV1 "github.com/RomaticDOG/GCR/FastGO/internal/biz/v1/post"
	userV1 "github.com/RomaticDOG/GCR/FastGO/internal/biz/v1/user"
	"github.com/RomaticDOG/GCR/FastGO/internal/store"
)

// IBiz 定义了业务层需要实现的方法
type IBiz interface {
	// UserV1 获取用户业务接口
	UserV1() userV1.UserBiz
	// PostV1 获取博文业务接口
	PostV1() postV1.PostBiz
}

// biz 业务层接口的具体实现
type biz struct {
	store store.IStore
}

// 确保 biz 实现了 IBiz 中的所有方法
var _ IBiz = (*biz)(nil)

// NewBiz 创建一个 biz 的实例
func NewBiz(store store.IStore) IBiz {
	return &biz{store: store}
}

// UserV1 获取用户业务层接口实例
func (b *biz) UserV1() userV1.UserBiz {
	return userV1.New(b.store)
}

// PostV1 获取用户业务层接口实例
func (b *biz) PostV1() postV1.PostBiz {
	return postV1.New(b.store)
}

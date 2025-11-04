package store

import (
	"context"
	"errors"
	"log/slog"

	"github.com/RomaticDOG/GCR/FastGO/internal/model"
	"github.com/RomaticDOG/GCR/FastGO/internal/pkg/errorsx"
	"github.com/onexstack/onexstack/pkg/store/where"
	"gorm.io/gorm"
)

// PostStore 定义了 Post 在 Store 层需要实现的主要方法
type PostStore interface {
	Create(ctx context.Context, obj *model.Post) error
	Update(ctx context.Context, obj *model.Post) error
	Delete(ctx context.Context, opts *where.Options) error
	Get(ctx context.Context, opts *where.Options) (*model.Post, error)
	List(ctx context.Context, opts *where.Options) (int64, []*model.Post, error)
	PostExpansion
}

// PostExpansion 定义了 Post 在 Store 层需要实现的额外方法
type PostExpansion interface {
}

// postStore 接口 PostStore 的具体实现
type postStore struct {
	store *dataStore
}

// 确保 postStore 实现了 PostStore 接口的所有方法
var _ PostStore = (*postStore)(nil)

func newPostStore(ds *dataStore) *postStore {
	return &postStore{ds}
}

// Create 创建一条博文记录
func (ps *postStore) Create(ctx context.Context, obj *model.Post) error {
	if err := ps.store.DB(ctx).Create(obj).Error; err != nil {
		slog.Error("Failed to create post", "error", err, "obj", obj)
		return errorsx.ErrDBWrite.WithMessage(err.Error())
	}
	return nil
}

// Update 更新博文记录
func (ps *postStore) Update(ctx context.Context, obj *model.Post) error {
	if err := ps.store.DB(ctx).Save(obj).Error; err != nil {
		slog.Error("Failed to update post", "error", err, "obj", obj)
		return errorsx.ErrDBWrite.WithMessage(err.Error())
	}
	return nil
}

// Delete 根据条件删除博文记录
func (ps *postStore) Delete(ctx context.Context, opts *where.Options) error {
	if err := ps.store.DB(ctx, opts).Delete(&model.Post{}).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		slog.Error("Failed to delete post", "error", err, "obj", opts)
		return errorsx.ErrDBWrite.WithMessage(err.Error())
	}
	return nil
}

// Get 根据条件获取博文记录
func (ps *postStore) Get(ctx context.Context, opts *where.Options) (*model.Post, error) {
	var post model.Post
	if err := ps.store.DB(ctx, opts).First(&post).Error; err != nil {
		slog.Error("Failed to get post", "error", err, "obj", opts)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorsx.ErrNotFound
		}
		return nil, errorsx.ErrDBRead.WithMessage(err.Error())
	}
	return &post, nil
}

// List 根据条件获取博文列表
func (ps *postStore) List(ctx context.Context, opts *where.Options) (cnt int64, ret []*model.Post, err error) {
	err = ps.store.DB(ctx, opts).Find(&ret).Offset(-1).Limit(-1).Count(&cnt).Error
	if err != nil {
		slog.Error("Failed to get posts", "error", err)
		err = errorsx.ErrDBRead.WithMessage(err.Error())
	}
	return
}

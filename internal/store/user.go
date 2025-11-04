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

// UserStore 定义了 User 模块在 Store 层需要实现的主要方法
type UserStore interface {
	Create(ctx context.Context, obj *model.User) error
	Update(ctx context.Context, obj *model.User) error
	Delete(ctx context.Context, opts *where.Options) error
	Get(ctx context.Context, opts *where.Options) (*model.User, error)
	List(ctx context.Context, opts *where.Options) (int64, []*model.User, error)

	UserExpansion
}

// UserExpansion 定义了 User 模块在 Store 层需要实现的其他方法
type UserExpansion interface {
}

// userStore 是 UserStore 的具体实现
type userStore struct {
	store *dataStore
}

// 确保实现 UserStore 的所有方法
var _ UserStore = (*userStore)(nil)

// newUserStore 创建 userStore 实例
func newUserStore(ds *dataStore) *userStore {
	return &userStore{ds}
}

// Create 插入一条用户记录
func (us *userStore) Create(ctx context.Context, obj *model.User) error {
	if err := us.store.DB(ctx).Create(obj).Error; err != nil {
		slog.Error("Failed to create user:", "error", err, "user", obj)
		return errorsx.ErrDBWrite.WithMessage(err.Error())
	}
	return nil
}

// Update 更新一条用户记录
func (us *userStore) Update(ctx context.Context, obj *model.User) error {
	if err := us.store.DB(ctx).Save(obj).Error; err != nil {
		slog.Error("Failed to update user:", "error", err, "user", obj)
		return errorsx.ErrDBWrite.WithMessage(err.Error())
	}
	return nil
}

// Delete 根据条件删除用户记录
func (us *userStore) Delete(ctx context.Context, opts *where.Options) error {
	err := us.store.DB(ctx, opts).Delete(&model.User{}).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		slog.Error("Failed to delete user:", "error", err, "conditions", opts)
		return errorsx.ErrDBWrite.WithMessage(err.Error())
	}
	return nil
}

// Get 根据条件获取用户记录
func (us *userStore) Get(ctx context.Context, opts *where.Options) (*model.User, error) {
	var user model.User
	if err := us.store.DB(ctx, opts).First(&user).Error; err != nil {
		slog.Error("Failed to get user:", "error", err, "conditions", opts)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorsx.ErrNotFound
		}
		return nil, errorsx.ErrDBRead.WithMessage(err.Error())
	}
	return &user, nil
}

// List 返回用户列表和总数
func (us *userStore) List(ctx context.Context, opts *where.Options) (cnt int64, ret []*model.User, err error) {
	err = us.store.DB(ctx, opts).Find(&ret).Offset(-1).Limit(-1).Count(&cnt).Error
	if err != nil {
		slog.Error("Failed to list user:", "error", err, "conditions", opts)
		err = errorsx.ErrDBRead.WithMessage(err.Error())
	}
	return
}

package store

import (
	"context"
	"sync"

	"github.com/onexstack/onexstack/pkg/store/where"
	"gorm.io/gorm"
)

var (
	once sync.Once
	S    *dataStore // 全局变量，方便其他包直接调用初始化好的 dataStore 实例
)

// IStore 层定义了 store 层需要实现的方法
type IStore interface {
	DB(ctx context.Context, wheres ...where.Where) *gorm.DB
	TX(ctx context.Context, fn func(tx *gorm.DB) error) error

	User() UserStore
	Post() PostStore
}

// transactionKey 用于在 context.Context 上下文中存储事务的键
type transactionKey struct{}

// dataStore 是 IStore 的具体实现
type dataStore struct {
	postgres *gorm.DB

	// TODO 可以根据需要添加其他数据库实例
}

var _ IStore = (*dataStore)(nil)

// NewStore 创建一个 IStore 类型的实例
func NewStore(db *gorm.DB) IStore {
	once.Do(func() {
		S = &dataStore{db}
	})
	return S
}

// DB 尝试获取事务实例，并根据传入的条件对上下文中的 db 进行筛选，若无条件，则返回上下文中的核心数据库实例
func (store *dataStore) DB(ctx context.Context, wheres ...where.Where) *gorm.DB {
	db := store.postgres
	// 从上下文中提取事务实例
	if tx, ok := ctx.Value(transactionKey{}).(*gorm.DB); ok {
		db = tx
	}
	// 遍历所有传入的条件并叠加到数据库查询实例上
	for _, w := range wheres {
		db = w.Where(db)
	}
	return db
}

// TX 返回一个新的事务实例
func (store *dataStore) TX(ctx context.Context, fn func(tx *gorm.DB) error) error {
	return store.postgres.WithContext(ctx).Transaction(
		func(tx *gorm.DB) error {
			ctx = context.WithValue(ctx, transactionKey{}, tx)
			return fn(tx)
		})
}

// User 返回一个实现了 UserStore 接口的实例
func (store *dataStore) User() UserStore {
	return newUserStore(store)
}

// Post 返回一个实现了 PostStore 接口的实例
func (store *dataStore) Post() PostStore {
	return newPostStore(store)
}

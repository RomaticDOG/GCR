package validation

import "github.com/RomaticDOG/GCR/FastGO/internal/store"

// Validator 验证逻辑的结构体
type Validator struct {
	// 有些复杂的验证逻辑，就需要查询结构体
	store store.IStore
}

// NewValidator 返回一个新的 Validator 实例
func NewValidator(store store.IStore) *Validator {
	return &Validator{store: store}
}

package rid

import "github.com/onexstack/onexstack/pkg/id"

const defaultABC = "abcdefghijklmnopqrstuvwxyz1234567890"

type ResourceID string

const (
	UserID ResourceID = "user" // 用户资源标识符
	PostID ResourceID = "post" // 博文资源标识符
)

// String 将资源标识符转换为字符串
func (rid ResourceID) String() string {
	return string(rid)
}

// New 创建带前缀的唯一标识符
func (rid ResourceID) New(counter uint64) string {
	uniqueStr := id.NewCode(
		counter,
		id.WithCodeChars([]rune(defaultABC)),
		id.WithCodeL(6),
		id.WithCodeSalt(Salt()))
	return rid.String() + "-" + uniqueStr
}

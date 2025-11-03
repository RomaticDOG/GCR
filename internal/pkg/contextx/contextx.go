package contextx

import "context"

type (
	requestIDKey struct{} // 上下文中存储 ID 的键
)

// WithRequestID 将请求 ID 存放到上下文中
func WithRequestID(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, requestIDKey{}, requestID)
}

// RequestID 从上下文中读取 ID
func RequestID(ctx context.Context) string {
	requestID, _ := ctx.Value(requestIDKey{}).(string)
	return requestID
}

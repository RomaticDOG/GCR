package options

import "time"

// SystemOptions 系统配置
type SystemOptions struct {
	DBMode string      `json:"dbMode" mapstructure:"db-mode"`
	JWT    *JWTOptions `json:"jwt" mapstructure:"jwt"`
}

// JWTOptions 认证配置项
type JWTOptions struct {
	Key        string        `json:"key" mapstructure:"key"`
	Expiration time.Duration `json:"expiration" mapstructure:"expiration"`
}

// NewSystem 返回一个默认值
func NewSystem() *SystemOptions {
	return &SystemOptions{
		DBMode: "mysql",
	}
}

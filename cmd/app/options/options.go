package options

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type ServerOptions struct {
	MySQL *MySQLOptions `json:"mysql" mapstructure:"mysql"`
}

// MySQLOptions MySQL 配置项结构体
type MySQLOptions struct {
	Addr             string        `json:"addr,omitempty" mapstructure:"addr" validate:"required,hostname_port"`
	Username         string        `json:"username,omitempty" mapstructure:"username" validate:"required"`
	Password         string        `json:"-" mapstructure:"password"`
	Database         string        `json:"database,omitempty" mapstructure:"database" validate:"required"`
	MaxIdleConns     int           `json:"max-idle-conns,omitempty" mapstructure:"max-idle-conns" validate:"gte=0"`
	MaxOpenConns     int           `json:"max-open-conns,omitempty" mapstructure:"max-open-conns" validate:"gte=0"`
	MaxConnsLifeTime time.Duration `json:"max-conns-life-time,omitempty" mapstructure:"max-conns-life-time" validate:"gte=0"`
}

// NewServerOptions 创建带有默认值的 ServerOptions 实例
func NewServerOptions() *ServerOptions {
	return &ServerOptions{MySQL: NewMySQLOptions()}
}

// NewMySQLOptions 返回一个零值 MySQL 配置项
func NewMySQLOptions() *MySQLOptions {
	return &MySQLOptions{
		Addr:             "127.0.0.1:3306",
		Username:         "RomanticDOG",
		Password:         "RomanticDOG",
		Database:         "fast-go",
		MaxIdleConns:     100,
		MaxOpenConns:     100,
		MaxConnsLifeTime: time.Duration(10) * time.Second,
	}
}

// Validate 验证 ServerOptions 中的各个配置项是否规范
func (so *ServerOptions) Validate() error {
	return so.MySQL.mySQLValidate()
}

// mySQLValidate 验证 ServerOptions 中的 MySQL 配置项是否规范
func (mo *MySQLOptions) mySQLValidate() error {
	v := validator.New()
	return v.Struct(mo)
}

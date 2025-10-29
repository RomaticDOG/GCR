package options

import (
	"github.com/RomaticDOG/GCR/FastGO/internal"
	genericOptions "github.com/RomaticDOG/GCR/FastGO/pkg/options"
)

type ServerOptions struct {
	MySQL *genericOptions.MySQLOptions `json:"mysql" mapstructure:"mysql"`
}

// NewServerOptions 创建带有默认值的 ServerOptions 实例
func NewServerOptions() *ServerOptions {
	return &ServerOptions{MySQL: genericOptions.NewMySQLOptions()}
}

// Validate 验证 ServerOptions 中的各个配置项是否规范
func (so *ServerOptions) Validate() error {
	return so.MySQL.Validate()
}

func (so *ServerOptions) Config() (*internal.Config, error) {
	return &internal.Config{
		MySqlOptions: so.MySQL,
	}, nil
}

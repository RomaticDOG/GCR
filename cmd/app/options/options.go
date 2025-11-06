package options

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/RomaticDOG/GCR/FastGO/internal"
	genericOptions "github.com/RomaticDOG/GCR/FastGO/pkg/options"
)

type ServerOptions struct {
	System   *genericOptions.SystemOptions   `json:"system" mapstructure:"system"`
	MySQL    *genericOptions.MySQLOptions    `json:"mysql" mapstructure:"mysql"`
	Postgres *genericOptions.PostgresOptions `json:"postgres" mapstructure:"postgres"`
	Addr     string                          `json:"addr" mapstructure:"addr"`
}

// NewServerOptions 创建带有默认值的 ServerOptions 实例
func NewServerOptions() *ServerOptions {
	return &ServerOptions{
		MySQL:    genericOptions.NewMySQLOptions(),
		Postgres: genericOptions.NewPostgres(),
		Addr:     "0.0.0.0:6666",
		System:   genericOptions.NewSystem(),
	}
}

// Validate 验证 ServerOptions 中的各个配置项是否规范
func (so *ServerOptions) Validate() (err error) {
	if so.System.DBMode == "mysql" {
		if err = so.MySQL.Validate(); err != nil {
			return
		}
	} else {
		if err = so.Postgres.Validate(); err != nil {
			return
		}
	}
	if err = addrValidate(so.Addr); err != nil {
		return
	}
	return nil
}

func (so *ServerOptions) Config() (*internal.Config, error) {
	return &internal.Config{
		MySqlOptions:    so.MySQL,
		PostgresOptions: so.Postgres,
		System:          so.System,
		Addr:            so.Addr,
	}, nil
}

// addrValidate 验证Gin的Addr配置是否合法
func addrValidate(addr string) error {
	// 1. 允许空地址（Gin会默认使用 ":8080"）
	if addr == "" {
		return nil
	}

	// 2. 分割host和port（格式必须包含 ":"）
	parts := strings.Split(addr, ":")
	if len(parts) != 2 {
		return errors.New("地址格式错误，必须包含 ':'（如 ':8080' 或 'localhost:8080'）")
	}

	host := parts[0]
	portStr := parts[1]

	// 3. 验证port是否为有效整数（1-65535）
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return fmt.Errorf("端口必须是整数，当前值: %s", portStr)
	}
	if port < 1 || port > 65535 {
		return fmt.Errorf("端口必须在1-65535之间，当前值: %d", port)
	}

	// 4. 验证host是否合法（可选，允许空host，如 ":8080"）
	// 空host表示监听所有网络接口，合法
	if host == "" {
		return nil
	}

	// 非空host需验证是否为合法IP或域名（通过尝试解析验证）
	// 注意：域名可能包含字母、数字、"-"、"."，且不能以"-"开头/结尾
	if net.ParseIP(host) == nil {
		// 不是IP地址，尝试验证是否为合法域名格式
		if !isValidDomain(host) {
			return fmt.Errorf("host格式不合法（不是有效IP或域名），当前值: %s", host)
		}
	}

	// 所有验证通过
	return nil
}

// isValidDomain 辅助函数：验证域名格式是否合法
func isValidDomain(domain string) bool {
	// 域名长度限制（1-255字符）
	if len(domain) == 0 || len(domain) > 255 {
		return false
	}
	// 不能以"."开头或结尾
	if strings.HasPrefix(domain, ".") || strings.HasSuffix(domain, ".") {
		return false
	}
	// 拆分域名标签（如 "example.com" 拆分为 ["example", "com"]）
	labels := strings.Split(domain, ".")
	for _, label := range labels {
		// 标签长度限制（1-63字符）
		if len(label) == 0 || len(label) > 63 {
			return false
		}
		// 标签只能包含字母、数字、"-"，且不能以"-"开头/结尾
		for i, c := range label {
			if !((c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') || c == '-') {
				return false
			}
			if c == '-' && (i == 0 || i == len(label)-1) {
				return false
			}
		}
	}
	return true
}

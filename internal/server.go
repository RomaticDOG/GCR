package internal

import (
	"log/slog"

	genericOptions "github.com/RomaticDOG/GCR/FastGO/pkg/options"
)

// Config 配置结构体，用于存储应用相关配置
type Config struct {
	MySqlOptions *genericOptions.MySQLOptions
}

// Server 服务器结构体类型
type Server struct {
	cfg *Config
}

// NewServer 根据配置创建服务器
func (cfg *Config) NewServer() (*Server, error) {
	return &Server{cfg: cfg}, nil
}

func (s *Server) Run() error {
	slog.Info("Slog Read Mysql host from config", "mysql.addr", s.cfg.MySqlOptions.Addr)
	return nil
}

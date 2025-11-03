package internal

import (
	"errors"
	"log/slog"
	"net/http"

	genericOptions "github.com/RomaticDOG/GCR/FastGO/pkg/options"
	"github.com/gin-gonic/gin"
)

// Config 配置结构体，用于存储应用相关配置
type Config struct {
	MySqlOptions *genericOptions.MySQLOptions
	Addr         string
}

// Server 服务器结构体类型
type Server struct {
	cfg *Config
	srv *http.Server
}

// NewServer 根据配置创建服务器
func (cfg *Config) NewServer() (*Server, error) {
	engine := gin.New()
	engine.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"code": "PageNotFound", "message": "Page not found."})
	})
	engine.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "OK"})
	})
	httpSrv := &http.Server{Addr: cfg.Addr, Handler: engine}
	return &Server{cfg: cfg, srv: httpSrv}, nil
}

func (s *Server) Run() error {
	// 运行 http 服务器
	slog.Info("Start to listening the incoming requests on http address", "addr", s.cfg.Addr)
	if err := s.srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

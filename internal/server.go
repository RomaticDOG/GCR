package internal

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/RomaticDOG/GCR/FastGO/internal/pkg/core"
	"github.com/RomaticDOG/GCR/FastGO/internal/pkg/errorsx"
	mw "github.com/RomaticDOG/GCR/FastGO/internal/pkg/middleware"
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
	// 添加 gin.Recovery() 中间件，用来捕获任何 panic，并恢复
	mws := []gin.HandlerFunc{gin.Recovery(), mw.NoCache, mw.Cors, mw.RequestID()}
	engine.Use(mws...)
	engine.NoRoute(func(c *gin.Context) {
		core.WriteResponse(c, errorsx.ErrNotFound.WithMessage("Page Not Found."), nil)
	})
	engine.GET("/healthz", func(c *gin.Context) {
		core.WriteResponse(c, nil, map[string]string{"status": "ok"})
	})
	httpSrv := &http.Server{Addr: cfg.Addr, Handler: engine}
	return &Server{cfg: cfg, srv: httpSrv}, nil
}

func (s *Server) Run() error {
	// 运行 http 服务器
	slog.Info("Start to listening the incoming requests on http address", "addr", s.cfg.Addr)
	go func() {
		if err := s.srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error(err.Error())
			os.Exit(1)
		}
	}()
	quit := make(chan os.Signal, 1)
	// 当执行 kill 命令时（不带参数），默认会发送 syscall.SIGTERM 信号
	// 当执行 kill -2 命令时，会发送 syscall.SIGINT 信号（例如 CTRL + C）
	// 当执行 kill -9 命令时，会发送 syscall.SIGKILL 信号，此信号无法被捕获，因此不用管
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	// 阻塞程序，从 quit channel 中接收信号
	<-quit
	slog.Info("receive signal, shutting down server...")
	// 优雅关停
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// 先关闭依赖的服务，再关闭被依赖的服务，限时 10 秒，超时则退出
	if err := s.srv.Shutdown(ctx); err != nil {
		slog.Error("Insecure server shutdown error", "error", err)
		return err
	}
	slog.Info("Server exited")
	return nil
}

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

	"github.com/RomaticDOG/GCR/FastGO/internal/biz"
	"github.com/RomaticDOG/GCR/FastGO/internal/handler"
	"github.com/RomaticDOG/GCR/FastGO/internal/pkg/core"
	"github.com/RomaticDOG/GCR/FastGO/internal/pkg/errorsx"
	mw "github.com/RomaticDOG/GCR/FastGO/internal/pkg/middleware"
	"github.com/RomaticDOG/GCR/FastGO/internal/pkg/validation"
	store "github.com/RomaticDOG/GCR/FastGO/internal/store"
	genericOptions "github.com/RomaticDOG/GCR/FastGO/pkg/options"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Config 配置结构体，用于存储应用相关配置
type Config struct {
	MySqlOptions    *genericOptions.MySQLOptions
	PostgresOptions *genericOptions.PostgresOptions
	System          *genericOptions.System
	Addr            string
}

// Server 服务器结构体类型
type Server struct {
	cfg *Config
	srv *http.Server
}

// initDB 初始化数据库连接
func (cfg *Config) initDB() (db *gorm.DB, err error) {
	if cfg.System.DBMode == "mysql" {
		db, err = cfg.MySqlOptions.NewDB()
		if err != nil {
			return nil, err
		}
		return
	} else if cfg.System.DBMode == "postgres" {
		db, err = cfg.PostgresOptions.NewDB()
		if err != nil {
			return nil, err
		}
		return
	}
	return
}

// NewServer 根据配置创建服务器
func (cfg *Config) NewServer() (*Server, error) {
	db, err := cfg.initDB()
	if err != nil {
		return nil, err
	}
	s := store.NewStore(db)
	engine := gin.New()
	// 添加 gin.Recovery() 中间件，用来捕获任何 panic，并恢复
	mws := []gin.HandlerFunc{gin.Recovery(), mw.NoCache, mw.Cors, mw.RequestID()}
	engine.Use(mws...)
	cfg.installRESTAPI(engine, s)
	httpSrv := &http.Server{Addr: cfg.Addr, Handler: engine}
	return &Server{cfg: cfg, srv: httpSrv}, nil
}

// installRESTAPI 注册 API 路由
func (cfg *Config) installRESTAPI(engine *gin.Engine, s store.IStore) {
	engine.NoRoute(func(c *gin.Context) {
		core.WriteResponse(c, errorsx.ErrNotFound.WithMessage("Page Not Found."), nil)
	})
	engine.GET("/health", func(c *gin.Context) {
		core.WriteResponse(c, nil, map[string]string{"status": "ok"})
	})
	// 创建核心业务处理器
	h := handler.NewHandler(biz.NewBiz(s), validation.NewValidator(s))

	authMWs := []gin.HandlerFunc{}

	// 注册 V1 版本的 api 路由
	v1 := engine.Group("/v1")
	{
		userV1 := v1.Group("/user")
		{
			userV1.POST("", h.CreateUser)          // 创建用户
			userV1.PUT(":userID", h.UpdateUser)    // 更新用户
			userV1.DELETE(":userID", h.DeleteUser) // 删除用户
			userV1.GET(":userID", h.GetUser)       // 获取用户
			userV1.GET("", h.ListUser)             // 获取用户列表
		}

		postV1 := v1.Group("/post", authMWs...)
		{
			postV1.POST("", h.CreatePost)       // 创建博文
			postV1.PUT(":postID", h.UpdatePost) // 更新博文
			postV1.DELETE("", h.DeletePost)     // 删除博文
			postV1.GET(":postID", h.GetPost)    // 获取博文
			postV1.GET("", h.ListPost)          // 获取博文列表
		}
	}
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

package options

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// PostgresOptions postgres 配置
type PostgresOptions struct {
	Addr             string        `json:"addr,omitempty" mapstructure:"addr" validation:"required,hostname|ip"`
	Port             string        `json:"port,omitempty" mapstructure:"port" validation:"required,gte=1,lte=65535"`
	Username         string        `json:"username,omitempty" mapstructure:"username" validation:"required,min=1"`
	Password         string        `json:"-" mapstructure:"password" validation:"omitempty"`
	Database         string        `json:"database,omitempty" mapstructure:"database" validation:"required,min=1"`
	SSLMode          string        `json:"sslmode,omitempty" mapstructure:"ssl-mode" validation:"required,oneof=disable"`
	MaxIdleConns     int           `json:"max-idle-conns,omitempty" mapstructure:"max-idle-conns" validation:"gte=0"`
	MaxOpenConns     int           `json:"max-open-conns,omitempty" mapstructure:"max-open-conns" validation:"gte=0"`
	MaxConnsLifeTime time.Duration `json:"max-conns-life-time,omitempty" mapstructure:"max-conns-life-time" validation:"gte=0"`
}

// Validate 验证 ServerOptions 中的 MySQL 配置项是否规范
func (po *PostgresOptions) Validate() error {
	v := validator.New()
	return v.Struct(po)
}

// NewPostgres 返回一个默认值配置
func NewPostgres() *PostgresOptions {
	return &PostgresOptions{
		Addr:             "127.0.0.1",
		Port:             "5432",
		Username:         "jing",
		Password:         "",
		Database:         "fastgo",
		MaxIdleConns:     100,
		MaxOpenConns:     100,
		MaxConnsLifeTime: time.Duration(10) * time.Second,
	}
}

// dsn 生成连接语句
func (po *PostgresOptions) dsn() string {
	if po.Password == "" {
		return fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s sslmode=%s",
			po.Addr,
			po.Port,
			po.Username,
			po.Database,
			po.SSLMode,
		)
	}
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		po.Addr,
		po.Port,
		po.Username,
		po.Password,
		po.Database,
		po.SSLMode,
	)
}

// NewDB 返回一个数据库链接
func (po *PostgresOptions) NewDB() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(po.dsn()), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		return nil, err
	}
	// 可选：获取底层 sql.DB 实例，设置连接池参数
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	// 设置连接池最大空闲连接数
	sqlDB.SetMaxIdleConns(po.MaxIdleConns)
	// 设置连接池最大打开连接数
	sqlDB.SetMaxOpenConns(po.MaxOpenConns)
	// 设置连接的最大存活时间
	sqlDB.SetConnMaxLifetime(po.MaxConnsLifeTime)
	return db, nil
}

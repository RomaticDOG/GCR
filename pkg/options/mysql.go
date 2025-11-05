package options

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// MySQLOptions MySQL 配置项结构体
type MySQLOptions struct {
	Addr             string        `json:"addr,omitempty" mapstructure:"addr" validation:"required,hostname_port"`
	Username         string        `json:"username,omitempty" mapstructure:"username" validation:"required"`
	Password         string        `json:"-" mapstructure:"password"`
	Database         string        `json:"database,omitempty" mapstructure:"database" validation:"required"`
	MaxIdleConns     int           `json:"max-idle-conns,omitempty" mapstructure:"max-idle-conns" validation:"gte=0"`
	MaxOpenConns     int           `json:"max-open-conns,omitempty" mapstructure:"max-open-conns" validation:"gte=0"`
	MaxConnsLifeTime time.Duration `json:"max-conns-life-time,omitempty" mapstructure:"max-conns-life-time" validation:"gte=0"`
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

// Validate 验证 ServerOptions 中的 MySQL 配置项是否规范
func (mo *MySQLOptions) Validate() error {
	v := validator.New()
	return v.Struct(mo)
}

func (mo *MySQLOptions) dsn() string {
	return fmt.Sprintf(`%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=%t&loc=%s`,
		mo.Username,
		mo.Password,
		mo.Addr,
		mo.Database,
		true,
		"Local")
}

func (mo *MySQLOptions) NewDB() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(mo.dsn()), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxOpenConns(mo.MaxOpenConns)
	sqlDB.SetMaxIdleConns(mo.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(mo.MaxConnsLifeTime)
	return db, nil
}

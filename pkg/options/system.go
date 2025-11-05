package options

// System 系统配置
type System struct {
	DBMode string `json:"dbMode" mapstructure:"db-mode"`
}

// NewSystem 返回一个默认值
func NewSystem() *System {
	return &System{
		DBMode: "mysql",
	}
}

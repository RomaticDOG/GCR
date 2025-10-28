package app

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	defaultConfigFile = "sys.yaml"
	defaultConfigDir  = "config"
	defaultRootDir    = "../../" // 项目所在的默认根目录
)

// onInitialize 设置需要读取的配置文件名、环境变量，并将其内容读取到 viper 中
func onInitialize() {
	if configFileLookUpFlag {
		// 从命令行指定的配置文件目录中读取配置项
		viper.SetConfigFile(configFile)
	} else {
		// 使用默认配置文件搜索目录搜索默认配置文件名称
		for _, dir := range searchDirs() {
			viper.AddConfigPath(dir)
		}
		viper.SetConfigType("yaml")
		viper.SetConfigName(defaultConfigFile)
	}
	// 读取环境变量并配置前缀
	setupEnvironmentVariables()
	err := viper.ReadInConfig()
	if err != nil {
		cobra.CheckErr(err)
		os.Exit(2)
	}
}

// searchDirs 返回默认的配置文件搜索目录
func searchDirs() []string {
	homeDir, err := os.Executable()
	cobra.CheckErr(err)
	return []string{filepath.Join(filepath.Dir(homeDir), defaultRootDir, defaultConfigDir), "."}
}

// setupEnvironmentVariables 配置 viper 的环境变量规则
func setupEnvironmentVariables() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("FAST_GO")
	// 替换环境变量中的 "." "-" 为 "_"
	replacer := strings.NewReplacer(".", "_", "-", "_")
	viper.SetEnvKeyReplacer(replacer)
}

// configDir 获取默认配置文件的完整路径
func configDir() string {
	home, err := os.Executable()
	cobra.CheckErr(err)
	return filepath.Join(filepath.Dir(home), defaultRootDir, defaultConfigDir, defaultConfigFile)
}

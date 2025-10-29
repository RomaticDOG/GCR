package app

import (
	"github.com/RomaticDOG/GCR/FastGO/cmd/app/options"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	configFileLookUpFlag bool   // 判断 configFile 是否有接受值
	configFile           string // 配置文件路径
)

// NewCommand 创建一个 *cobra.Command 对象，用于启动应用程序
func NewCommand() *cobra.Command {
	opts := options.NewServerOptions()
	cmd := &cobra.Command{
		// 指定命令的名字，该名字会出现在帮助信息中
		Use: "FastGO",
		// 命令的简短描述
		Short: "A very lightweight go project.",
		Long:  "A very lightweight go project, designed to help beginners quickly learn Go project development.",
		// 命令出错时，不打印帮助信息。设置为 true 时可以确保一眼就能看到错误信息
		SilenceUsage: true,
		// 指定调用 cmd.Execute() 时执行的 run 函数
		RunE: func(cmd *cobra.Command, args []string) error {
			err := run(opts)
			cobra.CheckErr(err)
			return err
		},
		// 设置命令行运行时的参数检查，不需要指定命令行参数
		Args: cobra.NoArgs,
	}

	// 初始化配置项，将配置内容读取到 viper 中
	cobra.OnInitialize(onInitialize)

	cmd.PersistentFlags().StringVarP(&configFile, "config", "c", configDir(), "path to the config file.")
	if flag := cmd.Flags().Lookup("config"); flag != nil {
		configFileLookUpFlag = true
	}
	return cmd
}

// run 主要运行逻辑，负责初始化日志、解析配置、校验选项并启动服务器
func run(opts *options.ServerOptions) error {
	// 将读取到的配置项解析到 opts 中
	if err := viper.Unmarshal(&opts); err != nil {
		return err
	}
	if err := opts.Validate(); err != nil {
		return err
	}
	// 获取应用配置，将命令行配置和应用配置分开，更加灵活处理 2 种不同的配置
	cfg, err := opts.Config()
	if err != nil {
		return err
	}

	server, err := cfg.NewServer()
	if err != nil {
		return err
	}

	// 启动服务器
	return server.Run()
}

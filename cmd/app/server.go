package app

import (
	"encoding/json"
	"fmt"

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
			// 将读取到的配置项解析到 opts 中
			if err := viper.Unmarshal(&opts); err != nil {
				cobra.CheckErr(err)
				return err
			}
			if err := opts.Validate(); err != nil {
				cobra.CheckErr(err)
				return err
			}
			// 输出一个 json 看看读取到的配置项是什么样的
			j, _ := json.MarshalIndent(opts, "", "  ")
			fmt.Println(string(j))
			return nil
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

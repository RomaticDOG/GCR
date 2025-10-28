package app

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewCommand 创建一个 *cobra.Command 对象，用于启动应用程序
func NewCommand() *cobra.Command {
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
			fmt.Println("Hello FastGO")
			return nil
		},
		// 设置命令行运行时的参数检查，不需要指定命令行参数
		Args: cobra.NoArgs,
	}
	return cmd
}

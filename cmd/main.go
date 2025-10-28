package main

import (
	"os"

	"github.com/RomaticDOG/GCR/FastGO/cmd/app"
	// 可以在程序启动时自动配置 GOMAXPROCS 的值，使其与 CPU 配额数相同
	// 避免了在容器中因默认 GOMAXPROCS 值不合适导致的性能问题，确保程序能够充分利用可用的资源
	_ "go.uber.org/automaxprocs"
)

func main() {
	cmd := app.NewCommand()
	if err := cmd.Execute(); err != nil {
		// 使用退出码可以让其他程序判断该程序的运行状态
		os.Exit(1)
	}
}

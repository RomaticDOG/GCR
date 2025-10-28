# FastGO

FastGo 是一个快速完整的 Go 项目实例，从今天开始，我们将从零开始构建 FastGo 项目。

# 预知识：GO 项目结构



# 1. 使用 Cobra 包来构建项目

对于一个项目，我们经常需要使用命令行带参数的形式来启动程序，因此我们需要获取这些参数。
传统的方式，我们通常在 `main` 函数中依次读取这些参数，这就会使得你的 `main` 函数非常庞大，不易阅读。

```go
func main() {
    // 解析命令行参数
    opt1 := flag.String("opt1", "default_value", "Description of opt 1")
    opt2 := flag.String("opt2", 0, "Description of opt 2")
    flag.Parse()
    
    fmt.Println("opt1's value is: ", *opt1)
    fmt.Println("opt2's value is: ", *opt2)
    
    // 接下来执行主逻辑。。。
}
```
因此，我们可以使用社区提供的优秀框架来实现这个功能，比如 **spf13/cobra** 就是一个优秀的主函数启动框架。

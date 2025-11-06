# MiniBlog

# 1. 初始工作

## 1.1 热加载应用
我们可以添加 Air 工具来实现 Go 应用的热加载

```shell
    $ go install github.com/air-verse/air@latest  
```

安装完成后，需要对 Air 进行配置，我们可以使用 Air 提供的默认配置文件 [air_example.toml](https://github.com/air-verse/air/blob/master/air_example.toml) 即可，
并根据注释对需要修改的部分进行修改就好，主要的修改内容在 [build] 下

```shell
# Working directory
# . or absolute path, please note that the directories following must be under root.
root = "."
tmp_dir = "tmp/air"

[build]
# Array of commands to run before each build
pre_cmd = ["echo 'hello air' > tmp/air/pre_cmd.txt"]
# Just plain old shell command. You could use `make` as well.
cmd = "go build -o target/bin/mb-api-server -v cmd/mb-api-server/main.go"
# Array of commands to run after ^C
post_cmd = ["echo 'hello air' > tmp/air/post_cmd.txt"]
# Binary file yields from `cmd`.
bin = "target/bin/mb-api-server"
# Customize binary, can setup environment variables when run your app.
# full_bin = "APP_ENV=dev APP_USER=air ./tmp/main"
# Add additional arguments when running binary (bin/full_bin). Will run './tmp/main hello world'.
args_bin = []
```

# 1.2 添加版权声明

一般项目的根目录下会存放一个 LICENSE 文件，用于声明开源项目所遵循的协议：

```shell
    $ go install github.com/nishanths/license/v5@latest
    $ license -list
    $ license -n 'RomanticDOG(浪漫的土狗) <ginwithouta@gmail.com>' -o LICENSE mit
```

# 1.3 给源文件添加版本声明

除了添加开源协议声明，还需要为每个文件添加一个版权头信息，以声明文件所遵循的开源协议。通常版权头信息保存的文件名称会命名为 boilerplate，为了防止漏加，
我们可以使用代码的方式为所有文件添加版权头信息：
```shell
    $ go install github.com/marmotedu/addlicense@latest
    $ addlicense -v -f ./scripts/boilerplate.txt --skip-dirs=third_party,target .
```

这样，在编写完代码后，执行上述命令，就能够在每个文件中添加对应的声明。

# 1.4 构建 Makefile 脚本

在构建项目包的时候，我们经常会执行如下的指令：
```shell
    $ go build -o target/bin/mb-api-server -v cmd/mb-api-server/main.go
```
随着项目的开发，上述的编译命令可能会不断加长。如果每次都手动执行的话，很容易出错，并且效率低下。
因此，最佳实践方式是使用构建工具来管理项目，比如 make 方式。

在构建完 MAKEFILE 文件后，还需要将 .air.toml 中的相关代码修改成 make build 执行：
```shell
[build]
# Array of commands to run before each build
pre_cmd = ["echo 'hello air' > tmp/air/pre_cmd.txt"]
# Just plain old shell command. You could use `make` as well.
cmd = "make build"
```
#!/bin/zsh

# 获取脚本所在目录作为项目根目录
PROJ_ROOT_DIR=$(dirname "${BASH_SOURCE[0]}")

# 定义构建产物的输出目录为项目根目录下的 target/bin 文件夹中
OUTPUT_DIR=${PROJ_ROOT_DIR}/target/bin

# 指定版本信息包的路径，后续会通过 -ldflags 将版本信息注入到这个包的变量中
VERSION_PACKAGE=github.com/RomaticDOG/GCR/pkg/version

# 确定 VERSION 的值，如果环境变量中没有使用 VERSION，则使用 git 标签作为版本号
if [[ -z "${VERSION}" ]]; then
  VERSION=$(git describe --tags --always --match='v*' )
fi

# 检查代码仓库状态，判断工作目录是否干净
# 默认状态设为 dirty 有未提交更改
GIT_TREE_STATE="dirty"
is_clean=$(git status --porcelain 2>/dev/null)
if [[ -z ${is_clean} ]]; then
  GIT_TREE_STATE="clean"
fi

# 获取当前 git commit 的完整哈希值
GIT_COMMIT=$(git rev-parse HEAD)

# 构建链接器标志（ldflags）
# 通过 -X 选项向 VERSION_PACKAGE 包中注入以下变量的值：
# - gitVersion: 版本号
# - gitCommit: 构建时的 commit 哈希
# - gitTreeState: 代码仓库状态（clean或dirty）
# - buildDate: 构建日期和时间（UTC格式）
GO_LDFALGS="-X ${VERSION_PACKAGE}.gitVersion=${VERSION} \
    -X ${VERSION_PACKAGE}.gitCommit=${GIT_COMMIT} \
    -X ${VERSION_PACKAGE}.gitTreeState=${GIT_TREE_STATE} \
    -X ${VERSION_PACKAGE}.buildDate=$(date -u + '%Y-%m-%dT%H:%M:%SZ')"

go build -v -ldflags "${GO_LDFALGS}" -o ${OUTPUT_DIR} -v cmd/main.go

package version

import (
	"encoding/json"
	"fmt"
	"runtime"

	"github.com/gosuri/uitable"
)

var (
	gitVersion   = "v0.0.0-master+$Format:%H$"
	gitCommit    = "$Format:%H$"
	gitTreeState = ""
	buildDate    = "1970-01-01T00:00:00Z"
)

// Info 包含版本信息
type Info struct {
	GitVersion   string `json:"gitVersion"`
	GitCommit    string `json:"gitCommit"`
	GitTreeState string `json:"gitTreeState"`
	BuildDate    string `json:"buildDate"`
	GoVersion    string `json:"goVersion"`
	Compiler     string `json:"compiler"`
	Platform     string `json:"platform"`
}

// String 返回友好的版本信息字符串
func (i Info) String() string {
	return i.GitVersion
}

// ToJson 返回 json 版本信息
func (i Info) ToJson() string {
	b, _ := json.Marshal(i)
	return string(b)
}

// Text 返回经过 UTF-8 编码的文本字符串
func (i Info) Text() string {
	table := uitable.New()
	table.RightAlign(0)
	table.MaxColWidth = 80
	table.Separator = " "
	table.AddRow("GitVersion:", i.GitVersion)
	table.AddRow("GitCommit:", i.GitCommit)
	table.AddRow("GitTreeState:", i.GitTreeState)
	table.AddRow("BuildDate:", i.BuildDate)
	table.AddRow("GoVersion:", i.GoVersion)
	table.AddRow("Compiler:", i.Compiler)
	table.AddRow("Platform:", i.Platform)
	return table.String()
}

// Get 返回详尽的代码库版本信息
func Get() Info {
	return Info{
		GitVersion:   gitVersion,
		GitCommit:    gitCommit,
		GitTreeState: gitTreeState,
		BuildDate:    buildDate,
		GoVersion:    runtime.Version(),
		Compiler:     runtime.Compiler,
		Platform:     fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}

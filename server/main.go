package main

import (
	"github.com/weilaim/wm-admin/server/core"
	"github.com/weilaim/wm-admin/server/global"
	"github.com/weilaim/wm-admin/server/initialize"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download
func main() {
	global.WM_VP = core.Viper() // 初始化Viper
	initialize.OtherInit()
	global.WM_LOG = core.Zap()
	global.WM_LOG.Error("测试失败!")
}

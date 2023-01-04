package main

import (
	"fmt"

	"github.com/weilaim/wm-admin/server/core"
	"github.com/weilaim/wm-admin/server/global"
	"github.com/weilaim/wm-admin/server/initialize"
	"go.uber.org/zap"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download
func main() {

	fmt.Println("我走到这里")
	global.WM_VP = core.Viper() // 初始化Viper
	initialize.OtherInit()
	global.WM_LOG = core.Zap() // 初始化zap日志库
	zap.ReplaceGlobals(global.WM_LOG)
	global.WM_DB = initialize.Gorm() // gorm连接数据库
	initialize.Timer()

	initialize.DBList()
	if global.WM_DB != nil {
		initialize.RegisterTables(global.WM_DB) // 初始化表
		// 程序结束前关闭数据库连接
		db, _ := global.WM_DB.DB()
		defer db.Close()

	}

	core.RunWindowsServer()

}

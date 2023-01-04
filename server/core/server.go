package core

import (
	"fmt"
	"time"

	"github.com/weilaim/wm-admin/server/global"
	"github.com/weilaim/wm-admin/server/initialize"
	"github.com/weilaim/wm-admin/server/service/system"
	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	if global.WM_CONFIG.System.UseMultipoint || global.WM_CONFIG.System.UseRedis {
		// 初始化redis服务
		initialize.Redis()
	}

	// 从db加载jwt数据
	if global.WM_DB != nil {
		system.LoadAll()
	}

	Router := initialize.Routers()

	address := fmt.Sprintf(":%d", global.WM_CONFIG.System.Addr)
	s := initServer(address, Router)
	//保证文本顺序输出
	time.Sleep(10 * time.Microsecond)
	global.WM_LOG.Info("serever run success on", zap.String("address", address))

	fmt.Printf(`
	欢迎使用wm-admin
	默认自动化文档地址:http://127.0.0.1%s/swagger/index.html
	默认前端文件运行地址:http://127.0.0.1:8080
	`, address)
	global.WM_LOG.Error(s.ListenAndServe().Error())
}

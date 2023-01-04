package system

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/weilaim/wm-admin/server/api/v1"
)


type InitRouter struct {}

func (s *InitRouter)InitInitRouter(Router *gin.RouterGroup) {
	initRouter := Router.Group("init")
	dbApi := v1.ApiGroupApp.SystemApiGroup.DBApi
	{
		initRouter.POST("initdb",dbApi.InitDB) //初始化数据库
		initRouter.POST("checkdb",dbApi.CheckDB) // 检查是否需要初始化数据库
	}
}
package system

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/weilaim/wm-admin/server/api/v1"
	"github.com/weilaim/wm-admin/server/middleware"
)

type UserRouter struct{}

func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user").Use(middleware.OperationRecord())
	// userRouterWithoutRecord := Router.Group("user")
	baseApi := v1.ApiGroupApp.SystemApiGroup.BaseApi
	{
		userRouter.POST("admin_register", baseApi.Register) // 管理员注册账号
	}
}

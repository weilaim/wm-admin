package middleware

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/weilaim/wm-admin/server/global"
	"github.com/weilaim/wm-admin/server/model/common/response"
	"github.com/weilaim/wm-admin/server/service"
	"github.com/weilaim/wm-admin/server/utils"
)

var casbinService = service.ServiceGroupApp.SystemServiceGroup.CasbinService

// 拦截器
func CasbinHandler()gin.HandlerFunc {
	return func(c *gin.Context) {
		if global.WM_CONFIG.System.Env != "develop"{
			waitUse, _ := utils.GetClaims(c)
			// 请求的PATH
			obj := c.Request.URL.Path
			// 获取请求方法
			act := c.Request.Method
			// 获取用户的角色
			sub := strconv.Itoa(int(waitUse.AuthorityID))
			e := casbinService.Casbin() // 判断策略中是否存在

			success, _ := e.Enforce(sub,obj,act)
			if !success {
				response.FailWithDetailed(gin.H{},"权限不足",c)
				c.Abort()
				return 
			}
		}
		c.Next()
	}
}
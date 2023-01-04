package middleware

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/weilaim/wm-admin/server/global"
	"github.com/weilaim/wm-admin/server/model/common/response"
	"github.com/weilaim/wm-admin/server/model/system"
	"github.com/weilaim/wm-admin/server/service"
	"github.com/weilaim/wm-admin/server/utils"
	"go.uber.org/zap"
)

var jwtService = service.ServiceGroupApp.SystemServiceGroup.JwtService

func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 我们从这里jwt鉴权取头部信息 x-token 登录时返回token信息 这里前端需要把token存储到
		// cookie 或者localStorage中 不过需要和后端商定过期时间 可以约定刷新令牌或者重新登录
		token := ctx.Request.Header.Get("x-token")
		if token == "" {
			response.FailWithDetailed(gin.H{"reload": true}, "未登录或非法访问", ctx)
			ctx.Abort()
			return
		}
		if jwtService.IsBlacklist(token) {
			response.FailWithDetailed(gin.H{"reload": true}, "您的账户异地登录或令牌失效", ctx)
			ctx.Abort()
			return
		}
		j := utils.NewJWT()
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == utils.TokenExpired {
				response.FailWithDetailed(gin.H{"reload": true}, "授权已过期", ctx)
				ctx.Abort()
				return
			}
			response.FailWithDetailed(gin.H{"reload": true}, err.Error(), ctx)
			ctx.Abort()
			return
		}
		// 已登录用户被管理员禁用 需要使该用户的jwt失效 此处比较消耗性能 如果需要 请自行打开
		// 用户被删除的逻辑 需要优化 此处比较消耗性能 如果需要 请自行打开

		//if user, err := userService.FindUserByUuid(claims.UUID.String()); err != nil || user.Enable == 2 {
		//	_ = jwtService.JsonInBlacklist(system.JwtBlacklist{Jwt: token})
		//	response.FailWithDetailed(gin.H{"reload": true}, err.Error(), c)
		//	c.Abort()
		//}
		if claims.ExpiresAt-time.Now().Unix() < claims.BufferTime {
			dr, _ := utils.ParseDuration(global.WM_CONFIG.JWT.ExpiresTime)
			claims.ExpiresAt = time.Now().Add(dr).Unix()
			newToken, _ := j.CreateTokenByOldToken(token, *claims)
			newClaims, _ := j.ParseToken(newToken)
			ctx.Header("new-token", newToken)
			ctx.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt, 10))
			if global.WM_CONFIG.System.UseMultipoint {
				RedisJwtToken, err := jwtService.GetRedisJWT(newClaims.UserName)
				if err != nil {
					global.WM_LOG.Error("get redis jwt failed", zap.Error(err))
				} else { // 当之前的取成功时才进行拉黑操作
					_ = jwtService.JsonInBlacklist(system.JwtBlacklist{Jwt: RedisJwtToken})
				}
				// 无论如何都要记录当前的活跃状态
				_ = jwtService.SetRedisJWT(newToken, newClaims.UserName)
			}
		}
		ctx.Set("claims", claims)
		ctx.Next()
	}
}

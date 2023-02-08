package system

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/weilaim/wm-admin/server/global"
	"github.com/weilaim/wm-admin/server/model/common/response"
	"github.com/weilaim/wm-admin/server/model/system"
	systemReq "github.com/weilaim/wm-admin/server/model/system/request"
	systemRes "github.com/weilaim/wm-admin/server/model/system/response"
	"github.com/weilaim/wm-admin/server/utils"
	"go.uber.org/zap"
)

func (b *BaseApi) Login(c *gin.Context) {
	fmt.Println("hhhhhhhh---hhh")
}

// Register
// 用户注册账号。
func (b *BaseApi) Register(c *gin.Context) {
	var r systemReq.Register
	fmt.Println(r.Phone)
	err := c.ShouldBindJSON(&r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = utils.Verify(r, utils.RegisterVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	var authorities []system.SysAuthority
	for _, v := range r.AuthorityIds {
		authorities = append(authorities, system.SysAuthority{
			AuthorityId: v,
		})
	}

	user := &system.SysUser{Username: r.UserName, NickName: r.NickName, Password: r.Password, HeaderImg: r.HeaderImg, AuthorityId: r.AuthorityId, Authorities: authorities, Enable: int(r.Enable), Phone: r.Phone, Email: r.Email}
	userReturn, err := userService.Register(*user)
	if err != nil {
		global.WM_LOG.Error("注册失败!", zap.Error(err))
		response.FailWithDetailed(systemRes.SysUserResponse{User: userReturn}, "注册失败", c)
		return
	}

	response.OkWithDetailed(systemRes.SysUserResponse{User: userReturn}, "注册成功", c)

}

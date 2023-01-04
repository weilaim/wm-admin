package system

import (
	"errors"

	uuid "github.com/satori/go.uuid"
	"github.com/weilaim/wm-admin/server/global"
	"github.com/weilaim/wm-admin/server/model/system"
	"github.com/weilaim/wm-admin/server/utils"
	"gorm.io/gorm"
)

type UserService struct{}

func (userService *UserService) Register(u system.SysUser) (userInter system.SysUser, err error) {
	var user system.SysUser
	if !errors.Is(global.WM_DB.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return userInter, errors.New("用户名已注册")
	}
	// 否则 附加uuid 密码hash 加密注册
	u.Password = utils.BcryptHash(u.Password)
	u.UUID = uuid.NewV4()
	err = global.WM_DB.Create(&u).Error
	return u,err
}

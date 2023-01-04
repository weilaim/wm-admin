package request

import "github.com/weilaim/wm-admin/server/model/system"

type Register struct{
	UserName string `json:"user_name" example:"用户名"`
	Password string `json:"password" example:"密码"`
	NickName string `json:"nick_name" example:"昵称"`
	HeaderImg string `json:"header_img" example:"头像连接"`
	AuthorityId uint `json:"authority_id" example:"int 角色id"`
	Enable uint `json:"enable" example:"int 是否启用"`
	AuthorityIds  []uint `json:"authority_ids" example:"[]int 角色id"`
	Phone string `json:"phone" example:"电话号码"`
	Email string `json:"email" example:"电子邮箱"`

}

type Login struct{
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Captcha string `json:"captcha"`
	CaptchaId string `json:"captchaId"` //验证码id
}

type ChangePasswordReq struct {
	ID uint `json:"-"` // 从JWT中提取出user id 避免越权
	Password string `json:"password"`
	NewPassword string `json:"new_password"`
}


type SetUserAuth struct {
	AuthorityId uint `json:"authorityId"` // 角色ID
}

type SetUserAuthorities struct {
	ID uint
	AuthorityIds []uint `json:"authorityIds"` // 角色id
}

type ChangeUserInfo struct {
	ID uint `gorm:"primarykey"`
	NickName string `json:"nick_name" gorm:"default:系统用户;comment:用户昵称"`
	Phone string `json:"phone" gorm:"comment:用户手机号"`
	AuthorityIds []uint `json:"authority_ids" gorm:"-"`
	Email string `json:"email" gorm:"comment:用户邮箱"`
	HeaderImg string `json:"header_img" gorm:"default:https://miniwx.arf-to.cn/YNrUv2.jpg;comment:用户头像连接"`
	SideMode string `json:"side_mode" gorm:"comment:用户侧边栏主题"`
	Enable int `json:"enable" grom:"comment:冻结用户"`
	Authorities []system.SysAuthority `json:"-" gorm:"many2many:sys_user_authority"`

}
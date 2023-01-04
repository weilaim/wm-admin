package system

import "github.com/weilaim/wm-admin/server/global"

type JwtBlacklist struct {
	global.WM_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}

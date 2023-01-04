package system

import "github.com/weilaim/wm-admin/server/service"


type ApiGroup struct {
	BaseApi
	DBApi
}

var (
	initDBService = service.ServiceGroupApp.SystemServiceGroup.InitDBService
	userService = service.ServiceGroupApp.SystemServiceGroup.UserService
)

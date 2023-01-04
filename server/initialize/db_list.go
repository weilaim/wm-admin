package initialize

import (
	"github.com/weilaim/wm-admin/server/config"
	"github.com/weilaim/wm-admin/server/global"
	"gorm.io/gorm"
)

const sys = "system"

func DBList() {
	dbMap := make(map[string]*gorm.DB)
	for _, info := range global.WM_CONFIG.DBList {
		if info.Disable {
			continue
		}

		switch info.Type {
		case "mysql":
			dbMap[info.AliasName] = GormMysqlByConfig(config.Mysql{GeneralDB: info.GeneralDB})
		case "pgsql":
			dbMap[info.AliasName] = GormPgSqlByConfig(config.Pgsql{GeneralDB: info.GeneralDB})
		default:
			continue
		}
	}

	// 做特色判断，是否有迁移
	if sysDB, ok := dbMap[sys]; ok {
		global.WM_DB = sysDB
	}

	global.WM_DBlist = dbMap
}

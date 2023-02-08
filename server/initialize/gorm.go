package initialize

import (
	"os"

	"github.com/weilaim/wm-admin/server/global"
	"github.com/weilaim/wm-admin/server/model/system"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// 初始化数据库并产生数据库全局变量
func Gorm() *gorm.DB {
	switch global.WM_CONFIG.System.DbType {
	case "mysql":
		return GormMysql()
	case "pgsql":
		return GormPgSql()
	default:
		return GormMysql()
	}
}

// RegisterTables 注册数据库表专用
func RegisterTables(db *gorm.DB) {
	err := db.AutoMigrate(
		// 系统模块
		system.SysApi{},
		system.SysUser{},
		system.SysBaseMenu{},
		system.JwtBlacklist{},
		system.SysAuthority{},
		// system.SysDictionary{},
		system.SysOperationRecord{},
		// system.SysAutoCodeHistory{},
		// system.SysDictionaryDetail{},
		system.SysBaseMenuParameter{},
		system.SysBaseMenuBtn{},
		// system.SysAutoCode{},

	)
	if err != nil {
		global.WM_LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}

	global.WM_LOG.Info("register table success")
}

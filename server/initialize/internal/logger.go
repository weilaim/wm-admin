package internal

import (
	"fmt"

	"github.com/weilaim/wm-admin/server/global"
	"gorm.io/gorm/logger"
)

type writer struct {
	logger.Writer
}

// NewWriter 构造函数
func NewWriter(w logger.Writer) *writer {
	return &writer{Writer: w}
}

// Printf 格式化打印日志
func (w *writer) Printf(message string, data ...interface{}) {
	var logZap bool
	switch global.WM_CONFIG.System.DbType {
	case "mysql":
		logZap = global.WM_CONFIG.Mysql.LogZap
	case "pgsql":
		logZap = global.WM_CONFIG.Pgsql.LogZap
	}

	if logZap {
		global.WM_LOG.Info(fmt.Sprintf(message+"\n", data...))
	} else {
		w.Writer.Printf(message, data...)
	}
}

package core

import (
	"fmt"
	"os"

	"github.com/weilaim/wm-admin/server/core/internal"
	"github.com/weilaim/wm-admin/server/global"
	"github.com/weilaim/wm-admin/server/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Zap 获取zap.Logger
func Zap() (logger *zap.Logger) {
	if ok, _ := utils.PathExists(global.WM_CONFIG.Zap.Director); !ok {
		// 判断是否有director 文件
		fmt.Printf("crate %v directory\n", global.WM_CONFIG.Zap.Director)
		_ = os.Mkdir(global.WM_CONFIG.Zap.Director, os.ModePerm)
	}

	cores := internal.Zap.GetZapCores()
	logger = zap.New(zapcore.NewTee(cores...))
	if global.WM_CONFIG.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}

	return logger
}

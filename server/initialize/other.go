package initialize

import (
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/weilaim/wm-admin/server/global"
	"github.com/weilaim/wm-admin/server/utils"
)

func OtherInit() {
	dr, err := utils.ParseDuration(global.WM_CONFIG.JWT.ExpiresTime)
	if err != nil {
		panic(err)
	}

	_, err = utils.ParseDuration(global.WM_CONFIG.JWT.BufferTime)
	if err != nil {
		panic(err)
	}

	global.BlackCach = local_cache.NewCache(
		local_cache.SetDefaultExpire(dr),
	)
}

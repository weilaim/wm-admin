package initialize

import (
	"fmt"

	"github.com/robfig/cron/v3"
	"github.com/weilaim/wm-admin/server/config"
	"github.com/weilaim/wm-admin/server/global"
	"github.com/weilaim/wm-admin/server/utils"
)

func Timer() {
	if global.WM_CONFIG.Timer.Start {
		for i := range global.WM_CONFIG.Timer.Detail {
			go func(detail config.Detail) {
				var option []cron.Option
				if global.WM_CONFIG.Timer.WithSeconds {
					option = append(option, cron.WithChain())
				}

				_, err := global.WM_Timer.AddTaskByFunc("ClearDB", global.WM_CONFIG.Timer.Spec, func() {
					err := utils.ClearTable(global.WM_DB, detail.TableName, detail.CompareField, detail.Interval)
					if err != nil {
						fmt.Println("timer error:", err)
					}
				}, option...)
				if err != nil {
					fmt.Println("add timer error:", err)
				}
			}(global.WM_CONFIG.Timer.Detail[i])
		}
	}
}

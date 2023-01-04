package initialize

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/weilaim/wm-admin/server/global"
	"go.uber.org/zap"
)

func Redis() {
	redisCfg := global.WM_CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password,
		DB:       redisCfg.DB,
	})

	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.WM_LOG.Error("redis connect ping failed,err:", zap.Error(err))
	} else {
		global.WM_LOG.Info("redis connect ping response:", zap.String("pong", pong))
		global.WM_REDIS = client
	}
}

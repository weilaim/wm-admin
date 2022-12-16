package global

import (
	"sync"

	"github.com/go-redis/redis/v8"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"github.com/weilaim/wm-admin/server/config"
	"github.com/weilaim/wm-admin/server/utils/timer"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
)

var (
	WM_DB     *gorm.DB
	WM_DBlist map[string]*gorm.DB
	WM_REDIS  *redis.Client
	WM_CONFIG config.Server
	WM_VP     *viper.Viper
	// WM_LOG *oplogging.Logger

	WM_LOG                 *zap.Logger
	WM_Timer               timer.Timer = timer.NewTimerTask()
	WM_Concurrency_Control             = &singleflight.Group{}

	BlackCach local_cache.Cache
	lock      sync.RWMutex
)

// GetGlobalDBByDBname 通过名称获取db list中的db
func GetGlobalDBByDBname(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RLock()
	return WM_DBlist[dbname]
}

// 通过名称获取db 如果不存在则panic
func MustGetGlobalDBByName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	db, ok := WM_DBlist[dbname]
	if !ok || db == nil {
		panic("db no init")
	}
	return db
}

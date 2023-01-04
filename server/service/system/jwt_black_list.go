package system

import (
	"context"

	"github.com/weilaim/wm-admin/server/global"
	"github.com/weilaim/wm-admin/server/model/system"
	"github.com/weilaim/wm-admin/server/utils"
	"go.uber.org/zap"
)

type JwtService struct{}

// 拉黑jwt
func (jwtService *JwtService) JsonInBlacklist(jwtList system.JwtBlacklist) (err error) {
	err = global.WM_DB.Create(&jwtList).Error
	if err != nil {
		return
	}
	global.BlackCache.SetDefault(jwtList.Jwt, struct{}{})
	return
}

// 判断JWT是否在黑名单内部
func (jwtService *JwtService) IsBlacklist(jwt string) bool {
	_, ok := global.BlackCache.Get(jwt)
	return ok
}

// 从redis取jwt
func (jwtService *JwtService) GetRedisJWT(userName string) (redisJWT string, err error) {
	redisJWT, err = global.WM_REDIS.Get(context.Background(), userName).Result()
	return redisJWT, err
}


// jtd存入redis并设置过期时间
func (jwtService *JwtService)SetRedisJWT(jwt string,userName string) (err error){
	dr, err := utils.ParseDuration(global.WM_CONFIG.JWT.ExpiresTime)
	if err != nil {
		return err
	}
	timer := dr
	err = global.WM_REDIS.Set(context.Background(),userName,jwt,timer).Err()
	return err
}
// 加载jwt
func LoadAll() {
	var data []string
	err := global.WM_DB.Model(&system.JwtBlacklist{}).Select("jwt").Find(&data).Error
	if err != nil {
		global.WM_LOG.Error("加载数据库jwt黑名单失败!", zap.Error(err))
		return
	}
	for i := 0; i < len(data); i++ {
		global.BlackCache.SetDefault(data[i], struct{}{})
	} // jwt 黑名单加入BlackCache中

}

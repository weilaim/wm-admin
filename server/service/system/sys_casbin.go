package system

import (
	"errors"
	"strconv"
	"sync"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/weilaim/wm-admin/server/global"
	"github.com/weilaim/wm-admin/server/model/system/request"
	"go.uber.org/zap"
)

type CasbinService struct{}

var casbinServiceApp = new(CasbinService)

// 更新casbin 权限
func (casbinService *CasbinService) UpdateCasbin(AuthorityID uint, casbinInfos []request.CasbinInfo) error {
	authorityID := strconv.Itoa(int(AuthorityID))
	casbinService.ClearCasbin(0, authorityID)
	rules := [][]string{}
	for _, v := range casbinInfos {
		rules = append(rules, []string{authorityID, v.Path, v.Method})
	}

	e := casbinService.Casbin()
	success, _ := e.AddPolicies(rules)
	if !success {
		return errors.New("存在相同api,添加失败,请联系管理员")
	}
	err := e.InvalidateCache()
	if err != nil {
		return err
	}

	return nil
}

// 清除匹配的权限
func (casbinService *CasbinService) ClearCasbin(v int, p ...string) bool {
	e := casbinService.Casbin()
	success, _ := e.RemoveFilteredPolicy(v, p...)
	return success
}

// 持久化到数据库，引入自定义规则
var (
	cachedEnforcer *casbin.CachedEnforcer
	once           sync.Once
)

func (casbinService *CasbinService) Casbin() *casbin.CachedEnforcer {
	once.Do(func() {
		a, _ := gormadapter.NewAdapterByDB(global.WM_DB)
		text := `
		[request_definition]
		r = sub, obj, act
		
		[policy_definition]
		p = sub, obj, act
		
		[role_definition]
		g = _, _
		
		[policy_effect]
		e = some(where (p.eft == allow))
		
		[matchers]
		m = r.sub == p.sub && keyMatch2(r.obj,p.obj) && r.act == p.act
		`
		m, err := model.NewModelFromString(text)
		if err != nil {
			zap.L().Error("字符串加载模型失败！", zap.Error(err))
			return
		}

		cachedEnforcer, _ = casbin.NewCachedEnforcer(m, a)
		cachedEnforcer.SetExpireTime(60 * 60)
	})

	return cachedEnforcer
}

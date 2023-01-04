package router

import (
	"github.com/weilaim/wm-admin/server/router/example"
	"github.com/weilaim/wm-admin/server/router/system"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
}

var RouterGroupApp = new(RouterGroup)

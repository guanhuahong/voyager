package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/club"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
	"github.com/flipped-aurora/gin-vue-admin/server/router/voyager"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
	Club    club.RouterGroup
	Voyager voyager.RouterGroup
}

var RouterGroupApp = new(RouterGroup)

package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/club"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/voyager"
)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
	ExampleApiGroup example.ApiGroup
	ClubApiGroup    club.ApiGroup
	VoyagerApiGroup voyager.ApiGroup
}

var ApiGroupApp = new(ApiGroup)

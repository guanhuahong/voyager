package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/club"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service/voyager"
)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
	ClubServiceGroup    club.ServiceGroup
	VoyagerServiceGroup voyager.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)

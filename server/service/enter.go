package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/club"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
	ClubServiceGroup    club.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)

package voyager

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type PlayerProfileRouter struct {
}

// InitPlayerProfileRouter 初始化 PlayerProfile 路由信息
func (s *PlayerProfileRouter) InitPlayerProfileRouter(Router *gin.RouterGroup) {
	playerProfileRouter := Router.Group("playerProfile").Use(middleware.OperationRecord())
	playerProfileRouterWithoutRecord := Router.Group("playerProfile")
	var playerProfileApi = v1.ApiGroupApp.VoyagerApiGroup.PlayerProfileApi
	{
		playerProfileRouter.POST("createPlayerProfile", playerProfileApi.CreatePlayerProfile)   // 新建PlayerProfile
		playerProfileRouter.DELETE("deletePlayerProfile", playerProfileApi.DeletePlayerProfile) // 删除PlayerProfile
		playerProfileRouter.DELETE("deletePlayerProfileByIds", playerProfileApi.DeletePlayerProfileByIds) // 批量删除PlayerProfile
		playerProfileRouter.PUT("updatePlayerProfile", playerProfileApi.UpdatePlayerProfile)    // 更新PlayerProfile
	}
	{
		playerProfileRouterWithoutRecord.GET("findPlayerProfile", playerProfileApi.FindPlayerProfile)        // 根据ID获取PlayerProfile
		playerProfileRouterWithoutRecord.GET("getPlayerProfileList", playerProfileApi.GetPlayerProfileList)  // 获取PlayerProfile列表
	}
}

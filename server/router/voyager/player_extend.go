package voyager

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type PlayerExtendRouter struct {
}

// InitPlayerExtendRouter 初始化 PlayerExtend 路由信息
func (s *PlayerExtendRouter) InitPlayerExtendRouter(Router *gin.RouterGroup) {
	playerExtendRouter := Router.Group("playerExtend").Use(middleware.OperationRecord())
	playerExtendRouterWithoutRecord := Router.Group("playerExtend")
	var playerExtendApi = v1.ApiGroupApp.VoyagerApiGroup.PlayerExtendApi
	{
		playerExtendRouter.POST("createPlayerExtend", playerExtendApi.CreatePlayerExtend)   // 新建PlayerExtend
		playerExtendRouter.DELETE("deletePlayerExtend", playerExtendApi.DeletePlayerExtend) // 删除PlayerExtend
		playerExtendRouter.DELETE("deletePlayerExtendByIds", playerExtendApi.DeletePlayerExtendByIds) // 批量删除PlayerExtend
		playerExtendRouter.PUT("updatePlayerExtend", playerExtendApi.UpdatePlayerExtend)    // 更新PlayerExtend
	}
	{
		playerExtendRouterWithoutRecord.GET("findPlayerExtend", playerExtendApi.FindPlayerExtend)        // 根据ID获取PlayerExtend
		playerExtendRouterWithoutRecord.GET("getPlayerExtendList", playerExtendApi.GetPlayerExtendList)  // 获取PlayerExtend列表
	}
}

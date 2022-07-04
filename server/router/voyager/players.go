package voyager

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type PlayerRouter struct {
}

// InitPlayerRouter 初始化 Player 路由信息
func (s *PlayerRouter) InitPlayerRouter(Router *gin.RouterGroup) {
	playerRouter := Router.Group("player").Use(middleware.OperationRecord())
	playerRouterWithoutRecord := Router.Group("player")
	var playerApi = v1.ApiGroupApp.VoyagerApiGroup.PlayerApi
	{
		playerRouter.POST("createPlayer", playerApi.CreatePlayer)   // 新建Player
		playerRouter.DELETE("deletePlayer", playerApi.DeletePlayer) // 删除Player
		playerRouter.DELETE("deletePlayerByIds", playerApi.DeletePlayerByIds) // 批量删除Player
		playerRouter.PUT("updatePlayer", playerApi.UpdatePlayer)    // 更新Player
	}
	{
		playerRouterWithoutRecord.GET("findPlayer", playerApi.FindPlayer)        // 根据ID获取Player
		playerRouterWithoutRecord.GET("getPlayerList", playerApi.GetPlayerList)  // 获取Player列表
	}
}

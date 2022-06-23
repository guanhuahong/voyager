package club

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TbClubMemRouter struct {
}

// InitTbClubMemRouter 初始化 TbClubMem 路由信息
func (s *TbClubMemRouter) InitTbClubMemRouter(Router *gin.RouterGroup) {
	tbClubMemRouter := Router.Group("tbClubMem").Use(middleware.OperationRecord())
	tbClubMemRouterWithoutRecord := Router.Group("tbClubMem")
	var tbClubMemApi = v1.ApiGroupApp.ClubApiGroup.TbClubMemApi
	{
		tbClubMemRouter.POST("createTbClubMem", tbClubMemApi.CreateTbClubMem)   // 新建TbClubMem
		tbClubMemRouter.DELETE("deleteTbClubMem", tbClubMemApi.DeleteTbClubMem) // 删除TbClubMem
		tbClubMemRouter.DELETE("deleteTbClubMemByIds", tbClubMemApi.DeleteTbClubMemByIds) // 批量删除TbClubMem
		tbClubMemRouter.PUT("updateTbClubMem", tbClubMemApi.UpdateTbClubMem)    // 更新TbClubMem
	}
	{
		tbClubMemRouterWithoutRecord.GET("findTbClubMem", tbClubMemApi.FindTbClubMem)        // 根据ID获取TbClubMem
		tbClubMemRouterWithoutRecord.GET("getTbClubMemList", tbClubMemApi.GetTbClubMemList)  // 获取TbClubMem列表
	}
}

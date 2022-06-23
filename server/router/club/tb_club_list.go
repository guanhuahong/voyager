package club

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TbClubListRouter struct {
}

// InitTbClubListRouter 初始化 TbClubList 路由信息
func (s *TbClubListRouter) InitTbClubListRouter(Router *gin.RouterGroup) {
	tbClubListRouter := Router.Group("tbClubList").Use(middleware.OperationRecord())
	tbClubListRouterWithoutRecord := Router.Group("tbClubList")
	var tbClubListApi = v1.ApiGroupApp.ClubApiGroup.TbClubListApi
	{
		tbClubListRouter.POST("createTbClubList", tbClubListApi.CreateTbClubList)             // 新建TbClubList
		tbClubListRouter.DELETE("deleteTbClubList", tbClubListApi.DeleteTbClubList)           // 删除TbClubList
		tbClubListRouter.DELETE("deleteTbClubListByIds", tbClubListApi.DeleteTbClubListByIds) // 批量删除TbClubList
		tbClubListRouter.PUT("updateTbClubList", tbClubListApi.UpdateTbClubList)              // 更新TbClubList
	}
	{
		tbClubListRouterWithoutRecord.GET("findTbClubList", tbClubListApi.FindTbClubList)       // 根据ID获取TbClubList
		tbClubListRouterWithoutRecord.GET("getTbClubListList", tbClubListApi.GetTbClubListList) // 获取TbClubList列表
		tbClubListRouterWithoutRecord.GET("getClubOptions", tbClubListApi.GetClubOptions)       // 获取Option列表
	}
}

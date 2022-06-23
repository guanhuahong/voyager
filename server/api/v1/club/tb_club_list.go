package club

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/club"
	clubReq "github.com/flipped-aurora/gin-vue-admin/server/model/club/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type TbClubListApi struct {
}

var tbClubListService = service.ServiceGroupApp.ClubServiceGroup.TbClubListService

// CreateTbClubList 创建TbClubList
// @Tags TbClubList
// @Summary 创建TbClubList
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body club.TbClubList true "创建TbClubList"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tbClubList/createTbClubList [post]
func (tbClubListApi *TbClubListApi) CreateTbClubList(c *gin.Context) {
	var tbClubList club.TbClubList
	_ = c.ShouldBindJSON(&tbClubList)
	if err := tbClubListService.CreateTbClubList(tbClubList); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteTbClubList 删除TbClubList
// @Tags TbClubList
// @Summary 删除TbClubList
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body club.TbClubList true "删除TbClubList"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tbClubList/deleteTbClubList [delete]
func (tbClubListApi *TbClubListApi) DeleteTbClubList(c *gin.Context) {
	var tbClubList club.TbClubList
	_ = c.ShouldBindJSON(&tbClubList)
	if err := tbClubListService.DeleteTbClubList(tbClubList); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTbClubListByIds 批量删除TbClubList
// @Tags TbClubList
// @Summary 批量删除TbClubList
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TbClubList"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /tbClubList/deleteTbClubListByIds [delete]
func (tbClubListApi *TbClubListApi) DeleteTbClubListByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := tbClubListService.DeleteTbClubListByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTbClubList 更新TbClubList
// @Tags TbClubList
// @Summary 更新TbClubList
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body club.TbClubList true "更新TbClubList"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /tbClubList/updateTbClubList [put]
func (tbClubListApi *TbClubListApi) UpdateTbClubList(c *gin.Context) {
	var tbClubList club.TbClubList
	_ = c.ShouldBindJSON(&tbClubList)
	if err := tbClubListService.UpdateTbClubList(tbClubList); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindTbClubList 用id查询TbClubList
// @Tags TbClubList
// @Summary 用id查询TbClubList
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query club.TbClubList true "用id查询TbClubList"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /tbClubList/findTbClubList [get]
func (tbClubListApi *TbClubListApi) FindTbClubList(c *gin.Context) {
	var tbClubList club.TbClubList
	_ = c.ShouldBindQuery(&tbClubList)
	if retbClubList, err := tbClubListService.GetTbClubList(tbClubList.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"retbClubList": retbClubList}, c)
	}
}

// GetTbClubListList 分页获取TbClubList列表
// @Tags TbClubList
// @Summary 分页获取TbClubList列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query clubReq.TbClubListSearch true "分页获取TbClubList列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tbClubList/getTbClubListList [get]
func (tbClubListApi *TbClubListApi) GetTbClubListList(c *gin.Context) {
	var pageInfo clubReq.TbClubListSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := tbClubListService.GetTbClubListInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

func (TbClubListApi *TbClubListApi) GetClubOptions(c *gin.Context) {
	if clubs, err := tbClubListService.GetClubOptions(); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(gin.H{"clubs": clubs}, "获取成功", c)
	}
}

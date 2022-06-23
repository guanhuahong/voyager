package club

import (
	"log"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/club"
	clubReq "github.com/flipped-aurora/gin-vue-admin/server/model/club/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type TbClubMemApi struct {
}

var tbClubMemService = service.ServiceGroupApp.ClubServiceGroup.TbClubMemService

// CreateTbClubMem 创建TbClubMem
// @Tags TbClubMem
// @Summary 创建TbClubMem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body club.TbClubMem true "创建TbClubMem"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tbClubMem/createTbClubMem [post]
func (tbClubMemApi *TbClubMemApi) CreateTbClubMem(c *gin.Context) {
	var tbClubMem club.TbClubMem
	_ = c.ShouldBindJSON(&tbClubMem)
	if err := tbClubMemService.CreateTbClubMem(tbClubMem); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteTbClubMem 删除TbClubMem
// @Tags TbClubMem
// @Summary 删除TbClubMem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body club.TbClubMem true "删除TbClubMem"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tbClubMem/deleteTbClubMem [delete]
func (tbClubMemApi *TbClubMemApi) DeleteTbClubMem(c *gin.Context) {
	var tbClubMem club.TbClubMem
	_ = c.ShouldBindJSON(&tbClubMem)
	if err := tbClubMemService.DeleteTbClubMem(tbClubMem); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTbClubMemByIds 批量删除TbClubMem
// @Tags TbClubMem
// @Summary 批量删除TbClubMem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TbClubMem"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /tbClubMem/deleteTbClubMemByIds [delete]
func (tbClubMemApi *TbClubMemApi) DeleteTbClubMemByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := tbClubMemService.DeleteTbClubMemByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTbClubMem 更新TbClubMem
// @Tags TbClubMem
// @Summary 更新TbClubMem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body club.TbClubMem true "更新TbClubMem"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /tbClubMem/updateTbClubMem [put]
func (tbClubMemApi *TbClubMemApi) UpdateTbClubMem(c *gin.Context) {
	var tbClubMem club.TbClubMem
	_ = c.ShouldBindJSON(&tbClubMem)
	if err := tbClubMemService.UpdateTbClubMem(tbClubMem); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindTbClubMem 用id查询TbClubMem
// @Tags TbClubMem
// @Summary 用id查询TbClubMem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query club.TbClubMem true "用id查询TbClubMem"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /tbClubMem/findTbClubMem [get]
func (tbClubMemApi *TbClubMemApi) FindTbClubMem(c *gin.Context) {
	var tbClubMem club.TbClubMem
	_ = c.ShouldBindQuery(&tbClubMem)
	if retbClubMem, err := tbClubMemService.GetTbClubMem(tbClubMem.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"retbClubMem": retbClubMem}, c)
	}
}

// GetTbClubMemList 分页获取TbClubMem列表
// @Tags TbClubMem
// @Summary 分页获取TbClubMem列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query clubReq.TbClubMemSearch true "分页获取TbClubMem列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tbClubMem/getTbClubMemList [get]
func (tbClubMemApi *TbClubMemApi) GetTbClubMemList(c *gin.Context) {

	log.Println("DTCS", c.Query("dtcs"))
	var pageInfo clubReq.TbClubMemSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := tbClubMemService.GetTbClubMemInfoList(pageInfo); err != nil {
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

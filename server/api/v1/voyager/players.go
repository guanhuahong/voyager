package voyager

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/voyager"
	voyagerReq "github.com/flipped-aurora/gin-vue-admin/server/model/voyager/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PlayerApi struct {
}

var playerService = service.ServiceGroupApp.VoyagerServiceGroup.PlayerService

// CreatePlayer 创建Player
// @Tags Player
// @Summary 创建Player
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body voyager.Player true "创建Player"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /player/createPlayer [post]
func (playerApi *PlayerApi) CreatePlayer(c *gin.Context) {
	var playerForm voyagerReq.PlayerForm
	_ = c.ShouldBindJSON(&playerForm)
	playerCreated := voyager.PlayerCreated{}
	playerCreated.Player = playerForm.Player
	playerAuth := voyager.PlayerAuth{
		Username: playerForm.Player.Username,
		Password: playerForm.Password,
	}

	playerCreated.PlayerAuth = playerAuth

	if err := playerService.CreatePlayer(playerCreated); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}

}

// DeletePlayer 删除Player
// @Tags Player
// @Summary 删除Player
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body voyager.Player true "删除Player"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /player/deletePlayer [delete]
func (playerApi *PlayerApi) DeletePlayer(c *gin.Context) {
	var player voyager.Player
	_ = c.ShouldBindJSON(&player)
	if err := playerService.DeletePlayer(player); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeletePlayerByIds 批量删除Player
// @Tags Player
// @Summary 批量删除Player
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Player"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /player/deletePlayerByIds [delete]
func (playerApi *PlayerApi) DeletePlayerByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := playerService.DeletePlayerByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdatePlayer 更新Player
// @Tags Player
// @Summary 更新Player
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body voyager.Player true "更新Player"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /player/updatePlayer [put]
func (playerApi *PlayerApi) UpdatePlayer(c *gin.Context) {
	var player voyager.Player
	_ = c.ShouldBindJSON(&player)
	if err := playerService.UpdatePlayer(player); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindPlayer 用id查询Player
// @Tags Player
// @Summary 用id查询Player
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query voyager.Player true "用id查询Player"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /player/findPlayer [get]
func (playerApi *PlayerApi) FindPlayer(c *gin.Context) {
	var player voyager.Player
	_ = c.ShouldBindQuery(&player)
	if replayer, err := playerService.GetPlayer(player.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"replayer": replayer}, c)
	}
}

// GetPlayerList 分页获取Player列表
// @Tags Player
// @Summary 分页获取Player列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query voyagerReq.PlayerSearch true "分页获取Player列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /player/getPlayerList [get]
func (playerApi *PlayerApi) GetPlayerList(c *gin.Context) {
	var pageInfo voyagerReq.PlayerSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := playerService.GetPlayerInfoList(pageInfo); err != nil {
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

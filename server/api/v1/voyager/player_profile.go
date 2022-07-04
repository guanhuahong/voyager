package voyager

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/voyager"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    voyagerReq "github.com/flipped-aurora/gin-vue-admin/server/model/voyager/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type PlayerProfileApi struct {
}

var playerProfileService = service.ServiceGroupApp.VoyagerServiceGroup.PlayerProfileService


// CreatePlayerProfile 创建PlayerProfile
// @Tags PlayerProfile
// @Summary 创建PlayerProfile
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body voyager.PlayerProfile true "创建PlayerProfile"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /playerProfile/createPlayerProfile [post]
func (playerProfileApi *PlayerProfileApi) CreatePlayerProfile(c *gin.Context) {
	var playerProfile voyager.PlayerProfile
	_ = c.ShouldBindJSON(&playerProfile)
	if err := playerProfileService.CreatePlayerProfile(playerProfile); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeletePlayerProfile 删除PlayerProfile
// @Tags PlayerProfile
// @Summary 删除PlayerProfile
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body voyager.PlayerProfile true "删除PlayerProfile"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /playerProfile/deletePlayerProfile [delete]
func (playerProfileApi *PlayerProfileApi) DeletePlayerProfile(c *gin.Context) {
	var playerProfile voyager.PlayerProfile
	_ = c.ShouldBindJSON(&playerProfile)
	if err := playerProfileService.DeletePlayerProfile(playerProfile); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeletePlayerProfileByIds 批量删除PlayerProfile
// @Tags PlayerProfile
// @Summary 批量删除PlayerProfile
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PlayerProfile"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /playerProfile/deletePlayerProfileByIds [delete]
func (playerProfileApi *PlayerProfileApi) DeletePlayerProfileByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := playerProfileService.DeletePlayerProfileByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdatePlayerProfile 更新PlayerProfile
// @Tags PlayerProfile
// @Summary 更新PlayerProfile
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body voyager.PlayerProfile true "更新PlayerProfile"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /playerProfile/updatePlayerProfile [put]
func (playerProfileApi *PlayerProfileApi) UpdatePlayerProfile(c *gin.Context) {
	var playerProfile voyager.PlayerProfile
	_ = c.ShouldBindJSON(&playerProfile)
	if err := playerProfileService.UpdatePlayerProfile(playerProfile); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindPlayerProfile 用id查询PlayerProfile
// @Tags PlayerProfile
// @Summary 用id查询PlayerProfile
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query voyager.PlayerProfile true "用id查询PlayerProfile"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /playerProfile/findPlayerProfile [get]
func (playerProfileApi *PlayerProfileApi) FindPlayerProfile(c *gin.Context) {
	var playerProfile voyager.PlayerProfile
	_ = c.ShouldBindQuery(&playerProfile)
	if replayerProfile, err := playerProfileService.GetPlayerProfile(playerProfile.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"replayerProfile": replayerProfile}, c)
	}
}

// GetPlayerProfileList 分页获取PlayerProfile列表
// @Tags PlayerProfile
// @Summary 分页获取PlayerProfile列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query voyagerReq.PlayerProfileSearch true "分页获取PlayerProfile列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /playerProfile/getPlayerProfileList [get]
func (playerProfileApi *PlayerProfileApi) GetPlayerProfileList(c *gin.Context) {
	var pageInfo voyagerReq.PlayerProfileSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := playerProfileService.GetPlayerProfileInfoList(pageInfo); err != nil {
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

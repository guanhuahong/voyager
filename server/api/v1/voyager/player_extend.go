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

type PlayerExtendApi struct {
}

var playerExtendService = service.ServiceGroupApp.VoyagerServiceGroup.PlayerExtendService


// CreatePlayerExtend 创建PlayerExtend
// @Tags PlayerExtend
// @Summary 创建PlayerExtend
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body voyager.PlayerExtend true "创建PlayerExtend"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /playerExtend/createPlayerExtend [post]
func (playerExtendApi *PlayerExtendApi) CreatePlayerExtend(c *gin.Context) {
	var playerExtend voyager.PlayerExtend
	_ = c.ShouldBindJSON(&playerExtend)
	if err := playerExtendService.CreatePlayerExtend(playerExtend); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeletePlayerExtend 删除PlayerExtend
// @Tags PlayerExtend
// @Summary 删除PlayerExtend
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body voyager.PlayerExtend true "删除PlayerExtend"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /playerExtend/deletePlayerExtend [delete]
func (playerExtendApi *PlayerExtendApi) DeletePlayerExtend(c *gin.Context) {
	var playerExtend voyager.PlayerExtend
	_ = c.ShouldBindJSON(&playerExtend)
	if err := playerExtendService.DeletePlayerExtend(playerExtend); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeletePlayerExtendByIds 批量删除PlayerExtend
// @Tags PlayerExtend
// @Summary 批量删除PlayerExtend
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PlayerExtend"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /playerExtend/deletePlayerExtendByIds [delete]
func (playerExtendApi *PlayerExtendApi) DeletePlayerExtendByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := playerExtendService.DeletePlayerExtendByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdatePlayerExtend 更新PlayerExtend
// @Tags PlayerExtend
// @Summary 更新PlayerExtend
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body voyager.PlayerExtend true "更新PlayerExtend"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /playerExtend/updatePlayerExtend [put]
func (playerExtendApi *PlayerExtendApi) UpdatePlayerExtend(c *gin.Context) {
	var playerExtend voyager.PlayerExtend
	_ = c.ShouldBindJSON(&playerExtend)
	if err := playerExtendService.UpdatePlayerExtend(playerExtend); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindPlayerExtend 用id查询PlayerExtend
// @Tags PlayerExtend
// @Summary 用id查询PlayerExtend
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query voyager.PlayerExtend true "用id查询PlayerExtend"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /playerExtend/findPlayerExtend [get]
func (playerExtendApi *PlayerExtendApi) FindPlayerExtend(c *gin.Context) {
	var playerExtend voyager.PlayerExtend
	_ = c.ShouldBindQuery(&playerExtend)
	if replayerExtend, err := playerExtendService.GetPlayerExtend(playerExtend.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"replayerExtend": replayerExtend}, c)
	}
}

// GetPlayerExtendList 分页获取PlayerExtend列表
// @Tags PlayerExtend
// @Summary 分页获取PlayerExtend列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query voyagerReq.PlayerExtendSearch true "分页获取PlayerExtend列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /playerExtend/getPlayerExtendList [get]
func (playerExtendApi *PlayerExtendApi) GetPlayerExtendList(c *gin.Context) {
	var pageInfo voyagerReq.PlayerExtendSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := playerExtendService.GetPlayerExtendInfoList(pageInfo); err != nil {
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

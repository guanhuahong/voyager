import service from '@/utils/request'

// @Tags PlayerProfile
// @Summary 创建PlayerProfile
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PlayerProfile true "创建PlayerProfile"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /playerProfile/createPlayerProfile [post]
export const createPlayerProfile = (data) => {
  return service({
    url: '/playerProfile/createPlayerProfile',
    method: 'post',
    data
  })
}

// @Tags PlayerProfile
// @Summary 删除PlayerProfile
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PlayerProfile true "删除PlayerProfile"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /playerProfile/deletePlayerProfile [delete]
export const deletePlayerProfile = (data) => {
  return service({
    url: '/playerProfile/deletePlayerProfile',
    method: 'delete',
    data
  })
}

// @Tags PlayerProfile
// @Summary 删除PlayerProfile
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PlayerProfile"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /playerProfile/deletePlayerProfile [delete]
export const deletePlayerProfileByIds = (data) => {
  return service({
    url: '/playerProfile/deletePlayerProfileByIds',
    method: 'delete',
    data
  })
}

// @Tags PlayerProfile
// @Summary 更新PlayerProfile
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PlayerProfile true "更新PlayerProfile"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /playerProfile/updatePlayerProfile [put]
export const updatePlayerProfile = (data) => {
  return service({
    url: '/playerProfile/updatePlayerProfile',
    method: 'put',
    data
  })
}

// @Tags PlayerProfile
// @Summary 用id查询PlayerProfile
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.PlayerProfile true "用id查询PlayerProfile"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /playerProfile/findPlayerProfile [get]
export const findPlayerProfile = (params) => {
  return service({
    url: '/playerProfile/findPlayerProfile',
    method: 'get',
    params
  })
}

// @Tags PlayerProfile
// @Summary 分页获取PlayerProfile列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取PlayerProfile列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /playerProfile/getPlayerProfileList [get]
export const getPlayerProfileList = (params) => {
  return service({
    url: '/playerProfile/getPlayerProfileList',
    method: 'get',
    params
  })
}

import service from '@/utils/request'

// @Tags PlayerExtend
// @Summary 创建PlayerExtend
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PlayerExtend true "创建PlayerExtend"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /playerExtend/createPlayerExtend [post]
export const createPlayerExtend = (data) => {
  return service({
    url: '/playerExtend/createPlayerExtend',
    method: 'post',
    data
  })
}

// @Tags PlayerExtend
// @Summary 删除PlayerExtend
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PlayerExtend true "删除PlayerExtend"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /playerExtend/deletePlayerExtend [delete]
export const deletePlayerExtend = (data) => {
  return service({
    url: '/playerExtend/deletePlayerExtend',
    method: 'delete',
    data
  })
}

// @Tags PlayerExtend
// @Summary 删除PlayerExtend
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PlayerExtend"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /playerExtend/deletePlayerExtend [delete]
export const deletePlayerExtendByIds = (data) => {
  return service({
    url: '/playerExtend/deletePlayerExtendByIds',
    method: 'delete',
    data
  })
}

// @Tags PlayerExtend
// @Summary 更新PlayerExtend
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PlayerExtend true "更新PlayerExtend"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /playerExtend/updatePlayerExtend [put]
export const updatePlayerExtend = (data) => {
  return service({
    url: '/playerExtend/updatePlayerExtend',
    method: 'put',
    data
  })
}

// @Tags PlayerExtend
// @Summary 用id查询PlayerExtend
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.PlayerExtend true "用id查询PlayerExtend"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /playerExtend/findPlayerExtend [get]
export const findPlayerExtend = (params) => {
  return service({
    url: '/playerExtend/findPlayerExtend',
    method: 'get',
    params
  })
}

// @Tags PlayerExtend
// @Summary 分页获取PlayerExtend列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取PlayerExtend列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /playerExtend/getPlayerExtendList [get]
export const getPlayerExtendList = (params) => {
  return service({
    url: '/playerExtend/getPlayerExtendList',
    method: 'get',
    params
  })
}

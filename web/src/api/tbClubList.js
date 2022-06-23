import service from '@/utils/request'

// @Tags TbClubList
// @Summary 创建TbClubList
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TbClubList true "创建TbClubList"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tbClubList/createTbClubList [post]
export const createTbClubList = (data) => {
  return service({
    url: '/tbClubList/createTbClubList',
    method: 'post',
    data
  })
}

// @Tags TbClubList
// @Summary 删除TbClubList
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TbClubList true "删除TbClubList"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tbClubList/deleteTbClubList [delete]
export const deleteTbClubList = (data) => {
  return service({
    url: '/tbClubList/deleteTbClubList',
    method: 'delete',
    data
  })
}

// @Tags TbClubList
// @Summary 删除TbClubList
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TbClubList"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tbClubList/deleteTbClubList [delete]
export const deleteTbClubListByIds = (data) => {
  return service({
    url: '/tbClubList/deleteTbClubListByIds',
    method: 'delete',
    data
  })
}

// @Tags TbClubList
// @Summary 更新TbClubList
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TbClubList true "更新TbClubList"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /tbClubList/updateTbClubList [put]
export const updateTbClubList = (data) => {
  return service({
    url: '/tbClubList/updateTbClubList',
    method: 'put',
    data
  })
}

// @Tags TbClubList
// @Summary 用id查询TbClubList
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.TbClubList true "用id查询TbClubList"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /tbClubList/findTbClubList [get]
export const findTbClubList = (params) => {
  return service({
    url: '/tbClubList/findTbClubList',
    method: 'get',
    params
  })
}

// @Tags TbClubList
// @Summary 分页获取TbClubList列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取TbClubList列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tbClubList/getTbClubListList [get]
export const getTbClubListList = () => {
  return service({
    url: '/tbClubList/getTbClubListList',
    method: 'get'
  })
}


export const getClubOptions = (params) => {
  return service({
    url: '/tbClubList/getClubOptions',
    method: 'get',
    params
  })
}
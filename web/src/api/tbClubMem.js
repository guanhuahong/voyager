import service from '@/utils/request'

// @Tags TbClubMem
// @Summary 创建TbClubMem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TbClubMem true "创建TbClubMem"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tbClubMem/createTbClubMem [post]
export const createTbClubMem = (data) => {
  return service({
    url: '/tbClubMem/createTbClubMem',
    method: 'post',
    data
  })
}

// @Tags TbClubMem
// @Summary 删除TbClubMem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TbClubMem true "删除TbClubMem"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tbClubMem/deleteTbClubMem [delete]
export const deleteTbClubMem = (data) => {
  return service({
    url: '/tbClubMem/deleteTbClubMem',
    method: 'delete',
    data
  })
}

// @Tags TbClubMem
// @Summary 删除TbClubMem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TbClubMem"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tbClubMem/deleteTbClubMem [delete]
export const deleteTbClubMemByIds = (data) => {
  return service({
    url: '/tbClubMem/deleteTbClubMemByIds',
    method: 'delete',
    data
  })
}

// @Tags TbClubMem
// @Summary 更新TbClubMem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TbClubMem true "更新TbClubMem"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /tbClubMem/updateTbClubMem [put]
export const updateTbClubMem = (data) => {
  return service({
    url: '/tbClubMem/updateTbClubMem',
    method: 'put',
    data
  })
}

// @Tags TbClubMem
// @Summary 用id查询TbClubMem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.TbClubMem true "用id查询TbClubMem"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /tbClubMem/findTbClubMem [get]
export const findTbClubMem = (params) => {
  return service({
    url: '/tbClubMem/findTbClubMem',
    method: 'get',
    params
  })
}

// @Tags TbClubMem
// @Summary 分页获取TbClubMem列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取TbClubMem列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tbClubMem/getTbClubMemList [get]
export const getTbClubMemList = (params) => {
  return service({
    url: '/tbClubMem/getTbClubMemList',
    method: 'get',
    params
  })
}

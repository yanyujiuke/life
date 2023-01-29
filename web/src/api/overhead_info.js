import service from '@/utils/request'

// @Tags OverheadInfo
// @Summary 创建OverheadInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OverheadInfo true "创建OverheadInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Info/createOverheadInfo [post]
export const createOverheadInfo = (data) => {
  return service({
    url: '/Info/createOverheadInfo',
    method: 'post',
    data
  })
}

// @Tags OverheadInfo
// @Summary 删除OverheadInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OverheadInfo true "删除OverheadInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Info/deleteOverheadInfo [delete]
export const deleteOverheadInfo = (data) => {
  return service({
    url: '/Info/deleteOverheadInfo',
    method: 'delete',
    data
  })
}

// @Tags OverheadInfo
// @Summary 删除OverheadInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除OverheadInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Info/deleteOverheadInfo [delete]
export const deleteOverheadInfoByIds = (data) => {
  return service({
    url: '/Info/deleteOverheadInfoByIds',
    method: 'delete',
    data
  })
}

// @Tags OverheadInfo
// @Summary 更新OverheadInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OverheadInfo true "更新OverheadInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /Info/updateOverheadInfo [put]
export const updateOverheadInfo = (data) => {
  return service({
    url: '/Info/updateOverheadInfo',
    method: 'put',
    data
  })
}

// @Tags OverheadInfo
// @Summary 用id查询OverheadInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.OverheadInfo true "用id查询OverheadInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /Info/findOverheadInfo [get]
export const findOverheadInfo = (params) => {
  return service({
    url: '/Info/findOverheadInfo',
    method: 'get',
    params
  })
}

// @Tags OverheadInfo
// @Summary 分页获取OverheadInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取OverheadInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Info/getOverheadInfoList [get]
export const getOverheadInfoList = (params) => {
  return service({
    url: '/Info/getOverheadInfoList',
    method: 'get',
    params
  })
}

import service from '@/utils/request'

// @Tags OverheadMonthStatistics
// @Summary 创建OverheadMonthStatistics
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OverheadMonthStatistics true "创建OverheadMonthStatistics"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /MonthStatistics/createOverheadMonthStatistics [post]
export const createOverheadMonthStatistics = (data) => {
  return service({
    url: '/MonthStatistics/createOverheadMonthStatistics',
    method: 'post',
    data
  })
}

// @Tags OverheadMonthStatistics
// @Summary 删除OverheadMonthStatistics
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OverheadMonthStatistics true "删除OverheadMonthStatistics"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /MonthStatistics/deleteOverheadMonthStatistics [delete]
export const deleteOverheadMonthStatistics = (data) => {
  return service({
    url: '/MonthStatistics/deleteOverheadMonthStatistics',
    method: 'delete',
    data
  })
}

// @Tags OverheadMonthStatistics
// @Summary 删除OverheadMonthStatistics
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除OverheadMonthStatistics"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /MonthStatistics/deleteOverheadMonthStatistics [delete]
export const deleteOverheadMonthStatisticsByIds = (data) => {
  return service({
    url: '/MonthStatistics/deleteOverheadMonthStatisticsByIds',
    method: 'delete',
    data
  })
}

// @Tags OverheadMonthStatistics
// @Summary 更新OverheadMonthStatistics
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OverheadMonthStatistics true "更新OverheadMonthStatistics"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /MonthStatistics/updateOverheadMonthStatistics [put]
export const updateOverheadMonthStatistics = (data) => {
  return service({
    url: '/MonthStatistics/updateOverheadMonthStatistics',
    method: 'put',
    data
  })
}

// @Tags OverheadMonthStatistics
// @Summary 用id查询OverheadMonthStatistics
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.OverheadMonthStatistics true "用id查询OverheadMonthStatistics"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /MonthStatistics/findOverheadMonthStatistics [get]
export const findOverheadMonthStatistics = (params) => {
  return service({
    url: '/MonthStatistics/findOverheadMonthStatistics',
    method: 'get',
    params
  })
}

// @Tags OverheadMonthStatistics
// @Summary 分页获取OverheadMonthStatistics列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取OverheadMonthStatistics列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /MonthStatistics/getOverheadMonthStatisticsList [get]
export const getOverheadMonthStatisticsList = (params) => {
  return service({
    url: '/MonthStatistics/getOverheadMonthStatisticsList',
    method: 'get',
    params
  })
}

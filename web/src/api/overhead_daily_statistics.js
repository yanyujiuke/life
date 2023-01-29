import service from '@/utils/request'

// @Tags OverheadDailyStatistics
// @Summary 创建OverheadDailyStatistics
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OverheadDailyStatistics true "创建OverheadDailyStatistics"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /DailyStatistics/createOverheadDailyStatistics [post]
export const createOverheadDailyStatistics = (data) => {
  return service({
    url: '/DailyStatistics/createOverheadDailyStatistics',
    method: 'post',
    data
  })
}

// @Tags OverheadDailyStatistics
// @Summary 删除OverheadDailyStatistics
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OverheadDailyStatistics true "删除OverheadDailyStatistics"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /DailyStatistics/deleteOverheadDailyStatistics [delete]
export const deleteOverheadDailyStatistics = (data) => {
  return service({
    url: '/DailyStatistics/deleteOverheadDailyStatistics',
    method: 'delete',
    data
  })
}

// @Tags OverheadDailyStatistics
// @Summary 删除OverheadDailyStatistics
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除OverheadDailyStatistics"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /DailyStatistics/deleteOverheadDailyStatistics [delete]
export const deleteOverheadDailyStatisticsByIds = (data) => {
  return service({
    url: '/DailyStatistics/deleteOverheadDailyStatisticsByIds',
    method: 'delete',
    data
  })
}

// @Tags OverheadDailyStatistics
// @Summary 更新OverheadDailyStatistics
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OverheadDailyStatistics true "更新OverheadDailyStatistics"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /DailyStatistics/updateOverheadDailyStatistics [put]
export const updateOverheadDailyStatistics = (data) => {
  return service({
    url: '/DailyStatistics/updateOverheadDailyStatistics',
    method: 'put',
    data
  })
}

// @Tags OverheadDailyStatistics
// @Summary 用id查询OverheadDailyStatistics
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.OverheadDailyStatistics true "用id查询OverheadDailyStatistics"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /DailyStatistics/findOverheadDailyStatistics [get]
export const findOverheadDailyStatistics = (params) => {
  return service({
    url: '/DailyStatistics/findOverheadDailyStatistics',
    method: 'get',
    params
  })
}

// @Tags OverheadDailyStatistics
// @Summary 分页获取OverheadDailyStatistics列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取OverheadDailyStatistics列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /DailyStatistics/getOverheadDailyStatisticsList [get]
export const getOverheadDailyStatisticsList = (params) => {
  return service({
    url: '/DailyStatistics/getOverheadDailyStatisticsList',
    method: 'get',
    params
  })
}

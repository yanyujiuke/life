import service from '@/utils/request'

// @Tags OverheadCategory
// @Summary 创建OverheadCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OverheadCategory true "创建OverheadCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /category/createOverheadCategory [post]
export const createOverheadCategory = (data) => {
  return service({
    url: '/category/createOverheadCategory',
    method: 'post',
    data
  })
}

// @Tags OverheadCategory
// @Summary 删除OverheadCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OverheadCategory true "删除OverheadCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /category/deleteOverheadCategory [delete]
export const deleteOverheadCategory = (data) => {
  return service({
    url: '/category/deleteOverheadCategory',
    method: 'delete',
    data
  })
}

// @Tags OverheadCategory
// @Summary 删除OverheadCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除OverheadCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /category/deleteOverheadCategory [delete]
export const deleteOverheadCategoryByIds = (data) => {
  return service({
    url: '/category/deleteOverheadCategoryByIds',
    method: 'delete',
    data
  })
}

// @Tags OverheadCategory
// @Summary 更新OverheadCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OverheadCategory true "更新OverheadCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /category/updateOverheadCategory [put]
export const updateOverheadCategory = (data) => {
  return service({
    url: '/category/updateOverheadCategory',
    method: 'put',
    data
  })
}

// @Tags OverheadCategory
// @Summary 用id查询OverheadCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.OverheadCategory true "用id查询OverheadCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /category/findOverheadCategory [get]
export const findOverheadCategory = (params) => {
  return service({
    url: '/category/findOverheadCategory',
    method: 'get',
    params
  })
}

// @Tags OverheadCategory
// @Summary 分页获取OverheadCategory列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取OverheadCategory列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /category/getOverheadCategoryList [get]
export const getOverheadCategoryList = (params) => {
  return service({
    url: '/category/getOverheadCategoryList',
    method: 'get',
    params
  })
}

package overhead

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"life/global"
	"life/model/common/request"
	"life/model/common/response"
	"life/model/overhead"
	overheadReq "life/model/overhead/request"
	"life/service"
	"life/utils"
)

type OverheadCategoryApi struct {
}

var categoryService = service.ServiceGroupApp.OverheadServiceGroup.OverheadCategoryService

// CreateOverheadCategory 创建OverheadCategory
// @Tags OverheadCategory
// @Summary 创建OverheadCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body overhead.OverheadCategory true "创建OverheadCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /category/createOverheadCategory [post]
func (categoryApi *OverheadCategoryApi) CreateOverheadCategory(c *gin.Context) {
	var category overhead.OverheadCategory
	err := c.ShouldBindJSON(&category)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Name": {utils.NotEmpty()},
	}
	if err := utils.Verify(category, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := categoryService.CreateOverheadCategory(category); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteOverheadCategory 删除OverheadCategory
// @Tags OverheadCategory
// @Summary 删除OverheadCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body overhead.OverheadCategory true "删除OverheadCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /category/deleteOverheadCategory [delete]
func (categoryApi *OverheadCategoryApi) DeleteOverheadCategory(c *gin.Context) {
	var category overhead.OverheadCategory
	err := c.ShouldBindJSON(&category)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := categoryService.DeleteOverheadCategory(category); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteOverheadCategoryByIds 批量删除OverheadCategory
// @Tags OverheadCategory
// @Summary 批量删除OverheadCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除OverheadCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /category/deleteOverheadCategoryByIds [delete]
func (categoryApi *OverheadCategoryApi) DeleteOverheadCategoryByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := categoryService.DeleteOverheadCategoryByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateOverheadCategory 更新OverheadCategory
// @Tags OverheadCategory
// @Summary 更新OverheadCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body overhead.OverheadCategory true "更新OverheadCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /category/updateOverheadCategory [put]
func (categoryApi *OverheadCategoryApi) UpdateOverheadCategory(c *gin.Context) {
	var category overhead.OverheadCategory
	err := c.ShouldBindJSON(&category)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Name": {utils.NotEmpty()},
	}
	if err := utils.Verify(category, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := categoryService.UpdateOverheadCategory(category); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindOverheadCategory 用id查询OverheadCategory
// @Tags OverheadCategory
// @Summary 用id查询OverheadCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query overhead.OverheadCategory true "用id查询OverheadCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /category/findOverheadCategory [get]
func (categoryApi *OverheadCategoryApi) FindOverheadCategory(c *gin.Context) {
	var category overhead.OverheadCategory
	err := c.ShouldBindQuery(&category)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if recategory, err := categoryService.GetOverheadCategory(category.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recategory": recategory}, c)
	}
}

// GetOverheadCategoryList 分页获取OverheadCategory列表
// @Tags OverheadCategory
// @Summary 分页获取OverheadCategory列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query overheadReq.OverheadCategorySearch true "分页获取OverheadCategory列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /category/getOverheadCategoryList [get]
func (categoryApi *OverheadCategoryApi) GetOverheadCategoryList(c *gin.Context) {
	var pageInfo overheadReq.OverheadCategorySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := categoryService.GetOverheadCategoryInfoList(pageInfo); err != nil {
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

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

type OverheadMonthStatisticsApi struct {
}

var MonthStatisticsService = service.ServiceGroupApp.OverheadServiceGroup.OverheadMonthStatisticsService

// CreateOverheadMonthStatistics 创建OverheadMonthStatistics
// @Tags OverheadMonthStatistics
// @Summary 创建OverheadMonthStatistics
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body overhead.OverheadMonthStatistics true "创建OverheadMonthStatistics"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /MonthStatistics/createOverheadMonthStatistics [post]
func (MonthStatisticsApi *OverheadMonthStatisticsApi) CreateOverheadMonthStatistics(c *gin.Context) {
	var MonthStatistics overhead.OverheadMonthStatistics
	err := c.ShouldBindJSON(&MonthStatistics)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Uid":    {utils.NotEmpty()},
		"Month":  {utils.NotEmpty()},
		"Amount": {utils.NotEmpty()},
		"Num":    {utils.NotEmpty()},
	}
	if err := utils.Verify(MonthStatistics, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := MonthStatisticsService.CreateOverheadMonthStatistics(MonthStatistics); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteOverheadMonthStatistics 删除OverheadMonthStatistics
// @Tags OverheadMonthStatistics
// @Summary 删除OverheadMonthStatistics
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body overhead.OverheadMonthStatistics true "删除OverheadMonthStatistics"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /MonthStatistics/deleteOverheadMonthStatistics [delete]
func (MonthStatisticsApi *OverheadMonthStatisticsApi) DeleteOverheadMonthStatistics(c *gin.Context) {
	var MonthStatistics overhead.OverheadMonthStatistics
	err := c.ShouldBindJSON(&MonthStatistics)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := MonthStatisticsService.DeleteOverheadMonthStatistics(MonthStatistics); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteOverheadMonthStatisticsByIds 批量删除OverheadMonthStatistics
// @Tags OverheadMonthStatistics
// @Summary 批量删除OverheadMonthStatistics
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除OverheadMonthStatistics"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /MonthStatistics/deleteOverheadMonthStatisticsByIds [delete]
func (MonthStatisticsApi *OverheadMonthStatisticsApi) DeleteOverheadMonthStatisticsByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := MonthStatisticsService.DeleteOverheadMonthStatisticsByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateOverheadMonthStatistics 更新OverheadMonthStatistics
// @Tags OverheadMonthStatistics
// @Summary 更新OverheadMonthStatistics
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body overhead.OverheadMonthStatistics true "更新OverheadMonthStatistics"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /MonthStatistics/updateOverheadMonthStatistics [put]
func (MonthStatisticsApi *OverheadMonthStatisticsApi) UpdateOverheadMonthStatistics(c *gin.Context) {
	var MonthStatistics overhead.OverheadMonthStatistics
	err := c.ShouldBindJSON(&MonthStatistics)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Uid":    {utils.NotEmpty()},
		"Month":  {utils.NotEmpty()},
		"Amount": {utils.NotEmpty()},
		"Num":    {utils.NotEmpty()},
	}
	if err := utils.Verify(MonthStatistics, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := MonthStatisticsService.UpdateOverheadMonthStatistics(MonthStatistics); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindOverheadMonthStatistics 用id查询OverheadMonthStatistics
// @Tags OverheadMonthStatistics
// @Summary 用id查询OverheadMonthStatistics
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query overhead.OverheadMonthStatistics true "用id查询OverheadMonthStatistics"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /MonthStatistics/findOverheadMonthStatistics [get]
func (MonthStatisticsApi *OverheadMonthStatisticsApi) FindOverheadMonthStatistics(c *gin.Context) {
	var MonthStatistics overhead.OverheadMonthStatistics
	err := c.ShouldBindQuery(&MonthStatistics)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reMonthStatistics, err := MonthStatisticsService.GetOverheadMonthStatistics(MonthStatistics.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reMonthStatistics": reMonthStatistics}, c)
	}
}

// GetOverheadMonthStatisticsList 分页获取OverheadMonthStatistics列表
// @Tags OverheadMonthStatistics
// @Summary 分页获取OverheadMonthStatistics列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query overheadReq.OverheadMonthStatisticsSearch true "分页获取OverheadMonthStatistics列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /MonthStatistics/getOverheadMonthStatisticsList [get]
func (MonthStatisticsApi *OverheadMonthStatisticsApi) GetOverheadMonthStatisticsList(c *gin.Context) {
	var pageInfo overheadReq.OverheadMonthStatisticsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := MonthStatisticsService.GetOverheadMonthStatisticsInfoList(pageInfo); err != nil {
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

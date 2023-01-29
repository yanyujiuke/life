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

type OverheadDailyStatisticsApi struct {
}

var DailyStatisticsService = service.ServiceGroupApp.OverheadServiceGroup.OverheadDailyStatisticsService

// CreateOverheadDailyStatistics 创建OverheadDailyStatistics
// @Tags OverheadDailyStatistics
// @Summary 创建OverheadDailyStatistics
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body overhead.OverheadDailyStatistics true "创建OverheadDailyStatistics"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /DailyStatistics/createOverheadDailyStatistics [post]
func (DailyStatisticsApi *OverheadDailyStatisticsApi) CreateOverheadDailyStatistics(c *gin.Context) {
	var DailyStatistics overhead.OverheadDailyStatistics
	err := c.ShouldBindJSON(&DailyStatistics)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Uid":    {utils.NotEmpty()},
		"Day":    {utils.NotEmpty()},
		"Amount": {utils.NotEmpty()},
		"Num":    {utils.NotEmpty()},
	}
	if err := utils.Verify(DailyStatistics, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := DailyStatisticsService.CreateOverheadDailyStatistics(DailyStatistics); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteOverheadDailyStatistics 删除OverheadDailyStatistics
// @Tags OverheadDailyStatistics
// @Summary 删除OverheadDailyStatistics
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body overhead.OverheadDailyStatistics true "删除OverheadDailyStatistics"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /DailyStatistics/deleteOverheadDailyStatistics [delete]
func (DailyStatisticsApi *OverheadDailyStatisticsApi) DeleteOverheadDailyStatistics(c *gin.Context) {
	var DailyStatistics overhead.OverheadDailyStatistics
	err := c.ShouldBindJSON(&DailyStatistics)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := DailyStatisticsService.DeleteOverheadDailyStatistics(DailyStatistics); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteOverheadDailyStatisticsByIds 批量删除OverheadDailyStatistics
// @Tags OverheadDailyStatistics
// @Summary 批量删除OverheadDailyStatistics
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除OverheadDailyStatistics"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /DailyStatistics/deleteOverheadDailyStatisticsByIds [delete]
func (DailyStatisticsApi *OverheadDailyStatisticsApi) DeleteOverheadDailyStatisticsByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := DailyStatisticsService.DeleteOverheadDailyStatisticsByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateOverheadDailyStatistics 更新OverheadDailyStatistics
// @Tags OverheadDailyStatistics
// @Summary 更新OverheadDailyStatistics
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body overhead.OverheadDailyStatistics true "更新OverheadDailyStatistics"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /DailyStatistics/updateOverheadDailyStatistics [put]
func (DailyStatisticsApi *OverheadDailyStatisticsApi) UpdateOverheadDailyStatistics(c *gin.Context) {
	var DailyStatistics overhead.OverheadDailyStatistics
	err := c.ShouldBindJSON(&DailyStatistics)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Uid":    {utils.NotEmpty()},
		"Day":    {utils.NotEmpty()},
		"Amount": {utils.NotEmpty()},
		"Num":    {utils.NotEmpty()},
	}
	if err := utils.Verify(DailyStatistics, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := DailyStatisticsService.UpdateOverheadDailyStatistics(DailyStatistics); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindOverheadDailyStatistics 用id查询OverheadDailyStatistics
// @Tags OverheadDailyStatistics
// @Summary 用id查询OverheadDailyStatistics
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query overhead.OverheadDailyStatistics true "用id查询OverheadDailyStatistics"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /DailyStatistics/findOverheadDailyStatistics [get]
func (DailyStatisticsApi *OverheadDailyStatisticsApi) FindOverheadDailyStatistics(c *gin.Context) {
	var DailyStatistics overhead.OverheadDailyStatistics
	err := c.ShouldBindQuery(&DailyStatistics)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reDailyStatistics, err := DailyStatisticsService.GetOverheadDailyStatistics(DailyStatistics.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reDailyStatistics": reDailyStatistics}, c)
	}
}

// GetOverheadDailyStatisticsList 分页获取OverheadDailyStatistics列表
// @Tags OverheadDailyStatistics
// @Summary 分页获取OverheadDailyStatistics列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query overheadReq.OverheadDailyStatisticsSearch true "分页获取OverheadDailyStatistics列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /DailyStatistics/getOverheadDailyStatisticsList [get]
func (DailyStatisticsApi *OverheadDailyStatisticsApi) GetOverheadDailyStatisticsList(c *gin.Context) {
	var pageInfo overheadReq.OverheadDailyStatisticsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := DailyStatisticsService.GetOverheadDailyStatisticsInfoList(pageInfo); err != nil {
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

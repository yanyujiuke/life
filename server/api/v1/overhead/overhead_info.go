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

type OverheadInfoApi struct {
}

var InfoService = service.ServiceGroupApp.OverheadServiceGroup.OverheadInfoService

// CreateOverheadInfo 创建OverheadInfo
// @Tags OverheadInfo
// @Summary 创建OverheadInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body overhead.OverheadInfo true "创建OverheadInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Info/createOverheadInfo [post]
func (InfoApi *OverheadInfoApi) CreateOverheadInfo(c *gin.Context) {
	var Info overhead.OverheadInfo
	err := c.ShouldBindJSON(&Info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Uid":        {utils.NotEmpty()},
		"Day":        {utils.NotEmpty()},
		"CategoryId": {utils.NotEmpty()},
		"Name":       {utils.NotEmpty()},
		"Amount":     {utils.NotEmpty()},
	}
	if err := utils.Verify(Info, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := InfoService.CreateOverheadInfo(Info); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteOverheadInfo 删除OverheadInfo
// @Tags OverheadInfo
// @Summary 删除OverheadInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body overhead.OverheadInfo true "删除OverheadInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Info/deleteOverheadInfo [delete]
func (InfoApi *OverheadInfoApi) DeleteOverheadInfo(c *gin.Context) {
	var Info overhead.OverheadInfo
	err := c.ShouldBindJSON(&Info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := InfoService.DeleteOverheadInfo(Info); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteOverheadInfoByIds 批量删除OverheadInfo
// @Tags OverheadInfo
// @Summary 批量删除OverheadInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除OverheadInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /Info/deleteOverheadInfoByIds [delete]
func (InfoApi *OverheadInfoApi) DeleteOverheadInfoByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := InfoService.DeleteOverheadInfoByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateOverheadInfo 更新OverheadInfo
// @Tags OverheadInfo
// @Summary 更新OverheadInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body overhead.OverheadInfo true "更新OverheadInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /Info/updateOverheadInfo [put]
func (InfoApi *OverheadInfoApi) UpdateOverheadInfo(c *gin.Context) {
	var Info overhead.OverheadInfo
	err := c.ShouldBindJSON(&Info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Uid":        {utils.NotEmpty()},
		"Day":        {utils.NotEmpty()},
		"CategoryId": {utils.NotEmpty()},
		"Name":       {utils.NotEmpty()},
		"Amount":     {utils.NotEmpty()},
	}
	if err := utils.Verify(Info, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := InfoService.UpdateOverheadInfo(Info); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindOverheadInfo 用id查询OverheadInfo
// @Tags OverheadInfo
// @Summary 用id查询OverheadInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query overhead.OverheadInfo true "用id查询OverheadInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /Info/findOverheadInfo [get]
func (InfoApi *OverheadInfoApi) FindOverheadInfo(c *gin.Context) {
	var Info overhead.OverheadInfo
	err := c.ShouldBindQuery(&Info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reInfo, err := InfoService.GetOverheadInfo(Info.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reInfo": reInfo}, c)
	}
}

// GetOverheadInfoList 分页获取OverheadInfo列表
// @Tags OverheadInfo
// @Summary 分页获取OverheadInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query overheadReq.OverheadInfoSearch true "分页获取OverheadInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Info/getOverheadInfoList [get]
func (InfoApi *OverheadInfoApi) GetOverheadInfoList(c *gin.Context) {
	var pageInfo overheadReq.OverheadInfoSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := InfoService.GetOverheadInfoInfoList(pageInfo); err != nil {
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

package overhead

import (
	"github.com/gin-gonic/gin"
	"life/api/v1"
	"life/middleware"
)

type OverheadMonthStatisticsRouter struct {
}

// InitOverheadMonthStatisticsRouter 初始化 OverheadMonthStatistics 路由信息
func (s *OverheadMonthStatisticsRouter) InitOverheadMonthStatisticsRouter(Router *gin.RouterGroup) {
	MonthStatisticsRouter := Router.Group("MonthStatistics").Use(middleware.OperationRecord())
	MonthStatisticsRouterWithoutRecord := Router.Group("MonthStatistics")
	var MonthStatisticsApi = v1.ApiGroupApp.OverheadApiGroup.OverheadMonthStatisticsApi
	{
		MonthStatisticsRouter.POST("createOverheadMonthStatistics", MonthStatisticsApi.CreateOverheadMonthStatistics)             // 新建OverheadMonthStatistics
		MonthStatisticsRouter.DELETE("deleteOverheadMonthStatistics", MonthStatisticsApi.DeleteOverheadMonthStatistics)           // 删除OverheadMonthStatistics
		MonthStatisticsRouter.DELETE("deleteOverheadMonthStatisticsByIds", MonthStatisticsApi.DeleteOverheadMonthStatisticsByIds) // 批量删除OverheadMonthStatistics
		MonthStatisticsRouter.PUT("updateOverheadMonthStatistics", MonthStatisticsApi.UpdateOverheadMonthStatistics)              // 更新OverheadMonthStatistics
	}
	{
		MonthStatisticsRouterWithoutRecord.GET("findOverheadMonthStatistics", MonthStatisticsApi.FindOverheadMonthStatistics)       // 根据ID获取OverheadMonthStatistics
		MonthStatisticsRouterWithoutRecord.GET("getOverheadMonthStatisticsList", MonthStatisticsApi.GetOverheadMonthStatisticsList) // 获取OverheadMonthStatistics列表
	}
}

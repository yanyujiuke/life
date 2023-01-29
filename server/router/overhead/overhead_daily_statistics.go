package overhead

import (
	"github.com/gin-gonic/gin"
	"life/api/v1"
	"life/middleware"
)

type OverheadDailyStatisticsRouter struct {
}

// InitOverheadDailyStatisticsRouter 初始化 OverheadDailyStatistics 路由信息
func (s *OverheadDailyStatisticsRouter) InitOverheadDailyStatisticsRouter(Router *gin.RouterGroup) {
	DailyStatisticsRouter := Router.Group("DailyStatistics").Use(middleware.OperationRecord())
	DailyStatisticsRouterWithoutRecord := Router.Group("DailyStatistics")
	var DailyStatisticsApi = v1.ApiGroupApp.OverheadApiGroup.OverheadDailyStatisticsApi
	{
		DailyStatisticsRouter.POST("createOverheadDailyStatistics", DailyStatisticsApi.CreateOverheadDailyStatistics)             // 新建OverheadDailyStatistics
		DailyStatisticsRouter.DELETE("deleteOverheadDailyStatistics", DailyStatisticsApi.DeleteOverheadDailyStatistics)           // 删除OverheadDailyStatistics
		DailyStatisticsRouter.DELETE("deleteOverheadDailyStatisticsByIds", DailyStatisticsApi.DeleteOverheadDailyStatisticsByIds) // 批量删除OverheadDailyStatistics
		DailyStatisticsRouter.PUT("updateOverheadDailyStatistics", DailyStatisticsApi.UpdateOverheadDailyStatistics)              // 更新OverheadDailyStatistics
	}
	{
		DailyStatisticsRouterWithoutRecord.GET("findOverheadDailyStatistics", DailyStatisticsApi.FindOverheadDailyStatistics)       // 根据ID获取OverheadDailyStatistics
		DailyStatisticsRouterWithoutRecord.GET("getOverheadDailyStatisticsList", DailyStatisticsApi.GetOverheadDailyStatisticsList) // 获取OverheadDailyStatistics列表
	}
}

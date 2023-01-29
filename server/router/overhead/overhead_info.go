package overhead

import (
	"github.com/gin-gonic/gin"
	"life/api/v1"
	"life/middleware"
)

type OverheadInfoRouter struct {
}

// InitOverheadInfoRouter 初始化 OverheadInfo 路由信息
func (s *OverheadInfoRouter) InitOverheadInfoRouter(Router *gin.RouterGroup) {
	InfoRouter := Router.Group("Info").Use(middleware.OperationRecord())
	InfoRouterWithoutRecord := Router.Group("Info")
	var InfoApi = v1.ApiGroupApp.OverheadApiGroup.OverheadInfoApi
	{
		InfoRouter.POST("createOverheadInfo", InfoApi.CreateOverheadInfo)             // 新建OverheadInfo
		InfoRouter.DELETE("deleteOverheadInfo", InfoApi.DeleteOverheadInfo)           // 删除OverheadInfo
		InfoRouter.DELETE("deleteOverheadInfoByIds", InfoApi.DeleteOverheadInfoByIds) // 批量删除OverheadInfo
		InfoRouter.PUT("updateOverheadInfo", InfoApi.UpdateOverheadInfo)              // 更新OverheadInfo
	}
	{
		InfoRouterWithoutRecord.GET("findOverheadInfo", InfoApi.FindOverheadInfo)       // 根据ID获取OverheadInfo
		InfoRouterWithoutRecord.GET("getOverheadInfoList", InfoApi.GetOverheadInfoList) // 获取OverheadInfo列表
	}
}

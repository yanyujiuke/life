package overhead

import (
	"github.com/gin-gonic/gin"
	"life/api/v1"
	"life/middleware"
)

type OverheadCategoryRouter struct {
}

// InitOverheadCategoryRouter 初始化 OverheadCategory 路由信息
func (s *OverheadCategoryRouter) InitOverheadCategoryRouter(Router *gin.RouterGroup) {
	categoryRouter := Router.Group("category").Use(middleware.OperationRecord())
	categoryRouterWithoutRecord := Router.Group("category")
	var categoryApi = v1.ApiGroupApp.OverheadApiGroup.OverheadCategoryApi
	{
		categoryRouter.POST("createOverheadCategory", categoryApi.CreateOverheadCategory)             // 新建OverheadCategory
		categoryRouter.DELETE("deleteOverheadCategory", categoryApi.DeleteOverheadCategory)           // 删除OverheadCategory
		categoryRouter.DELETE("deleteOverheadCategoryByIds", categoryApi.DeleteOverheadCategoryByIds) // 批量删除OverheadCategory
		categoryRouter.PUT("updateOverheadCategory", categoryApi.UpdateOverheadCategory)              // 更新OverheadCategory
	}
	{
		categoryRouterWithoutRecord.GET("findOverheadCategory", categoryApi.FindOverheadCategory)       // 根据ID获取OverheadCategory
		categoryRouterWithoutRecord.GET("getOverheadCategoryList", categoryApi.GetOverheadCategoryList) // 获取OverheadCategory列表
	}
}

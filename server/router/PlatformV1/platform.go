package PlatformV1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type PlatformRouter struct{}

// InitPlatformRouter 初始化 平台 路由信息
func (s *PlatformRouter) InitPlatformRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	platformRouter := Router.Group("platform").Use(middleware.OperationRecord())
	platformRouterWithoutRecord := Router.Group("platform")
	platformRouterWithoutAuth := PublicRouter.Group("platform")
	{
		platformRouter.POST("createPlatform", platformApi.CreatePlatform)             // 新建平台
		platformRouter.DELETE("deletePlatform", platformApi.DeletePlatform)           // 删除平台
		platformRouter.DELETE("deletePlatformByIds", platformApi.DeletePlatformByIds) // 批量删除平台
		platformRouter.PUT("updatePlatform", platformApi.UpdatePlatform)              // 更新平台
	}
	{
		platformRouterWithoutRecord.GET("findPlatform", platformApi.FindPlatform)       // 根据ID获取平台
		platformRouterWithoutRecord.GET("getPlatformList", platformApi.GetPlatformList) // 获取平台列表
	}
	{
		platformRouterWithoutAuth.GET("getPlatformPublic", platformApi.GetPlatformPublic) // 平台开放接口
	}

	platformMobileRouterWithoutAuth := PublicRouter.Group("platformMobile")
	{
		platformMobileRouterWithoutAuth.GET("getPlatformList", platformApi.GetPlatformList) // 获取空投项目列表
	}
}

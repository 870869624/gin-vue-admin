package NavigationBar

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type NavigationBarRouter struct{}

// InitNavigationBarRouter 初始化 导航栏分类 路由信息
func (s *NavigationBarRouter) InitNavigationBarRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	NARouter := Router.Group("NA").Use(middleware.OperationRecord())
	NARouterWithoutRecord := Router.Group("NA")
	NARouterWithoutAuth := PublicRouter.Group("NA")
	{
		NARouter.POST("createNavigationBar", NAApi.CreateNavigationBar)             // 新建导航栏分类
		NARouter.DELETE("deleteNavigationBar", NAApi.DeleteNavigationBar)           // 删除导航栏分类
		NARouter.DELETE("deleteNavigationBarByIds", NAApi.DeleteNavigationBarByIds) // 批量删除导航栏分类
		NARouter.PUT("updateNavigationBar", NAApi.UpdateNavigationBar)              // 更新导航栏分类
	}
	{
		NARouterWithoutRecord.GET("findNavigationBar", NAApi.FindNavigationBar)       // 根据ID获取导航栏分类
		NARouterWithoutRecord.GET("getNavigationBarList", NAApi.GetNavigationBarList) // 获取导航栏分类列表
	}
	{
		NARouterWithoutAuth.GET("getNavigationBarPublic", NAApi.GetNavigationBarPublic) // 导航栏分类开放接口
	}

	navigationMobileRouterWithoutAuth := PublicRouter.Group("navigationBarMobile")
	{
		navigationMobileRouterWithoutAuth.GET("find", NAApi.FindNavigationBar)       // 根据ID获取预售项目
		navigationMobileRouterWithoutAuth.GET("getList", NAApi.GetNavigationBarList) // 获取预售项目列表
	}
}

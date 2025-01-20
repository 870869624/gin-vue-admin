package NavigationProject

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type NavigationProjectRouter struct{}

// InitNavigationProjectRouter 初始化 导航栏项目 路由信息
func (s *NavigationProjectRouter) InitNavigationProjectRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	NPRouter := Router.Group("NP").Use(middleware.OperationRecord())
	NPRouterWithoutRecord := Router.Group("NP")
	NPRouterWithoutAuth := PublicRouter.Group("NP")
	{
		NPRouter.POST("createNavigationProject", NPApi.CreateNavigationProject)             // 新建导航栏项目
		NPRouter.DELETE("deleteNavigationProject", NPApi.DeleteNavigationProject)           // 删除导航栏项目
		NPRouter.DELETE("deleteNavigationProjectByIds", NPApi.DeleteNavigationProjectByIds) // 批量删除导航栏项目
		NPRouter.PUT("updateNavigationProject", NPApi.UpdateNavigationProject)              // 更新导航栏项目
	}
	{
		NPRouterWithoutRecord.GET("findNavigationProject", NPApi.FindNavigationProject)       // 根据ID获取导航栏项目
		NPRouterWithoutRecord.GET("getNavigationProjectList", NPApi.GetNavigationProjectList) // 获取导航栏项目列表
	}
	{
		NPRouterWithoutAuth.GET("getNavigationProjectPublic", NPApi.GetNavigationProjectPublic) // 导航栏项目开放接口
	}

	NPRouterMobileRouterWithoutAuth := PublicRouter.Group("navigationMobile")
	{
		NPRouterMobileRouterWithoutAuth.GET("find", NPApi.FindNavigationProject)       // 根据ID获取预售项目
		NPRouterMobileRouterWithoutAuth.GET("getList", NPApi.GetNavigationProjectList) // 获取预售项目列表
	}
}

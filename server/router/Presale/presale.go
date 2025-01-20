package Presale

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type PresaleRouter struct{}

// InitPresaleRouter 初始化 预售项目 路由信息
func (s *PresaleRouter) InitPresaleRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	presaleRouter := Router.Group("presale").Use(middleware.OperationRecord())
	presaleRouterWithoutRecord := Router.Group("presale")
	presaleRouterWithoutAuth := PublicRouter.Group("presale")
	{
		presaleRouter.POST("createPresale", presaleApi.CreatePresale)             // 新建预售项目
		presaleRouter.DELETE("deletePresale", presaleApi.DeletePresale)           // 删除预售项目
		presaleRouter.DELETE("deletePresaleByIds", presaleApi.DeletePresaleByIds) // 批量删除预售项目
		presaleRouter.PUT("updatePresale", presaleApi.UpdatePresale)              // 更新预售项目
	}
	{
		presaleRouterWithoutRecord.GET("findPresale", presaleApi.FindPresale)       // 根据ID获取预售项目
		presaleRouterWithoutRecord.GET("getPresaleList", presaleApi.GetPresaleList) // 获取预售项目列表
	}
	{
		presaleRouterWithoutAuth.GET("getPresalePublic", presaleApi.GetPresalePublic) // 预售项目开放接口
	}

	presaleMobileRouterWithoutAuth := PublicRouter.Group("presaleMobile")
	{
		presaleMobileRouterWithoutAuth.GET("findPresale", presaleApi.FindPresaleMobile)       // 根据ID获取预售项目
		presaleMobileRouterWithoutAuth.GET("getPresaleList", presaleApi.MobileGetPresaleList) // 获取预售项目列表
	}

	presaleMobileRouterAuth := PublicRouter.Group("presaleMobileAuth")
	presaleMobileRouterAuth.Use(middleware.MoblileJWTAuth()).Use(middleware.OperationRecord())
	{
		presaleMobileRouterAuth.POST("createPresale", presaleApi.MobileCreatePresale) // 新建预售项目
	}

}

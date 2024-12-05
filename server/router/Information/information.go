package Information

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type InformationRouter struct{}

// InitInformationRouter 初始化 资讯 路由信息
func (s *InformationRouter) InitInformationRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	informationRouter := Router.Group("information").Use(middleware.OperationRecord())
	informationRouterWithoutRecord := Router.Group("information")
	informationRouterWithoutAuth := PublicRouter.Group("information")
	{
		informationRouter.POST("createInformation", informationApi.CreateInformation)             // 新建资讯
		informationRouter.DELETE("deleteInformation", informationApi.DeleteInformation)           // 删除资讯
		informationRouter.DELETE("deleteInformationByIds", informationApi.DeleteInformationByIds) // 批量删除资讯
		informationRouter.PUT("updateInformation", informationApi.UpdateInformation)              // 更新资讯
	}
	{
		informationRouterWithoutRecord.GET("findInformation", informationApi.FindInformation)       // 根据ID获取资讯
		informationRouterWithoutRecord.GET("getInformationList", informationApi.GetInformationList) // 获取资讯列表
	}
	{
		informationRouterWithoutAuth.GET("getInformationPublic", informationApi.GetInformationPublic) // 资讯开放接口
	}

	//用户使用
	informationMobileRouter := PublicRouter.Group("infoMobile")
	{
		informationMobileRouter.GET("getInformationList", informationApi.GetInformationList)
		informationMobileRouter.GET("findInformation", informationApi.FindInformation) // 根据ID获取资讯
	}

}

package PublicChain

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type PublicChainRouter struct{}

// InitPublicChainRouter 初始化 公链 路由信息
func (s *PublicChainRouter) InitPublicChainRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	PCRouter := Router.Group("PC").Use(middleware.OperationRecord())
	PCRouterWithoutRecord := Router.Group("PC")
	PCRouterWithoutAuth := PublicRouter.Group("PC")
	{
		PCRouter.POST("createPublicChain", PCApi.CreatePublicChain)             // 新建公链
		PCRouter.DELETE("deletePublicChain", PCApi.DeletePublicChain)           // 删除公链
		PCRouter.DELETE("deletePublicChainByIds", PCApi.DeletePublicChainByIds) // 批量删除公链
		PCRouter.PUT("updatePublicChain", PCApi.UpdatePublicChain)              // 更新公链
	}
	{
		PCRouterWithoutRecord.GET("findPublicChain", PCApi.FindPublicChain)       // 根据ID获取公链
		PCRouterWithoutRecord.GET("getPublicChainList", PCApi.GetPublicChainList) // 获取公链列表
	}
	{
		PCRouterWithoutAuth.GET("getPublicChainPublic", PCApi.GetPublicChainPublic) // 公链开放接口
	}

	PCMobileRouter := PublicRouter.Group("PCMobile")
	{
		PCMobileRouter.GET("getPublicChainList", PCApi.GetPublicChainListMobile)
		// PCMobilePriRouter := PCMobileRouter.Use(middleware.MoblileJWTAuth())
		// {
		// }

	}
}

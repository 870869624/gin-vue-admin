package Airdrop

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AirdropRouter struct{}

// InitAirdropRouter 初始化 空投项目 路由信息
func (s *AirdropRouter) InitAirdropRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	airdropRouter := Router.Group("airdrop").Use(middleware.OperationRecord())
	airdropRouterWithoutRecord := Router.Group("airdrop")
	airdropRouterWithoutAuth := PublicRouter.Group("airdrop")
	{
		airdropRouter.POST("createAirdrop", airdropApi.CreateAirdrop)             // 新建空投项目
		airdropRouter.DELETE("deleteAirdrop", airdropApi.DeleteAirdrop)           // 删除空投项目
		airdropRouter.DELETE("deleteAirdropByIds", airdropApi.DeleteAirdropByIds) // 批量删除空投项目
		airdropRouter.PUT("updateAirdrop", airdropApi.UpdateAirdrop)              // 更新空投项目
	}
	{
		airdropRouterWithoutRecord.GET("findAirdrop", airdropApi.FindAirdrop)       // 根据ID获取空投项目
		airdropRouterWithoutRecord.GET("getAirdropList", airdropApi.GetAirdropList) // 获取空投项目列表
	}
	{
		airdropRouterWithoutAuth.GET("getAirdropPublic", airdropApi.GetAirdropPublic) // 空投项目开放接口
	}

	airdropMobileRouter := PublicRouter.Group("airdropMobile").Use(middleware.MoblileJWTAuth()).Use(middleware.OperationRecord())
	{
		airdropMobileRouter.POST("createAirdrop", airdropApi.MobileCreateAirdrop)
	}
	airdropMobileRouterWithoutAuth := PublicRouter.Group("airdropMobile")
	{
		airdropMobileRouterWithoutAuth.GET("getAirdropList", airdropApi.MobileGetAirdropList) // 获取空投项目列表
	}
}

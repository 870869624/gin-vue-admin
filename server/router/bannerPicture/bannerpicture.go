package bannerPicture

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type BannerPictureRouter struct{}

// InitBannerPictureRouter 初始化 横幅图 路由信息
func (s *BannerPictureRouter) InitBannerPictureRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	BPRouter := Router.Group("BP").Use(middleware.OperationRecord())
	BPRouterWithoutRecord := Router.Group("BP")
	BPRouterWithoutAuth := PublicRouter.Group("BP")
	{
		BPRouter.POST("createBannerPicture", BPApi.CreateBannerPicture)             // 新建横幅图
		BPRouter.DELETE("deleteBannerPicture", BPApi.DeleteBannerPicture)           // 删除横幅图
		BPRouter.DELETE("deleteBannerPictureByIds", BPApi.DeleteBannerPictureByIds) // 批量删除横幅图
		BPRouter.PUT("updateBannerPicture", BPApi.UpdateBannerPicture)              // 更新横幅图
	}
	{
		BPRouterWithoutRecord.GET("findBannerPicture", BPApi.FindBannerPicture)       // 根据ID获取横幅图
		BPRouterWithoutRecord.GET("getBannerPictureList", BPApi.GetBannerPictureList) // 获取横幅图列表
	}
	{
		BPRouterWithoutAuth.GET("getBannerPicturePublic", BPApi.GetBannerPictureList) // 横幅图开放接口
	}
}

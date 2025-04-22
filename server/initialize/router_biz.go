package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router"
	"github.com/gin-gonic/gin"
)

func holder(routers ...*gin.RouterGroup) {
	_ = routers
	_ = router.RouterGroupApp
}
func initBizRouter(routers ...*gin.RouterGroup) {
	privateGroup := routers[0]
	publicGroup := routers[1]
	holder(publicGroup, privateGroup)
	{
		PublicChainRouter := router.RouterGroupApp.PublicChain
		PublicChainRouter.InitPublicChainRouter(privateGroup, publicGroup)
	}
	{
		PlatformV1Router := router.RouterGroupApp.PlatformV1
		PlatformV1Router.InitPlatformRouter(privateGroup, publicGroup)
	}
	{
		VoteRecordRouter := router.RouterGroupApp.VoteRecord
		VoteRecordRouter.InitVoteRecordRouter(privateGroup, publicGroup)
	}
	{
		InformationRouter := router.RouterGroupApp.Information
		InformationRouter.InitInformationRouter(privateGroup, publicGroup)
	}
	{
		VipRouter := router.RouterGroupApp.Vip
		VipRouter.InitVipRecordRouter(privateGroup, publicGroup)
	}
	{
		UsersRouter := router.RouterGroupApp.Users
		UsersRouter.InitUsersRouter(privateGroup, publicGroup)
	}
	{
		AirdropRouter := router.RouterGroupApp.Airdrop
		AirdropRouter.InitAirdropRouter(privateGroup, publicGroup)
	}
	{
		VoteRouter := router.RouterGroupApp.Vote
		VoteRouter.InitVoteRouter(privateGroup, publicGroup)
	}
	{
		PresaleRouter := router.RouterGroupApp.Presale
		PresaleRouter.InitPresaleRouter(privateGroup, publicGroup)
	}
	{
		NavigationBarRouter := router.RouterGroupApp.NavigationBar
		NavigationBarRouter.InitNavigationBarRouter(privateGroup, publicGroup)
	}
	{
		NavigationProjectRouter := router.RouterGroupApp.NavigationProject
		NavigationProjectRouter.InitNavigationProjectRouter(privateGroup, publicGroup)
	}
	{
		describeRouter := router.RouterGroupApp.Describe
		describeRouter.InitDescribeRouter(privateGroup, publicGroup)
	}
	{
		InvitationRecordRouter := router.RouterGroupApp.InvitationRecord
		InvitationRecordRouter.InitInvitationRecordRouter(privateGroup, publicGroup)
	} // 占位方法，保证文件可以正确加载，避免go空变量检测报错，请勿删除。
	{
		bannerPictureRouter := router.RouterGroupApp.BannerPicture
		bannerPictureRouter.InitBannerPictureRouter(privateGroup, publicGroup)
	}
}

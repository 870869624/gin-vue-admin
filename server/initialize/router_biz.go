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
		InvitationRecordRouter := router.RouterGroupApp.InvitationRecord
		InvitationRecordRouter.InitInvitationRecordRouter(privateGroup, publicGroup)
	}
	{
		AirdropRouter := router.RouterGroupApp.Airdrop
		AirdropRouter.InitAirdropRouter(privateGroup, publicGroup)
	}
	{
		PlatformV1Router := router.RouterGroupApp.PlatformV1
		PlatformV1Router.InitPlatformRouter(privateGroup, publicGroup)
	}
	{
		VoteRouter := router.RouterGroupApp.Vote
		VoteRouter.InitVoteRouter(privateGroup, publicGroup)
	}
	{
		VoteRecordRouter := router.RouterGroupApp.VoteRecord
		VoteRecordRouter.InitVoteRecordRouter(privateGroup, publicGroup)
	}
	{
		PresaleRouter := router.RouterGroupApp.Presale
		PresaleRouter.InitPresaleRouter(privateGroup, publicGroup)
	}
	{
		InformationRouter := router.RouterGroupApp.Information
		InformationRouter.InitInformationRouter(privateGroup, publicGroup)
	}
	{
		VipRouter := router.RouterGroupApp.Vip
		VipRouter.InitVipRecordRouter(privateGroup, publicGroup)
	} // 占位方法，保证文件可以正确加载，避免go空变量检测报错，请勿删除。
	{
		UsersRouter := router.RouterGroupApp.Users
		UsersRouter.InitUsersRouter(privateGroup, publicGroup)
	}
}

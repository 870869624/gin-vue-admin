package Vote

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type VoteRouter struct{}

// InitVoteRouter 初始化 投票 路由信息
func (s *VoteRouter) InitVoteRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	voteRouter := Router.Group("vote").Use(middleware.OperationRecord())
	voteRouterWithoutRecord := Router.Group("vote")
	voteRouterWithoutAuth := PublicRouter.Group("vote")
	{
		voteRouter.POST("createVote", voteApi.CreateVote)             // 新建投票
		voteRouter.DELETE("deleteVote", voteApi.DeleteVote)           // 删除投票
		voteRouter.DELETE("deleteVoteByIds", voteApi.DeleteVoteByIds) // 批量删除投票
		voteRouter.PUT("updateVote", voteApi.UpdateVote)              // 更新投票
	}
	{
		voteRouterWithoutRecord.GET("findVote", voteApi.FindVote)       // 根据ID获取投票
		voteRouterWithoutRecord.GET("getVoteList", voteApi.GetVoteList) // 获取投票列表
	}
	{
		voteRouterWithoutAuth.GET("getVotePublic", voteApi.GetVotePublic) // 投票开放接口
	}

	voteMobileRouterWithoutAuth := PublicRouter.Group("voteMobile")
	{
		voteMobileRouterWithoutAuth.GET("getVoteList", voteApi.MobileGetVoteList) // 获取投票列表
		voteMobileRouterWithoutAuth.GET("findVote", voteApi.FindVote)             // 根据ID获取投票
	}

	voteMobileRouterAuth := PublicRouter.Group("voteMobileAuth")
	voteMobileRouterAuth.Use(middleware.MoblileJWTAuth()).Use(middleware.OperationRecord())
	{
		voteMobileRouterAuth.PUT("vote", voteApi.Vote)                    // 投票
		voteMobileRouterAuth.POST("createVote", voteApi.MobileCreateVote) // 新建投票
	}
}

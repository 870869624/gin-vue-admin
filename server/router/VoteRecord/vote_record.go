package VoteRecord

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type VoteRecordRouter struct{}

// InitVoteRecordRouter 初始化 投票记录 路由信息
func (s *VoteRecordRouter) InitVoteRecordRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	voteRecordRouter := Router.Group("voteRecord").Use(middleware.OperationRecord())
	voteRecordRouterWithoutRecord := Router.Group("voteRecord")
	voteRecordRouterWithoutAuth := PublicRouter.Group("voteRecord")
	{
		voteRecordRouter.POST("createVoteRecord", voteRecordApi.CreateVoteRecord)             // 新建投票记录
		voteRecordRouter.DELETE("deleteVoteRecord", voteRecordApi.DeleteVoteRecord)           // 删除投票记录
		voteRecordRouter.DELETE("deleteVoteRecordByIds", voteRecordApi.DeleteVoteRecordByIds) // 批量删除投票记录
		voteRecordRouter.PUT("updateVoteRecord", voteRecordApi.UpdateVoteRecord)              // 更新投票记录
	}
	{
		voteRecordRouterWithoutRecord.GET("findVoteRecord", voteRecordApi.FindVoteRecord)       // 根据ID获取投票记录
		voteRecordRouterWithoutRecord.GET("getVoteRecordList", voteRecordApi.GetVoteRecordList) // 获取投票记录列表
	}
	{
		voteRecordRouterWithoutAuth.GET("getVoteRecordPublic", voteRecordApi.GetVoteRecordPublic) // 投票记录开放接口
	}

	voteRecordMobileRouterWithoutAuth := PublicRouter.Group("voteRecordMobile")
	{
		voteRecordMobileRouterWithoutAuth.GET("findVoteRecord", voteRecordApi.FindVoteRecord) // 根据ID获取投票记录
	}
}

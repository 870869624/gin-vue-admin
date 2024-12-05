package InvitationRecord

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type InvitationRecordRouter struct{}

// InitInvitationRecordRouter 初始化 邀请记录 路由信息
func (s *InvitationRecordRouter) InitInvitationRecordRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	IRRouter := Router.Group("IR").Use(middleware.OperationRecord())
	IRRouterWithoutRecord := Router.Group("IR")
	IRRouterWithoutAuth := PublicRouter.Group("IR")
	{
		IRRouter.POST("createInvitationRecord", IRApi.CreateInvitationRecord)             // 新建邀请记录
		IRRouter.DELETE("deleteInvitationRecord", IRApi.DeleteInvitationRecord)           // 删除邀请记录
		IRRouter.DELETE("deleteInvitationRecordByIds", IRApi.DeleteInvitationRecordByIds) // 批量删除邀请记录
		IRRouter.PUT("updateInvitationRecord", IRApi.UpdateInvitationRecord)              // 更新邀请记录
	}
	{
		IRRouterWithoutRecord.GET("findInvitationRecord", IRApi.FindInvitationRecord)       // 根据ID获取邀请记录
		IRRouterWithoutRecord.GET("getInvitationRecordList", IRApi.GetInvitationRecordList) // 获取邀请记录列表
	}
	{
		IRRouterWithoutAuth.GET("getInvitationRecordPublic", IRApi.GetInvitationRecordPublic) // 邀请记录开放接口
	}
	IRMobileRouterWithoutAuth := PublicRouter.Group("IRMobile")
	{
		IRMobileRouterWithoutAuth.POST("createInvitationRecord", IRApi.CreateInvitationRecord) // 新建邀请记录
	}
}

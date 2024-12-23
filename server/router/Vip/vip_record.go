package Vip

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type VipRecordRouter struct{}

// InitVipRecordRouter 初始化 会员记录 路由信息
func (s *VipRecordRouter) InitVipRecordRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	vipRecordRouter := Router.Group("vipRecord").Use(middleware.OperationRecord())
	vipRecordRouterWithoutRecord := Router.Group("vipRecord")
	vipRecordRouterWithoutAuth := PublicRouter.Group("vipRecord")
	{
		vipRecordRouter.POST("createVipRecord", vipRecordApi.CreateVipRecord)             // 新建会员记录
		vipRecordRouter.DELETE("deleteVipRecord", vipRecordApi.DeleteVipRecord)           // 删除会员记录
		vipRecordRouter.DELETE("deleteVipRecordByIds", vipRecordApi.DeleteVipRecordByIds) // 批量删除会员记录
		vipRecordRouter.PUT("updateVipRecord", vipRecordApi.UpdateVipRecord)              // 更新会员记录
	}
	{
		vipRecordRouterWithoutRecord.GET("findVipRecord", vipRecordApi.FindVipRecord)       // 根据ID获取会员记录
		vipRecordRouterWithoutRecord.GET("getVipRecordList", vipRecordApi.GetVipRecordList) // 获取会员记录列表
	}
	{
		vipRecordRouterWithoutAuth.GET("getVipRecordPublic", vipRecordApi.GetVipRecordPublic) // 会员记录开放接口
	}

	vipRecordRouterAuth := PublicRouter.Group("vipRecordMobileAuth")
	vipRecordRouterAuth.Use(middleware.MoblileJWTAuth()).Use(middleware.OperationRecord())
	{
		vipRecordRouterAuth.PUT("create/v1", vipRecordApi.CreateVipRecordMobile) // 创建vip
	}
}

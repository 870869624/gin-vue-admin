package describe

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DescribeRouter struct {}

// InitDescribeRouter 初始化 文档说明 路由信息
func (s *DescribeRouter) InitDescribeRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	DRouter := Router.Group("D").Use(middleware.OperationRecord())
	DRouterWithoutRecord := Router.Group("D")
	DRouterWithoutAuth := PublicRouter.Group("D")
	{
		DRouter.POST("createDescribe", DApi.CreateDescribe)   // 新建文档说明
		DRouter.DELETE("deleteDescribe", DApi.DeleteDescribe) // 删除文档说明
		DRouter.DELETE("deleteDescribeByIds", DApi.DeleteDescribeByIds) // 批量删除文档说明
		DRouter.PUT("updateDescribe", DApi.UpdateDescribe)    // 更新文档说明
	}
	{
		DRouterWithoutRecord.GET("findDescribe", DApi.FindDescribe)        // 根据ID获取文档说明
		DRouterWithoutRecord.GET("getDescribeList", DApi.GetDescribeList)  // 获取文档说明列表
	}
	{
	    DRouterWithoutAuth.GET("getDescribePublic", DApi.GetDescribePublic)  // 文档说明开放接口
	}
}

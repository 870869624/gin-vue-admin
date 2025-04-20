package Users

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type UsersRouter struct{}

// InitUsersRouter 初始化 用户列表 路由信息
func (s *UsersRouter) InitUsersRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	usersRouter := Router.Group("users").Use(middleware.OperationRecord())
	usersRouterWithoutRecord := Router.Group("users")
	usersRouterWithoutAuth := PublicRouter.Group("users")
	{
		usersRouter.POST("createUsers", usersApi.CreateUsers)             // 新建用户列表
		usersRouter.DELETE("deleteUsers", usersApi.DeleteUsers)           // 删除用户列表
		usersRouter.DELETE("deleteUsersByIds", usersApi.DeleteUsersByIds) // 批量删除用户列表
		usersRouter.PUT("updateUsers", usersApi.UpdateUsers)              // 更新用户列表
	}
	{
		usersRouterWithoutRecord.GET("findUsers", usersApi.FindUsers)       // 根据ID获取用户列表
		usersRouterWithoutRecord.GET("getUsersList", usersApi.GetUsersList) // 获取用户列表列表
	}
	{
		usersRouterWithoutAuth.GET("getUsersPublic", usersApi.GetUsersPublic) // 用户列表开放接口
	}

	usersRouterWithoutAuth.POST("login", usersApi.Login)
	usersRouterWithoutAuth.POST("import/v1", usersApi.Import)

	PrivateGroup := usersRouterWithoutAuth.Use(middleware.MoblileJWTAuth())
	{
		PrivateGroup.GET("balance/v1", usersApi.Balance)
	}

}

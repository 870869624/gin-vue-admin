package Users

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/Users"
    UsersReq "github.com/flipped-aurora/gin-vue-admin/server/model/Users/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type UsersApi struct {}



// CreateUsers 创建用户列表
// @Tags Users
// @Summary 创建用户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Users.Users true "创建用户列表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /users/createUsers [post]
func (usersApi *UsersApi) CreateUsers(c *gin.Context) {
	var users Users.Users
	err := c.ShouldBindJSON(&users)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = usersService.CreateUsers(&users)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteUsers 删除用户列表
// @Tags Users
// @Summary 删除用户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Users.Users true "删除用户列表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /users/deleteUsers [delete]
func (usersApi *UsersApi) DeleteUsers(c *gin.Context) {
	ID := c.Query("ID")
	err := usersService.DeleteUsers(ID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteUsersByIds 批量删除用户列表
// @Tags Users
// @Summary 批量删除用户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /users/deleteUsersByIds [delete]
func (usersApi *UsersApi) DeleteUsersByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := usersService.DeleteUsersByIds(IDs)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateUsers 更新用户列表
// @Tags Users
// @Summary 更新用户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Users.Users true "更新用户列表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /users/updateUsers [put]
func (usersApi *UsersApi) UpdateUsers(c *gin.Context) {
	var users Users.Users
	err := c.ShouldBindJSON(&users)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = usersService.UpdateUsers(users)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindUsers 用id查询用户列表
// @Tags Users
// @Summary 用id查询用户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query Users.Users true "用id查询用户列表"
// @Success 200 {object} response.Response{data=Users.Users,msg=string} "查询成功"
// @Router /users/findUsers [get]
func (usersApi *UsersApi) FindUsers(c *gin.Context) {
	ID := c.Query("ID")
	reusers, err := usersService.GetUsers(ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(reusers, c)
}

// GetUsersList 分页获取用户列表列表
// @Tags Users
// @Summary 分页获取用户列表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query UsersReq.UsersSearch true "分页获取用户列表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /users/getUsersList [get]
func (usersApi *UsersApi) GetUsersList(c *gin.Context) {
	var pageInfo UsersReq.UsersSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := usersService.GetUsersInfoList(pageInfo)
	if err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败:" + err.Error(), c)
        return
    }
    response.OkWithDetailed(response.PageResult{
        List:     list,
        Total:    total,
        Page:     pageInfo.Page,
        PageSize: pageInfo.PageSize,
    }, "获取成功", c)
}

// GetUsersPublic 不需要鉴权的用户列表接口
// @Tags Users
// @Summary 不需要鉴权的用户列表接口
// @accept application/json
// @Produce application/json
// @Param data query UsersReq.UsersSearch true "分页获取用户列表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /users/getUsersPublic [get]
func (usersApi *UsersApi) GetUsersPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    usersService.GetUsersPublic()
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的用户列表接口信息",
    }, "获取成功", c)
}

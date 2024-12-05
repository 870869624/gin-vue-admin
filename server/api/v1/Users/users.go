package Users

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Users"
	UsersReq "github.com/flipped-aurora/gin-vue-admin/server/model/Users/request"
	usersRes "github.com/flipped-aurora/gin-vue-admin/server/model/Users/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

type UsersApi struct{}

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
		response.FailWithMessage("创建失败:"+err.Error(), c)
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
		response.FailWithMessage("删除失败:"+err.Error(), c)
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
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
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
		response.FailWithMessage("更新失败:"+err.Error(), c)
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
		response.FailWithMessage("查询失败:"+err.Error(), c)
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
		response.FailWithMessage("获取失败:"+err.Error(), c)
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

// Login
// @Tags     Base
// @Summary  用户登录
// @Produce   application/json
// @Param    data  body      systemReq.Login                                             true  "用户名, 密码, 验证码"
// @Success  200   {object}  response.Response{data=systemRes.LoginResponse,msg=string}  "返回包括用户信息,token,过期时间"
// @Router   /base/login [post]
func (usersApi *UsersApi) Login(c *gin.Context) {
	var l UsersReq.Login
	err := c.ShouldBindJSON(&l)
	key := c.ClientIP()

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(l, utils.LoginVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 判断验证码是否开启
	openCaptcha := global.GVA_CONFIG.Captcha.OpenCaptcha               // 是否开启防爆次数
	openCaptchaTimeOut := global.GVA_CONFIG.Captcha.OpenCaptchaTimeOut // 缓存超时时间
	v, ok := global.BlackCache.Get(key)
	if !ok {
		global.BlackCache.Set(key, 1, time.Second*time.Duration(openCaptchaTimeOut))
	}

	var oc bool = openCaptcha == 0 || openCaptcha < interfaceToInt(v)

	if !oc || (l.CaptchaId != "" && l.Captcha != "" && store.Verify(l.CaptchaId, l.Captcha, true)) {
		u := &Users.Users{Username: &l.Username, Password: &l.Password}
		user, err := usersService.Login(u)
		if err != nil {
			global.GVA_LOG.Error("登陆失败! 用户名不存在或者密码错误!", zap.Error(err))
			// 验证码次数+1
			global.BlackCache.Increment(key, 1)
			response.FailWithMessage("用户名不存在或者密码错误", c)
			return
		}
		usersApi.TokenNext(c, *user)
		return
	}
	// 验证码次数+1
	global.BlackCache.Increment(key, 1)
	response.FailWithMessage("验证码错误", c)
}

// TokenNext 登录以后签发jwt
func (usersApi *UsersApi) TokenNext(c *gin.Context, user Users.Users) {
	token, claims, err := utils.MobileLoginToken(&user)
	if err != nil {
		global.GVA_LOG.Error("获取token失败!", zap.Error(err))
		response.FailWithMessage("获取token失败", c)
		return
	}
	if !global.GVA_CONFIG.System.UseMultipoint {
		utils.SetToken(c, token, int(claims.RegisteredClaims.ExpiresAt.Unix()-time.Now().Unix()))
		response.OkWithDetailed(usersRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
		}, "登录成功", c)
		return
	}

	if jwtStr, err := jwtService.GetRedisJWT(*user.Username); err == redis.Nil {
		if err := jwtService.SetRedisJWT(token, *user.Username); err != nil {
			global.GVA_LOG.Error("设置登录状态失败!", zap.Error(err))
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		utils.SetToken(c, token, int(claims.RegisteredClaims.ExpiresAt.Unix()-time.Now().Unix()))
		response.OkWithDetailed(usersRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
		}, "登录成功", c)
	} else if err != nil {
		global.GVA_LOG.Error("设置登录状态失败!", zap.Error(err))
		response.FailWithMessage("设置登录状态失败", c)
	} else {
		var blackJWT system.JwtBlacklist
		blackJWT.Jwt = jwtStr
		if err := jwtService.JsonInBlacklist(blackJWT); err != nil {
			response.FailWithMessage("jwt作废失败", c)
			return
		}
		if err := jwtService.SetRedisJWT(token, user.GetUsername()); err != nil {
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		utils.SetToken(c, token, int(claims.RegisteredClaims.ExpiresAt.Unix()-time.Now().Unix()))
		response.OkWithDetailed(usersRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
		}, "登录成功", c)
	}
}

// Register
// @Tags     SysUser
// @Summary  用户注册账号
// @Produce   application/json
// @Param    data  body      systemReq.Register                                            true  "用户名, 昵称, 密码, 角色ID"
// @Success  200   {object}  response.Response{data=systemRes.SysUserResponse,msg=string}  "用户注册账号,返回包括用户信息"
// @Router   /user/admin_register [post]
func (usersApi *UsersApi) Register(c *gin.Context) {
	var r UsersReq.Register
	err := c.ShouldBindJSON(&r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(r, utils.RegisterVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	user := &Users.Users{Username: &r.Username, Password: &r.Password, Avatar: r.Avatar, Gender: &r.Gender, PhoneNumber: &r.PhoneNumber, Email: &r.Email, MetaMask: &r.MetaMask, Brief: &r.Brief, TokenPocket: &r.TokenPocket}
	userReturn, err := usersService.Register(user)
	if err != nil {
		global.GVA_LOG.Error("注册失败!", zap.Error(err))
		response.FailWithDetailed(usersRes.UsersResponse{User: *userReturn}, "注册失败", c)
		return
	}
	response.OkWithDetailed(usersRes.UsersResponse{User: *userReturn}, "注册成功", c)
}

// ChangePassword
// @Tags      SysUser
// @Summary   用户修改密码
// @Security  ApiKeyAuth
// @Produce  application/json
// @Param     data  body      systemReq.ChangePasswordReq    true  "用户名, 原密码, 新密码"
// @Success   200   {object}  response.Response{msg=string}  "用户修改密码"
// @Router    /user/changePassword [post]
func (usersApi *UsersApi) ChangePassword(c *gin.Context) {
	var req UsersReq.ChangePasswordReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(req, utils.ChangePasswordVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	uid := utils.GetMobileUserID(c)
	u := &Users.Users{GVA_MODEL: global.GVA_MODEL{ID: uid}, Password: &req.Password}
	_, err = usersService.ChangePassword(u, req.NewPassword)
	if err != nil {
		global.GVA_LOG.Error("修改失败!", zap.Error(err))
		response.FailWithMessage("修改失败，原密码与当前账户不符", c)
		return
	}
	response.OkWithMessage("修改成功", c)
}

// // DeleteUser
// // @Tags      SysUser
// // @Summary   删除用户
// // @Security  ApiKeyAuth
// // @accept    application/json
// // @Produce   application/json
// // @Param     data  body      request.GetById                true  "用户ID"
// // @Success   200   {object}  response.Response{msg=string}  "删除用户"
// // @Router    /user/deleteUser [delete]
// func (usersApi *UsersApi) DeleteUser(c *gin.Context) {
// 	var reqId request.GetById
// 	err := c.ShouldBindJSON(&reqId)
// 	if err != nil {
// 		response.FailWithMessage(err.Error(), c)
// 		return
// 	}
// 	err = utils.Verify(reqId, utils.IdVerify)
// 	if err != nil {
// 		response.FailWithMessage(err.Error(), c)
// 		return
// 	}
// 	jwtId := utils.GetUserID(c)
// 	if jwtId == uint(reqId.ID) {
// 		response.FailWithMessage("删除失败, 无法删除自己。", c)
// 		return
// 	}
// 	err = userService.DeleteUser(reqId.ID)
// 	if err != nil {
// 		global.GVA_LOG.Error("删除失败!", zap.Error(err))
// 		response.FailWithMessage("删除失败", c)
// 		return
// 	}
// 	response.OkWithMessage("删除成功", c)
// }

// // SetSelfInfo
// // @Tags      SysUser
// // @Summary   设置用户信息
// // @Security  ApiKeyAuth
// // @accept    application/json
// // @Produce   application/json
// // @Param     data  body      system.SysUser                                             true  "ID, 用户名, 昵称, 头像链接"
// // @Success   200   {object}  response.Response{data=map[string]interface{},msg=string}  "设置用户信息"
// // @Router    /user/SetSelfInfo [put]
// func (usersApi *UsersApi) SetSelfInfo(c *gin.Context) {
// 	var user systemReq.ChangeUserInfo
// 	err := c.ShouldBindJSON(&user)
// 	if err != nil {
// 		response.FailWithMessage(err.Error(), c)
// 		return
// 	}
// 	user.ID = utils.GetUserID(c)
// 	err = userService.SetSelfInfo(system.SysUser{
// 		GVA_MODEL: global.GVA_MODEL{
// 			ID: user.ID,
// 		},
// 		NickName:  user.NickName,
// 		HeaderImg: user.HeaderImg,
// 		Phone:     user.Phone,
// 		Email:     user.Email,
// 		Enable:    user.Enable,
// 	})
// 	if err != nil {
// 		global.GVA_LOG.Error("设置失败!", zap.Error(err))
// 		response.FailWithMessage("设置失败", c)
// 		return
// 	}
// 	response.OkWithMessage("设置成功", c)
// }

// // ResetPassword
// // @Tags      SysUser
// // @Summary   重置用户密码
// // @Security  ApiKeyAuth
// // @Produce  application/json
// // @Param     data  body      system.SysUser                 true  "ID"
// // @Success   200   {object}  response.Response{msg=string}  "重置用户密码"
// // @Router    /user/resetPassword [post]
// func (usersApi *UsersApi) ResetPassword(c *gin.Context) {
// 	var user system.SysUser
// 	err := c.ShouldBindJSON(&user)
// 	if err != nil {
// 		response.FailWithMessage(err.Error(), c)
// 		return
// 	}
// 	err = userService.ResetPassword(user.ID)
// 	if err != nil {
// 		global.GVA_LOG.Error("重置失败!", zap.Error(err))
// 		response.FailWithMessage("重置失败"+err.Error(), c)
// 		return
// 	}
// 	response.OkWithMessage("重置成功", c)
// }

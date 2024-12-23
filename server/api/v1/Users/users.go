package Users

import (
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Users"
	UsersReq "github.com/flipped-aurora/gin-vue-admin/server/model/Users/request"
	usersRes "github.com/flipped-aurora/gin-vue-admin/server/model/Users/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/contract"
	"github.com/gin-gonic/gin"
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

	// 判断验证码是否开启
	openCaptcha := global.GVA_CONFIG.Captcha.OpenCaptcha               // 是否开启防爆次数
	openCaptchaTimeOut := global.GVA_CONFIG.Captcha.OpenCaptchaTimeOut // 缓存超时时间
	v, ok := global.BlackCache.Get(key)
	if !ok {
		global.BlackCache.Set(key, 1, time.Second*time.Duration(openCaptchaTimeOut))
	}

	var oc bool = openCaptcha == 0 || openCaptcha < interfaceToInt(v)

	if !oc || (l.CaptchaId != "" && l.Captcha != "" && store.Verify(l.CaptchaId, l.Captcha, true)) {
		var u Users.Users
		if l.MetaMask != "" {
			u = Users.Users{MetaMask: &l.MetaMask, MetaMaskMoney: &l.MetaMaskMoney, Username: &l.Username}
		} else if l.TokenPocket != "" {
			u = Users.Users{TokenPocket: &l.TokenPocket, TokenPocketMoney: &l.TokenPocketMoney, Username: &l.Username}
		}
		user, err := usersService.Login(&u)
		if err != nil {
			global.GVA_LOG.Error("登陆失败! 用户名不存在或者密码错误!", zap.Error(err))
			// 验证码次数+1
			global.BlackCache.Increment(key, 1)
			response.FailWithMessage("用户名不存在或者密码错误"+err.Error(), c)
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
	token, claims, err := utils.MobileLoginToken(user)
	if err != nil {
		global.GVA_LOG.Error("获取token失败!", zap.Error(err))
		response.FailWithMessage("获取token失败", c)
		return
	}
	response.OkWithDetailed(usersRes.LoginResponse{
		User:      user,
		Token:     token,
		ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
	}, "登录成功", c)
}

func (usersApi *UsersApi) Balance(c *gin.Context) {
	utils.ClearToken(c)
	address := c.Query("address")
	uid := utils.GetMobileUserID(c)
	if uid == 0 {
		response.FailWithMessage("查询失败:用户ID为空", c)
		return
	}
	userId := int(uid)

	user := &Users.Users{
		GVA_MODEL: global.GVA_MODEL{ID: uint(userId)},
	}
	if err := global.GVA_DB.Debug().Table(user.TableName()).Where(user).First(user).Error; err != nil {
		response.FailWithMessage("用户信息错误"+err.Error(), c)
		return
	}

	if *user.MetaMask != address && *user.TokenPocket != address {
		fmt.Println(user.MetaMask, address)
		response.FailWithMessage("用户信息错误", c)
		return
	}

	fromAddress, instance, _, err := contract.GetServe(address)
	if err != nil {
		response.FailWithMessage("获取网络错误"+err.Error(), c)
		return
	}

	opts := &bind.CallOpts{
		From: *fromAddress,
	}
	balance, err := instance.BalanceOf(opts, *fromAddress)
	if err != nil {
		response.FailWithMessage("获取金额错误"+err.Error(), c)
		return
	}

	if *user.MetaMask != address {
		if err := global.GVA_DB.Debug().Table(user.TableName()).Where(user).Update("meta_mask_money", balance).Error; err != nil {
			response.FailWithMessage("用户信息错误"+err.Error(), c)
			return
		}
	} else if *user.TokenPocket != address {
		if err := global.GVA_DB.Debug().Table(user.TableName()).Where(user).Update("token_pocket_money", balance).Error; err != nil {
			response.FailWithMessage("用户信息错误"+err.Error(), c)
			return
		}
	}
	response.OkWithDetailed(gin.H{
		"balance": balance,
	}, "获取成功", c)
}

func ChangeNumber(balance string) string {
	number, _ := new(big.Int).SetString(balance, 10)

	// 定义除数
	divisor := new(big.Int).SetInt64(1000000000000000000)

	// 转换为big.Float以保持精度
	numberFloat := new(big.Float).SetInt(number)
	divisorFloat := new(big.Float).SetInt(divisor)

	// 执行除法
	resultFloat := new(big.Float).Quo(numberFloat, divisorFloat)

	// 格式化输出，保留一位小数
	// big.Float的Text方法可以按照指定的格式输出，这里是'f'表示定点格式，0表示没有小数位
	resultStr := resultFloat.Text('f', 1)
	return resultStr
}

type req struct {
	Address string `json:"address" form:"address"`
}

func (usersApi *UsersApi) Import(c *gin.Context) {
	var req req
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	fromAddress, err := contract.GetAddress(req.Address)
	if err != nil {
		response.FailWithMessage("获取网络错误"+err.Error(), c)
		return
	}

	response.OkWithDetailed(gin.H{
		"address": fromAddress,
	}, "获取成功", c)
}

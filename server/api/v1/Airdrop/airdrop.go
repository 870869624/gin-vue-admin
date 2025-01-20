package Airdrop

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Airdrop"
	AirdropReq "github.com/flipped-aurora/gin-vue-admin/server/model/Airdrop/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AirdropApi struct{}

// CreateAirdrop 创建空投项目
// @Tags Airdrop
// @Summary 创建空投项目
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Airdrop.Airdrop true "创建空投项目"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /airdrop/createAirdrop [post]
func (airdropApi *AirdropApi) CreateAirdrop(c *gin.Context) {
	var airdrop Airdrop.Airdrop
	err := c.ShouldBindJSON(&airdrop)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = airdropService.CreateAirdrop(&airdrop)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteAirdrop 删除空投项目
// @Tags Airdrop
// @Summary 删除空投项目
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Airdrop.Airdrop true "删除空投项目"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /airdrop/deleteAirdrop [delete]
func (airdropApi *AirdropApi) DeleteAirdrop(c *gin.Context) {
	ID := c.Query("ID")
	err := airdropService.DeleteAirdrop(ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteAirdropByIds 批量删除空投项目
// @Tags Airdrop
// @Summary 批量删除空投项目
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /airdrop/deleteAirdropByIds [delete]
func (airdropApi *AirdropApi) DeleteAirdropByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := airdropService.DeleteAirdropByIds(IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateAirdrop 更新空投项目
// @Tags Airdrop
// @Summary 更新空投项目
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Airdrop.Airdrop true "更新空投项目"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /airdrop/updateAirdrop [put]
func (airdropApi *AirdropApi) UpdateAirdrop(c *gin.Context) {
	var airdrop Airdrop.Airdrop
	err := c.ShouldBindJSON(&airdrop)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = airdropService.UpdateAirdrop(airdrop)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindAirdrop 用id查询空投项目
// @Tags Airdrop
// @Summary 用id查询空投项目
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query Airdrop.Airdrop true "用id查询空投项目"
// @Success 200 {object} response.Response{data=Airdrop.Airdrop,msg=string} "查询成功"
// @Router /airdrop/findAirdrop [get]
func (airdropApi *AirdropApi) FindAirdrop(c *gin.Context) {
	ID := c.Query("ID")
	reairdrop, err := airdropService.GetAirdrop(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reairdrop, c)
}

// GetAirdropList 分页获取空投项目列表
// @Tags Airdrop
// @Summary 分页获取空投项目列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query AirdropReq.AirdropSearch true "分页获取空投项目列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /airdrop/getAirdropList [get]
func (airdropApi *AirdropApi) GetAirdropList(c *gin.Context) {
	var pageInfo AirdropReq.AirdropSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := airdropService.GetAirdropInfoList(pageInfo)
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

// GetAirdropPublic 不需要鉴权的空投项目接口
// @Tags Airdrop
// @Summary 不需要鉴权的空投项目接口
// @accept application/json
// @Produce application/json
// @Param data query AirdropReq.AirdropSearch true "分页获取空投项目列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /airdrop/getAirdropPublic [get]
func (airdropApi *AirdropApi) GetAirdropPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	airdropService.GetAirdropPublic()
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的空投项目接口信息",
	}, "获取成功", c)
}

func (airdropApi *AirdropApi) MobileCreateAirdrop(c *gin.Context) {
	var isShow = false
	var isPass = false
	var airdrop Airdrop.MobileAirdropCre
	err := c.ShouldBindJSON(&airdrop)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	uid := utils.GetMobileUserID(c)
	if uid == 0 {
		response.FailWithMessage("查询失败:用户ID为空", c)
		return
	}
	userId := int(uid)
	airdrop.UserId = &userId
	airdrop.AirdropIsShow = &isShow
	airdrop.AirdropIsPass = &isPass
	err = airdropService.CreateAirdropMobile(&airdrop)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

func (airdropApi *AirdropApi) MobileGetAirdropList(c *gin.Context) {
	var pageInfo AirdropReq.AirdropSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := airdropService.MobileGetAirdropInfoList(pageInfo)
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

func (airdropApi *AirdropApi) FindAirdropMobile(c *gin.Context) {
	ID := c.Query("ID")
	reairdrop, err := airdropService.GetAirdropMobile(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reairdrop, c)
}

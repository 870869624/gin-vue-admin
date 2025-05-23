package Vip

import (
	"strconv"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Vip"
	VipReq "github.com/flipped-aurora/gin-vue-admin/server/model/Vip/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type VipRecordApi struct{}

// CreateVipRecord 创建会员记录
// @Tags VipRecord
// @Summary 创建会员记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Vip.VipRecord true "创建会员记录"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /vipRecord/createVipRecord [post]
func (vipRecordApi *VipRecordApi) CreateVipRecord(c *gin.Context) {
	var vipRecord Vip.VipRecord
	err := c.ShouldBindJSON(&vipRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = vipRecordService.CreateVipRecord(&vipRecord)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteVipRecord 删除会员记录
// @Tags VipRecord
// @Summary 删除会员记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Vip.VipRecord true "删除会员记录"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /vipRecord/deleteVipRecord [delete]
func (vipRecordApi *VipRecordApi) DeleteVipRecord(c *gin.Context) {
	ID := c.Query("ID")
	err := vipRecordService.DeleteVipRecord(ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteVipRecordByIds 批量删除会员记录
// @Tags VipRecord
// @Summary 批量删除会员记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /vipRecord/deleteVipRecordByIds [delete]
func (vipRecordApi *VipRecordApi) DeleteVipRecordByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := vipRecordService.DeleteVipRecordByIds(IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateVipRecord 更新会员记录
// @Tags VipRecord
// @Summary 更新会员记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Vip.VipRecord true "更新会员记录"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /vipRecord/updateVipRecord [put]
func (vipRecordApi *VipRecordApi) UpdateVipRecord(c *gin.Context) {
	var vipRecord Vip.VipRecord
	err := c.ShouldBindJSON(&vipRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = vipRecordService.UpdateVipRecord(vipRecord)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindVipRecord 用id查询会员记录
// @Tags VipRecord
// @Summary 用id查询会员记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query Vip.VipRecord true "用id查询会员记录"
// @Success 200 {object} response.Response{data=Vip.VipRecord,msg=string} "查询成功"
// @Router /vipRecord/findVipRecord [get]
func (vipRecordApi *VipRecordApi) FindVipRecord(c *gin.Context) {
	ID := c.Query("ID")
	revipRecord, err := vipRecordService.GetVipRecord(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(revipRecord, c)
}

// GetVipRecordList 分页获取会员记录列表
// @Tags VipRecord
// @Summary 分页获取会员记录列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query VipReq.VipRecordSearch true "分页获取会员记录列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /vipRecord/getVipRecordList [get]
func (vipRecordApi *VipRecordApi) GetVipRecordList(c *gin.Context) {
	var pageInfo VipReq.VipRecordSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := vipRecordService.GetVipRecordInfoList(pageInfo)
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

// GetVipRecordPublic 不需要鉴权的会员记录接口
// @Tags VipRecord
// @Summary 不需要鉴权的会员记录接口
// @accept application/json
// @Produce application/json
// @Param data query VipReq.VipRecordSearch true "分页获取会员记录列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /vipRecord/getVipRecordPublic [get]
func (vipRecordApi *VipRecordApi) GetVipRecordPublic(c *gin.Context) {
	bab := c.Query("bab")

	uid := utils.GetMobileUserID(c)
	if uid == 0 {
		response.FailWithMessage("查询失败:用户ID为空", c)
		return
	}

	// userId := int(uid)

	_, err := vipRecordService.GetVipRecordBabMobile(bab)
	if err == nil {
		response.FailWithMessage("bab已使用", c)
		return
	}

	response.OkWithMessage("bab未使用", c)
}

func (vipRecordApi *VipRecordApi) CreateVipRecordMobile(c *gin.Context) {
	var vipRecord Vip.VipRecordMobile
	if err := c.ShouldBindJSON(&vipRecord); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var IsEffective = true
	uid := utils.GetMobileUserID(c)
	if uid == 0 {

		response.FailWithMessage("查询失败:用户ID为空", c)
		return
	}
	userId := int(uid)
	vipRecord.UserId = &userId
	vipRecord.IsEffective = &IsEffective

	_, err := vipRecordService.GetVipRecordMobile(strconv.Itoa(userId))
	if err == nil {
		response.OkWithMessage("会员记录已存在", c)
		return
	}

	err = vipRecordService.CreateVipRecordMobile(&vipRecord)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

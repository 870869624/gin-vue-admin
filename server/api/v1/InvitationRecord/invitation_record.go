package InvitationRecord

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/InvitationRecord"
    InvitationRecordReq "github.com/flipped-aurora/gin-vue-admin/server/model/InvitationRecord/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type InvitationRecordApi struct {}



// CreateInvitationRecord 创建邀请记录
// @Tags InvitationRecord
// @Summary 创建邀请记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body InvitationRecord.InvitationRecord true "创建邀请记录"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /IR/createInvitationRecord [post]
func (IRApi *InvitationRecordApi) CreateInvitationRecord(c *gin.Context) {
	var IR InvitationRecord.InvitationRecord
	err := c.ShouldBindJSON(&IR)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = IRService.CreateInvitationRecord(&IR)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteInvitationRecord 删除邀请记录
// @Tags InvitationRecord
// @Summary 删除邀请记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body InvitationRecord.InvitationRecord true "删除邀请记录"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /IR/deleteInvitationRecord [delete]
func (IRApi *InvitationRecordApi) DeleteInvitationRecord(c *gin.Context) {
	ID := c.Query("ID")
	err := IRService.DeleteInvitationRecord(ID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteInvitationRecordByIds 批量删除邀请记录
// @Tags InvitationRecord
// @Summary 批量删除邀请记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /IR/deleteInvitationRecordByIds [delete]
func (IRApi *InvitationRecordApi) DeleteInvitationRecordByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := IRService.DeleteInvitationRecordByIds(IDs)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateInvitationRecord 更新邀请记录
// @Tags InvitationRecord
// @Summary 更新邀请记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body InvitationRecord.InvitationRecord true "更新邀请记录"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /IR/updateInvitationRecord [put]
func (IRApi *InvitationRecordApi) UpdateInvitationRecord(c *gin.Context) {
	var IR InvitationRecord.InvitationRecord
	err := c.ShouldBindJSON(&IR)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = IRService.UpdateInvitationRecord(IR)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindInvitationRecord 用id查询邀请记录
// @Tags InvitationRecord
// @Summary 用id查询邀请记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query InvitationRecord.InvitationRecord true "用id查询邀请记录"
// @Success 200 {object} response.Response{data=InvitationRecord.InvitationRecord,msg=string} "查询成功"
// @Router /IR/findInvitationRecord [get]
func (IRApi *InvitationRecordApi) FindInvitationRecord(c *gin.Context) {
	ID := c.Query("ID")
	reIR, err := IRService.GetInvitationRecord(ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(reIR, c)
}

// GetInvitationRecordList 分页获取邀请记录列表
// @Tags InvitationRecord
// @Summary 分页获取邀请记录列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query InvitationRecordReq.InvitationRecordSearch true "分页获取邀请记录列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /IR/getInvitationRecordList [get]
func (IRApi *InvitationRecordApi) GetInvitationRecordList(c *gin.Context) {
	var pageInfo InvitationRecordReq.InvitationRecordSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := IRService.GetInvitationRecordInfoList(pageInfo)
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

// GetInvitationRecordPublic 不需要鉴权的邀请记录接口
// @Tags InvitationRecord
// @Summary 不需要鉴权的邀请记录接口
// @accept application/json
// @Produce application/json
// @Param data query InvitationRecordReq.InvitationRecordSearch true "分页获取邀请记录列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /IR/getInvitationRecordPublic [get]
func (IRApi *InvitationRecordApi) GetInvitationRecordPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    IRService.GetInvitationRecordPublic()
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的邀请记录接口信息",
    }, "获取成功", c)
}

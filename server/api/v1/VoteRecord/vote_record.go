package VoteRecord

import (
	"strconv"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/VoteRecord"
	VoteRecordReq "github.com/flipped-aurora/gin-vue-admin/server/model/VoteRecord/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type VoteRecordApi struct{}

// CreateVoteRecord 创建投票记录
// @Tags VoteRecord
// @Summary 创建投票记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body VoteRecord.VoteRecord true "创建投票记录"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /voteRecord/createVoteRecord [post]
func (voteRecordApi *VoteRecordApi) CreateVoteRecord(c *gin.Context) {
	var voteRecord VoteRecord.VoteRecord
	err := c.ShouldBindJSON(&voteRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = voteRecordService.CreateVoteRecord(&voteRecord)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteVoteRecord 删除投票记录
// @Tags VoteRecord
// @Summary 删除投票记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body VoteRecord.VoteRecord true "删除投票记录"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /voteRecord/deleteVoteRecord [delete]
func (voteRecordApi *VoteRecordApi) DeleteVoteRecord(c *gin.Context) {
	ID := c.Query("ID")
	err := voteRecordService.DeleteVoteRecord(ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteVoteRecordByIds 批量删除投票记录
// @Tags VoteRecord
// @Summary 批量删除投票记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /voteRecord/deleteVoteRecordByIds [delete]
func (voteRecordApi *VoteRecordApi) DeleteVoteRecordByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := voteRecordService.DeleteVoteRecordByIds(IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateVoteRecord 更新投票记录
// @Tags VoteRecord
// @Summary 更新投票记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body VoteRecord.VoteRecord true "更新投票记录"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /voteRecord/updateVoteRecord [put]
func (voteRecordApi *VoteRecordApi) UpdateVoteRecord(c *gin.Context) {
	var voteRecord VoteRecord.VoteRecord
	err := c.ShouldBindJSON(&voteRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = voteRecordService.UpdateVoteRecord(voteRecord)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindVoteRecord 用id查询投票记录
// @Tags VoteRecord
// @Summary 用id查询投票记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query VoteRecord.VoteRecord true "用id查询投票记录"
// @Success 200 {object} response.Response{data=VoteRecord.VoteRecord,msg=string} "查询成功"
// @Router /voteRecord/findVoteRecord [get]
func (voteRecordApi *VoteRecordApi) FindVoteRecord(c *gin.Context) {
	ID := c.Query("ID")
	revoteRecord, err := voteRecordService.MobileGetVoteRecord(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(revoteRecord, c)
}

// GetVoteRecordList 分页获取投票记录列表
// @Tags VoteRecord
// @Summary 分页获取投票记录列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query VoteRecordReq.VoteRecordSearch true "分页获取投票记录列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /voteRecord/getVoteRecordList [get]
func (voteRecordApi *VoteRecordApi) GetVoteRecordList(c *gin.Context) {
	var pageInfo VoteRecordReq.VoteRecordSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := voteRecordService.GetVoteRecordInfoList(pageInfo)
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

// GetVoteRecordPublic 不需要鉴权的投票记录接口
// @Tags VoteRecord
// @Summary 不需要鉴权的投票记录接口
// @accept application/json
// @Produce application/json
// @Param data query VoteRecordReq.VoteRecordSearch true "分页获取投票记录列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /voteRecord/getVoteRecordPublic [get]
func (voteRecordApi *VoteRecordApi) GetVoteRecordPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	voteRecordService.GetVoteRecordPublic()
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的投票记录接口信息",
	}, "获取成功", c)
}

func (voteRecordApi *VoteRecordApi) GetVoteRecordNumber(c *gin.Context) {
	uid := utils.GetMobileUserID(c)
	userId := int(uid)
	revoteRecord, err := voteRecordService.MobileCountMyVoteRecord(strconv.Itoa(userId))
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(revoteRecord, c)
}

func (voteRecordApi *VoteRecordApi) GetAllVoteRecordNumber(c *gin.Context) {
	var revoteRecord1 int64
	var err error
	uid := utils.GetMobileUserID(c)
	if uid != 0 {
		userId := int(uid)
		revoteRecord1, err = voteRecordService.MobileCountMyVoteRecord(strconv.Itoa(userId))
		if err != nil {
			global.GVA_LOG.Error("查询失败!", zap.Error(err))
			response.FailWithMessage("查询失败:"+err.Error(), c)
			return
		}
	} else {
		response.FailWithMessage("查询失败:用户ID为空", c)
		return
	}

	ID := c.Query("ID")
	revoteRecord, err := voteRecordService.MobileCountAllVoteRecord(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	revoteRecord.MyVoteRecord = revoteRecord1
	response.OkWithData(revoteRecord, c)
}

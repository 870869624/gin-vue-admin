package Vote

import (
	"strconv"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Vote"
	VoteReq "github.com/flipped-aurora/gin-vue-admin/server/model/Vote/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/VoteRecord"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type VoteApi struct{}

// CreateVote 创建投票
// @Tags Vote
// @Summary 创建投票
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Vote.Vote true "创建投票"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /vote/createVote [post]
func (voteApi *VoteApi) CreateVote(c *gin.Context) {
	var vote Vote.Vote
	err := c.ShouldBindJSON(&vote)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = voteService.CreateVote(&vote)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteVote 删除投票
// @Tags Vote
// @Summary 删除投票
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Vote.Vote true "删除投票"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /vote/deleteVote [delete]
func (voteApi *VoteApi) DeleteVote(c *gin.Context) {
	ID := c.Query("ID")
	err := voteService.DeleteVote(ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteVoteByIds 批量删除投票
// @Tags Vote
// @Summary 批量删除投票
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /vote/deleteVoteByIds [delete]
func (voteApi *VoteApi) DeleteVoteByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := voteService.DeleteVoteByIds(IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateVote 更新投票
// @Tags Vote
// @Summary 更新投票
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Vote.Vote true "更新投票"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /vote/updateVote [put]
func (voteApi *VoteApi) UpdateVote(c *gin.Context) {
	var vote Vote.Vote
	err := c.ShouldBindJSON(&vote)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = voteService.UpdateVote(vote)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindVote 用id查询投票
// @Tags Vote
// @Summary 用id查询投票
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query Vote.Vote true "用id查询投票"
// @Success 200 {object} response.Response{data=Vote.Vote,msg=string} "查询成功"
// @Router /vote/findVote [get]
func (voteApi *VoteApi) FindVote(c *gin.Context) {
	ID := c.Query("ID")
	revote, err := voteService.GetVote(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(revote, c)
}

// GetVoteList 分页获取投票列表
// @Tags Vote
// @Summary 分页获取投票列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query VoteReq.VoteSearch true "分页获取投票列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /vote/getVoteList [get]
func (voteApi *VoteApi) GetVoteList(c *gin.Context) {
	var pageInfo VoteReq.VoteSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := voteService.GetVoteInfoList(pageInfo)
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

// GetVotePublic 不需要鉴权的投票接口
// @Tags Vote
// @Summary 不需要鉴权的投票接口
// @accept application/json
// @Produce application/json
// @Param data query VoteReq.VoteSearch true "分页获取投票列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /vote/getVotePublic [get]
func (voteApi *VoteApi) GetVotePublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	voteService.GetVotePublic()
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的投票接口信息",
	}, "获取成功", c)
}

func (voteApi *VoteApi) MobileCreateVote(c *gin.Context) {
	var isShow = false
	var isPass = false
	var voteNum = 0
	var vote Vote.Vote
	err := c.ShouldBindJSON(&vote)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	uid := utils.GetMobileUserID(c)
	if uid == 0 {
		response.FailWithMessage("查询失败:用户ID为空", c)
		return
	}
	vote.VoteIsShow = &isShow
	vote.VoteIsPass = &isPass
	vote.VoteNum = &voteNum
	err = voteService.CreateVote(&vote)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

type Req struct {
	ID string `json:"ID"`
}

func (voteApi *VoteApi) Vote(c *gin.Context) {
	var req Req
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	revote, err := voteService.GetVote(req.ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}

	uid := utils.GetMobileUserID(c)
	if uid == 0 {
		response.FailWithMessage("查询失败:用户ID为空", c)
		return
	}
	userid := int(uid)
	revoteId := int(revote.ID)

	id := strconv.Itoa(int(uid))

	revoteRecord, err := voteRecordService.MobileCountMyVoteRecord(strconv.Itoa(userid))
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}

	if revoteRecord > 0 {
		response.FailWithMessage("每天只能投票一次", c)
		return
	}

	_, err = usersService.GetUsers(id)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}

	*revote.VoteNum++
	err = voteService.SaveVote(&revote)
	if err != nil {
		global.GVA_LOG.Error("投票失败!", zap.Error(err))
		response.FailWithMessage("投票失败:"+err.Error(), c)
		return
	}

	var voteRecord VoteRecord.VoteRecord
	voteRecord.UserId = &userid
	voteRecord.VoteId = &revoteId
	err = voteRecordService.CreateVoteRecord(&voteRecord)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("投票成功", c)
}

func (voteApi *VoteApi) MobileGetVoteList(c *gin.Context) {
	var pageInfo VoteReq.VoteSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := voteService.MobileGetVoteInfoList(pageInfo)
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

func (voteApi *VoteApi) FindVoteMobile(c *gin.Context) {
	ID := c.Query("ID")
	revote, err := voteService.GetVoteMobile(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(revote, c)
}

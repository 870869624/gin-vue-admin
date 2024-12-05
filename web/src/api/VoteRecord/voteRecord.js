import service from '@/utils/request'
// @Tags VoteRecord
// @Summary 创建投票记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.VoteRecord true "创建投票记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /voteRecord/createVoteRecord [post]
export const createVoteRecord = (data) => {
  return service({
    url: '/voteRecord/createVoteRecord',
    method: 'post',
    data
  })
}

// @Tags VoteRecord
// @Summary 删除投票记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.VoteRecord true "删除投票记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /voteRecord/deleteVoteRecord [delete]
export const deleteVoteRecord = (params) => {
  return service({
    url: '/voteRecord/deleteVoteRecord',
    method: 'delete',
    params
  })
}

// @Tags VoteRecord
// @Summary 批量删除投票记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除投票记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /voteRecord/deleteVoteRecord [delete]
export const deleteVoteRecordByIds = (params) => {
  return service({
    url: '/voteRecord/deleteVoteRecordByIds',
    method: 'delete',
    params
  })
}

// @Tags VoteRecord
// @Summary 更新投票记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.VoteRecord true "更新投票记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /voteRecord/updateVoteRecord [put]
export const updateVoteRecord = (data) => {
  return service({
    url: '/voteRecord/updateVoteRecord',
    method: 'put',
    data
  })
}

// @Tags VoteRecord
// @Summary 用id查询投票记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.VoteRecord true "用id查询投票记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /voteRecord/findVoteRecord [get]
export const findVoteRecord = (params) => {
  return service({
    url: '/voteRecord/findVoteRecord',
    method: 'get',
    params
  })
}

// @Tags VoteRecord
// @Summary 分页获取投票记录列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取投票记录列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /voteRecord/getVoteRecordList [get]
export const getVoteRecordList = (params) => {
  return service({
    url: '/voteRecord/getVoteRecordList',
    method: 'get',
    params
  })
}

// @Tags VoteRecord
// @Summary 不需要鉴权的投票记录接口
// @accept application/json
// @Produce application/json
// @Param data query VoteRecordReq.VoteRecordSearch true "分页获取投票记录列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /voteRecord/getVoteRecordPublic [get]
export const getVoteRecordPublic = () => {
  return service({
    url: '/voteRecord/getVoteRecordPublic',
    method: 'get',
  })
}

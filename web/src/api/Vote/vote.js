import service from '@/utils/request'
// @Tags Vote
// @Summary 创建投票
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Vote true "创建投票"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /vote/createVote [post]
export const createVote = (data) => {
  return service({
    url: '/vote/createVote',
    method: 'post',
    data
  })
}

// @Tags Vote
// @Summary 删除投票
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Vote true "删除投票"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /vote/deleteVote [delete]
export const deleteVote = (params) => {
  return service({
    url: '/vote/deleteVote',
    method: 'delete',
    params
  })
}

// @Tags Vote
// @Summary 批量删除投票
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除投票"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /vote/deleteVote [delete]
export const deleteVoteByIds = (params) => {
  return service({
    url: '/vote/deleteVoteByIds',
    method: 'delete',
    params
  })
}

// @Tags Vote
// @Summary 更新投票
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Vote true "更新投票"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /vote/updateVote [put]
export const updateVote = (data) => {
  return service({
    url: '/vote/updateVote',
    method: 'put',
    data
  })
}

// @Tags Vote
// @Summary 用id查询投票
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Vote true "用id查询投票"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /vote/findVote [get]
export const findVote = (params) => {
  return service({
    url: '/vote/findVote',
    method: 'get',
    params
  })
}

// @Tags Vote
// @Summary 分页获取投票列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取投票列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /vote/getVoteList [get]
export const getVoteList = (params) => {
  return service({
    url: '/vote/getVoteList',
    method: 'get',
    params
  })
}

// @Tags Vote
// @Summary 不需要鉴权的投票接口
// @accept application/json
// @Produce application/json
// @Param data query VoteReq.VoteSearch true "分页获取投票列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /vote/getVotePublic [get]
export const getVotePublic = () => {
  return service({
    url: '/vote/getVotePublic',
    method: 'get',
  })
}

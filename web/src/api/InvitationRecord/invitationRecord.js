import service from '@/utils/request'
// @Tags InvitationRecord
// @Summary 创建邀请记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.InvitationRecord true "创建邀请记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /IR/createInvitationRecord [post]
export const createInvitationRecord = (data) => {
  return service({
    url: '/IR/createInvitationRecord',
    method: 'post',
    data
  })
}

// @Tags InvitationRecord
// @Summary 删除邀请记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.InvitationRecord true "删除邀请记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /IR/deleteInvitationRecord [delete]
export const deleteInvitationRecord = (params) => {
  return service({
    url: '/IR/deleteInvitationRecord',
    method: 'delete',
    params
  })
}

// @Tags InvitationRecord
// @Summary 批量删除邀请记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除邀请记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /IR/deleteInvitationRecord [delete]
export const deleteInvitationRecordByIds = (params) => {
  return service({
    url: '/IR/deleteInvitationRecordByIds',
    method: 'delete',
    params
  })
}

// @Tags InvitationRecord
// @Summary 更新邀请记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.InvitationRecord true "更新邀请记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /IR/updateInvitationRecord [put]
export const updateInvitationRecord = (data) => {
  return service({
    url: '/IR/updateInvitationRecord',
    method: 'put',
    data
  })
}

// @Tags InvitationRecord
// @Summary 用id查询邀请记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.InvitationRecord true "用id查询邀请记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /IR/findInvitationRecord [get]
export const findInvitationRecord = (params) => {
  return service({
    url: '/IR/findInvitationRecord',
    method: 'get',
    params
  })
}

// @Tags InvitationRecord
// @Summary 分页获取邀请记录列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取邀请记录列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /IR/getInvitationRecordList [get]
export const getInvitationRecordList = (params) => {
  return service({
    url: '/IR/getInvitationRecordList',
    method: 'get',
    params
  })
}

// @Tags InvitationRecord
// @Summary 不需要鉴权的邀请记录接口
// @accept application/json
// @Produce application/json
// @Param data query InvitationRecordReq.InvitationRecordSearch true "分页获取邀请记录列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /IR/getInvitationRecordPublic [get]
export const getInvitationRecordPublic = () => {
  return service({
    url: '/IR/getInvitationRecordPublic',
    method: 'get',
  })
}

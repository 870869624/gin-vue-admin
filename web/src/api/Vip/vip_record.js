import service from '@/utils/request'
// @Tags VipRecord
// @Summary 创建会员记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.VipRecord true "创建会员记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /vipRecord/createVipRecord [post]
export const createVipRecord = (data) => {
  return service({
    url: '/vipRecord/createVipRecord',
    method: 'post',
    data
  })
}

// @Tags VipRecord
// @Summary 删除会员记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.VipRecord true "删除会员记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /vipRecord/deleteVipRecord [delete]
export const deleteVipRecord = (params) => {
  return service({
    url: '/vipRecord/deleteVipRecord',
    method: 'delete',
    params
  })
}

// @Tags VipRecord
// @Summary 批量删除会员记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除会员记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /vipRecord/deleteVipRecord [delete]
export const deleteVipRecordByIds = (params) => {
  return service({
    url: '/vipRecord/deleteVipRecordByIds',
    method: 'delete',
    params
  })
}

// @Tags VipRecord
// @Summary 更新会员记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.VipRecord true "更新会员记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /vipRecord/updateVipRecord [put]
export const updateVipRecord = (data) => {
  return service({
    url: '/vipRecord/updateVipRecord',
    method: 'put',
    data
  })
}

// @Tags VipRecord
// @Summary 用id查询会员记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.VipRecord true "用id查询会员记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /vipRecord/findVipRecord [get]
export const findVipRecord = (params) => {
  return service({
    url: '/vipRecord/findVipRecord',
    method: 'get',
    params
  })
}

// @Tags VipRecord
// @Summary 分页获取会员记录列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取会员记录列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /vipRecord/getVipRecordList [get]
export const getVipRecordList = (params) => {
  return service({
    url: '/vipRecord/getVipRecordList',
    method: 'get',
    params
  })
}

// @Tags VipRecord
// @Summary 不需要鉴权的会员记录接口
// @accept application/json
// @Produce application/json
// @Param data query VipReq.VipRecordSearch true "分页获取会员记录列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /vipRecord/getVipRecordPublic [get]
export const getVipRecordPublic = () => {
  return service({
    url: '/vipRecord/getVipRecordPublic',
    method: 'get',
  })
}

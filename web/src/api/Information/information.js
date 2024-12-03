import service from '@/utils/request'
// @Tags Information
// @Summary 创建资讯
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Information true "创建资讯"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /information/createInformation [post]
export const createInformation = (data) => {
  return service({
    url: '/information/createInformation',
    method: 'post',
    data
  })
}

// @Tags Information
// @Summary 删除资讯
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Information true "删除资讯"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /information/deleteInformation [delete]
export const deleteInformation = (params) => {
  return service({
    url: '/information/deleteInformation',
    method: 'delete',
    params
  })
}

// @Tags Information
// @Summary 批量删除资讯
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除资讯"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /information/deleteInformation [delete]
export const deleteInformationByIds = (params) => {
  return service({
    url: '/information/deleteInformationByIds',
    method: 'delete',
    params
  })
}

// @Tags Information
// @Summary 更新资讯
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Information true "更新资讯"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /information/updateInformation [put]
export const updateInformation = (data) => {
  return service({
    url: '/information/updateInformation',
    method: 'put',
    data
  })
}

// @Tags Information
// @Summary 用id查询资讯
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Information true "用id查询资讯"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /information/findInformation [get]
export const findInformation = (params) => {
  return service({
    url: '/information/findInformation',
    method: 'get',
    params
  })
}

// @Tags Information
// @Summary 分页获取资讯列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取资讯列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /information/getInformationList [get]
export const getInformationList = (params) => {
  return service({
    url: '/information/getInformationList',
    method: 'get',
    params
  })
}

// @Tags Information
// @Summary 不需要鉴权的资讯接口
// @accept application/json
// @Produce application/json
// @Param data query InformationReq.InformationSearch true "分页获取资讯列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /information/getInformationPublic [get]
export const getInformationPublic = () => {
  return service({
    url: '/information/getInformationPublic',
    method: 'get',
  })
}

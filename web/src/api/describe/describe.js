import service from '@/utils/request'
// @Tags Describe
// @Summary 创建文档说明
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Describe true "创建文档说明"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /D/createDescribe [post]
export const createDescribe = (data) => {
  return service({
    url: '/D/createDescribe',
    method: 'post',
    data
  })
}

// @Tags Describe
// @Summary 删除文档说明
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Describe true "删除文档说明"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /D/deleteDescribe [delete]
export const deleteDescribe = (params) => {
  return service({
    url: '/D/deleteDescribe',
    method: 'delete',
    params
  })
}

// @Tags Describe
// @Summary 批量删除文档说明
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除文档说明"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /D/deleteDescribe [delete]
export const deleteDescribeByIds = (params) => {
  return service({
    url: '/D/deleteDescribeByIds',
    method: 'delete',
    params
  })
}

// @Tags Describe
// @Summary 更新文档说明
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Describe true "更新文档说明"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /D/updateDescribe [put]
export const updateDescribe = (data) => {
  return service({
    url: '/D/updateDescribe',
    method: 'put',
    data
  })
}

// @Tags Describe
// @Summary 用id查询文档说明
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Describe true "用id查询文档说明"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /D/findDescribe [get]
export const findDescribe = (params) => {
  return service({
    url: '/D/findDescribe',
    method: 'get',
    params
  })
}

// @Tags Describe
// @Summary 分页获取文档说明列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取文档说明列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /D/getDescribeList [get]
export const getDescribeList = (params) => {
  return service({
    url: '/D/getDescribeList',
    method: 'get',
    params
  })
}

// @Tags Describe
// @Summary 不需要鉴权的文档说明接口
// @accept application/json
// @Produce application/json
// @Param data query describeReq.DescribeSearch true "分页获取文档说明列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /D/getDescribePublic [get]
export const getDescribePublic = () => {
  return service({
    url: '/D/getDescribePublic',
    method: 'get',
  })
}

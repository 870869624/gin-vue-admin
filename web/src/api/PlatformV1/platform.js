import service from '@/utils/request'
// @Tags Platform
// @Summary 创建平台
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Platform true "创建平台"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /platform/createPlatform [post]
export const createPlatform = (data) => {
  return service({
    url: '/platform/createPlatform',
    method: 'post',
    data
  })
}

// @Tags Platform
// @Summary 删除平台
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Platform true "删除平台"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /platform/deletePlatform [delete]
export const deletePlatform = (params) => {
  return service({
    url: '/platform/deletePlatform',
    method: 'delete',
    params
  })
}

// @Tags Platform
// @Summary 批量删除平台
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除平台"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /platform/deletePlatform [delete]
export const deletePlatformByIds = (params) => {
  return service({
    url: '/platform/deletePlatformByIds',
    method: 'delete',
    params
  })
}

// @Tags Platform
// @Summary 更新平台
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Platform true "更新平台"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /platform/updatePlatform [put]
export const updatePlatform = (data) => {
  return service({
    url: '/platform/updatePlatform',
    method: 'put',
    data
  })
}

// @Tags Platform
// @Summary 用id查询平台
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Platform true "用id查询平台"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /platform/findPlatform [get]
export const findPlatform = (params) => {
  return service({
    url: '/platform/findPlatform',
    method: 'get',
    params
  })
}

// @Tags Platform
// @Summary 分页获取平台列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取平台列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /platform/getPlatformList [get]
export const getPlatformList = (params) => {
  return service({
    url: '/platform/getPlatformList',
    method: 'get',
    params
  })
}

// @Tags Platform
// @Summary 不需要鉴权的平台接口
// @accept application/json
// @Produce application/json
// @Param data query PlatformV1Req.PlatformSearch true "分页获取平台列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /platform/getPlatformPublic [get]
export const getPlatformPublic = () => {
  return service({
    url: '/platform/getPlatformPublic',
    method: 'get',
  })
}

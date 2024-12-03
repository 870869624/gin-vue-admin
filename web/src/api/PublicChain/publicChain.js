import service from '@/utils/request'
// @Tags PublicChain
// @Summary 创建公链
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PublicChain true "创建公链"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /PC/createPublicChain [post]
export const createPublicChain = (data) => {
  return service({
    url: '/PC/createPublicChain',
    method: 'post',
    data
  })
}

// @Tags PublicChain
// @Summary 删除公链
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PublicChain true "删除公链"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /PC/deletePublicChain [delete]
export const deletePublicChain = (params) => {
  return service({
    url: '/PC/deletePublicChain',
    method: 'delete',
    params
  })
}

// @Tags PublicChain
// @Summary 批量删除公链
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除公链"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /PC/deletePublicChain [delete]
export const deletePublicChainByIds = (params) => {
  return service({
    url: '/PC/deletePublicChainByIds',
    method: 'delete',
    params
  })
}

// @Tags PublicChain
// @Summary 更新公链
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PublicChain true "更新公链"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /PC/updatePublicChain [put]
export const updatePublicChain = (data) => {
  return service({
    url: '/PC/updatePublicChain',
    method: 'put',
    data
  })
}

// @Tags PublicChain
// @Summary 用id查询公链
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.PublicChain true "用id查询公链"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /PC/findPublicChain [get]
export const findPublicChain = (params) => {
  return service({
    url: '/PC/findPublicChain',
    method: 'get',
    params
  })
}

// @Tags PublicChain
// @Summary 分页获取公链列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取公链列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /PC/getPublicChainList [get]
export const getPublicChainList = (params) => {
  return service({
    url: '/PC/getPublicChainList',
    method: 'get',
    params
  })
}

// @Tags PublicChain
// @Summary 不需要鉴权的公链接口
// @accept application/json
// @Produce application/json
// @Param data query PublicChainReq.PublicChainSearch true "分页获取公链列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /PC/getPublicChainPublic [get]
export const getPublicChainPublic = () => {
  return service({
    url: '/PC/getPublicChainPublic',
    method: 'get',
  })
}

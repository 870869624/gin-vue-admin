import service from '@/utils/request'
// @Tags Airdrop
// @Summary 创建空投项目
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Airdrop true "创建空投项目"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /airdrop/createAirdrop [post]
export const createAirdrop = (data) => {
  return service({
    url: '/airdrop/createAirdrop',
    method: 'post',
    data
  })
}

// @Tags Airdrop
// @Summary 删除空投项目
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Airdrop true "删除空投项目"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /airdrop/deleteAirdrop [delete]
export const deleteAirdrop = (params) => {
  return service({
    url: '/airdrop/deleteAirdrop',
    method: 'delete',
    params
  })
}

// @Tags Airdrop
// @Summary 批量删除空投项目
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除空投项目"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /airdrop/deleteAirdrop [delete]
export const deleteAirdropByIds = (params) => {
  return service({
    url: '/airdrop/deleteAirdropByIds',
    method: 'delete',
    params
  })
}

// @Tags Airdrop
// @Summary 更新空投项目
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Airdrop true "更新空投项目"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /airdrop/updateAirdrop [put]
export const updateAirdrop = (data) => {
  return service({
    url: '/airdrop/updateAirdrop',
    method: 'put',
    data
  })
}

// @Tags Airdrop
// @Summary 用id查询空投项目
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Airdrop true "用id查询空投项目"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /airdrop/findAirdrop [get]
export const findAirdrop = (params) => {
  return service({
    url: '/airdrop/findAirdrop',
    method: 'get',
    params
  })
}

// @Tags Airdrop
// @Summary 分页获取空投项目列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取空投项目列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /airdrop/getAirdropList [get]
export const getAirdropList = (params) => {
  return service({
    url: '/airdrop/getAirdropList',
    method: 'get',
    params
  })
}

// @Tags Airdrop
// @Summary 不需要鉴权的空投项目接口
// @accept application/json
// @Produce application/json
// @Param data query AirdropReq.AirdropSearch true "分页获取空投项目列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /airdrop/getAirdropPublic [get]
export const getAirdropPublic = () => {
  return service({
    url: '/airdrop/getAirdropPublic',
    method: 'get',
  })
}

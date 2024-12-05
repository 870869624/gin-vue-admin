import service from '@/utils/request'
// @Tags Presale
// @Summary 创建预售项目
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Presale true "创建预售项目"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /presale/createPresale [post]
export const createPresale = (data) => {
  return service({
    url: '/presale/createPresale',
    method: 'post',
    data
  })
}

// @Tags Presale
// @Summary 删除预售项目
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Presale true "删除预售项目"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /presale/deletePresale [delete]
export const deletePresale = (params) => {
  return service({
    url: '/presale/deletePresale',
    method: 'delete',
    params
  })
}

// @Tags Presale
// @Summary 批量删除预售项目
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除预售项目"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /presale/deletePresale [delete]
export const deletePresaleByIds = (params) => {
  return service({
    url: '/presale/deletePresaleByIds',
    method: 'delete',
    params
  })
}

// @Tags Presale
// @Summary 更新预售项目
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Presale true "更新预售项目"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /presale/updatePresale [put]
export const updatePresale = (data) => {
  return service({
    url: '/presale/updatePresale',
    method: 'put',
    data
  })
}

// @Tags Presale
// @Summary 用id查询预售项目
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Presale true "用id查询预售项目"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /presale/findPresale [get]
export const findPresale = (params) => {
  return service({
    url: '/presale/findPresale',
    method: 'get',
    params
  })
}

// @Tags Presale
// @Summary 分页获取预售项目列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取预售项目列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /presale/getPresaleList [get]
export const getPresaleList = (params) => {
  return service({
    url: '/presale/getPresaleList',
    method: 'get',
    params
  })
}

// @Tags Presale
// @Summary 不需要鉴权的预售项目接口
// @accept application/json
// @Produce application/json
// @Param data query PresaleReq.PresaleSearch true "分页获取预售项目列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /presale/getPresalePublic [get]
export const getPresalePublic = () => {
  return service({
    url: '/presale/getPresalePublic',
    method: 'get',
  })
}

import service from '@/utils/request'
// @Tags NavigationBar
// @Summary 创建导航栏分类
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NavigationBar true "创建导航栏分类"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /NA/createNavigationBar [post]
export const createNavigationBar = (data) => {
  return service({
    url: '/NA/createNavigationBar',
    method: 'post',
    data
  })
}

// @Tags NavigationBar
// @Summary 删除导航栏分类
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NavigationBar true "删除导航栏分类"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /NA/deleteNavigationBar [delete]
export const deleteNavigationBar = (params) => {
  return service({
    url: '/NA/deleteNavigationBar',
    method: 'delete',
    params
  })
}

// @Tags NavigationBar
// @Summary 批量删除导航栏分类
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除导航栏分类"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /NA/deleteNavigationBar [delete]
export const deleteNavigationBarByIds = (params) => {
  return service({
    url: '/NA/deleteNavigationBarByIds',
    method: 'delete',
    params
  })
}

// @Tags NavigationBar
// @Summary 更新导航栏分类
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NavigationBar true "更新导航栏分类"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /NA/updateNavigationBar [put]
export const updateNavigationBar = (data) => {
  return service({
    url: '/NA/updateNavigationBar',
    method: 'put',
    data
  })
}

// @Tags NavigationBar
// @Summary 用id查询导航栏分类
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.NavigationBar true "用id查询导航栏分类"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /NA/findNavigationBar [get]
export const findNavigationBar = (params) => {
  return service({
    url: '/NA/findNavigationBar',
    method: 'get',
    params
  })
}

// @Tags NavigationBar
// @Summary 分页获取导航栏分类列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取导航栏分类列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /NA/getNavigationBarList [get]
export const getNavigationBarList = (params) => {
  return service({
    url: '/NA/getNavigationBarList',
    method: 'get',
    params
  })
}

// @Tags NavigationBar
// @Summary 不需要鉴权的导航栏分类接口
// @accept application/json
// @Produce application/json
// @Param data query NavigationBarReq.NavigationBarSearch true "分页获取导航栏分类列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /NA/getNavigationBarPublic [get]
export const getNavigationBarPublic = () => {
  return service({
    url: '/NA/getNavigationBarPublic',
    method: 'get',
  })
}

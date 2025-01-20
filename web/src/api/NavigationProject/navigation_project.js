import service from '@/utils/request'
// @Tags NavigationProject
// @Summary 创建导航栏项目
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NavigationProject true "创建导航栏项目"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /NP/createNavigationProject [post]
export const createNavigationProject = (data) => {
  return service({
    url: '/NP/createNavigationProject',
    method: 'post',
    data
  })
}

// @Tags NavigationProject
// @Summary 删除导航栏项目
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NavigationProject true "删除导航栏项目"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /NP/deleteNavigationProject [delete]
export const deleteNavigationProject = (params) => {
  return service({
    url: '/NP/deleteNavigationProject',
    method: 'delete',
    params
  })
}

// @Tags NavigationProject
// @Summary 批量删除导航栏项目
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除导航栏项目"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /NP/deleteNavigationProject [delete]
export const deleteNavigationProjectByIds = (params) => {
  return service({
    url: '/NP/deleteNavigationProjectByIds',
    method: 'delete',
    params
  })
}

// @Tags NavigationProject
// @Summary 更新导航栏项目
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NavigationProject true "更新导航栏项目"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /NP/updateNavigationProject [put]
export const updateNavigationProject = (data) => {
  return service({
    url: '/NP/updateNavigationProject',
    method: 'put',
    data
  })
}

// @Tags NavigationProject
// @Summary 用id查询导航栏项目
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.NavigationProject true "用id查询导航栏项目"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /NP/findNavigationProject [get]
export const findNavigationProject = (params) => {
  return service({
    url: '/NP/findNavigationProject',
    method: 'get',
    params
  })
}

// @Tags NavigationProject
// @Summary 分页获取导航栏项目列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取导航栏项目列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /NP/getNavigationProjectList [get]
export const getNavigationProjectList = (params) => {
  return service({
    url: '/NP/getNavigationProjectList',
    method: 'get',
    params
  })
}

// @Tags NavigationProject
// @Summary 不需要鉴权的导航栏项目接口
// @accept application/json
// @Produce application/json
// @Param data query NavigationProjectReq.NavigationProjectSearch true "分页获取导航栏项目列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /NP/getNavigationProjectPublic [get]
export const getNavigationProjectPublic = () => {
  return service({
    url: '/NP/getNavigationProjectPublic',
    method: 'get',
  })
}

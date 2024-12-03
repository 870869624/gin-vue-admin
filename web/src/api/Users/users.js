import service from '@/utils/request'
// @Tags Users
// @Summary 创建用户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Users true "创建用户列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /users/createUsers [post]
export const createUsers = (data) => {
  return service({
    url: '/users/createUsers',
    method: 'post',
    data
  })
}

// @Tags Users
// @Summary 删除用户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Users true "删除用户列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /users/deleteUsers [delete]
export const deleteUsers = (params) => {
  return service({
    url: '/users/deleteUsers',
    method: 'delete',
    params
  })
}

// @Tags Users
// @Summary 批量删除用户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除用户列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /users/deleteUsers [delete]
export const deleteUsersByIds = (params) => {
  return service({
    url: '/users/deleteUsersByIds',
    method: 'delete',
    params
  })
}

// @Tags Users
// @Summary 更新用户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Users true "更新用户列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /users/updateUsers [put]
export const updateUsers = (data) => {
  return service({
    url: '/users/updateUsers',
    method: 'put',
    data
  })
}

// @Tags Users
// @Summary 用id查询用户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Users true "用id查询用户列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /users/findUsers [get]
export const findUsers = (params) => {
  return service({
    url: '/users/findUsers',
    method: 'get',
    params
  })
}

// @Tags Users
// @Summary 分页获取用户列表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取用户列表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /users/getUsersList [get]
export const getUsersList = (params) => {
  return service({
    url: '/users/getUsersList',
    method: 'get',
    params
  })
}

// @Tags Users
// @Summary 不需要鉴权的用户列表接口
// @accept application/json
// @Produce application/json
// @Param data query UsersReq.UsersSearch true "分页获取用户列表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /users/getUsersPublic [get]
export const getUsersPublic = () => {
  return service({
    url: '/users/getUsersPublic',
    method: 'get',
  })
}

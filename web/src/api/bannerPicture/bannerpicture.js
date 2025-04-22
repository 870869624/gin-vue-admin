import service from '@/utils/request'
// @Tags BannerPicture
// @Summary 创建横幅图
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BannerPicture true "创建横幅图"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /BP/createBannerPicture [post]
export const createBannerPicture = (data) => {
  return service({
    url: '/BP/createBannerPicture',
    method: 'post',
    data
  })
}

// @Tags BannerPicture
// @Summary 删除横幅图
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BannerPicture true "删除横幅图"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /BP/deleteBannerPicture [delete]
export const deleteBannerPicture = (params) => {
  return service({
    url: '/BP/deleteBannerPicture',
    method: 'delete',
    params
  })
}

// @Tags BannerPicture
// @Summary 批量删除横幅图
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除横幅图"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /BP/deleteBannerPicture [delete]
export const deleteBannerPictureByIds = (params) => {
  return service({
    url: '/BP/deleteBannerPictureByIds',
    method: 'delete',
    params
  })
}

// @Tags BannerPicture
// @Summary 更新横幅图
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BannerPicture true "更新横幅图"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /BP/updateBannerPicture [put]
export const updateBannerPicture = (data) => {
  return service({
    url: '/BP/updateBannerPicture',
    method: 'put',
    data
  })
}

// @Tags BannerPicture
// @Summary 用id查询横幅图
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BannerPicture true "用id查询横幅图"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /BP/findBannerPicture [get]
export const findBannerPicture = (params) => {
  return service({
    url: '/BP/findBannerPicture',
    method: 'get',
    params
  })
}

// @Tags BannerPicture
// @Summary 分页获取横幅图列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取横幅图列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /BP/getBannerPictureList [get]
export const getBannerPictureList = (params) => {
  return service({
    url: '/BP/getBannerPictureList',
    method: 'get',
    params
  })
}

// @Tags BannerPicture
// @Summary 不需要鉴权的横幅图接口
// @accept application/json
// @Produce application/json
// @Param data query bannerPictureReq.BannerPictureSearch true "分页获取横幅图列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /BP/getBannerPicturePublic [get]
export const getBannerPicturePublic = () => {
  return service({
    url: '/BP/getBannerPicturePublic',
    method: 'get',
  })
}

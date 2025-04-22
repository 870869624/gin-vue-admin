package bannerPicture

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/bannerPicture"
    bannerPictureReq "github.com/flipped-aurora/gin-vue-admin/server/model/bannerPicture/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type BannerPictureApi struct {}



// CreateBannerPicture 创建横幅图
// @Tags BannerPicture
// @Summary 创建横幅图
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bannerPicture.BannerPicture true "创建横幅图"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /BP/createBannerPicture [post]
func (BPApi *BannerPictureApi) CreateBannerPicture(c *gin.Context) {
	var BP bannerPicture.BannerPicture
	err := c.ShouldBindJSON(&BP)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = BPService.CreateBannerPicture(&BP)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteBannerPicture 删除横幅图
// @Tags BannerPicture
// @Summary 删除横幅图
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bannerPicture.BannerPicture true "删除横幅图"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /BP/deleteBannerPicture [delete]
func (BPApi *BannerPictureApi) DeleteBannerPicture(c *gin.Context) {
	ID := c.Query("ID")
	err := BPService.DeleteBannerPicture(ID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteBannerPictureByIds 批量删除横幅图
// @Tags BannerPicture
// @Summary 批量删除横幅图
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /BP/deleteBannerPictureByIds [delete]
func (BPApi *BannerPictureApi) DeleteBannerPictureByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := BPService.DeleteBannerPictureByIds(IDs)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateBannerPicture 更新横幅图
// @Tags BannerPicture
// @Summary 更新横幅图
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bannerPicture.BannerPicture true "更新横幅图"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /BP/updateBannerPicture [put]
func (BPApi *BannerPictureApi) UpdateBannerPicture(c *gin.Context) {
	var BP bannerPicture.BannerPicture
	err := c.ShouldBindJSON(&BP)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = BPService.UpdateBannerPicture(BP)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindBannerPicture 用id查询横幅图
// @Tags BannerPicture
// @Summary 用id查询横幅图
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bannerPicture.BannerPicture true "用id查询横幅图"
// @Success 200 {object} response.Response{data=bannerPicture.BannerPicture,msg=string} "查询成功"
// @Router /BP/findBannerPicture [get]
func (BPApi *BannerPictureApi) FindBannerPicture(c *gin.Context) {
	ID := c.Query("ID")
	reBP, err := BPService.GetBannerPicture(ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(reBP, c)
}

// GetBannerPictureList 分页获取横幅图列表
// @Tags BannerPicture
// @Summary 分页获取横幅图列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bannerPictureReq.BannerPictureSearch true "分页获取横幅图列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /BP/getBannerPictureList [get]
func (BPApi *BannerPictureApi) GetBannerPictureList(c *gin.Context) {
	var pageInfo bannerPictureReq.BannerPictureSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := BPService.GetBannerPictureInfoList(pageInfo)
	if err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败:" + err.Error(), c)
        return
    }
    response.OkWithDetailed(response.PageResult{
        List:     list,
        Total:    total,
        Page:     pageInfo.Page,
        PageSize: pageInfo.PageSize,
    }, "获取成功", c)
}

// GetBannerPicturePublic 不需要鉴权的横幅图接口
// @Tags BannerPicture
// @Summary 不需要鉴权的横幅图接口
// @accept application/json
// @Produce application/json
// @Param data query bannerPictureReq.BannerPictureSearch true "分页获取横幅图列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /BP/getBannerPicturePublic [get]
func (BPApi *BannerPictureApi) GetBannerPicturePublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    BPService.GetBannerPicturePublic()
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的横幅图接口信息",
    }, "获取成功", c)
}

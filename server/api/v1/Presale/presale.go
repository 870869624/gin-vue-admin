package Presale

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Presale"
	PresaleReq "github.com/flipped-aurora/gin-vue-admin/server/model/Presale/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PresaleApi struct{}

// CreatePresale 创建预售项目
// @Tags Presale
// @Summary 创建预售项目
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Presale.Presale true "创建预售项目"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /presale/createPresale [post]
func (presaleApi *PresaleApi) CreatePresale(c *gin.Context) {
	var presale Presale.Presale
	err := c.ShouldBindJSON(&presale)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = presaleService.CreatePresale(&presale)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeletePresale 删除预售项目
// @Tags Presale
// @Summary 删除预售项目
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Presale.Presale true "删除预售项目"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /presale/deletePresale [delete]
func (presaleApi *PresaleApi) DeletePresale(c *gin.Context) {
	ID := c.Query("ID")
	err := presaleService.DeletePresale(ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeletePresaleByIds 批量删除预售项目
// @Tags Presale
// @Summary 批量删除预售项目
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /presale/deletePresaleByIds [delete]
func (presaleApi *PresaleApi) DeletePresaleByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := presaleService.DeletePresaleByIds(IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdatePresale 更新预售项目
// @Tags Presale
// @Summary 更新预售项目
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Presale.Presale true "更新预售项目"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /presale/updatePresale [put]
func (presaleApi *PresaleApi) UpdatePresale(c *gin.Context) {
	var presale Presale.Presale
	err := c.ShouldBindJSON(&presale)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = presaleService.UpdatePresale(presale)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindPresale 用id查询预售项目
// @Tags Presale
// @Summary 用id查询预售项目
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query Presale.Presale true "用id查询预售项目"
// @Success 200 {object} response.Response{data=Presale.Presale,msg=string} "查询成功"
// @Router /presale/findPresale [get]
func (presaleApi *PresaleApi) FindPresale(c *gin.Context) {
	ID := c.Query("ID")
	represale, err := presaleService.GetPresale(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(represale, c)
}

// GetPresaleList 分页获取预售项目列表
// @Tags Presale
// @Summary 分页获取预售项目列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query PresaleReq.PresaleSearch true "分页获取预售项目列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /presale/getPresaleList [get]
func (presaleApi *PresaleApi) GetPresaleList(c *gin.Context) {
	var pageInfo PresaleReq.PresaleSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := presaleService.GetPresaleInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// GetPresalePublic 不需要鉴权的预售项目接口
// @Tags Presale
// @Summary 不需要鉴权的预售项目接口
// @accept application/json
// @Produce application/json
// @Param data query PresaleReq.PresaleSearch true "分页获取预售项目列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /presale/getPresalePublic [get]
func (presaleApi *PresaleApi) GetPresalePublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	presaleService.GetPresalePublic()
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的预售项目接口信息",
	}, "获取成功", c)
}

func (presaleApi *PresaleApi) MobileCreatePresale(c *gin.Context) {
	var isShow = false
	var isPass = false
	var presale Presale.MobilePresaleCre
	err := c.ShouldBindJSON(&presale)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	uid := utils.GetMobileUserID(c)
	if uid == 0 {

		response.FailWithMessage("查询失败:用户ID为空", c)
		return
	}
	userId := int(uid)
	presale.UserId = &userId
	presale.PresaleIsShow = &isShow
	presale.PresaleIsPass = &isPass
	err = presaleService.CreatePresaleMobile(&presale)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

func (presaleApi *PresaleApi) MobileGetPresaleList(c *gin.Context) {
	var pageInfo PresaleReq.PresaleSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := presaleService.MobileGetPresaleInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

func (presaleApi *PresaleApi) MobileFindPresale(c *gin.Context) {
	ID := c.Query("ID")
	represale, err := presaleService.GetPresale(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(represale, c)
}

func (presaleApi *PresaleApi) GetPresaleListMobile(c *gin.Context) {
	var pageInfo PresaleReq.PresaleSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := presaleService.GetPresaleInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

func (presaleApi *PresaleApi) FindPresaleMobile(c *gin.Context) {
	ID := c.Query("ID")
	represale, err := presaleService.GetPresaleMobile(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(represale, c)
}

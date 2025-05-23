package Information

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Information"
	InformationReq "github.com/flipped-aurora/gin-vue-admin/server/model/Information/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type InformationApi struct{}

// CreateInformation 创建资讯
// @Tags Information
// @Summary 创建资讯
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Information.Information true "创建资讯"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /information/createInformation [post]
func (informationApi *InformationApi) CreateInformation(c *gin.Context) {
	var information Information.Information
	err := c.ShouldBindJSON(&information)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = informationService.CreateInformation(&information)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteInformation 删除资讯
// @Tags Information
// @Summary 删除资讯
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Information.Information true "删除资讯"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /information/deleteInformation [delete]
func (informationApi *InformationApi) DeleteInformation(c *gin.Context) {
	ID := c.Query("ID")
	err := informationService.DeleteInformation(ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteInformationByIds 批量删除资讯
// @Tags Information
// @Summary 批量删除资讯
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /information/deleteInformationByIds [delete]
func (informationApi *InformationApi) DeleteInformationByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := informationService.DeleteInformationByIds(IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateInformation 更新资讯
// @Tags Information
// @Summary 更新资讯
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Information.Information true "更新资讯"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /information/updateInformation [put]
func (informationApi *InformationApi) UpdateInformation(c *gin.Context) {
	var information Information.Information
	err := c.ShouldBindJSON(&information)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = informationService.UpdateInformation(information)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindInformation 用id查询资讯
// @Tags Information
// @Summary 用id查询资讯
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query Information.Information true "用id查询资讯"
// @Success 200 {object} response.Response{data=Information.Information,msg=string} "查询成功"
// @Router /information/findInformation [get]
func (informationApi *InformationApi) FindInformation(c *gin.Context) {
	ID := c.Query("ID")
	reinformation, err := informationService.GetInformation(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reinformation, c)
}

// GetInformationList 分页获取资讯列表
// @Tags Information
// @Summary 分页获取资讯列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query InformationReq.InformationSearch true "分页获取资讯列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /information/getInformationList [get]
func (informationApi *InformationApi) GetInformationList(c *gin.Context) {
	var pageInfo InformationReq.InformationSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := informationService.GetInformationInfoList(pageInfo)
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

// GetInformationPublic 不需要鉴权的资讯接口
// @Tags Information
// @Summary 不需要鉴权的资讯接口
// @accept application/json
// @Produce application/json
// @Param data query InformationReq.InformationSearch true "分页获取资讯列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /information/getInformationPublic [get]
func (informationApi *InformationApi) GetInformationPublic(c *gin.Context) {
	var pageInfo request.SysDictionaryDetailSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := dictionaryDetailService.GetSysDictionaryDetailInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

func (informationApi *InformationApi) FindInformationMobile(c *gin.Context) {
	ID := c.Query("ID")
	reinformation, err := informationService.GetInformationMobile(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reinformation, c)
}

func (informationApi *InformationApi) GetInformationListMobile(c *gin.Context) {
	var pageInfo InformationReq.InformationSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := informationService.GetInformationInfoListMobile(pageInfo)
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

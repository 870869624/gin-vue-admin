package PublicChain

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/PublicChain"
	PublicChainReq "github.com/flipped-aurora/gin-vue-admin/server/model/PublicChain/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PublicChainApi struct{}

// CreatePublicChain 创建公链
// @Tags PublicChain
// @Summary 创建公链
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body PublicChain.PublicChain true "创建公链"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /PC/createPublicChain [post]
func (PCApi *PublicChainApi) CreatePublicChain(c *gin.Context) {
	var PC PublicChain.PublicChain
	err := c.ShouldBindJSON(&PC)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = PCService.CreatePublicChain(&PC)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeletePublicChain 删除公链
// @Tags PublicChain
// @Summary 删除公链
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body PublicChain.PublicChain true "删除公链"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /PC/deletePublicChain [delete]
func (PCApi *PublicChainApi) DeletePublicChain(c *gin.Context) {
	ID := c.Query("ID")
	err := PCService.DeletePublicChain(ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeletePublicChainByIds 批量删除公链
// @Tags PublicChain
// @Summary 批量删除公链
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /PC/deletePublicChainByIds [delete]
func (PCApi *PublicChainApi) DeletePublicChainByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := PCService.DeletePublicChainByIds(IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdatePublicChain 更新公链
// @Tags PublicChain
// @Summary 更新公链
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body PublicChain.PublicChain true "更新公链"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /PC/updatePublicChain [put]
func (PCApi *PublicChainApi) UpdatePublicChain(c *gin.Context) {
	var PC PublicChain.PublicChain
	err := c.ShouldBindJSON(&PC)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = PCService.UpdatePublicChain(PC)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindPublicChain 用id查询公链
// @Tags PublicChain
// @Summary 用id查询公链
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query PublicChain.PublicChain true "用id查询公链"
// @Success 200 {object} response.Response{data=PublicChain.PublicChain,msg=string} "查询成功"
// @Router /PC/findPublicChain [get]
func (PCApi *PublicChainApi) FindPublicChain(c *gin.Context) {
	ID := c.Query("ID")
	rePC, err := PCService.GetPublicChain(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(rePC, c)
}

// GetPublicChainList 分页获取公链列表
// @Tags PublicChain
// @Summary 分页获取公链列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query PublicChainReq.PublicChainSearch true "分页获取公链列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /PC/getPublicChainList [get]
func (PCApi *PublicChainApi) GetPublicChainList(c *gin.Context) {
	var pageInfo PublicChainReq.PublicChainSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := PCService.GetPublicChainInfoList(pageInfo)
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

// GetPublicChainPublic 不需要鉴权的公链接口
// @Tags PublicChain
// @Summary 不需要鉴权的公链接口
// @accept application/json
// @Produce application/json
// @Param data query PublicChainReq.PublicChainSearch true "分页获取公链列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /PC/getPublicChainPublic [get]
func (PCApi *PublicChainApi) GetPublicChainPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	PCService.GetPublicChainPublic()
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的公链接口信息",
	}, "获取成功", c)
}

package describe

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/describe"
    describeReq "github.com/flipped-aurora/gin-vue-admin/server/model/describe/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type DescribeApi struct {}



// CreateDescribe 创建文档说明
// @Tags Describe
// @Summary 创建文档说明
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body describe.Describe true "创建文档说明"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /D/createDescribe [post]
func (DApi *DescribeApi) CreateDescribe(c *gin.Context) {
	var D describe.Describe
	err := c.ShouldBindJSON(&D)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = DService.CreateDescribe(&D)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteDescribe 删除文档说明
// @Tags Describe
// @Summary 删除文档说明
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body describe.Describe true "删除文档说明"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /D/deleteDescribe [delete]
func (DApi *DescribeApi) DeleteDescribe(c *gin.Context) {
	ID := c.Query("ID")
	err := DService.DeleteDescribe(ID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteDescribeByIds 批量删除文档说明
// @Tags Describe
// @Summary 批量删除文档说明
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /D/deleteDescribeByIds [delete]
func (DApi *DescribeApi) DeleteDescribeByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := DService.DeleteDescribeByIds(IDs)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateDescribe 更新文档说明
// @Tags Describe
// @Summary 更新文档说明
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body describe.Describe true "更新文档说明"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /D/updateDescribe [put]
func (DApi *DescribeApi) UpdateDescribe(c *gin.Context) {
	var D describe.Describe
	err := c.ShouldBindJSON(&D)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = DService.UpdateDescribe(D)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindDescribe 用id查询文档说明
// @Tags Describe
// @Summary 用id查询文档说明
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query describe.Describe true "用id查询文档说明"
// @Success 200 {object} response.Response{data=describe.Describe,msg=string} "查询成功"
// @Router /D/findDescribe [get]
func (DApi *DescribeApi) FindDescribe(c *gin.Context) {
	ID := c.Query("ID")
	reD, err := DService.GetDescribe(ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(reD, c)
}

// GetDescribeList 分页获取文档说明列表
// @Tags Describe
// @Summary 分页获取文档说明列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query describeReq.DescribeSearch true "分页获取文档说明列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /D/getDescribeList [get]
func (DApi *DescribeApi) GetDescribeList(c *gin.Context) {
	var pageInfo describeReq.DescribeSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := DService.GetDescribeInfoList(pageInfo)
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

// GetDescribePublic 不需要鉴权的文档说明接口
// @Tags Describe
// @Summary 不需要鉴权的文档说明接口
// @accept application/json
// @Produce application/json
// @Param data query describeReq.DescribeSearch true "分页获取文档说明列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /D/getDescribePublic [get]
func (DApi *DescribeApi) GetDescribePublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    DService.GetDescribePublic()
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的文档说明接口信息",
    }, "获取成功", c)
}

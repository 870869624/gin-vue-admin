package NavigationBar

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/NavigationBar"
    NavigationBarReq "github.com/flipped-aurora/gin-vue-admin/server/model/NavigationBar/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type NavigationBarApi struct {}



// CreateNavigationBar 创建导航栏分类
// @Tags NavigationBar
// @Summary 创建导航栏分类
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body NavigationBar.NavigationBar true "创建导航栏分类"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /NA/createNavigationBar [post]
func (NAApi *NavigationBarApi) CreateNavigationBar(c *gin.Context) {
	var NA NavigationBar.NavigationBar
	err := c.ShouldBindJSON(&NA)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = NAService.CreateNavigationBar(&NA)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteNavigationBar 删除导航栏分类
// @Tags NavigationBar
// @Summary 删除导航栏分类
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body NavigationBar.NavigationBar true "删除导航栏分类"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /NA/deleteNavigationBar [delete]
func (NAApi *NavigationBarApi) DeleteNavigationBar(c *gin.Context) {
	ID := c.Query("ID")
	err := NAService.DeleteNavigationBar(ID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteNavigationBarByIds 批量删除导航栏分类
// @Tags NavigationBar
// @Summary 批量删除导航栏分类
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /NA/deleteNavigationBarByIds [delete]
func (NAApi *NavigationBarApi) DeleteNavigationBarByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := NAService.DeleteNavigationBarByIds(IDs)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateNavigationBar 更新导航栏分类
// @Tags NavigationBar
// @Summary 更新导航栏分类
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body NavigationBar.NavigationBar true "更新导航栏分类"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /NA/updateNavigationBar [put]
func (NAApi *NavigationBarApi) UpdateNavigationBar(c *gin.Context) {
	var NA NavigationBar.NavigationBar
	err := c.ShouldBindJSON(&NA)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = NAService.UpdateNavigationBar(NA)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindNavigationBar 用id查询导航栏分类
// @Tags NavigationBar
// @Summary 用id查询导航栏分类
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query NavigationBar.NavigationBar true "用id查询导航栏分类"
// @Success 200 {object} response.Response{data=NavigationBar.NavigationBar,msg=string} "查询成功"
// @Router /NA/findNavigationBar [get]
func (NAApi *NavigationBarApi) FindNavigationBar(c *gin.Context) {
	ID := c.Query("ID")
	reNA, err := NAService.GetNavigationBar(ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(reNA, c)
}

// GetNavigationBarList 分页获取导航栏分类列表
// @Tags NavigationBar
// @Summary 分页获取导航栏分类列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query NavigationBarReq.NavigationBarSearch true "分页获取导航栏分类列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /NA/getNavigationBarList [get]
func (NAApi *NavigationBarApi) GetNavigationBarList(c *gin.Context) {
	var pageInfo NavigationBarReq.NavigationBarSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := NAService.GetNavigationBarInfoList(pageInfo)
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

// GetNavigationBarPublic 不需要鉴权的导航栏分类接口
// @Tags NavigationBar
// @Summary 不需要鉴权的导航栏分类接口
// @accept application/json
// @Produce application/json
// @Param data query NavigationBarReq.NavigationBarSearch true "分页获取导航栏分类列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /NA/getNavigationBarPublic [get]
func (NAApi *NavigationBarApi) GetNavigationBarPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    NAService.GetNavigationBarPublic()
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的导航栏分类接口信息",
    }, "获取成功", c)
}

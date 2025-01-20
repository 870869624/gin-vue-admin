package NavigationProject

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/NavigationProject"
	NavigationProjectReq "github.com/flipped-aurora/gin-vue-admin/server/model/NavigationProject/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type NavigationProjectApi struct{}

// CreateNavigationProject 创建导航栏项目
// @Tags NavigationProject
// @Summary 创建导航栏项目
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body NavigationProject.NavigationProject true "创建导航栏项目"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /NP/createNavigationProject [post]
func (NPApi *NavigationProjectApi) CreateNavigationProject(c *gin.Context) {
	var NP NavigationProject.NavigationProject
	err := c.ShouldBindJSON(&NP)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = NPService.CreateNavigationProject(&NP)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteNavigationProject 删除导航栏项目
// @Tags NavigationProject
// @Summary 删除导航栏项目
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body NavigationProject.NavigationProject true "删除导航栏项目"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /NP/deleteNavigationProject [delete]
func (NPApi *NavigationProjectApi) DeleteNavigationProject(c *gin.Context) {
	ID := c.Query("ID")
	err := NPService.DeleteNavigationProject(ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteNavigationProjectByIds 批量删除导航栏项目
// @Tags NavigationProject
// @Summary 批量删除导航栏项目
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /NP/deleteNavigationProjectByIds [delete]
func (NPApi *NavigationProjectApi) DeleteNavigationProjectByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := NPService.DeleteNavigationProjectByIds(IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateNavigationProject 更新导航栏项目
// @Tags NavigationProject
// @Summary 更新导航栏项目
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body NavigationProject.NavigationProject true "更新导航栏项目"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /NP/updateNavigationProject [put]
func (NPApi *NavigationProjectApi) UpdateNavigationProject(c *gin.Context) {
	var NP NavigationProject.NavigationProject
	err := c.ShouldBindJSON(&NP)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = NPService.UpdateNavigationProject(NP)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindNavigationProject 用id查询导航栏项目
// @Tags NavigationProject
// @Summary 用id查询导航栏项目
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query NavigationProject.NavigationProject true "用id查询导航栏项目"
// @Success 200 {object} response.Response{data=NavigationProject.NavigationProject,msg=string} "查询成功"
// @Router /NP/findNavigationProject [get]
func (NPApi *NavigationProjectApi) FindNavigationProject(c *gin.Context) {
	ID := c.Query("ID")
	reNP, err := NPService.GetNavigationProject(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reNP, c)
}

// GetNavigationProjectList 分页获取导航栏项目列表
// @Tags NavigationProject
// @Summary 分页获取导航栏项目列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query NavigationProjectReq.NavigationProjectSearch true "分页获取导航栏项目列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /NP/getNavigationProjectList [get]
func (NPApi *NavigationProjectApi) GetNavigationProjectList(c *gin.Context) {
	var pageInfo NavigationProjectReq.NavigationProjectSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := NPService.GetNavigationProjectInfoList(pageInfo)
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

// GetNavigationProjectPublic 不需要鉴权的导航栏项目接口
// @Tags NavigationProject
// @Summary 不需要鉴权的导航栏项目接口
// @accept application/json
// @Produce application/json
// @Param data query NavigationProjectReq.NavigationProjectSearch true "分页获取导航栏项目列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /NP/getNavigationProjectPublic [get]
func (NPApi *NavigationProjectApi) GetNavigationProjectPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	NPService.GetNavigationProjectPublic()
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的导航栏项目接口信息",
	}, "获取成功", c)
}

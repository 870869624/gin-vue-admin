package PlatformV1

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/PlatformV1"
    PlatformV1Req "github.com/flipped-aurora/gin-vue-admin/server/model/PlatformV1/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type PlatformApi struct {}



// CreatePlatform 创建平台
// @Tags Platform
// @Summary 创建平台
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body PlatformV1.Platform true "创建平台"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /platform/createPlatform [post]
func (platformApi *PlatformApi) CreatePlatform(c *gin.Context) {
	var platform PlatformV1.Platform
	err := c.ShouldBindJSON(&platform)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = platformService.CreatePlatform(&platform)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeletePlatform 删除平台
// @Tags Platform
// @Summary 删除平台
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body PlatformV1.Platform true "删除平台"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /platform/deletePlatform [delete]
func (platformApi *PlatformApi) DeletePlatform(c *gin.Context) {
	ID := c.Query("ID")
	err := platformService.DeletePlatform(ID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeletePlatformByIds 批量删除平台
// @Tags Platform
// @Summary 批量删除平台
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /platform/deletePlatformByIds [delete]
func (platformApi *PlatformApi) DeletePlatformByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := platformService.DeletePlatformByIds(IDs)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdatePlatform 更新平台
// @Tags Platform
// @Summary 更新平台
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body PlatformV1.Platform true "更新平台"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /platform/updatePlatform [put]
func (platformApi *PlatformApi) UpdatePlatform(c *gin.Context) {
	var platform PlatformV1.Platform
	err := c.ShouldBindJSON(&platform)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = platformService.UpdatePlatform(platform)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindPlatform 用id查询平台
// @Tags Platform
// @Summary 用id查询平台
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query PlatformV1.Platform true "用id查询平台"
// @Success 200 {object} response.Response{data=PlatformV1.Platform,msg=string} "查询成功"
// @Router /platform/findPlatform [get]
func (platformApi *PlatformApi) FindPlatform(c *gin.Context) {
	ID := c.Query("ID")
	replatform, err := platformService.GetPlatform(ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(replatform, c)
}

// GetPlatformList 分页获取平台列表
// @Tags Platform
// @Summary 分页获取平台列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query PlatformV1Req.PlatformSearch true "分页获取平台列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /platform/getPlatformList [get]
func (platformApi *PlatformApi) GetPlatformList(c *gin.Context) {
	var pageInfo PlatformV1Req.PlatformSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := platformService.GetPlatformInfoList(pageInfo)
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

// GetPlatformPublic 不需要鉴权的平台接口
// @Tags Platform
// @Summary 不需要鉴权的平台接口
// @accept application/json
// @Produce application/json
// @Param data query PlatformV1Req.PlatformSearch true "分页获取平台列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /platform/getPlatformPublic [get]
func (platformApi *PlatformApi) GetPlatformPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    platformService.GetPlatformPublic()
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的平台接口信息",
    }, "获取成功", c)
}

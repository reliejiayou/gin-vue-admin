package toolsManager

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/toolsManager"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    toolsManagerReq "github.com/flipped-aurora/gin-vue-admin/server/model/toolsManager/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type ToolsManagerApi struct {
}

var toolsManagerService = service.ServiceGroupApp.ToolsmanagerServiceGroup.ToolsManagerService


// CreateToolsManager 创建ToolsManager
// @Tags ToolsManager
// @Summary 创建ToolsManager
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body toolsManager.ToolsManager true "创建ToolsManager"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /toolsManager/createToolsManager [post]
func (toolsManagerApi *ToolsManagerApi) CreateToolsManager(c *gin.Context) {
	var toolsManager toolsManager.ToolsManager
	err := c.ShouldBindJSON(&toolsManager)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := toolsManagerService.CreateToolsManager(toolsManager); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteToolsManager 删除ToolsManager
// @Tags ToolsManager
// @Summary 删除ToolsManager
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body toolsManager.ToolsManager true "删除ToolsManager"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /toolsManager/deleteToolsManager [delete]
func (toolsManagerApi *ToolsManagerApi) DeleteToolsManager(c *gin.Context) {
	var toolsManager toolsManager.ToolsManager
	err := c.ShouldBindJSON(&toolsManager)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := toolsManagerService.DeleteToolsManager(toolsManager); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteToolsManagerByIds 批量删除ToolsManager
// @Tags ToolsManager
// @Summary 批量删除ToolsManager
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ToolsManager"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /toolsManager/deleteToolsManagerByIds [delete]
func (toolsManagerApi *ToolsManagerApi) DeleteToolsManagerByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := toolsManagerService.DeleteToolsManagerByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateToolsManager 更新ToolsManager
// @Tags ToolsManager
// @Summary 更新ToolsManager
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body toolsManager.ToolsManager true "更新ToolsManager"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /toolsManager/updateToolsManager [put]
func (toolsManagerApi *ToolsManagerApi) UpdateToolsManager(c *gin.Context) {
	var toolsManager toolsManager.ToolsManager
	err := c.ShouldBindJSON(&toolsManager)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := toolsManagerService.UpdateToolsManager(toolsManager); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindToolsManager 用id查询ToolsManager
// @Tags ToolsManager
// @Summary 用id查询ToolsManager
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query toolsManager.ToolsManager true "用id查询ToolsManager"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /toolsManager/findToolsManager [get]
func (toolsManagerApi *ToolsManagerApi) FindToolsManager(c *gin.Context) {
	var toolsManager toolsManager.ToolsManager
	err := c.ShouldBindQuery(&toolsManager)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if retoolsManager, err := toolsManagerService.GetToolsManager(toolsManager.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"retoolsManager": retoolsManager}, c)
	}
}

// GetToolsManagerList 分页获取ToolsManager列表
// @Tags ToolsManager
// @Summary 分页获取ToolsManager列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query toolsManagerReq.ToolsManagerSearch true "分页获取ToolsManager列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /toolsManager/getToolsManagerList [get]
func (toolsManagerApi *ToolsManagerApi) GetToolsManagerList(c *gin.Context) {
	var pageInfo toolsManagerReq.ToolsManagerSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := toolsManagerService.GetToolsManagerInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}

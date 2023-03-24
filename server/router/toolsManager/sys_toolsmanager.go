package toolsManager

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ToolsManagerRouter struct {
}

// InitToolsManagerRouter 初始化 ToolsManager 路由信息
func (s *ToolsManagerRouter) InitToolsManagerRouter(Router *gin.RouterGroup) {
	toolsManagerRouter := Router.Group("toolsManager").Use(middleware.OperationRecord())
	toolsManagerRouterWithoutRecord := Router.Group("toolsManager")
	var toolsManagerApi = v1.ApiGroupApp.ToolsmanagerApiGroup.ToolsManagerApi
	{
		toolsManagerRouter.POST("createToolsManager", toolsManagerApi.CreateToolsManager)   // 新建ToolsManager
		toolsManagerRouter.DELETE("deleteToolsManager", toolsManagerApi.DeleteToolsManager) // 删除ToolsManager
		toolsManagerRouter.DELETE("deleteToolsManagerByIds", toolsManagerApi.DeleteToolsManagerByIds) // 批量删除ToolsManager
		toolsManagerRouter.PUT("updateToolsManager", toolsManagerApi.UpdateToolsManager)    // 更新ToolsManager
	}
	{
		toolsManagerRouterWithoutRecord.GET("findToolsManager", toolsManagerApi.FindToolsManager)        // 根据ID获取ToolsManager
		toolsManagerRouterWithoutRecord.GET("getToolsManagerList", toolsManagerApi.GetToolsManagerList)  // 获取ToolsManager列表
	}
}

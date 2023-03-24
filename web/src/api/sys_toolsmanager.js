import service from '@/utils/request'

// @Tags ToolsManager
// @Summary 创建ToolsManager
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ToolsManager true "创建ToolsManager"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /toolsManager/createToolsManager [post]
export const createToolsManager = (data) => {
  return service({
    url: '/toolsManager/createToolsManager',
    method: 'post',
    data
  })
}

// @Tags ToolsManager
// @Summary 删除ToolsManager
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ToolsManager true "删除ToolsManager"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /toolsManager/deleteToolsManager [delete]
export const deleteToolsManager = (data) => {
  return service({
    url: '/toolsManager/deleteToolsManager',
    method: 'delete',
    data
  })
}

// @Tags ToolsManager
// @Summary 删除ToolsManager
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ToolsManager"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /toolsManager/deleteToolsManager [delete]
export const deleteToolsManagerByIds = (data) => {
  return service({
    url: '/toolsManager/deleteToolsManagerByIds',
    method: 'delete',
    data
  })
}

// @Tags ToolsManager
// @Summary 更新ToolsManager
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ToolsManager true "更新ToolsManager"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /toolsManager/updateToolsManager [put]
export const updateToolsManager = (data) => {
  return service({
    url: '/toolsManager/updateToolsManager',
    method: 'put',
    data
  })
}

// @Tags ToolsManager
// @Summary 用id查询ToolsManager
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ToolsManager true "用id查询ToolsManager"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /toolsManager/findToolsManager [get]
export const findToolsManager = (params) => {
  return service({
    url: '/toolsManager/findToolsManager',
    method: 'get',
    params
  })
}

// @Tags ToolsManager
// @Summary 分页获取ToolsManager列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取ToolsManager列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /toolsManager/getToolsManagerList [get]
export const getToolsManagerList = (params) => {
  return service({
    url: '/toolsManager/getToolsManagerList',
    method: 'get',
    params
  })
}

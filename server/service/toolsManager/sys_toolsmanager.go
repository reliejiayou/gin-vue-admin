package toolsManager

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/toolsManager"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    toolsManagerReq "github.com/flipped-aurora/gin-vue-admin/server/model/toolsManager/request"
)

type ToolsManagerService struct {
}

// CreateToolsManager 创建ToolsManager记录
// Author [piexlmax](https://github.com/piexlmax)
func (toolsManagerService *ToolsManagerService) CreateToolsManager(toolsManager toolsManager.ToolsManager) (err error) {
	err = global.GVA_DB.Create(&toolsManager).Error
	return err
}

// DeleteToolsManager 删除ToolsManager记录
// Author [piexlmax](https://github.com/piexlmax)
func (toolsManagerService *ToolsManagerService)DeleteToolsManager(toolsManager toolsManager.ToolsManager) (err error) {
	err = global.GVA_DB.Delete(&toolsManager).Error
	return err
}

// DeleteToolsManagerByIds 批量删除ToolsManager记录
// Author [piexlmax](https://github.com/piexlmax)
func (toolsManagerService *ToolsManagerService)DeleteToolsManagerByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]toolsManager.ToolsManager{},"id in ?",ids.Ids).Error
	return err
}

// UpdateToolsManager 更新ToolsManager记录
// Author [piexlmax](https://github.com/piexlmax)
func (toolsManagerService *ToolsManagerService)UpdateToolsManager(toolsManager toolsManager.ToolsManager) (err error) {
	err = global.GVA_DB.Save(&toolsManager).Error
	return err
}

// GetToolsManager 根据id获取ToolsManager记录
// Author [piexlmax](https://github.com/piexlmax)
func (toolsManagerService *ToolsManagerService)GetToolsManager(id uint) (toolsManager toolsManager.ToolsManager, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&toolsManager).Error
	return
}

// GetToolsManagerInfoList 分页获取ToolsManager记录
// Author [piexlmax](https://github.com/piexlmax)
func (toolsManagerService *ToolsManagerService)GetToolsManagerInfoList(info toolsManagerReq.ToolsManagerSearch) (list []toolsManager.ToolsManager, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&toolsManager.ToolsManager{})
    var toolsManagers []toolsManager.ToolsManager
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Parent_id != nil {
        db = db.Where("parent_id = ?",info.Parent_id)
    }
    if info.Path != "" {
        db = db.Where("path = ?",info.Path)
    }
    if info.Name != "" {
        db = db.Where("name = ?",info.Name)
    }
    if info.Component != "" {
        db = db.Where("component = ?",info.Component)
    }
    if info.Title != "" {
        db = db.Where("title = ?",info.Title)
    }
    if info.Description != "" {
        db = db.Where("description = ?",info.Description)
    }
    if info.Icon != "" {
        db = db.Where("icon = ?",info.Icon)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&toolsManagers).Error
	return  toolsManagers, total, err
}

// 自动生成模板ToolsManager
package toolsManager

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// ToolsManager 结构体
type ToolsManager struct {
      global.GVA_MODEL
      Parent_id  *int `json:"parent_id" form:"parent_id" gorm:"column:parent_id;comment:父工具id;"`
      Path  string `json:"path" form:"path" gorm:"column:path;comment:工具路径;"`
      Name  string `json:"name" form:"name" gorm:"column:name;comment:工具名;"`
      Component  string `json:"component" form:"component" gorm:"column:component;comment:组件;"`
      Title  string `json:"title" form:"title" gorm:"column:title;comment:标题;"`
      Description  string `json:"description" form:"description" gorm:"column:description;comment:工具描述;"`
      Icon  string `json:"icon" form:"icon" gorm:"column:icon;comment:图标;"`
}


// TableName ToolsManager 表名
func (ToolsManager) TableName() string {
  return "tools_list"
}


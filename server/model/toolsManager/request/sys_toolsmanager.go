package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/toolsManager"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type ToolsManagerSearch struct{
    toolsManager.ToolsManager
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    request.PageInfo
}

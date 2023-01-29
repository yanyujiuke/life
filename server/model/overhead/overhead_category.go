// 自动生成模板OverheadCategory
package overhead

import (
	"life/global"
)

// OverheadCategory 结构体
type OverheadCategory struct {
	global.GVA_MODEL
	Name string `json:"name" form:"name" gorm:"column:name;comment:名称;size:32;"`
	Desc string `json:"desc" form:"desc" gorm:"column:desc;comment:描述;"`
}

// TableName OverheadCategory 表名
func (OverheadCategory) TableName() string {
	return "overhead_category"
}

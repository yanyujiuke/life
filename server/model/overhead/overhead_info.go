// 自动生成模板OverheadInfo
package overhead

import (
	"github.com/songzhibin97/gkit/timeout"
	"life/global"
)

// OverheadInfo 结构体
type OverheadInfo struct {
	global.GVA_MODEL
	Uid        *int         `json:"uid" form:"uid" gorm:"column:uid;comment:用户id;"`
	Day        timeout.Date `json:"day" form:"day" gorm:"column:day;comment:创建日期;"`
	CategoryId *int         `json:"category_id" form:"category_id" gorm:"column:category_id;comment:分类id;"`
	Name       string       `json:"name" form:"name" gorm:"column:name;comment:名称;size:32;"`
	Amount     *float64     `json:"amount" form:"amount" gorm:"column:amount;comment:消费金额;"`
	Desc       string       `json:"desc" form:"desc" gorm:"column:desc;comment:描述;"`
}

// TableName OverheadInfo 表名
func (OverheadInfo) TableName() string {
	return "overhead_info"
}

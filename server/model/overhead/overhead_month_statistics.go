// 自动生成模板OverheadMonthStatistics
package overhead

import (
	"life/global"
	"time"
)

// OverheadMonthStatistics 结构体
type OverheadMonthStatistics struct {
	global.GVA_MODEL
	Uid    *int       `json:"uid" form:"uid" gorm:"column:uid;comment:用户id;"`
	Month  time.Month `json:"month" form:"month" gorm:"column:month;comment:月份;"`
	Amount *float64   `json:"amount" form:"amount" gorm:"column:amount;comment:每月开销金额;"`
	Num    *int       `json:"num" form:"num" gorm:"column:num;comment:每月开销单数;"`
	Rank   *int       `json:"rank" form:"rank" gorm:"column:rank;comment:每月消费金额排名;"`
	Desc   string     `json:"desc" form:"desc" gorm:"column:desc;comment:描述;"`
}

// TableName OverheadMonthStatistics 表名
func (OverheadMonthStatistics) TableName() string {
	return "overhead_month_statistics"
}

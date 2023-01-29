// 自动生成模板OverheadDailyStatistics
package overhead

import (
	"github.com/songzhibin97/gkit/timeout"
	"life/global"
)

// OverheadDailyStatistics 结构体
type OverheadDailyStatistics struct {
	global.GVA_MODEL
	Uid    *int         `json:"uid" form:"uid" gorm:"column:uid;comment:用户id;"`
	Day    timeout.Date `json:"day" form:"day" gorm:"column:day;comment:开销日期;"`
	Amount *float64     `json:"amount" form:"amount" gorm:"column:amount;comment:每日开销金额;"`
	Num    *int         `json:"num" form:"num" gorm:"column:num;comment:每日消费单数;"`
	Rank   *int         `json:"rank" form:"rank" gorm:"column:rank;comment:消费金额排名;"`
	Desc   string       `json:"desc" form:"desc" gorm:"column:desc;comment:描述;"`
}

// TableName OverheadDailyStatistics 表名
func (OverheadDailyStatistics) TableName() string {
	return "overhead_daily_statistics"
}

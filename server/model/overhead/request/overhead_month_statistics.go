package request

import (
	"life/model/common/request"
	"life/model/overhead"
	"time"
)

type OverheadMonthStatisticsSearch struct {
	overhead.OverheadMonthStatistics
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}

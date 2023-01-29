package request

import (
	"life/model/common/request"
	"life/model/overhead"
	"time"
)

type OverheadDailyStatisticsSearch struct {
	overhead.OverheadDailyStatistics
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}

package request

import (
	"life/model/common/request"
	"life/model/overhead"
	"time"
)

type OverheadInfoSearch struct {
	overhead.OverheadInfo
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}

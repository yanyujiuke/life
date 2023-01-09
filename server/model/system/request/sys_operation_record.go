package request

import (
	"life/model/common/request"
	"life/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}

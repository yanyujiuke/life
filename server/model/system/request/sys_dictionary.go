package request

import (
	"life/model/common/request"
	"life/model/system"
)

type SysDictionarySearch struct {
	system.SysDictionary
	request.PageInfo
}

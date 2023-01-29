package v1

import (
	"life/api/v1/example"
	"life/api/v1/overhead"
	"life/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup   system.ApiGroup
	ExampleApiGroup  example.ApiGroup
	OverheadApiGroup overhead.ApiGroup
}

var ApiGroupApp = new(ApiGroup)

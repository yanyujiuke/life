package service

import (
	"life/service/example"
	"life/service/overhead"
	"life/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup   system.ServiceGroup
	ExampleServiceGroup  example.ServiceGroup
	OverheadServiceGroup overhead.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)

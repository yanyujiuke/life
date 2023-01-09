package initialize

import (
	_ "life/source/example"
	_ "life/source/system"
)

func init() {
	// do nothing,only import source package so that inits can be registered
}

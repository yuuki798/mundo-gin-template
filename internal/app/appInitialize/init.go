package appInitialize

import (
	"github.com/trancecho/mundo-be-template/internal/app"
)

var (
	apps = make([]app.Module, 0)
)

func GetApps() []app.Module {
	return apps
}

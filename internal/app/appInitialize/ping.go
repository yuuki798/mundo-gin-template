package appInitialize

import "github.com/trancecho/mundo-be-template/internal/app/ping"

func init() {
	apps = append(apps, &ping.Ping{Name: "ping module"})
}

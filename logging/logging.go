package logging

import "fmt"

var mkStr = fmt.Sprintf

func SystemStart(poweredOnTime, poweredOnDate string) string {
	return mkStr("System powered on at: %s on %s\n", poweredOnTime, poweredOnDate)
}

func SystemShutdown(event, shutdownTime, date string) string {
	return mkStr("System %s at: %s on %s\n", event, shutdownTime, date)
}

func SystemUptime(uptime string) string {
	return mkStr("System has been %s\n", uptime)
}

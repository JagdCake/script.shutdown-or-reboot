package logging

import "fmt"

var mkStr = fmt.Sprintf

func SystemStart(poweredOnTime, poweredOnDate string) string {
	return mkStr("System powered on at: %s on %s\n", poweredOnTime, poweredOnDate)
}

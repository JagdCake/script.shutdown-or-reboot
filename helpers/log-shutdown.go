package helpers

import (
	"github.com/jagdcake/shutdown-or-reboot/execute"
	"github.com/jagdcake/shutdown-or-reboot/logging"
	"log"
	"strings"
)

func systemStart() (startTime, startDate string, err error) {
	cmd, err := execute.Command("uptime", "-s")

	var trimmedResult string = strings.TrimSpace(cmd.String())
	var ss []string = strings.Split(trimmedResult, " ")

	return ss[1], ss[0], err
}

func LogShutdown() (shutdownLogged bool) {
	var err error
	//
	var startTime, startDate string
	startTime, startDate, err = systemStart()
	if err != nil {
		log.Fatal(err)
	}

	var systemStart string = logging.SystemStart(startTime, startDate)
	println(systemStart)

	return true
}

package helpers

import (
	"fmt"
	"github.com/jagdcake/shutdown-or-reboot/execute"
	"github.com/jagdcake/shutdown-or-reboot/logging"
	"log"
	"strings"
	"time"
)

func systemStart() (startTime, startDate string, err error) {
	cmd, err := execute.Command("uptime", "-s")

	var trimmedResult string = strings.TrimSpace(cmd.String())
	var ss []string = strings.Split(trimmedResult, " ")

	return ss[1], ss[0], err
}

func shutdownTime(currentTime time.Time, timeToShutdown string) (string, error) {
	var shutdownDur time.Duration

	shutdownDur, err := time.ParseDuration(timeToShutdown)
	st := currentTime.Add(shutdownDur).Format("15:04:05")

	return st, err
}

func shutdownDate(currentTime time.Time) string {
	sd := fmt.Sprintf("%d-%02d-%02d", // "%02d" adds a 0 as a prefix if the number has less than 2 digits
		currentTime.Year(),
		currentTime.Month(),
		currentTime.Day(),
	)

	return sd
}

func LogShutdown(event, timeToShutdown string) (shutdownLogged bool) {
	var err error
	//
	var startTime, startDate string
	startTime, startDate, err = systemStart()
	if err != nil {
		log.Fatal(err)
	}

	var systemStart string = logging.SystemStart(startTime, startDate)
	println(systemStart)
	//
	var st, sd string
	var currentTime time.Time
	currentTime = time.Now()
	st, err = shutdownTime(currentTime, timeToShutdown)
	if err != nil {
		log.Fatal(err)
	}
	sd = shutdownDate(currentTime)

	var systemShutdown string = logging.SystemShutdown(event, st, sd)
	println(systemShutdown)

	return true
}

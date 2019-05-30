package helpers

import (
	"fmt"
	"github.com/jagdcake/shutdown-or-reboot/execute"
	"github.com/jagdcake/shutdown-or-reboot/logging"
	"github.com/jagdcake/shutdown-or-reboot/write"
	"log"
	"regexp"
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

func lessThanMinUptime(uptime string) (bool, error) {
	// if the system has been up for less than a minute, the uptime command outputs only "up", no numbers
	numbersInUptime, err := regexp.MatchString(`\d+`, uptime)
	if err != nil || numbersInUptime {
		return false, err
	}

	return true, err
}

func systemUptime() string {
	cmd, _ := execute.Command("uptime", "-p")

	var cmdResult string = cmd.String()
	trimmedResult := strings.TrimSpace(cmdResult)

	return trimmedResult
}

func LogShutdown(event, timeToShutdown, logFile string) (shutdownLogged bool) {
	var err error
	//
	var startTime, startDate string
	startTime, startDate, err = systemStart()
	if err != nil {
		log.Fatal(err)
	}

	var systemStart string = logging.SystemStart(startTime, startDate)
	err = write.String(systemStart, logFile)
	if err != nil {
		log.Fatal(err)
	}
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
	err = write.String(systemShutdown, logFile)
	if err != nil {
		log.Fatal(err)
	}
	//
	var uptime string = systemUptime()
	lessThanMin, err := lessThanMinUptime(uptime)
	if err != nil {
		log.Fatal(err)
	}

	if lessThanMin {
		uptime = "up for less than a minute"
	}
	var systemUptime string = logging.SystemUptime(uptime)

	err = write.String(systemUptime+"\n", logFile)
	if err != nil {
		log.Fatal(err)
	}

	return true
}

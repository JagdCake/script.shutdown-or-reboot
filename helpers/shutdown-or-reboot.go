package helpers

import (
	"github.com/jagdcake/script.shutdown-or-reboot/execute"
	"log"
	"time"
)

func sleepDuration(timeToSleep string) (time.Duration, error) {
	sd, err := time.ParseDuration(timeToSleep)
	if err != nil {
		return time.Duration(0), err
	}

	return sd, err
}

func ShutdownOrReboot(timeToShutdown, shutdownFlag string) {
	var err error
	//
	sd, err := sleepDuration(timeToShutdown)
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(sd)
	//
	_, err = execute.Command("shutdown", shutdownFlag, "now")
	if err != nil {
		log.Fatal(err)
	}
}

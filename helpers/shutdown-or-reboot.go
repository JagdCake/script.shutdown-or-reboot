package helpers

import (
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

func ShutdownOrReboot(timeToShutdown string) {
	var err error
	//
	sd, err := sleepDuration(timeToShutdown)
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(sd)
}

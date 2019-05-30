package helpers

import (
	"time"
)

func sleepDuration(timeToSleep string) (time.Duration, error) {
	sd, err := time.ParseDuration(timeToSleep)
	if err != nil {
		return time.Duration(0), err
	}

	return sd, err
}

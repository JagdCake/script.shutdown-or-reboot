package main

import (
	"fmt"
	"github.com/jagdcake/script.shutdown-or-reboot/config"
	"github.com/jagdcake/script.shutdown-or-reboot/helpers"
)

func main() {
	menuOptions := []string{
		"Shutdown",
		"Reboot",
		"Cancel",
	}
	var choice int = helpers.SelectMenu(menuOptions)

	var event string
	var timeToShutdown string
	var shutdownFlag string

	switch choice {
	case 1:
		event = "shutdown"
		timeToShutdown = config.TIME_TO_SHUTDOWN
		shutdownFlag = "-P"
	case 2:
		event = "reboot"
		timeToShutdown = config.TIME_TO_REBOOT
		shutdownFlag = "-r"
	case 3:
		return
	default:
		fmt.Println("Please enter option number!")
		main()
		return
	}

	var closedApps string = helpers.CloseApps(config.AppsToClose)
	fmt.Println(closedApps)

	var shutdownLogged bool = helpers.LogShutdown(event, timeToShutdown, config.LOG_FILE)
	if shutdownLogged {
		fmt.Println("Shutdown has been logged in", config.LOG_FILE)
	}

	helpers.ShutdownOrReboot(timeToShutdown, shutdownFlag)
}

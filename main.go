package main

import (
	"fmt"
	"github.com/jagdcake/shutdown-or-reboot/config"
	"github.com/jagdcake/shutdown-or-reboot/helpers"
)

func main() {
	menuOptions := []string{
		"Shutdown",
		"Reboot",
		"Cancel",
	}
	var choice int = helpers.SelectMenu(menuOptions)

	println(choice)

	var closedApps string = helpers.CloseApps(config.AppsToClose)
	fmt.Println(closedApps)
}

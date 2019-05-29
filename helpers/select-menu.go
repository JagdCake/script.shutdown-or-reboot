package helpers

import (
	"fmt"
	"github.com/jagdcake/shutdown-or-reboot/menu"
)

func printMenu(menuOptions []string) {
	sm := menu.CreateWith(menuOptions)

	fmt.Print(sm)
}

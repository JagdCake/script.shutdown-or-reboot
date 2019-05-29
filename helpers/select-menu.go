package helpers

import (
	"fmt"
	"github.com/jagdcake/shutdown-or-reboot/menu"
	"strconv"
)

func printMenu(menuOptions []string) {
	sm := menu.CreateWith(menuOptions)

	fmt.Print(sm)
}

func userChoice() int {
	var input string
	var choice int

	input = menu.Prompt()
	choice, _ = strconv.Atoi(input)

	return choice
}

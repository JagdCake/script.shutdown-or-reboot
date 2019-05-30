package helpers

import (
	"fmt"
	"github.com/jagdcake/script.shutdown-or-reboot/menu"
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
	// no need to handle string conversion errors, the switch statement in ../main.go is enough
	choice, _ = strconv.Atoi(input)

	return choice
}

func SelectMenu(menuOptions []string) int {
	printMenu(menuOptions)
	//
	choice := userChoice()

	return choice
}

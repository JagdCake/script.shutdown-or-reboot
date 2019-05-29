package menu

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func CreateWith(options []string) string {
	var formattedOpt string
	var formattedOpts []string
	var prompt string

	for i, option := range options {
		formattedOpt = fmt.Sprintf("%v) %s", i+1, option) // e.g. "2) Open"
		formattedOpts = append(formattedOpts, formattedOpt)
	}
	prompt = fmt.Sprintf("# ")
	formattedOpts = append(formattedOpts, prompt)

	menuOpts := strings.Join(formattedOpts, "\n")

	return menuOpts
}

func Prompt() string {
	var input string

	reader := bufio.NewScanner(os.Stdin)

	reader.Scan()
	input = reader.Text()

	return input
}

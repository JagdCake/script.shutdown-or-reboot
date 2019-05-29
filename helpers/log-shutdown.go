package helpers

import (
	"github.com/jagdcake/shutdown-or-reboot/execute"
	"strings"
)

func systemStart() (startTime, startDate string, err error) {
	cmd, err := execute.Command("uptime", "-s")

	var trimmedResult string = strings.TrimSpace(cmd.String())
	var ss []string = strings.Split(trimmedResult, " ")

	return ss[1], ss[0], err
}

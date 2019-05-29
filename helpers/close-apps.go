package helpers

import (
	"fmt"
	"github.com/jagdcake/shutdown-or-reboot/execute"
	"strings"
)

func CloseApps(apps map[string]string) string {
	var closedApp string
	var closedApps []string

	for name, process := range apps {
		_, err := execute.Command("pkill", process)
		if err == nil {
			closedApp = fmt.Sprintf("Closing %s", name)
			closedApps = append(closedApps, closedApp)
		}
	}

	return strings.Join(closedApps, "\n")
}

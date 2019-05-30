package config

var AppsToClose = map[string]string{
	// process name is "transmission-gtk" but calling pkill tab-completes without the "k"
	"Transmission": "transmission-gt",
	"Firefox":      "firefox",
	"Steam":        "steam",
}

const (
	TIME_TO_SHUTDOWN string = "10s"
	TIME_TO_REBOOT   string = "5s"
	LOG_FILE         string = "/home/jagdcake/Documents/text_files/my_logs/shutdown_log"
)

package config

var AppsToClose = map[string]string{
	"Transmission": "transmission-gtk",
	"Firefox":      "firefox",
	"Steam":        "steam",
}

const (
	TIME_TO_SHUTDOWN string = "10s"
	TIME_TO_REBOOT   string = "5s"
	LOG_FILE         string = "/home/jagdcake/Documents/text_files/my_logs/shutdown_log"
)

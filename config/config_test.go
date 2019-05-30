package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLogFileExists(t *testing.T) {
	logFileDir := filepath.Dir(LOG_FILE)
	_, err := os.Stat(logFileDir)

	if err != nil {
		t.Errorf("Directory: %s doesn't exist", logFileDir)
	}
}

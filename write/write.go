package write

import "os"

func String(s string, file string) error {
	var f *os.File

	f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	_, err = f.WriteString(s)

	return err
}

package execute

import "os/exec"

func Command(command string, argument ...string) (Result, error) {
	var cmd *exec.Cmd
	cmd = exec.Command(command, argument...)

	var result Result
	var err error
	result.output, err = cmd.Output()

	return result, err
}

package utils

import "os/exec"

//
// Run command and return response as string
//
func RunCommand(command string, args ...string) string  {
	output, err := exec.Command(command, args...).Output()
	if err != nil {
		return err.Error()
	}
	return string(output[:])
}

//
// Run command and return response as string
//
func RunShellCommand(command string) string  {
	return RunCommand("sh", "-c", command)
}
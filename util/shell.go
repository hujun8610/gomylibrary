package util

import "os/exec"

func ExecCmd(cmdString string) (string, error) {
	cmd := exec.Command("bash", "-c", cmdString)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

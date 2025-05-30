package connection

import (
	"os/exec"
)

type LocalConnection struct{}

func (c LocalConnection) RunCommand(cmd string) (string, error) {
	out, err := exec.Command("bash", "-c", cmd).CombinedOutput()
	return string(out), err
}

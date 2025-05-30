package module

import (
	"github.com/mabou-dev/go-ansible/pkg/connection"
)

type CommandModule struct{}

func (m CommandModule) Execute(c connection.Connection, args map[string]string) (string, error) {
	cmd := args["cmd"]
	return c.RunCommand(cmd)
}

func init() {
	Register("command", func() Module {
		return CommandModule{}
	})
}

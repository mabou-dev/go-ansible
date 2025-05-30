package module

import (
	"github.com/mabou-dev/go-ansible/pkg/connection"
)

type Module interface {
	Execute(connection.Connection, map[string]string) (string, error)
}

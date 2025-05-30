package connection

type Connection interface {
	RunCommand(cmd string) (string, error)
}

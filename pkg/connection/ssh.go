package connection

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
)

const (
	protocol = "tcp"
	port     = "22"
)

type SSHConnection struct {
	Client *ssh.Client
}

func NewSSHConnection(user string, ip string, keyPath string) (*SSHConnection, error) {
	key, err := ioutil.ReadFile(keyPath)
	if err != nil {
		return nil, err
	}

	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return nil, err
	}

	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	addr := fmt.Sprintf("%s:%s", ip, port)
	client, err := ssh.Dial(protocol, addr, config)
	if err != nil {
		return nil, err
	}

	return &SSHConnection{
		Client: client,
	}, nil
}

func (c *SSHConnection) RunCommand(cmd string) (string, error) {
	session, err := c.Client.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()

	out, err := session.CombinedOutput(cmd)
	return string(out), err
}

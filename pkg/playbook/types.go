package playbook

type Playbook struct {
	Connection Connection `yaml:"connection"`
	Tasks      []Task     `yaml:"tasks"`
}

type Connection struct {
	Address  string `yaml:"address"`
	Username string `yaml:"username"`
	Key      string `yaml:"key"`
}

type Task struct {
	Name   string            `yaml:"name"`
	Module string            `yaml:"module"`
	Args   map[string]string `yaml:"args"`
}

// Structure of a playbook
// ```
// connection:
//     type: local | ssh
//     address: <ip> | localhost
//     username: <username>
//     key: /home/user/.ssh/id_rsa
//
// tasks:
//     - name: hello world
//       module: command
//       args:
//           cmd: echo hello world
// ```

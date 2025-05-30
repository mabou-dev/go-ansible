package execute

import (
    "fmt"

	"github.com/mabou-dev/go-ansible/pkg/module"
	"github.com/mabou-dev/go-ansible/pkg/playbook"
	"github.com/mabou-dev/go-ansible/pkg/connection"
)

func RunPlaybook(pb *playbook.Playbook, c connection.Connection) error {
	for _, task := range pb.Tasks {
		fmt.Printf("running: %s\n", task.Name)

		mod, err := module.Get(task.Module)
		if err != nil {
			if _, ok := err.(*module.ModuleNotFoundError); ok {
				return fmt.Errorf("error: %v", err)
			}
			return fmt.Errorf("unexpected error: %v", err)
		}

		out, err := mod.Execute(c, task.Args)
		if err != nil {
			return fmt.Errorf("error executing module %s: %v", task.Module, err)
		}

		fmt.Printf("output: %s\n", out)
	}
        return nil
}

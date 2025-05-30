package cmd

import (
	"github.com/spf13/cobra"

	"github.com/mabou-dev/go-ansible/pkg/connection"
	"github.com/mabou-dev/go-ansible/pkg/playbook"
	"github.com/mabou-dev/go-ansible/pkg/execute"
)

var debugTaskCmd = newDebugTaskCmd()

func newDebugTaskCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "debug-task",
		Short: "test all workflow without connection",
		RunE:  debugTaskRunE,
	}

	cmd.Flags().StringVar(&playbookPath, "playbook", "", "path of the playbook")

	return cmd
}

func debugTaskRunE(cmd *cobra.Command, args []string) error {
	pb, err := playbook.Load(playbookPath)
	if err != nil {
		return err
	}

	c := &connection.LocalConnection{}

	return execute.RunPlaybook(pb, c)
}

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/mabou-dev/go-ansible/pkg/playbook"
)

var debugParserCmd = newDebugParserCmd()

var playbookPath string

func newDebugParserCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "debug-parser",
		Short: "print information to check that code is working",
		RunE:  debugParserRunE,
	}

	cmd.Flags().StringVar(&playbookPath, "playbook", "", "path of the playbook")

	return cmd
}

func debugParserRunE(cmd *cobra.Command, args []string) error {
	pb, err := playbook.Load(playbookPath)
	if err != nil {
		return err
	}

	fmt.Printf("%+v", pb)
	return nil
}

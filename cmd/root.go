package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = newRootCmd()

func newRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "go-ansible",
		Short: "Ansible written in golang",
	}

	cmd.AddCommand(helloCmd)
	cmd.AddCommand(debugParserCmd)
	cmd.AddCommand(debugTaskCmd)

	return cmd
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

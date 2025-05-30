package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var helloCmd = newHelloCmd()

func newHelloCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "hello",
		Short: "say hello",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("hello")
		},
	}

	return cmd
}

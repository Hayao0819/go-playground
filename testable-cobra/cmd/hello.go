package cmd

import (
	"github.com/spf13/cobra"
)

func newHelloCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "hello",
		Short: "Print hello world",
		Long:  "Print hello world",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Println("Hello world!")
			return nil
		},
	}
	return &cmd
}

func init() {
	addSubCmd(newHelloCmd)
}

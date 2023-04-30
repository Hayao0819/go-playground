package cmd

import (
	"github.com/spf13/cobra"
)

func newNameCmd() *cobra.Command {
	name := "john"

	cmd := cobra.Command{
		Use:   "name",
		Short: "say hello and your name",
		Long:  "say hello and your name",
		RunE: func(cmd *cobra.Command, args []string) error {
			RunSubCmdFromCmd("hello", cmd)
			cmd.Println(name)
			return nil
		},
	}

	cmd.Flags().StringVarP(&name, "name", "", name, "your name")
	return &cmd
}

func init() {
	addSubCmd(newNameCmd)
}

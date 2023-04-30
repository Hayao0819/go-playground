package cmd

import "github.com/spf13/cobra"

func rootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "myapp",
		Short: "myapp is a testable cobra command",
		Long:  "myapp is a testable cobra command",
		/*
			Run: func(cmd *cobra.Command, args []string){

			},
		*/
	}

	// add sub command
	for _, f := range subCmdList {
		c := f()
		c.SetOut(cmd.OutOrStdout())
		cmd.AddCommand(c)
	}

	// add global flags here
	//myname := cmd.PersistentFlags().StringP("name", "", "", "your name")

	return cmd
}

func Execute() error {
	cmd := rootCmd()
	return cmd.Execute()
}

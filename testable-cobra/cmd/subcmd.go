package cmd

import (
	"github.com/spf13/cobra"
	"io"
)

var subCmdList = map[string]func() *cobra.Command{}

func addSubCmd(f func() *cobra.Command) {
	subCmdList[f().Name()] = f
}

// 名前からコマンドを取得する
func GetSubCmd(name string) func() *cobra.Command {
	return subCmdList[name]
}

// コマンド内から別のコマンドを実行する
func RunSubCmdFromCmd(name string, cmd *cobra.Command, args ...string) error {
	subcmdfunc := GetSubCmd(name)
	if subcmdfunc == nil {
		return nil
	}
	subcmd := subcmdfunc()
	subcmd.SetArgs(args)
	subcmd.SetOut(cmd.OutOrStdout())
	return subcmd.Execute()
}

// 出力を指定してコマンドを実行する
func RunSubCmdWithIO(name string, stdout, stderr io.Writer, args ...string) error {
	cmdfunc := GetSubCmd(name)
	if cmdfunc == nil {
		return nil
	}
	cmd := cmdfunc()

	// argsのフラグをcmdに渡す
	cmd.SetArgs(args)

	// 出力を指定
	cmd.SetOut(stdout)
	cmd.SetErr(stderr)
	return cmd.Execute()
}

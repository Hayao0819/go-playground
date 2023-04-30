package cmd_test

import (
	"bytes"
	"testing"
	"github.com/Hayao0819/hayao-go-playground/testable-cobra/cmd"
)

func TestHelloCmd(t *testing.T) {
	// 入出力 (io.Writer)
	stdout := bytes.NewBuffer(nil)
	stderr := bytes.NewBuffer(nil)
	cmd.RunSubCmdWithIO("hello", stdout, stderr)

	if stdout.String() != "Hello world!\n" {
		t.Errorf("unexpected output: %s", stdout.String())
	}
}


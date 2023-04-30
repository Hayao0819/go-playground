package cmd_test

import (
	"bytes"
	"testing"
	"github.com/Hayao0819/hayao-go-playground/testable-cobra/cmd"
)



func TestNameCmd(t *testing.T) {
	stdout := bytes.NewBuffer(nil)
	stderr := bytes.NewBuffer(nil)

	// no arg
	cmd.RunSubCmdWithIO("name", stdout, stderr)
	if stdout.String() != "Hello world!\njohn\n" {
		t.Errorf("unexpected output: %s", stdout.String())
	}

	// with arg
	names := []string{"mike", "tom", "jane"}
	for _, name := range names {
		stdout.Reset()
		stderr.Reset()
		cmd.RunSubCmdWithIO("name", stdout, stderr, "--name", name)

		if stdout.String() != "Hello world!\n"+name+"\n" {
			t.Errorf("unexpected output: %s", stdout.String())
		}
	}
}

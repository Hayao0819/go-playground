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
	err := cmd.RunSubCmdWithIO("name", stdout, stderr)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if stdout.String() != "Hello world!\njohn\n" {
		t.Errorf("unexpected output: %s", stdout.String())
	}

	// with arg
	names := []string{"mike", "tom", "jane"}
	for _, name := range names {
		stdout.Reset()
		stderr.Reset()
		err := cmd.RunSubCmdWithIO("name", stdout, stderr, "--name", name)
		if err != nil{
			t.Errorf("unexpected error: %s", err)
		}

		if stdout.String() != "Hello world!\n"+name+"\n" {
			t.Errorf("unexpected output: %s", stdout.String())
		}
	}
}

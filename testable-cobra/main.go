package main

import (
	"os"

	"github.com/Hayao0819/hayao-go-playground/testable-cobra/cmd"
)

func main() {
	if cmd.Execute() != nil {
		os.Exit(-1)
	}
}

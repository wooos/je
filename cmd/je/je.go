package main

import (
	"os"
)

func main() {
	cmd := newRootCmd(os.Stdout)

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

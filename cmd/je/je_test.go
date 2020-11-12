package main

import (
	"bytes"
	"testing"

	shellwords "github.com/mattn/go-shellwords"
	"github.com/spf13/cobra"
	"github.com/wooos/je/internal/test"
)

type cmdTestCase struct {
	name   string
	cmd    string
	golden string
	repeat int
}

func runTestCmd(t *testing.T, tests []cmdTestCase) {
	t.Helper()
	for _, tt := range tests {
		_, out, err := executeActionCommandC(tt.cmd)

		if err != nil {
			t.Errorf("expected error, got '%v'", err)
		}

		if tt.golden != "" {
			test.AssertGoldenString(t, out, tt.golden)
		}
	}
}

func executeActionCommandC(cmd string) (*cobra.Command, string, error) {
	args, err := shellwords.Parse(cmd)
	if err != nil {
		return nil, "", err
	}

	buf := new(bytes.Buffer)

	root := newRootCmd(buf)
	root.SetOut(buf)
	root.SetArgs(args)

	c, err := root.ExecuteC()

	return c, buf.String(), err
}

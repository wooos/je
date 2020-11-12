package require

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

type testCase struct {
	args         []string
	validateFunc cobra.PositionalArgs
	wantError    string
}

func TestArgs(t *testing.T) {
	tests := []testCase{{
		validateFunc: NoArgs,
	}, {
		validateFunc: NoArgs,
		args:         []string{"one"},
		wantError:    `"root" accepts no arguments`,
	}, {
		validateFunc: ExactArgs(1),
		args:         []string{"one"},
	}, {
		validateFunc: ExactArgs(1),
		wantError:    `"root" requires 1 argument`,
	}}

	runTestCase(t, tests)
}

func runTestCase(t *testing.T, tests []testCase) {
	for i, tc := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			cmd := &cobra.Command{
				Use:  "root",
				Run:  func(*cobra.Command, []string) {},
				Args: tc.validateFunc,
			}

			cmd.SetOut(ioutil.Discard)
			cmd.SetArgs(tc.args)

			err := cmd.Execute()

			if tc.wantError == "" {
				if err != nil {
					t.Errorf("unexpected error, got '%v'", err)
				}
				return
			}

			if !strings.Contains(err.Error(), tc.wantError) {
				t.Errorf("unexpected error \n\nEXPECT:\n%q\n\nACTUAL:\n%q\n", tc.wantError, err)
			}

			if !strings.Contains(err.Error(), "Usage:") {
				t.Fatalf("unexpected error: want Usage string\n\nGOT:\n%q\n", err)
			}
		})
	}
}

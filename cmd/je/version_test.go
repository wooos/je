package main

import (
	"testing"
)

func TestVersion(t *testing.T) {
	tests := []cmdTestCase{{
		name:   "default",
		cmd:    "version",
		golden: "output/version.txt",
	}, {
		name:   "short",
		cmd:    "version --short",
		golden: "output/version-short.txt",
	}, {
		name:   "template",
		cmd:    "version --template='Version: {{.Version}}'",
		golden: "output/version-template.txt",
	}}

	runTestCmd(t, tests)
}

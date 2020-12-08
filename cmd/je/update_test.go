package main

import (
	"testing"
)

func TestUpdate(t *testing.T) {
	tests := []cmdTestCase{{
		name:   "default",
		cmd:    "update --set name=test testdata/input/update.txt --dry-run",
		golden: "output/update.txt",
	},{
		name:   "map",
		cmd:    "update --set detail.author=test testdata/input/update.txt --dry-run",
		golden: "output/update-map.txt",
	},{
		name:   "array",
		cmd:    "update --set detail.vendor[0]=test testdata/input/update.txt --dry-run",
		golden: "output/update-array.txt",
	}}

	runTestCmd(t, tests)
}

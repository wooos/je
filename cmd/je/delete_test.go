package main

import "testing"

func TestDelete(t *testing.T) {
	tests := []cmdTestCase{{
		name:   "default",
		cmd:    "delete --keys name --dry-run testdata/input/delete.txt",
		golden: "output/delete.txt",
	},{
		name:   "multiple",
		cmd:    "delete --keys name,version --dry-run testdata/input/delete.txt",
		golden: "output/delete-multiple.txt",
	},{
		name:   "multiple-1",
		cmd:    "delete --keys name,details.email --dry-run testdata/input/delete.txt",
		golden: "output/delete-multiple-1.txt",
	},{
		name:   "multiple-2",
		cmd:    "delete --keys name --keys details.email --dry-run testdata/input/delete.txt",
		golden: "output/delete-multiple-2.txt",
	}}

	runTestCmd(t, tests)
}
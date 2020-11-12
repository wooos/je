package test

import (
	"bytes"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/pkg/errors"
)

// AssertGoldenString asserts that the given string matches the contents of the given file.
func AssertGoldenString(t *testing.T, actual, filename string) {
	t.Helper()

	if err := compare([]byte(actual), path(filename)); err != nil {
		t.Fatalf("%v", err)
	}
}

func compare(actual []byte, filename string) error {
	expected, err := ioutil.ReadFile(filename)

	if err != nil {
		return errors.Errorf("unable to read testdata %s", filename)
	}

	if !bytes.Equal(expected, actual) {
		return errors.Errorf("does not match golden file %s\n\nEXPECT: \n'%s'\nACTUAL: \n'%s'\n", filename, expected, actual)
	}

	return nil
}

func path(filename string) string {
	if filepath.IsAbs(filename) {
		return filename
	}
	return filepath.Join("testdata", filename)
}

package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/wooos/je/cmd/je/require"
	"github.com/wooos/je/pkg/parser"
)

const updateDesc = `
This command update a json file.

`

var (
	setArrays []string
)

func newUpdateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update [FILENAME]",
		Short: "update a json file",
		Args:  require.ExactArgs(1),
		Long:  updateDesc,
		RunE:  runUpdate,
	}

	addUpdateFlags(cmd.Flags())

	return cmd
}

func addUpdateFlags(f *pflag.FlagSet) {
	f.StringSliceVar(&setArrays, "set", []string{}, "set values on the command line (can specify multiple or separate values with commas: key1=val1,key2=val2)")
}

func runUpdate(cmd *cobra.Command, args []string) error {
	filename := args[0]

	currentMap := map[string]interface{}{}
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(bytes, &currentMap); err != nil {
		return err
	}

	for _, val := range setArrays {
		if err := parser.ParseInto(val, currentMap); err != nil {
			return err
		}
	}

	data, err := json.MarshalIndent(currentMap, "", "  ")
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(filename, data, os.ModePerm); err != nil {
		return err
	}
	return nil
}

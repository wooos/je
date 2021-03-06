package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/wooos/je/cmd/je/require"
	"github.com/wooos/je/pkg/parser"
)

const updateDesc = `
This command update a json file.

`

type updateOptions struct {
	setArrays []string
	dryRun bool
}

var (
	PreBytes = []byte(`{"je":`)
	SubBytes = []byte(`}`)
)

func newUpdateCmd(out io.Writer) *cobra.Command {
	o := updateOptions{}
	cmd := &cobra.Command{
		Use:   "update [FILENAME]",
		Short: "Update a json file",
		Args:  require.ExactArgs(1),
		Long:  updateDesc,
		RunE: func(cmd *cobra.Command, args []string) error {
			return o.runUpdate(out, cmd, args)
		},
	}

	cmd.SetOut(out)

	flags := cmd.Flags()
	flags.StringSliceVar(&o.setArrays, "set", []string{}, "set values on the command line (can specify multiple or separate values with commas: key1=val1,key2=val2)")
	flags.BoolVar(&o.dryRun, "dry-run", false, "simulate a update")

	cmd.MarkFlagRequired("set")

	return cmd
}

func (o *updateOptions) runUpdate(out io.Writer, cmd *cobra.Command, args []string) error {
	filename := args[0]

	currentMap := map[string]interface{}{}
	_bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	var newBytes bytes.Buffer
	newBytes.Write(PreBytes)
	newBytes.Write(_bytes)
	newBytes.Write(SubBytes)

	if err := json.Unmarshal(newBytes.Bytes(), &currentMap); err != nil {
		return err
	}

	for _, val := range o.setArrays {
		if strings.HasPrefix(val, "[") {
			val = fmt.Sprintf("%s%s", "je", val)
		} else {
			val = fmt.Sprintf("%s.%s", "je", val)
		}

		if err := parser.ParseInto(val, currentMap); err != nil {
			return err
		}
	}

	data, err := json.MarshalIndent(currentMap["je"], "", "  ")
	if err != nil {
		return err
	}

	if o.dryRun {
		fmt.Fprintln(out, string(data))
		return nil
	}

	if err := ioutil.WriteFile(filename, data, os.ModePerm); err != nil {
		return err
	}
	return nil
}

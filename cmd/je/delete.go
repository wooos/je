package main

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/wooos/je/cmd/je/require"
	"io"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
)

type deleteOptions struct {
	keys []string
	dryRun bool
}

const deleteDesc = `
This command delete keys of the json for a specified json file.
`

func newDeleteCmd(out io.Writer) *cobra.Command {
	o := &deleteOptions{}

	cmd := &cobra.Command{
		Use:   "delete [FILENAME]",
		Short: "Delete the keys of json",
		Long:  deleteDesc,
		Args:  require.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return o.runDeleteCmd(out, cmd, args)
		},
	}

	flags := cmd.Flags()
	flags.StringArrayVar(&o.keys, "keys", []string{}, "specified keys of delete(can specify multiple or separate values with commas: key1,key2)")
	flags.BoolVar(&o.dryRun, "dry-run", false, "simulate a delete")
	cmd.MarkFlagRequired("keys")

	cmd.SetOut(out)

	return cmd
}

func (o *deleteOptions) runDeleteCmd(out io.Writer, cmd *cobra.Command, args []string) (err error) {
	filename := args[0]

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	var currentMap map[string]interface{}

	if err := json.Unmarshal(bytes, &currentMap); err != nil {
		return err
	}

	var newMap map[string]interface{}
	for _, key := range o.keys {
		for _, k := range strings.Split(key, ",") {
			if newMap, err = deleteMapKey(currentMap, strings.Split(k, ".")); err != nil {
				return err
			}
		}
	}

	data, err := json.MarshalIndent(newMap, "", "  ")
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

func deleteMapKey(current map[string]interface{}, key []string) (newMap map[string]interface{}, err error) {
	if _, ok := current[key[0]]; !ok {
		return nil, errors.Errorf("can not find specified key: %s", strings.Join(key, "."))
	}

	newMap = current
	if len(key) >= 2 {
		if reflect.TypeOf(current[key[0]]).String() == "map[string]interface {}" {
			newMap[key[0]], err = deleteMapKey(current[key[0]].(map[string]interface {}), append(key[1:]))
			if err != nil {
				return nil, err
			}
		} else {
			return current, errors.Errorf("Error: Key not found.")
		}
	} else {
		delete(current, key[0])
		newMap = current
	}

	return newMap, nil
}

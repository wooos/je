package main

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

const listener = `
# Copyright 2020 The Goe Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
`

const completionDesc = `
Generate auto completions script for Goe for the specified shell (bash).
This command can generate shell auto completions. e.g.
    $ goe completion bash
Can be sourced as such
    $ source <(goe completion bash)
`

var completionShells = map[string]func(cmd *cobra.Command) error{
	"bash": runCompletionBash,
}

func newCompletionCommand() *cobra.Command {
	var completionCmd = &cobra.Command{
		Use:   "completion",
		Short: "Generate auto completions script for goe for the specified shell (bash)",
		Long:  completionDesc,
		RunE:  runCompletion,
	}

	return completionCmd
}

func runCompletion(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return errors.New("shell not specified")
	}
	if len(args) > 1 {
		return errors.New("too many arguments, expected only the shell type")
	}
	run, found := completionShells[args[0]]
	if !found {
		return errors.Errorf("unsupported shell type %q", args[0])
	}

	return run(cmd)
}

func runCompletionBash(cmd *cobra.Command) error {
	fmt.Printf("%s", listener)
	err := cmd.Root().GenBashCompletion(os.Stdout)
	bashrc := `
if [[ $(type -t compopt) = "builtin" ]]; then
    complete -o default -F __start_goe goe
else
    complete -o default -o nospace -F __start_goe goe
fi
`
	fmt.Printf("%s", bashrc)
	return err
}

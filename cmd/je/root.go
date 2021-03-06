package main

import (
	"io"

	"github.com/spf13/cobra"
)

func newRootCmd(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:          "je",
		Short:        "Command line editor for json.",
		Long:         "",
		SilenceUsage: true,
	}

	cmd.AddCommand(
		newCompletionCommand(),
		newUpdateCmd(out),
		newVersionCmd(out),
		newDeleteCmd(out),
	)
	return cmd
}

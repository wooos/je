package main

import (
	"io"

	"github.com/spf13/cobra"
)

func newRootCmd(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:          "je",
		Short:        "Stream editor for json",
		Long:         "",
		SilenceUsage: true,
	}

	cmd.AddCommand(
		newCompletionCommand(),
		newUpdateCmd(),
		newVersionCmd(out),
		newDeleteCmd(out),
	)
	return cmd
}

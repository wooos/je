package main

import (
	"fmt"
	"io"
	"text/template"

	"github.com/spf13/cobra"
	"github.com/wooos/je/cmd/je/require"
	"github.com/wooos/je/internal/version"
)

const versionDesc = `
Show the version for Je.

This will print a representation the version of Je.
The output will look something like this:

version.BuildInfo{Version:"v1.0.0", GitCommit:"ff52399e51bb880526e9cd0ed8386f6433b74da1", GitTreeState:"clean"}

- Version is the semantic version of the release.
- GitCommit is the SHA for the commit that this version was built from.
- GitTreeState is "clean" if there are no local code changes when this binary was
  built, and "dirty" if the binary was built from locally modified code.
`

type versionOptions struct {
	short    bool
	template string
}

func newVersionCmd(out io.Writer) *cobra.Command {
	o := &versionOptions{}

	cmd := &cobra.Command{
		Use:   "version",
		Short: "print current version information",
		Long:  versionDesc,
		Args:  require.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return o.run(out)
		},
	}

	flags := cmd.Flags()
	flags.BoolVar(&o.short, "short", false, "print the version number")
	flags.StringVar(&o.template, "template", "", "template for version string format")

	return cmd
}

func (o *versionOptions) run(out io.Writer) error {
	if o.template != "" {
		t, err := template.New("_").Parse(o.template)
		if err != nil {
			return err
		}

		return t.Execute(out, version.Get())
	}

	fmt.Fprintln(out, formatVersion(o.short))
	return nil
}

func formatVersion(short bool) string {
	v := version.Get()
	if short {
		return version.GetVersion()
	}

	return fmt.Sprintf("%#v", v)
}

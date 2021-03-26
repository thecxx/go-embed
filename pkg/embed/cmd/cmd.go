package cmd

import (
	"github.com/spf13/cobra"
	"github.com/thecxx/embed/pkg/embed/cmd/build"
)

var (
	Version = "0.0.0"
)

// NewCommand returns root command.
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Version: Version,
		Use:     "session",
		Short:   "",
		Long:    "",
		// EntryPoint
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	// Sub commands
	cmd.AddCommand(build.NewCommand())

	return cmd
}

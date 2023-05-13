package cmd

import (
	"io"
	"log"
	"os"

	"github.com/ShibataTakao/worklog/cmd/show"
	"github.com/spf13/cobra"
)

// NewCommand return new `worklog` command instance.
func NewCommand() *cobra.Command {
	var (
		silent bool
	)

	command := &cobra.Command{
		Use:   "worklog",
		Short: "CLI tool to control worklog.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if silent {
				log.SetOutput(io.Discard)
			}
		},
	}

	command.AddCommand(show.NewShowCommand())

	command.Flags().BoolVar(&silent, "silent", os.Getenv("WL_SILENT") == "true", "If true, do not print logs. [$WL_SILENT]")

	return command
}

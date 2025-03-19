package cmd

import (
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// NewCommand return new `dailyreport` command instance.
func NewCommand() *cobra.Command {
	var (
		silent bool
	)

	command := &cobra.Command{
		Use:   "dailyreport",
		Short: "CLI tool to control daily report.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
		PersistentPreRun: func(_ *cobra.Command, _ []string) {
			if silent {
				log.SetOutput(io.Discard)
			}
		},
	}

	command.AddCommand(NewShowCommand())

	command.Flags().BoolVar(&silent, "silent", os.Getenv("DR_SILENT") == "true", "If true, do not print logs. [$DR_SILENT]")

	return command
}

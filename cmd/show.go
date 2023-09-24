package cmd

import (
	"github.com/spf13/cobra"
)

// NewShowCommand return new `dailyreport show` sub-command instance.
func NewShowCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "show",
		Short: "Show information in daily report.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}

	command.AddCommand(NewShowTasksCommand())
	command.AddCommand(NewShowWorkTimeCommand())

	return command
}

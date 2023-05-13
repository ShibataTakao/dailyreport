package show

import (
	"github.com/spf13/cobra"
)

// NewShowCommand return new `worklog show` sub-command instance.
func NewShowCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "show",
		Short: "Show worklog.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}

	command.AddCommand(NewShowTasksCommand())
	command.AddCommand(NewShowWorkTimeCommand())

	return command
}

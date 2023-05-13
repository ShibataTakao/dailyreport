package show

import (
	"fmt"
	"os"
	"time"

	"github.com/ShibataTakao/worklog/dailyreport"
	"github.com/ShibataTakao/worklog/task"
	"github.com/spf13/cobra"
)

// NewShowCommand return new `worklog show tasks` sub-command instance.
func NewShowWorkTimeCommand() *cobra.Command {
	var (
		dailyReportDir     string
		dailyReportStartAt string
		dailyReportEndAt   string
		filterByProject    string
	)

	command := &cobra.Command{
		Use:   "worktime",
		Short: "Show work time in worklog.",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Read from daily reports.
			dailyReportApp := dailyreport.NewApplicationService(dailyreport.NewRepository(dailyReportDir))
			start, err := time.Parse("20060102", dailyReportStartAt)
			if err != nil {
				return err
			}
			end, err := time.Parse("20060102", dailyReportEndAt)
			if err != nil {
				return err
			}
			reports, err := dailyReportApp.Read(start, end)
			if err != nil {
				return err
			}

			// Get tasks from daily report.
			tasks, err := reports.Tasks()
			if err != nil {
				return err
			}
			if filterByProject != "" {
				prj := task.NewProject(filterByProject)
				tasks = tasks.Filter(func(t task.Task) bool {
					return t.Project.Equals(prj)
				})
			}

			// Show worktime.
			fmt.Printf("Work Time\t\t%.2fh\n", reports.WorkTimes().Duration().Hours())
			fmt.Printf("Tasks (Estimated)\t%.2fh\n", tasks.Estimate().Hours())
			fmt.Printf("Tasks (Actual)\t\t%.2fh\n", tasks.Actual().Hours())

			return nil
		},
	}

	command.Flags().StringVar(&dailyReportDir, "daily-report-dir", os.Getenv("WL_DAILY_REPORT_DIR"), "Directory where daily report file exists. [$WL_DAILY_REPORT_DIR]")
	command.Flags().StringVarP(&dailyReportStartAt, "daily-report-start-at", "s", os.Getenv("WL_DAILY_REPORT_START_AT"), "Start of daily report date range. [$WL_DAILY_REPORT_START_AT]")
	command.Flags().StringVarP(&dailyReportEndAt, "daily-report-end-at", "e", os.Getenv("WL_DAILY_REPORT_END_AT"), "End of daily report date range. [$WL_DAILY_REPORT_END_AT]")
	command.Flags().StringVar(&filterByProject, "filter-by-project", os.Getenv("WL_FILTER_BY_PROJECT"), "Show only tasks which project name is this. [$WL_FILTER_BY_PROJECT]")

	return command
}

package show

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/ShibataTakao/worklog/backlog"
	"github.com/ShibataTakao/worklog/dailyreport"
	"github.com/ShibataTakao/worklog/task"
	"github.com/spf13/cobra"
)

// NewShowCommand return new `worklog show tasks` sub-command instance.
func NewShowTasksCommand() *cobra.Command {
	var (
		dailyReportDir          string
		dailyReportStartAt      string
		dailyReportEndAt        string
		backlogAPIKey           string
		backlogURL              string
		backlogQuery            string
		backlogProjectOverwrite string
		filterByProject         string
	)

	command := &cobra.Command{
		Use:   "tasks",
		Short: "Show tasks in worklog.",
		RunE: func(cmd *cobra.Command, args []string) error {
			tasks := task.Set{}

			// Get tasks from daily reports.
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
			dailyReportTasks, err := reports.Tasks()
			if err != nil {
				return err
			}
			tasks = append(tasks, dailyReportTasks...)

			// Get tasks from backlog.
			if backlogAPIKey != "" {
				backlogApp := backlog.NewApplicationService(backlog.NewRepository(backlogAPIKey, backlogURL))
				backlogTasks, err := backlogApp.GetTasks(backlogQuery)
				if err != nil {
					return err
				}
				if backlogProjectOverwrite != "" {
					prj := task.NewProject(backlogProjectOverwrite)
					for i := range backlogTasks {
						backlogTasks[i].Project = prj
					}
				}
				tasks = append(tasks, backlogTasks...)
			}

			// Filter tasks.
			if filterByProject != "" {
				prj := task.NewProject(filterByProject)
				tasks = tasks.Filter(func(t task.Task) bool {
					return t.Project.Equals(prj)
				})
			}

			// Union tasks.
			tasks, err = tasks.Union()
			if err != nil {
				return err
			}

			// Sort tasks.
			tasks = tasks.Sort(func(a task.Task, b task.Task) bool {
				if !a.Project.Equals(b.Project) {
					return a.Project.Name < b.Project.Name
				}
				reBacklogKey := regexp.MustCompile(`.+-(\d+)`)
				if reBacklogKey.MatchString(a.Key) && reBacklogKey.MatchString(b.Key) {
					matches1 := reBacklogKey.FindStringSubmatch(a.Key)
					if len(matches1) != 2 {
						panic(fmt.Errorf("fail to parse task key '%s'", a.Key))
					}
					key1, err := strconv.Atoi(matches1[1])
					if err != nil {
						panic(err)
					}
					matches2 := reBacklogKey.FindStringSubmatch(b.Key)
					if len(matches2) != 2 {
						panic(fmt.Errorf("fail to parse task key '%s'", b.Key))
					}
					key2, err := strconv.Atoi(matches2[1])
					if err != nil {
						panic(err)
					}
					return key1 < key2
				}
				if a.Key == "" && b.Key == "" {
					return a.Name < b.Name
				}
				if a.Key == "" || b.Key == "" {
					return a.Key > b.Key
				}
				return a.Key < b.Key
			})

			// Show tasks.
			for _, task := range tasks {
				var completionMark string
				if task.IsCompleted {
					completionMark = "x"
				} else {
					completionMark = " "
				}
				if task.Key == "" {
					fmt.Printf("- [%s] %.2fh / %.2fh [%s] %s\n", completionMark, task.Estimate.Hours(), task.Actual.Hours(), task.Project.Name, task.Name)
				} else {
					fmt.Printf("- [%s] %.2fh / %.2fh [%s] [%s] %s\n", completionMark, task.Estimate.Hours(), task.Actual.Hours(), task.Project.Name, task.Key, task.Name)
				}
			}

			return nil
		},
	}

	command.Flags().StringVar(&dailyReportDir, "daily-report-dir", os.Getenv("WL_DAILY_REPORT_DIR"), "Directory where daily report file exists. [$WL_DAILY_REPORT_DIR]")
	command.Flags().StringVarP(&dailyReportStartAt, "daily-report-start-at", "s", os.Getenv("WL_DAILY_REPORT_START_AT"), "Start of daily report date range. [$WL_DAILY_REPORT_START_AT]")
	command.Flags().StringVarP(&dailyReportEndAt, "daily-report-end-at", "e", os.Getenv("WL_DAILY_REPORT_END_AT"), "End of daily report date range. [$WL_DAILY_REPORT_END_AT]")
	command.Flags().StringVar(&filterByProject, "filter-by-project", os.Getenv("WL_FILTER_BY_PROJECT"), "Show only tasks which project name is this. [$WL_FILTER_BY_PROJECT]")
	command.Flags().StringVar(&backlogAPIKey, "backlog-api-key", os.Getenv("WL_BACKLOG_API_KEY"), "Backlog API key. [$WL_BACKLOG_API_KEY]")
	command.Flags().StringVar(&backlogURL, "backlog-url", os.Getenv("WL_BACKLOG_URL"), "Backlog URL. [$WL_BACKLOG_URL]")
	command.Flags().StringVar(&backlogQuery, "backlog-query", os.Getenv("WL_BACKLOG_QUERY"), "Query to get issues from backlog. [$WL_BACKLOG_QUERY]")
	command.Flags().StringVar(&backlogProjectOverwrite, "backlog-project-overwrite", os.Getenv("WL_BACKLOG_PROJECT_OVERWRITE"), "This overwrite project name of issues in backlog. [$WL_BACKLOG_PROJECT_OVERWRITE]")

	return command
}

package dailyreport

import (
	"bytes"
	"fmt"
	"time"
)

type Analysis []DailyReport

func (a Analysis) String() string {
	var b bytes.Buffer

	var worktime, estimate, actual time.Duration
	tasks := []Task{}
	for _, report := range a {
		worktime += report.EndAt.Sub(report.StartAt) - report.BreakTime
		tasks = append(tasks, report.Tasks...)
	}
	for _, task := range tasks {
		estimate += task.Estimate
		actual += task.Actual
	}

	fmt.Fprintf(&b, "# Summary\n\n")
	fmt.Fprintf(&b, "- Days: %d days\n", len(a))
	fmt.Fprintf(&b, "- Work time: %s\n", worktime)
	fmt.Fprintf(&b, "- Tasks (Estimated): %s\n", estimate)
	fmt.Fprintf(&b, "- Tasks (Actual): %s\n", actual)

	fmt.Fprintf(&b, "# Tasks\n\n")
	projects := []Project{}
	for _, task := range tasks {
		projects = append(projects, task.Project)
	}
	for _, project := range projects {
		for _, task := range tasks {
			if task.Project.Name == project.Name {
				estimate += task.Estimate
				actual += task.Actual
			}
		}
		fmt.Fprintf(&b, "- [ ] %d/%d %sProject: %s\n", project.Name)
	}

	return b.String()
}

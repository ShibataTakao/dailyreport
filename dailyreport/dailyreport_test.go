package dailyreport

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestDailyReportSetTasks(t *testing.T) {
	tests := []struct {
		name    string
		reports Set
		tasks   TaskSet
	}{
		{
			name: "",
			reports: Set{
				New(
					NewWorkTime(time.Date(2023, 1, 1, 9, 30, 0, 0, time.Local), time.Date(2023, 1, 1, 17, 30, 0, 0, time.Local), 1*time.Hour),
					TaskSet{
						NewTask("Task A", NewProject("Project A"), 1, 2, false),
						NewTask("Task B", NewProject("Project B"), 3, 4, false),
					},
				),
				New(
					NewWorkTime(time.Date(2023, 1, 1, 9, 30, 0, 0, time.Local), time.Date(2023, 1, 1, 17, 30, 0, 0, time.Local), 1*time.Hour),
					TaskSet{
						NewTask("Task C", NewProject("Project C"), 5, 6, false),
						NewTask("Task D", NewProject("Project D"), 7, 8, false),
					},
				),
			},
			tasks: TaskSet{
				NewTask("Task A", NewProject("Project A"), 1, 2, false),
				NewTask("Task B", NewProject("Project B"), 3, 4, false),
				NewTask("Task C", NewProject("Project C"), 5, 6, false),
				NewTask("Task D", NewProject("Project D"), 7, 8, false),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := require.New(t)
			tasks, err := tt.reports.Tasks()
			assert.NoError(err)
			assert.Equal(tt.tasks, tasks)
		})
	}
}

func TestDailyReportSetWorkTimes(t *testing.T) {
	tests := []struct {
		name      string
		reports   Set
		worktimes WorkTimeSet
	}{
		{
			name: "",
			reports: Set{
				New(
					NewWorkTime(time.Date(2023, 1, 1, 9, 30, 0, 0, time.Local), time.Date(2023, 1, 1, 17, 30, 0, 0, time.Local), 1*time.Hour),
					TaskSet{
						NewTask("Task A", NewProject("Project A"), 1, 2, false),
						NewTask("Task B", NewProject("Project B"), 3, 4, false),
					},
				),
				New(
					NewWorkTime(time.Date(2023, 1, 1, 9, 30, 0, 0, time.Local), time.Date(2023, 1, 1, 17, 30, 0, 0, time.Local), 1*time.Hour),
					TaskSet{
						NewTask("Task C", NewProject("Project C"), 5, 6, false),
						NewTask("Task D", NewProject("Project D"), 7, 8, false),
					},
				),
			},
			worktimes: WorkTimeSet{
				NewWorkTime(time.Date(2023, 1, 1, 9, 30, 0, 0, time.Local), time.Date(2023, 1, 1, 17, 30, 0, 0, time.Local), 1*time.Hour),
				NewWorkTime(time.Date(2023, 1, 1, 9, 30, 0, 0, time.Local), time.Date(2023, 1, 1, 17, 30, 0, 0, time.Local), 1*time.Hour),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := require.New(t)
			assert.Equal(tt.worktimes, tt.reports.WorkTimes())
		})
	}
}

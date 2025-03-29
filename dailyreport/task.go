package dailyreport

import (
	"strings"
	"time"
)

// Task.
type Task struct {
	Summary     string
	Project     Project
	Estimate    time.Duration
	Actual      time.Duration
	IsCompleted bool
}

// newTask returns a new task instance.
func newTask(summary string, project Project, estimate time.Duration, actual time.Duration, isCompleted bool) Task {
	return Task{
		Summary:     strings.TrimSpace(summary),
		Project:     project,
		Estimate:    estimate,
		Actual:      actual,
		IsCompleted: isCompleted,
	}
}

package task

import (
	"fmt"
	"strings"
	"time"
)

// Task.
type Task struct {
	Key         string
	Name        string
	Project     Project
	Estimate    time.Duration
	Actual      time.Duration
	IsCompleted bool
}

// NewTask return new task instance.
func NewTask(key string, name string, project Project, estimate time.Duration, actual time.Duration, isCompleted bool) Task {
	return Task{
		Key:         strings.TrimSpace(key),
		Name:        strings.TrimSpace(name),
		Project:     project,
		Estimate:    estimate,
		Actual:      actual,
		IsCompleted: isCompleted,
	}
}

// NewDailyReportTask return new task instance for daily report.
func NewDailyReportTask(name string, project Project, estimate time.Duration, actual time.Duration, isCompleted bool) Task {
	return NewTask("", name, project, estimate, actual, isCompleted)
}

// NewBacklogTask return new task instance for backlog.
func NewBacklogTask(key string, name string, project Project, isCompleted bool) Task {
	return NewTask(key, name, project, 0, 0, isCompleted)
}

// Equals return true if both are same task.
func (t Task) Equals(other Task) bool {
	isSameProject := t.Project.Equals(other.Project)
	isSameTaskName := t.Name == other.Name
	isDifferentKey := t.HasKey() && other.HasKey() && t.Key != other.Key
	return isSameProject && isSameTaskName && !isDifferentKey
}

// Merge two tasks into one task.
// Two tasks must be same.
func (t Task) Merge(other Task) (Task, error) {
	if !t.Equals(other) {
		return Task{}, fmt.Errorf("task %v and %v are not same task", t, other)
	}

	var key string
	if t.HasKey() {
		key = t.Key
	} else {
		key = other.Key
	}

	project := t.Project
	name := t.Name
	estimate := t.Estimate + other.Estimate
	actual := t.Actual + other.Actual
	isCompleted := t.IsCompleted || other.IsCompleted

	return NewTask(key, name, project, estimate, actual, isCompleted), nil
}

// Clone task instance
func (t Task) Clone() Task {
	return NewTask(t.Key, t.Name, t.Project, t.Estimate, t.Actual, t.IsCompleted)
}

// HasKey return true if task has key.
func (t Task) HasKey() bool {
	return t.Key != ""
}

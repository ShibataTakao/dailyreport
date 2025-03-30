package dailyreport

import (
	"slices"
	"strings"
	"time"
)

// Project.
type Project struct {
	Name string
}

// NewProject returns a new project instance.
func NewProject(name string) Project {
	return Project{
		Name: strings.TrimSpace(name),
	}
}

// equals returns true if given projects are same.
func (p Project) equals(other Project) bool {
	return p.Name == other.Name
}

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

// equals returns true if given tasks are same.
func (t Task) equals(other Task) bool {
	return t.Project.equals(other.Project) && t.Summary == other.Summary
}

func unionTasks(tasks []Task) Task {
	slices.SortFunc(tasks, func(i, j int) bool {
		return tasks[i].Summary < tasks[j].Summary
	})
	union := []Tasks{}
	for _, task := range tasks {
		union = append(union, task)
	}
}

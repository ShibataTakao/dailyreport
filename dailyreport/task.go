package dailyreport

import (
	"fmt"
	"strings"
	"time"

	"golang.org/x/exp/slices"
)

// Task.
type Task struct {
	Name        string
	Project     Project
	Estimate    time.Duration
	Actual      time.Duration
	IsCompleted bool
}

// TaskSet is set of tasks.
type TaskSet []Task

// NewTask return new task instance.
func NewTask(name string, project Project, estimate time.Duration, actual time.Duration, isCompleted bool) Task {
	return Task{
		Name:        strings.TrimSpace(name),
		Project:     project,
		Estimate:    estimate,
		Actual:      actual,
		IsCompleted: isCompleted,
	}
}

// Equals return true if both are same task.
func (t Task) Equals(other Task) bool {
	return t.Project.Equals(other.Project) && t.Name == other.Name
}

// Merge two tasks into one task.
// Two tasks must be same.
func (t Task) Merge(other Task) (Task, error) {
	if !t.Equals(other) {
		return Task{}, fmt.Errorf("task %v and %v are not same task", t, other)
	}

	project := t.Project
	name := t.Name
	estimate := t.Estimate + other.Estimate
	actual := t.Actual + other.Actual
	isCompleted := t.IsCompleted || other.IsCompleted

	return NewTask(name, project, estimate, actual, isCompleted), nil
}

// Clone task instance
func (t Task) Clone() Task {
	return NewTask(t.Name, t.Project, t.Estimate, t.Actual, t.IsCompleted)
}

// Union merge same tasks in task set.
func (s TaskSet) Union() (TaskSet, error) {
	set := TaskSet{}
	added := make([]bool, len(s))
	var err error

	for i, task := range s {
		if added[i] {
			continue
		}
		added[i] = true
		for j := i + 1; j < len(s); j++ {
			if added[j] {
				continue
			}
			other := s[j]
			if task.Equals(other) {
				task, err = task.Merge(other)
				if err != nil {
					return TaskSet{}, err
				}
				added[j] = true
			}
		}
		set = append(set, task.Clone())
	}

	if slices.Contains(added, false) || s.Estimate() != set.Estimate() || s.Actual() != set.Actual() {
		return TaskSet{}, fmt.Errorf("fail to merge task set")
	}

	return set, nil
}

// Estimate return sum of estimated duration in task set.
func (s TaskSet) Estimate() time.Duration {
	var t time.Duration
	for _, task := range s {
		t += task.Estimate
	}
	return t
}

// Actual return sum of actual duration in task set.
func (s TaskSet) Actual() time.Duration {
	var t time.Duration
	for _, task := range s {
		t += task.Actual
	}
	return t
}

// Filter tasks in set and return new set .
func (s TaskSet) Filter(filter func(Task) bool) TaskSet {
	set := TaskSet{}
	for _, task := range s {
		if filter(task) {
			set = append(set, task)
		}
	}
	return set
}

// Sort tasks in task set and return new task set instance.
func (s TaskSet) Sort(less func(Task, Task) int) TaskSet {
	set := s.Clone()
	slices.SortFunc(set, less)
	return set
}

// Clone task set instance.
func (s TaskSet) Clone() TaskSet {
	set := TaskSet{}
	return append(set, s...)
}

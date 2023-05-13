package task

import (
	"fmt"
	"time"

	"golang.org/x/exp/slices"
)

// Set of tasks.
type Set []Task

// Union merge same task in set.
func (s Set) Union() (Set, error) {
	set := Set{}
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
					return Set{}, err
				}
				added[j] = true
			}
		}
		set = append(set, task.Clone())
	}

	if slices.Contains(added, false) || s.Estimate() != set.Estimate() || s.Actual() != set.Actual() {
		return Set{}, fmt.Errorf("fail to merge task set")
	}

	return set, nil
}

// Estimate return sum of estimated duration for tasks in set.
func (s Set) Estimate() time.Duration {
	var t time.Duration
	for _, task := range s {
		t += task.Estimate
	}
	return t
}

// Actual return sum of actual duration for tasks in set.
func (s Set) Actual() time.Duration {
	var t time.Duration
	for _, task := range s {
		t += task.Actual
	}
	return t
}

// Filter tasks in set and return new set .
func (s Set) Filter(filter func(Task) bool) Set {
	set := Set{}
	for _, task := range s {
		if filter(task) {
			set = append(set, task)
		}
	}
	return set
}

// Sort tasks in set and return new set instance.
func (s Set) Sort(less func(Task, Task) bool) Set {
	set := s.Clone()
	slices.SortFunc(set, less)
	return set
}

// Clone set instance.
func (s Set) Clone() Set {
	set := Set{}
	return append(set, s...)
}

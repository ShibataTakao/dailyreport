package dailyreport

import (
	"sort"
	"strings"
	"time"
)

type taskItem struct {
	name string

	// fields from parsed daily report
	category   string
	expectTime time.Duration
	actualTime time.Duration

	// fields from issue
	done      bool
	createdAt time.Time
}

type taskItems []taskItem

func (tasks taskItems) expectWorktime() time.Duration {
	d := time.Duration(0)
	for _, task := range tasks {
		d += task.expectTime
	}
	return d
}

func (tasks taskItems) actualWorktime() time.Duration {
	d := time.Duration(0)
	for _, task := range tasks {
		d += task.actualTime
	}
	return d
}

func (tasks taskItems) filteredByCategory(category string) taskItems {
	newTasks := taskItems{}
	for _, task := range tasks {
		if task.category == category {
			newTasks = append(newTasks, task)
		}
	}
	return newTasks
}

func (tasks taskItems) categories() []string {
	categories := []string{}
	for _, task := range tasks {
		found := false
		for _, category := range categories {
			if task.category == category {
				found = true
				break
			}
		}
		if !found {
			categories = append(categories, task.category)
		}
	}
	return categories
}

func (tasks taskItems) aggregated() taskItems {
	newTasks := taskItems{}
	for _, task := range tasks {
		if task.expectTime == 0 && task.actualTime == 0 {
			continue
		}
		found := false
		for i, newTask := range newTasks {
			if task.category == newTask.category && task.name == newTask.name {
				newTasks[i].expectTime += task.expectTime
				newTasks[i].actualTime += task.actualTime
				found = true
				break
			}
		}
		if !found {
			newTasks = append(newTasks, task)
		}
	}
	sort.Slice(newTasks, func(i, j int) bool { return newTasks[i].actualTime > newTasks[j].actualTime })
	return newTasks
}

func (tasks taskItems) mergeIssues(issues []issueItem) taskItems {
	newTasks := taskItems{}
	isTaskPaired := make([]bool, len(tasks))
	for _, issue := range issues {
		newTask := taskItem{
			name:       issue.name,
			category:   "",
			expectTime: 0,
			actualTime: 0,
			done:       issue.isDone(),
			createdAt:  issue.createdAt,
		}
		for i, task := range tasks {
			if !isTaskPaired[i] && strings.Contains(issue.name, task.name) {
				newTask.category = task.category
				newTask.expectTime = task.expectTime
				newTask.actualTime = task.actualTime
				isTaskPaired[i] = true
				break
			}
		}
		newTasks = append(newTasks, newTask)
	}
	for i, task := range tasks {
		if !isTaskPaired[i] {
			task.createdAt = time.Now()
			newTasks = append(newTasks, task)
		}
	}
	sort.Slice(newTasks, func(i, j int) bool { return newTasks[i].createdAt.Before(newTasks[j].createdAt) })
	return newTasks
}

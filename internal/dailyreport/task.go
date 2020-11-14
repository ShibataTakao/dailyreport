package dailyreport

import "time"

type taskItem struct {
	category   string
	name       string
	expectTime time.Duration
	actualTime time.Duration
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
	return newTasks
}

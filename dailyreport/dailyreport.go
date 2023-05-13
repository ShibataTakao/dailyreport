package dailyreport

import (
	"github.com/ShibataTakao/worklog/task"
)

// DailyReport.
type DailyReport struct {
	WorkTime WorkTime
	Tasks    task.Set
}

// Set of daily reports.
type Set []DailyReport

// NewDailyReport return new daily report insntace.
func NewDailyReport(workTime WorkTime, tasks task.Set) DailyReport {
	return DailyReport{
		WorkTime: workTime,
		Tasks:    tasks,
	}
}

// Tasks return set of task in daily report set.
func (s Set) Tasks() (task.Set, error) {
	tasks := task.Set{}
	for _, r := range s {
		tasks = append(tasks, r.Tasks...)
	}
	return tasks.Union()
}

// WorkTimes return set of work time in daily report set.
func (s Set) WorkTimes() WorkTimeSet {
	worktimes := WorkTimeSet{}
	for _, r := range s {
		worktimes = append(worktimes, r.WorkTime)
	}
	return worktimes
}

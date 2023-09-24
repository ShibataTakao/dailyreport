package dailyreport

// DailyReport.
type DailyReport struct {
	WorkTime WorkTime
	Tasks    TaskSet
}

// Set of daily reports.
type Set []DailyReport

// New return new daily report insntace.
func New(workTime WorkTime, tasks TaskSet) DailyReport {
	return DailyReport{
		WorkTime: workTime,
		Tasks:    tasks,
	}
}

// Tasks return set of task in daily report set.
func (s Set) Tasks() (TaskSet, error) {
	tasks := TaskSet{}
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

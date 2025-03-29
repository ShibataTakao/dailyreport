package dailyreport

import "time"

// DailyReport represents a daily report.
type DailyReport struct {
	StartAt   time.Time
	EndAt     time.Time
	BreakTime time.Duration

	Tasks []Task
}

// new returns a new daily report object.
func new(startAt time.Time, endAt time.Time, breakTime time.Duration, tasks []Task) DailyReport {
	return DailyReport{
		StartAt:   startAt,
		EndAt:     endAt,
		BreakTime: breakTime,
		Tasks:     tasks,
	}
}

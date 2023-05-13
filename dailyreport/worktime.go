package dailyreport

import "time"

// WorkTime.
type WorkTime struct {
	StartAt   time.Time
	EndAt     time.Time
	BreakTime time.Duration
}

// WorkTimeSet is set of work time.
type WorkTimeSet []WorkTime

// NewWorkTime return new work time instance.
func NewWorkTime(startAt time.Time, endAt time.Time, breakTime time.Duration) WorkTime {
	return WorkTime{
		StartAt:   startAt,
		EndAt:     endAt,
		BreakTime: breakTime,
	}
}

// Duration of work time.
func (wt WorkTime) Duration() time.Duration {
	return wt.EndAt.Sub(wt.StartAt) - wt.BreakTime
}

// Duration return sum of work time duration in work time set.
func (s WorkTimeSet) Duration() time.Duration {
	var t time.Duration
	for _, wt := range s {
		t += wt.Duration()
	}
	return t
}

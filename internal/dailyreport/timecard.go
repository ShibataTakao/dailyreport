package dailyreport

import "time"

type timecardItem struct {
	begin time.Time
	end   time.Time
	rest  time.Duration
}

func (tc timecardItem) timecardWorktime() time.Duration {
	return tc.end.Sub(tc.begin.Add(tc.rest))
}

package dailyreport

import (
	"testing"
	"time"
)

func TestTimecardWorktime(t *testing.T) {
	tests := []struct {
		name     string
		timecard timecardItem
		out      time.Duration
	}{
		{
			name: "case01",
			timecard: timecardItem{
				begin: newTime(9, 30, 0),
				end:   newTime(17, 30, 0),
				rest:  time.Duration(1) * time.Hour,
			},
			out: time.Duration(7) * time.Hour,
		},
		{
			name: "case02",
			timecard: timecardItem{
				begin: newTime(9, 30, 0),
				end:   newTime(9, 30, 0),
				rest:  time.Duration(1) * time.Hour,
			},
			out: time.Duration(-1) * time.Hour,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.timecard.timecardWorktime()
			if actual != tt.out {
				t.Errorf("Expected is %v but actual is %v", tt.out, actual)
			}
		})
	}
}

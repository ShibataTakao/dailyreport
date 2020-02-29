package dailyreport

import (
	"testing"
	"time"
)

func TestTimecardWorktime(t *testing.T) {
	tests := []struct {
		name   string
		report dailyReport
		out    time.Duration
	}{
		{
			name: "case01",
			report: dailyReport{
				timecard: timecard{
					begin: newTime(9, 30, 0),
					end:   newTime(17, 30, 0),
					rest:  time.Duration(1) * time.Hour,
				},
			},
			out: time.Duration(7) * time.Hour,
		},
		{
			name: "case02",
			report: dailyReport{
				timecard: timecard{
					begin: newTime(9, 30, 0),
					end:   newTime(9, 30, 0),
					rest:  time.Duration(1) * time.Hour,
				},
			},
			out: time.Duration(-1) * time.Hour,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.report.timecardWorktime()
			if actual != tt.out {
				t.Errorf("Expected is %v but actual is %v", tt.out, actual)
			}
		})
	}
}

func TestExpectWorktime(t *testing.T) {
	tests := []struct {
		name   string
		report dailyReport
		out    time.Duration
	}{
		{
			name: "case01",
			report: dailyReport{
				tasks: []task{
					task{expectTime: time.Duration(1) * time.Hour},
					task{expectTime: time.Duration(1)*time.Hour + time.Duration(30)*time.Minute},
				},
			},
			out: time.Duration(2)*time.Hour + time.Duration(30)*time.Minute,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.report.expectWorktime()
			if actual != tt.out {
				t.Errorf("Expected is %v but actual is %v", tt.out, actual)
			}
		})
	}
}

func TestActualWorktime(t *testing.T) {
	tests := []struct {
		name   string
		report dailyReport
		out    time.Duration
	}{
		{
			name: "case01",
			report: dailyReport{
				tasks: []task{
					task{actualTime: time.Duration(1) * time.Hour},
					task{actualTime: time.Duration(1)*time.Hour + time.Duration(30)*time.Minute},
				},
			},
			out: time.Duration(2)*time.Hour + time.Duration(30)*time.Minute,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.report.actualWorktime()
			if actual != tt.out {
				t.Errorf("Expected is %v but actual is %v", tt.out, actual)
			}
		})
	}
}

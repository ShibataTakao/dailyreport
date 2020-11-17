package dailyreport

import (
	"testing"
	"time"
)

func TestLastDate(t *testing.T) {
	tests := []struct {
		name string
		t    time.Time
		d    time.Weekday
		out  time.Time
	}{
		{
			name: "case01",
			t:    newDate(2020, 11, 17),
			d:    time.Monday,
			out:  newDate(2020, 11, 16),
		},
		{
			name: "case02",
			t:    newDate(2020, 11, 16),
			d:    time.Monday,
			out:  newDate(2020, 11, 16),
		},
		{
			name: "case03",
			t:    newDate(2020, 11, 15),
			d:    time.Monday,
			out:  newDate(2020, 11, 9),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := lastDate(tt.t, tt.d)
			if actual != tt.out {
				t.Errorf("Expected is %v but actual is %v", tt.out, actual)
			}
		})
	}
}

func TestNextDate(t *testing.T) {
	tests := []struct {
		name string
		t    time.Time
		d    time.Weekday
		out  time.Time
	}{
		{
			name: "case01",
			t:    newDate(2020, 11, 19),
			d:    time.Friday,
			out:  newDate(2020, 11, 20),
		},
		{
			name: "case02",
			t:    newDate(2020, 11, 20),
			d:    time.Friday,
			out:  newDate(2020, 11, 20),
		},
		{
			name: "case03",
			t:    newDate(2020, 11, 21),
			d:    time.Friday,
			out:  newDate(2020, 11, 27),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := nextDate(tt.t, tt.d)
			if actual != tt.out {
				t.Errorf("Expected is %v but actual is %v", tt.out, actual)
			}
		})
	}
}

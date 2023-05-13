package dailyreport

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestWorkTimeDuration(t *testing.T) {
	tests := []struct {
		name     string
		worktime WorkTime
		duration time.Duration
	}{
		{
			name:     "7Hours",
			worktime: NewWorkTime(time.Date(2023, 1, 1, 9, 30, 0, 0, time.Local), time.Date(2023, 1, 1, 17, 30, 0, 0, time.Local), 1*time.Hour),
			duration: 7 * time.Hour,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := require.New(t)
			assert.Equal(tt.duration, tt.worktime.Duration())
		})
	}
}

func TestWorkTimeSetDuration(t *testing.T) {
	tests := []struct {
		name     string
		set      WorkTimeSet
		duration time.Duration
	}{
		{
			name: "14Hours",
			set: WorkTimeSet{
				NewWorkTime(time.Date(2023, 1, 1, 9, 30, 0, 0, time.Local), time.Date(2023, 1, 1, 17, 30, 0, 0, time.Local), 1*time.Hour),
				NewWorkTime(time.Date(2023, 1, 2, 9, 30, 0, 0, time.Local), time.Date(2023, 1, 2, 17, 30, 0, 0, time.Local), 1*time.Hour),
			},
			duration: 14 * time.Hour,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := require.New(t)
			assert.Equal(tt.duration, tt.set.Duration())
		})
	}
}

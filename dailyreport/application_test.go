package dailyreport

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestApplicationServiceRead(t *testing.T) {
	tests := []struct {
		name    string
		dir     string
		start   time.Time
		end     time.Time
		reports Set
	}{
		{
			name:  "Read",
			dir:   "testdata",
			start: time.Date(2022, 12, 1, 0, 0, 0, 0, time.Local),
			end:   time.Date(2023, 2, 1, 0, 0, 0, 0, time.Local),
			reports: Set{
				New(
					NewWorkTime(time.Date(2023, 1, 1, 9, 30, 0, 0, time.Local), time.Date(2023, 1, 1, 17, 30, 0, 0, time.Local), 1*time.Hour),
					TaskSet{
						NewTask("タスクA", NewProject("プロジェクトA"), 2*time.Hour, 2*time.Hour, false),
						NewTask("タスクB", NewProject("プロジェクトA"), 2*time.Hour, 2*time.Hour, true),
						NewTask("タスクA", NewProject("プロジェクトB"), 1*time.Hour+30*time.Minute, 1*time.Hour+30*time.Minute, false),
						NewTask("タスクB", NewProject("プロジェクトB"), 1*time.Hour+30*time.Minute, 1*time.Hour+30*time.Minute, false),
					},
				),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := require.New(t)
			repository := NewRepository(tt.dir)
			application := NewApplicationService(repository)
			reports, err := application.Read(tt.start, tt.end)
			assert.NoError(err)
			assert.Equal(tt.reports, reports)
		})
	}
}

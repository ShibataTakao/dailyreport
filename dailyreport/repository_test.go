package dailyreport

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestRepositoryRead(t *testing.T) {
	tests := []struct {
		name   string
		dir    string
		date   time.Time
		report DailyReport
	}{
		{
			name: "testdata/20230101.md",
			dir:  "testdata",
			date: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
			report: New(
				NewWorkTime(time.Date(2023, 1, 1, 9, 30, 0, 0, time.Local), time.Date(2023, 1, 1, 17, 30, 0, 0, time.Local), 1*time.Hour),
				TaskSet{
					NewTask("タスクA", NewProject("プロジェクトA"), 2*time.Hour, 2*time.Hour, false),
					NewTask("タスクB", NewProject("プロジェクトA"), 2*time.Hour, 2*time.Hour, true),
					NewTask("タスクA", NewProject("プロジェクトB"), 1*time.Hour+30*time.Minute, 1*time.Hour+30*time.Minute, false),
					NewTask("タスクB", NewProject("プロジェクトB"), 1*time.Hour+30*time.Minute, 1*time.Hour+30*time.Minute, false),
				},
			),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := require.New(t)
			repository := NewRepository(tt.dir)
			report, err := repository.Read(tt.date)
			assert.NoError(err)
			assert.Equal(tt.report, report)
		})
	}
}

func TestRepositoryExists(t *testing.T) {
	tests := []struct {
		name   string
		dir    string
		date   time.Time
		exists bool
	}{
		{
			name:   "Exist",
			dir:    "testdata",
			date:   time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
			exists: true,
		},
		{
			name:   "NotExist",
			dir:    "testdata",
			date:   time.Date(2023, 2, 1, 0, 0, 0, 0, time.Local),
			exists: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := require.New(t)
			repository := NewRepository(tt.dir)
			assert.Equal(tt.exists, repository.Exists(tt.date))
		})
	}
}

func TestRepositoryPath(t *testing.T) {
	tests := []struct {
		name string
		dir  string
		date time.Time
		path string
	}{
		{
			name: "Exist",
			dir:  "testdata",
			date: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
			path: "testdata/20230101.md",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := require.New(t)
			repository := NewRepository(tt.dir)
			assert.Equal(tt.path, repository.Path(tt.date))
		})
	}
}

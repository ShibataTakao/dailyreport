package dailyreport

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestRepositoryParse(t *testing.T) {
	tests := []struct {
		name   string
		text   string
		date   time.Time
		report DailyReport
	}{
		{
			name: "Prase",
			text: `
# 日報

## 業務時間
- 始業 09:30
- 終業 17:30
- 休憩 01:00

## 今日のタスク（予定/実績）
- [ ] プロジェクトA
	- [ ] 2.0h/2.0h タスクA
	- [x] 2.0h/2.0h タスクB
- [ ] プロジェクトB
	- [ ] 1.5h/1.5h タスクA
	- [ ] 1.5h/1.5h タスクB

---

# 業務記録
- [ ] プロジェクトA
	- [ ] 2.0h/2.0h タスクA
`,
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
			report, err := NewParser(tt.text, tt.date).Parse()
			assert.NoError(err)
			assert.Equal(tt.report, report)
		})
	}
}

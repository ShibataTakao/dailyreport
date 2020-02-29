package dailyreport

import (
	"reflect"
	"testing"
	"time"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name string
		in   string
		out  dailyReport
	}{
		{
			name: "case01",
			in: `
# 日報

## 業務時間
- 始業 09:30
- 終業 17:30
- 休憩 01:00

## 今日のタスク（予定/実績）
- [ ] カテゴリ1
	- [ ] 1.0h/1.0h タスク1
		- [ ] 1.0h/0.0h サブタスク1-1
	- [ ] 1.5h/0.5h タスク2
- [ ] カテゴリ2
	- [ ] 1.0h/1.0h タスク3

---

# 業務記録
`,
			out: dailyReport{
				timecard: timecard{
					begin: newTime(9, 30, 0),
					end:   newTime(17, 30, 0),
					rest:  time.Duration(1) * time.Hour,
				},
				tasks: []task{
					{
						category:   "カテゴリ1",
						name:       "タスク1",
						expectTime: time.Duration(1) * time.Hour,
						actualTime: time.Duration(1) * time.Hour,
					},
					{
						category:   "カテゴリ1",
						name:       "サブタスク1-1",
						expectTime: time.Duration(1) * time.Hour,
						actualTime: time.Duration(0),
					},
					{
						category:   "カテゴリ1",
						name:       "タスク2",
						expectTime: time.Duration(1)*time.Hour + time.Duration(30)*time.Minute,
						actualTime: time.Duration(30) * time.Minute,
					},
					{
						category:   "カテゴリ2",
						name:       "タスク3",
						expectTime: time.Duration(1) * time.Hour,
						actualTime: time.Duration(1) * time.Hour,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := parse(tt.in)
			if err != nil {
				t.Error(err)
			}
			if !reflect.DeepEqual(actual, tt.out) {
				t.Errorf("Expected is %v but actual is %v", tt.out, actual)
			}
		})
	}
}

func TestParseTime(t *testing.T) {
	tests := []struct {
		name string
		in   string
		out  time.Time
	}{
		{
			name: "- 始業 09:30",
			in:   "- 始業 09:30",
			out:  newTime(9, 30, 0),
		},
		{
			name: "- 終業 17:30",
			in:   "- 終業 17:30",
			out:  newTime(17, 30, 0),
		},
		{
			name: "- 休憩 01:00",
			in:   "- 休憩 01:00",
			out:  newTime(1, 0, 0),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := parseTime(tt.in)
			if err != nil {
				t.Error(err)
			}
			if actual != tt.out {
				t.Errorf("Expected is %v but actual is %v", tt.out, actual)
			}
		})
	}
}

func TestParseTask(t *testing.T) {
	tests := []struct {
		name     string
		category string
		s        string
		out      task
	}{
		{
			name:     "    - [ ] 1.0h/1.5h タスク",
			category: "カテゴリ",
			s:        "    - [ ] 1.0h/1.5h タスク",
			out: task{
				category:   "カテゴリ",
				name:       "タスク",
				expectTime: time.Duration(1) * time.Hour,
				actualTime: time.Duration(1)*time.Hour + time.Duration(30)*time.Minute,
			},
		},
		{
			name:     "    - [x] 1.0h/1.5h タスク",
			category: "カテゴリ",
			s:        "    - [x] 1.0h/1.5h タスク",
			out: task{
				category:   "カテゴリ",
				name:       "タスク",
				expectTime: time.Duration(1) * time.Hour,
				actualTime: time.Duration(1)*time.Hour + time.Duration(30)*time.Minute,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := parseTask(tt.s, tt.category)
			if err != nil {
				t.Error(err)
			}
			if actual != tt.out {
				t.Errorf("Expected is %v but actual is %v", tt.out, actual)
			}
		})
	}
}

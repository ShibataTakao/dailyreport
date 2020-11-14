package dailyreport

import (
	"reflect"
	"testing"
	"time"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name     string
		in       string
		timecard timecardItem
		tasks    taskItems
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
			timecard: timecardItem{
				begin: newTime(9, 30, 0),
				end:   newTime(17, 30, 0),
				rest:  time.Duration(1) * time.Hour,
			},
			tasks: taskItems{
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualTimecard, actualTasks, err := parse(tt.in)
			if err != nil {
				t.Error(err)
			}
			if !reflect.DeepEqual(actualTimecard, tt.timecard) {
				t.Errorf("Expected is %v but actual is %v", tt.timecard, actualTimecard)
			}
			if !reflect.DeepEqual(actualTasks, tt.tasks) {
				t.Errorf("Expected is %v but actual is %v", tt.tasks, actualTasks)
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
		in       string
		out      taskItem
	}{
		{
			name:     "    - [ ] 1.0h/1.5h タスク",
			category: "カテゴリ",
			in:       "    - [ ] 1.0h/1.5h タスク",
			out: taskItem{
				category:   "カテゴリ",
				name:       "タスク",
				expectTime: time.Duration(1) * time.Hour,
				actualTime: time.Duration(1)*time.Hour + time.Duration(30)*time.Minute,
			},
		},
		{
			name:     "    - [x] 1.0h/1.5h タスク",
			category: "カテゴリ",
			in:       "    - [x] 1.0h/1.5h タスク",
			out: taskItem{
				category:   "カテゴリ",
				name:       "タスク",
				expectTime: time.Duration(1) * time.Hour,
				actualTime: time.Duration(1)*time.Hour + time.Duration(30)*time.Minute,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := parseTask(tt.in, tt.category)
			if err != nil {
				t.Error(err)
			}
			if actual != tt.out {
				t.Errorf("Expected is %v but actual is %v", tt.out, actual)
			}
		})
	}
}

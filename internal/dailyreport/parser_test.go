package dailyreport

import (
	"reflect"
	"testing"
	"time"
)

type parseTestData struct {
	text   string
	expect dailyReport
}

func TestParse(t *testing.T) {
	testData := []parseTestData{
		{
			text: `
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
			expect: dailyReport{
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

	for _, test := range testData {
		actual, err := parse(test.text)
		if err != nil {
			t.Error(err)
		}
		if !reflect.DeepEqual(test.expect, actual) {
			t.Errorf("Expected is %v but actual is %v", test.expect, actual)
		}
	}
}

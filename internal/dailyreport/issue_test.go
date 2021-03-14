package dailyreport

import "testing"

func TestIsDone(t *testing.T) {
	tests := []struct {
		name  string
		issue issueItem
		out   bool
	}{
		{
			name: "case01",
			issue: issueItem{
				name:   "name1",
				status: "処理中",
			},
			out: false,
		},
		{
			name: "case02",
			issue: issueItem{
				name:   "name2",
				status: "完了",
			},
			out: true,
		},
		{
			name: "case03",
			issue: issueItem{
				name:   "name3",
				status: "処理済み",
			},
			out: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.issue.isDone()
			if tt.out != actual {
				t.Errorf("Expected is %v but actual is %v", tt.out, actual)
			}
		})
	}
}

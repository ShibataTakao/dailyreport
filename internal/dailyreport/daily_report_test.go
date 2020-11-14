package dailyreport

import (
	"testing"
	"time"
)

func TestGetDailyReportFilePath(t *testing.T) {
	tests := []struct {
		name               string
		dailyReportDirPath string
		date               time.Time
		out                string
	}{
		{
			name:               "case01",
			dailyReportDirPath: "/home/user/dailyreport",
			date:               newDate(2020, 11, 1),
			out:                "/home/user/dailyreport/20201101.md",
		},
	}

	for _, tt := range tests {
		t.Run(tt.out, func(t *testing.T) {
			dailyReportFilePath := getDailyReportFilePath(tt.dailyReportDirPath, tt.date)
			if dailyReportFilePath != tt.out {
				t.Errorf("Expected is %v but actual is %v", tt.out, dailyReportFilePath)
			}
		})
	}

}

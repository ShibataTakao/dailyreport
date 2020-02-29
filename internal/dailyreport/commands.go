package dailyreport

import (
	"fmt"
)

// Create today's daily report
func Create(dailyReportDirPath string) error {
	dst := getDailyReportFilePath(dailyReportDirPath)
	if isFileExists(dst) {
		return fmt.Errorf("今日の日報 %s は既に存在します。", dst)
	}
	createDailyReport(dailyReportDirPath)
	fmt.Printf("今日の日報 %s を作成しました。\n", dst)
	return nil
}

// Validate worktime in today's daily report
func Validate(dailyReportDirPath string) error {
	report, err := readDailyReport(dailyReportDirPath)
	if err != nil {
		return err
	}
	fmt.Printf("業務時間 = %.2fh\n", report.timecardWorktime().Hours())
	fmt.Printf("今日のタスク（予定） = %.2fh\n", report.expectWorktime().Hours())
	fmt.Printf("今日のタスク（実績） = %.2fh\n", report.actualWorktime().Hours())
	return nil
}

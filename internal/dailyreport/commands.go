package dailyreport

import (
	"fmt"
	"time"
)

// Create today's daily report
func Create(dailyReportDirPath string) error {
	path := getDailyReportFilePath(dailyReportDirPath, time.Now())
	if isFileExists(path) {
		return fmt.Errorf("今日の日報 %s は既に存在します。", path)
	}
	createDailyReport(path)
	fmt.Printf("今日の日報 %s を作成しました。\n", path)
	return nil
}

// Validate worktime in today's daily report
func Validate(dailyReportDirPath string) error {
	path := getDailyReportFilePath(dailyReportDirPath, time.Now())
	report, err := readDailyReport(path)
	if err != nil {
		return err
	}
	fmt.Printf("業務時間 = %.2fh\n", report.timecardWorktime().Hours())
	fmt.Printf("今日のタスク（予定） = %.2fh\n", report.expectWorktime().Hours())
	fmt.Printf("今日のタスク（実績） = %.2fh\n", report.actualWorktime().Hours())
	return nil
}

// Aggregate tasks in some daily reports
func Aggregate(dailyReportDirPath, from, to string) error {
	fromDate, err := time.Parse("20060102", from)
	if err != nil {
		return err
	}
	toDate, err := time.Parse("20060102", to)
	if err != nil {
		return err
	}
	if fromDate.After(toDate) {
		fromDate, toDate = toDate, fromDate
	}
	reports, err := readDailyReports(dailyReportDirPath, fromDate, toDate)
	if err != nil {
		return err
	}
	for _, report := range reports {
		fmt.Printf("Found: %s\n", report.path)
	}
	for _, c := range reports.categories() {
		fmt.Printf("\n[%s]\n", c)
		for _, t := range reports.tasksByCategory(c).aggregated() {
			fmt.Printf("%.2fh / %.2fh  %s\n", t.expectTime.Hours(), t.actualTime.Hours(), t.name)
		}
	}
	return nil
}

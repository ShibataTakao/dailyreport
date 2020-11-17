package dailyreport

import (
	"fmt"
	"time"
)

// Create today's daily report
func Create(dirPath string) error {
	path := getDailyReportFilePath(dirPath, time.Now())
	if isFileExists(path) {
		return fmt.Errorf("今日の日報 %s は既に存在します。", path)
	}
	createDailyReport(path)
	fmt.Printf("今日の日報 %s を作成しました。\n", path)
	return nil
}

// Validate worktime in today's daily report
func Validate(dirPath, filePath string) error {
	path := getDailyReportFilePath(dirPath, time.Now())
	if filePath != "" {
		path = filePath
	}
	report, err := readDailyReport(path)
	if err != nil {
		return err
	}
	fmt.Printf("業務時間 = %.2fh\n", report.timecardWorktime().Hours())
	fmt.Printf("今日のタスク（予定） = %.2fh\n", report.expectWorktime().Hours())
	fmt.Printf("今日のタスク（実績） = %.2fh\n", report.actualWorktime().Hours())
	return nil
}

// Analyze show aggregation of daily reports
func Analyze(dirPath, from, to string) error {
	now := time.Now()
	fromDate := lastDate(now, time.Monday)
	if from != "" {
		var err error
		fromDate, err = time.Parse("20060102", from)
		if err != nil {
			return err
		}
	}
	toDate := nextDate(now, time.Friday)
	if to != "" {
		var err error
		toDate, err = time.Parse("20060102", to)
		if err != nil {
			return err
		}
	}
	if fromDate.After(toDate) {
		fromDate, toDate = toDate, fromDate
	}
	reports, err := readDailyReports(dirPath, fromDate, toDate)
	if err != nil {
		return err
	}

	for _, report := range reports {
		fmt.Printf("Found: %s\n", report.path)
	}

	fmt.Printf("\n## 業務時間\n")
	fmt.Printf("- 実績 %.2fh\n", reports.tasks().actualWorktime().Hours())

	fmt.Printf("\n## タスク（予定/実績）\n")
	for _, c := range reports.categories() {
		tasks := reports.tasksByCategory(c).aggregated()
		fmt.Printf("- [ ] %.2fh / %.2fh  %s\n", tasks.expectWorktime().Hours(), tasks.actualWorktime().Hours(), c)
		for _, t := range tasks {
			fmt.Printf("    - [ ] %.2fh / %.2fh  %s\n", t.expectTime.Hours(), t.actualTime.Hours(), t.name)
		}
	}
	return nil
}

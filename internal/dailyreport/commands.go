package dailyreport

import (
	"fmt"
	"strings"
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

// Report tasks in specific category from daily reports and issues
func Report(dirPath, category, startStr, endStr, trelloAppKey, trelloToken, trelloQueries string) error {
	start, err := time.Parse("20060102", startStr)
	if err != nil {
		return err
	}
	end, err := time.Parse("20060102", endStr)
	if err != nil {
		return err
	}
	if start.After(end) {
		start, end = end, start
	}

	reports, err := readDailyReports(dirPath, start, end)
	if err != nil {
		return err
	}
	for _, report := range reports {
		fmt.Printf("Found: %s\n", report.path)
	}
	tasks := reports.tasksByCategory(category).aggregated()

	issueClient := newIssueClient(trelloAppKey, trelloToken)
	issues, err := issueClient.fetchIssuesbyQueries(strings.Split(trelloQueries, ","))
	if err != nil {
		return err
	}
	tasks = tasks.mergeIssues(issues)

	fmt.Printf("\n## 工数（実績）\n")
	fmt.Printf("- %.2fh\n", tasks.actualWorktime().Hours())

	fmt.Printf("\n## タスク（予定/実績）\n")
	for _, task := range tasks {
		status := " "
		if task.isDone() {
			status = "x"
		}
		fmt.Printf("- [%s] %.2fh / %.2fh  %s\n", status, task.expectTime.Hours(), task.actualTime.Hours(), task.name)
	}
	return nil
}

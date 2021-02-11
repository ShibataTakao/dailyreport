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
	fmt.Println("")
	tasks := reports.tasksByCategory(category).aggregated()

	issueClient, err := newIssueClient(trelloAppKey, trelloToken)
	if err != nil {
		return err
	}
	issues, err := issueClient.fetchIssuesbyQueries(strings.Split(trelloQueries, ","))
	if err != nil {
		return err
	}

	pairs, err := zipTasksAndIssues(tasks, issues)
	if err != nil {
		return err
	}

	fmt.Printf("## 工数（実績）\n")
	fmt.Printf("- %.2fh\n", tasks.actualWorktime().Hours())

	fmt.Printf("\n## タスク（予定/実績）\n")
	for _, pair := range pairs {
		status := " "
		if pair.isDone() {
			status = "x"
		}
		fmt.Printf("- [%s] %.2fh / %.2fh  %s\n", status, pair.expectTime.Hours(), pair.actualTime.Hours(), pair.name)
	}
	return nil
}

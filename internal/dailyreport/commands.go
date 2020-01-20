package dailyreport

import (
	"fmt"
	"io/ioutil"
	"os"
)

// Create today's daily report
func Create(templateFilePath string, dailyReportDirPath string) error {
	dst := todayDailyReportFilePath(dailyReportDirPath)
	if isFileExists(dst) {
		fmt.Printf("Today's daily report %s already exists.\n", dst)
	} else {
		os.Link(templateFilePath, dst)
		fmt.Printf("Today's daily report %s is created.\n", dst)
	}
	return nil
}

// Validate worktime in today's daily report
func Validate(dailyReportDirPath string) error {
	filepath := todayDailyReportFilePath(dailyReportDirPath)
	text, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}
	report, err := parse(string(text))
	if err != nil {
		return err
	}
	fmt.Printf("Timecard:\t%.2fh\n", worktimeInTimecard(report).Hours())
	fmt.Printf("Expect Task:\t%.2fh\n", expectWorktimeInTask(report).Hours())
	fmt.Printf("Actual Task:\t%.2fh\n", actualWorktimeInTask(report).Hours())
	return nil
}

// KKMS do something
func KKMS() error {
	return nil
}

package dailyreport

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

type dailyReports []dailyReport

type dailyReport struct {
	path     string
	timecard timecardItem
	tasks    taskItems
}

func getDailyReportFilePath(dirPath string, date time.Time) string {
	filename := date.Format("20060102") + ".md"
	return filepath.Join(dirPath, filename)
}

func createDailyReport(path string) error {
	out, err := os.Create(path)
	if err != nil {
		return err
	}
	io.WriteString(out, dailyReportTemplate)
	return nil
}

func readDailyReport(path string) (dailyReport, error) {
	text, err := ioutil.ReadFile(path)
	if err != nil {
		return dailyReport{}, err
	}
	timecard, tasks, err := parse(string(text))
	if err != nil {
		return dailyReport{}, err
	}
	return dailyReport{
		path:     path,
		timecard: timecard,
		tasks:    tasks,
	}, nil
}

func readDailyReports(dirPath string, from time.Time, to time.Time) (dailyReports, error) {
	reports := dailyReports{}
	current := from
	to = to.AddDate(0, 0, 1)
	for current.Before(to) {
		path := getDailyReportFilePath(dirPath, current)
		if isFileExists(path) {
			report, err := readDailyReport(path)
			if err != nil {
				return dailyReports{}, err
			}
			reports = append(reports, report)
		}
		current = current.AddDate(0, 0, 1)
	}
	return reports, nil
}

func (report dailyReport) timecardWorktime() time.Duration {
	return report.timecard.timecardWorktime()
}

func (report dailyReport) expectWorktime() time.Duration {
	return report.tasks.expectWorktime()
}

func (report dailyReport) actualWorktime() time.Duration {
	return report.tasks.actualWorktime()
}

func (reports dailyReports) tasks() taskItems {
	tasks := taskItems{}
	for _, report := range reports {
		tasks = append(tasks, report.tasks...)
	}
	return tasks
}

func (reports dailyReports) categories() []string {
	return reports.tasks().categories()
}

func (reports dailyReports) tasksByCategory(category string) taskItems {
	return reports.tasks().filteredByCategory(category)
}

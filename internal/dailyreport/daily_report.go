package dailyreport

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

type dailyReport struct {
	timecard timecard
	tasks    []task
}

type timecard struct {
	begin time.Time
	end   time.Time
	rest  time.Duration
}

type task struct {
	category   string
	name       string
	expectTime time.Duration
	actualTime time.Duration
}

func getDailyReportFilePath(dailyReportDirPath string) string {
	now := time.Now()
	filename := now.Format("20060102") + ".md"
	return filepath.Join(dailyReportDirPath, filename)
}

func createDailyReport(dailyReportDirPath string) error {
	filepath := getDailyReportFilePath(dailyReportDirPath)
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	io.WriteString(out, dailyReportTemplate)
	return nil
}

func readDailyReport(dailyReportDirPath string) (dailyReport, error) {
	filepath := getDailyReportFilePath(dailyReportDirPath)
	text, err := ioutil.ReadFile(filepath)
	if err != nil {
		return dailyReport{}, err
	}
	return parse(string(text))
}

func (report dailyReport) timecardWorktime() time.Duration {
	return report.timecard.end.Sub(report.timecard.begin.Add(report.timecard.rest))
}

func (report dailyReport) expectWorktime() time.Duration {
	d := time.Duration(0)
	for _, task := range report.tasks {
		d += task.expectTime
	}
	return d
}

func (report dailyReport) actualWorktime() time.Duration {
	d := time.Duration(0)
	for _, task := range report.tasks {
		d += task.actualTime
	}
	return d
}

package dailyreport

import (
	"os"
	"path/filepath"
	"time"
)

func todayDailyReportFilePath(dailyReportDirPath string) string {
	now := time.Now()
	filename := now.Format("20060102") + ".md"
	return filepath.Join(dailyReportDirPath, filename)
}

func isFileExists(filepath string) bool {
	_, err := os.Stat(filepath)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}

func newTime(hour, min, sec int) time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), hour, min, sec, 0, now.Location())
}

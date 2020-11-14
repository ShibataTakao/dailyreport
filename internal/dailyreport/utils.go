package dailyreport

import (
	"os"
	"time"
)

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

func newDate(year int, month time.Month, day int) time.Time {
	now := time.Now()
	return time.Date(year, month, day, 0, 0, 0, 0, now.Location())
}

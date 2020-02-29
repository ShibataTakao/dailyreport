package dailyreport

import (
	"bufio"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func parse(text string) (dailyReport, error) {
	reBegin := regexp.MustCompile(`- 始業 \d{2}:\d{2}`)
	reEnd := regexp.MustCompile(`- 終業 \d{2}:\d{2}`)
	reRest := regexp.MustCompile(`- 休憩 \d{2}:\d{2}`)
	reCategory := regexp.MustCompile(`- \[.\] (.+)`)
	reTask := regexp.MustCompile(`- \[.\] \d+\.\d+h/\d+\.\d+h .+`)

	report := dailyReport{}
	scanner := bufio.NewScanner(strings.NewReader(text))
	currentCategory := ""
	for scanner.Scan() {
		line := scanner.Text()
		if reBegin.MatchString(line) {
			t, err := parseTime(line)
			if err != nil {
				return dailyReport{}, err
			}
			report.timecard.begin = t
		} else if reEnd.MatchString(line) {
			t, err := parseTime(line)
			if err != nil {
				return dailyReport{}, err
			}
			report.timecard.end = t
		} else if reRest.MatchString(line) {
			t, err := parseTime(line)
			if err != nil {
				return dailyReport{}, err
			}
			report.timecard.rest = time.Duration(t.Hour())*time.Hour + time.Duration(t.Minute())*time.Minute
		} else if reTask.MatchString(line) {
			task, err := parseTask(line, currentCategory)
			if err != nil {
				return dailyReport{}, err
			}
			report.tasks = append(report.tasks, task)
		} else if reCategory.MatchString(line) {
			currentCategory = reCategory.FindStringSubmatch(line)[1]
		} else if line == "---" {
			break
		}
	}
	return report, nil
}

func parseTime(s string) (time.Time, error) {
	re := regexp.MustCompile(`(\d{2}):(\d{2})`)
	match := re.FindStringSubmatch(s)
	hour, err := strconv.Atoi(match[1])
	if err != nil {
		return time.Now(), err
	}
	min, err := strconv.Atoi(match[2])
	if err != nil {
		return time.Now(), err
	}
	return newTime(hour, min, 0), nil
}

func parseTask(s string, category string) (task, error) {
	re := regexp.MustCompile(`- \[.\] (\d+.\d+h)/(\d+.\d+h) (.+)`)
	match := re.FindStringSubmatch(s)
	expect, err := time.ParseDuration(match[1])
	if err != nil {
		return task{}, nil
	}
	actual, err := time.ParseDuration(match[2])
	if err != nil {
		return task{}, nil
	}
	name := match[3]
	return task{
		category:   category,
		name:       name,
		expectTime: expect,
		actualTime: actual,
	}, nil
}

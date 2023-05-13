package dailyreport

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/ShibataTakao/worklog/task"
)

const (
	// StartAtLine means that current line in daily report has starting work time.
	StartAtLine LineKind = iota
	// EndAtLine means that current line in daily report has ending work time.
	EndAtLine
	// BreakTimeLine means that current line in daily report has break time.
	BreakTimeLine
	// BreakTimeLine means that current line in daily report has project of tasks.
	ProjectLine
	// BreakTimeLine means that current line in daily report has task.
	TaskLine
	// BreakTimeLine means that current line in daily report has horizontal rule.
	HorizontalRuleLine
	// BreakTimeLine means that kind of current line in daily report is unknown.
	UnknownLine
)

var (
	reStartAt        = regexp.MustCompile(`- 始業 (\d{2}):(\d{2})`)
	reEndAt          = regexp.MustCompile(`- 終業 (\d{2}):(\d{2})`)
	reBreakTime      = regexp.MustCompile(`- 休憩 (\d{2}):(\d{2})`)
	rePproject       = regexp.MustCompile(`^- \[.\] (.+)`)
	reTask           = regexp.MustCompile(`- \[(.)\] (\d+\.\d+h)/(\d+\.\d+h) (.+)`)
	reHorizontalRule = regexp.MustCompile(`---`)
)

// Parser parse daily report.
type Parser struct {
	scanner *bufio.Scanner
	date    time.Time
}

// LineKind is kind of line in daily report.
type LineKind int

// NewParser return new parser instance.
func NewParser(text string, date time.Time) *Parser {
	return &Parser{
		scanner: bufio.NewScanner(strings.NewReader(text)),
		date:    date,
	}
}

// Parse daily report.
func (p *Parser) Parse() (DailyReport, error) {
	var startAt, endAt time.Time
	var breakTime time.Duration
	var project task.Project
	var tasks task.Set
	for p.Next() && p.Kind() != HorizontalRuleLine {
		var err error
		switch p.Kind() {
		case StartAtLine:
			startAt, err = p.StartAt()
		case EndAtLine:
			endAt, err = p.EndAt()
		case BreakTimeLine:
			breakTime, err = p.BreakTime()
		case ProjectLine:
			project, err = p.Project()
		case TaskLine:
			var task task.Task
			task, err = p.Task(project)
			tasks = append(tasks, task)
		}
		if err != nil {
			return DailyReport{}, err
		}
	}
	return NewDailyReport(NewWorkTime(startAt, endAt, breakTime), tasks), nil
}

// Next advance parser to next line.
func (p *Parser) Next() bool {
	return p.scanner.Scan()
}

// Line return current line.
func (p *Parser) Line() string {
	return p.scanner.Text()
}

// Kind return kind of current line.
func (p *Parser) Kind() LineKind {
	switch {
	case reStartAt.MatchString(p.Line()):
		return StartAtLine
	case reEndAt.MatchString(p.Line()):
		return EndAtLine
	case reBreakTime.MatchString(p.Line()):
		return BreakTimeLine
	case rePproject.MatchString(p.Line()):
		return ProjectLine
	case reTask.MatchString(p.Line()):
		return TaskLine
	case reHorizontalRule.MatchString(p.Line()):
		return HorizontalRuleLine
	default:
		return UnknownLine
	}
}

// StartAt parse current line and return starting work time in daily report.
func (p *Parser) StartAt() (time.Time, error) {
	if p.Kind() != StartAtLine {
		return time.Time{}, fmt.Errorf("'%s' is not StartAtLine", p.Line())
	}
	matches := reStartAt.FindStringSubmatch(p.Line())
	if len(matches) != 3 {
		return time.Time{}, fmt.Errorf("matches for startAt must have 3 elements but actual is %v", matches)
	}
	hour, err := strconv.Atoi(matches[1])
	if err != nil {
		return time.Time{}, err
	}
	minute, err := strconv.Atoi(matches[2])
	if err != nil {
		return time.Time{}, err
	}
	return time.Date(p.date.Year(), p.date.Month(), p.date.Day(), hour, minute, 0, 0, time.Local), nil
}

// EndAt parse current line and return ending work time in daily report.
func (p *Parser) EndAt() (time.Time, error) {
	if p.Kind() != EndAtLine {
		return time.Time{}, fmt.Errorf("'%s' is not EndAtLine", p.Line())
	}
	matches := reEndAt.FindStringSubmatch(p.Line())
	if len(matches) != 3 {
		return time.Time{}, fmt.Errorf("matches for endAt must have 3 elements but actual is %v", matches)
	}
	hour, err := strconv.Atoi(matches[1])
	if err != nil {
		return time.Time{}, err
	}
	minute, err := strconv.Atoi(matches[2])
	if err != nil {
		return time.Time{}, err
	}
	return time.Date(p.date.Year(), p.date.Month(), p.date.Day(), hour, minute, 0, 0, time.Local), nil
}

// StartAt parse current line and return break time duration in daily report.
func (p *Parser) BreakTime() (time.Duration, error) {
	if p.Kind() != BreakTimeLine {
		return 0, fmt.Errorf("'%s' is not BreakTimeLine", p.Line())
	}
	matches := reBreakTime.FindStringSubmatch(p.Line())
	if len(matches) != 3 {
		return 0, fmt.Errorf("matches for breakTime must have 3 elements but actual is %v", matches)
	}
	hour, err := strconv.Atoi(matches[1])
	if err != nil {
		return 0, err
	}
	minute, err := strconv.Atoi(matches[2])
	if err != nil {
		return 0, err
	}
	return time.Duration(hour)*time.Hour + time.Duration(minute)*time.Minute, nil
}

// Project parse current line and return project of tasks in daily report.
func (p *Parser) Project() (task.Project, error) {
	if p.Kind() != ProjectLine {
		return task.Project{}, fmt.Errorf("'%s' is not ProjectLine", p.Line())
	}
	matches := rePproject.FindStringSubmatch(p.Line())
	if len(matches) != 2 {
		return task.Project{}, fmt.Errorf("matches for project must have 2 elements but actual is %v", matches)
	}
	return newProject(matches[1]), nil
}

// Task parse current line and return task in daily report.
func (p *Parser) Task(project task.Project) (task.Task, error) {
	if p.Kind() != TaskLine {
		return task.Task{}, fmt.Errorf("'%s' is not TaskLine", p.Line())
	}
	matches := reTask.FindStringSubmatch(p.Line())
	if len(matches) != 5 {
		return task.Task{}, fmt.Errorf("matches for task must have 6 elements but actual is %v", matches)
	}
	isCompleted := matches[1] == "x"
	estimate, err := time.ParseDuration(matches[2])
	if err != nil {
		return task.Task{}, err
	}
	actual, err := time.ParseDuration(matches[3])
	if err != nil {
		return task.Task{}, err
	}
	taskName := matches[4]
	return newTask(taskName, project, estimate, actual, isCompleted), nil
}

// newProject return new project instance.
func newProject(name string) task.Project {
	return task.NewProject(name)
}

// newTask return new task instance.
func newTask(name string, project task.Project, estimate time.Duration, actual time.Duration, isCompleted bool) task.Task {
	return task.NewTask("", name, project, estimate, actual, isCompleted)
}

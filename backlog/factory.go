package backlog

import (
	"github.com/ShibataTakao/worklog/task"
	"github.com/kenzo0107/backlog"
	"golang.org/x/exp/slices"
)

// Factory.
type Factory struct{}

// NewFactory return new factory instance.
func NewFactory() *Factory {
	return &Factory{}
}

// NewProject return new project instance from backlog project.
func (f *Factory) NewProject(prj *backlog.Project) task.Project {
	return task.NewProject(*prj.Name)
}

// NewTask return new task instance from backlog issue.
func (f *Factory) NewTask(issue *backlog.Issue, prj task.Project) task.Task {
	isCompleted := slices.Contains(CompletedStatusIDs, *issue.Status.ID)
	return task.NewBacklogTask(*issue.IssueKey, *issue.Summary, prj, isCompleted)
}

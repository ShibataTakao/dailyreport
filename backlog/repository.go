package backlog

import (
	"encoding/json"

	"github.com/ShibataTakao/worklog/task"
	"github.com/kenzo0107/backlog"
)

var (
	// CompletedStatusIDs is array of backlog status ID which means that issue is completed.
	CompletedStatusIDs = []int{3, 4}
)

// Repository for backlog.
type Repository struct {
	client  *backlog.Client
	factory *Factory
	cache   map[int]task.Project
}

// GetTasksQuery is query to get tasks.
type GetTasksQuery []backlog.GetIssuesOptions

// NewRepository return new repository instance.
func NewRepository(apiKey string, endpoint string, factory *Factory) *Repository {
	return &Repository{
		client:  backlog.New(apiKey, endpoint),
		factory: factory,
		cache:   map[int]task.Project{},
	}
}

// NewGetTasksQuery unmarshal query string and return new GetTasksQuery instance.
func NewGetTasksQuery(query string) (GetTasksQuery, error) {
	var opts []backlog.GetIssuesOptions
	err := json.Unmarshal([]byte(query), &opts)
	if err != nil {
		return nil, err
	}
	return GetTasksQuery(opts), nil
}

// GetTasks get tasks from backlog.
func (r *Repository) GetTasks(query GetTasksQuery) (task.Set, error) {
	tasks := task.Set{}
	opts := []backlog.GetIssuesOptions(query)
	for _, opt := range opts {
		issues, err := r.client.GetIssues(&opt)
		if err != nil {
			return task.Set{}, err
		}
		for _, issue := range issues {
			prj, err := r.GetProject(*issue.ProjectID)
			if err != nil {
				return task.Set{}, err
			}
			t := r.factory.NewTask(issue, prj)
			tasks = append(tasks, t)
		}
	}
	return tasks, nil
}

// GetProject get project from backlog.
func (r *Repository) GetProject(id int) (task.Project, error) {
	if value, found := r.cache[id]; found {
		return value, nil
	}

	prj, err := r.client.GetProject(id)
	if err != nil {
		return task.Project{}, err
	}
	r.cache[id] = r.factory.NewProject(prj)
	return r.cache[id], nil
}

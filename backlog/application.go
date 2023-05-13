package backlog

import "github.com/ShibataTakao/worklog/task"

// ApplicationService for backlog.
type ApplicationService struct {
	repository *Repository
}

// NewApplicationService return new application service instance.
func NewApplicationService(repo *Repository) *ApplicationService {
	return &ApplicationService{
		repository: repo,
	}
}

// GetTasks get tasks from backlog.
func (a *ApplicationService) GetTasks(query string) (task.Set, error) {
	q, err := NewGetTasksQuery(query)
	if err != nil {
		return task.Set{}, err
	}
	return a.repository.GetTasks(q)
}

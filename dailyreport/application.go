package dailyreport

import (
	"time"
)

// ApplicationService for a daily report.
type ApplicationService struct {
	repository *Repository
}

// NewApplicationService returns a new application service instance.
func NewApplicationService(repo *Repository) *ApplicationService {
	return &ApplicationService{
		repository: repo,
	}
}

func (a *ApplicationService) Analyze(start time.Time, end time.Time) (Analysis, error) {
	reports := Analysis{}
	for now := start; now.Equal(end) || now.Before(end); now = now.Add(24 * time.Hour) {
		if !a.repository.Exists(now) {
			continue
		}
		report, err := a.repository.Read(now)
		if err != nil {
			return Analysis{}, err
		}
		reports = append(reports, report)
	}
	return reports, nil
}

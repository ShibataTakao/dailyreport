package dailyreport

import (
	"log"
	"time"
)

// ApplicationService for daily report.
type ApplicationService struct {
	repository *Repository
}

// NewApplicationService return new application service instance.
func NewApplicationService(repo *Repository) *ApplicationService {
	return &ApplicationService{
		repository: repo,
	}
}

// Read daily reports.
func (a *ApplicationService) Read(start time.Time, end time.Time) (Set, error) {
	reports := Set{}
	for now := start; now.Equal(end) || now.Before(end); now = now.Add(24 * time.Hour) {
		if !a.repository.Exists(now) {
			continue
		}
		report, err := a.repository.Read(now)
		log.Printf("Read daily report from %s", a.repository.Path(now))
		if err != nil {
			return Set{}, err
		}
		reports = append(reports, report)
	}
	return reports, nil
}

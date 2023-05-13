package dailyreport

import (
	"os"
	"path"
	"time"
)

// Repository for daily report.
type Repository struct {
	dir string
}

// NewRepository return new repository instance.
func NewRepository(dir string) *Repository {
	return &Repository{
		dir: dir,
	}
}

// Read daily report.
func (r *Repository) Read(date time.Time) (DailyReport, error) {
	path := r.Path(date)
	body, err := os.ReadFile(path)
	if err != nil {
		return DailyReport{}, err
	}
	return NewParser(string(body), date).Parse()
}

// Exists return true if daily report on provided date exists.
func (r *Repository) Exists(date time.Time) bool {
	path := r.Path(date)
	return exists(path)
}

// Path of daily report file on provided date.
func (r *Repository) Path(date time.Time) string {
	filename := date.Format("20060102.md")
	return path.Join(r.dir, filename)
}

// exists return true if file or directory exists in provided path.
func exists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

package dailyreport

import "strings"

// Project.
type Project struct {
	Name string
}

// NewProject return new project instance.
func NewProject(name string) Project {
	return Project{
		Name: strings.TrimSpace(name),
	}
}

// Equals return true if both are same project.
func (p Project) Equals(other Project) bool {
	return p.Name == other.Name
}

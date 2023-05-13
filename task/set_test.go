package task

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestSetUnion(t *testing.T) {
	tests := []struct {
		name string
		in   Set
		out  Set
	}{
		{
			name: "Union",
			in: Set{
				NewTask("A-1", "Task A", NewProject("Project A"), 1, 2, false),
				NewTask("B-1", "Task A", NewProject("Project B"), 3, 4, false),
				NewTask("A-1", "Task A", NewProject("Project A"), 5, 6, true),
			},
			out: Set{
				NewTask("A-1", "Task A", NewProject("Project A"), 6, 8, true),
				NewTask("B-1", "Task A", NewProject("Project B"), 3, 4, false),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := require.New(t)
			union, err := tt.in.Union()
			assert.NoError(err)
			assert.Equal(tt.out, union)
		})
	}
}

func TestSetEstimate(t *testing.T) {
	tests := []struct {
		name     string
		set      Set
		estimate time.Duration
	}{
		{
			name: "Estimate",
			set: Set{
				NewTask("A-1", "Task A", NewProject("Project A"), 1, 2, false),
				NewTask("B-1", "Task A", NewProject("Project B"), 3, 4, false),
				NewTask("A-1", "Task A", NewProject("Project A"), 5, 6, true),
			},
			estimate: 9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := require.New(t)
			assert.Equal(tt.estimate, tt.set.Estimate())
		})
	}
}

func TestSetActual(t *testing.T) {
	tests := []struct {
		name   string
		set    Set
		actual time.Duration
	}{
		{
			name: "Actual",
			set: Set{
				NewTask("A-1", "Task A", NewProject("Project A"), 1, 2, false),
				NewTask("B-1", "Task A", NewProject("Project B"), 3, 4, false),
				NewTask("A-1", "Task A", NewProject("Project A"), 5, 6, true),
			},
			actual: 12,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := require.New(t)
			assert.Equal(tt.actual, tt.set.Actual())
		})
	}
}

func TestSetFilter(t *testing.T) {
	tests := []struct {
		name   string
		in     Set
		filter func(Task) bool
		out    Set
	}{
		{
			name: "Filter",
			in: Set{
				NewTask("A-1", "Task A", NewProject("Project A"), 1, 2, false),
				NewTask("B-1", "Task A", NewProject("Project B"), 3, 4, false),
				NewTask("A-1", "Task A", NewProject("Project A"), 5, 6, true),
			},
			filter: func(task Task) bool {
				prj := NewProject("Project A")
				return task.Project.Equals(prj)
			},
			out: Set{
				NewTask("A-1", "Task A", NewProject("Project A"), 1, 2, false),
				NewTask("A-1", "Task A", NewProject("Project A"), 5, 6, true),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := require.New(t)
			assert.Equal(tt.out, tt.in.Filter(tt.filter))
		})
	}

}

func TestSetSort(t *testing.T) {
	tests := []struct {
		name string
		in   Set
		less func(Task, Task) bool
		out  Set
	}{
		{
			name: "Filter",
			in: Set{
				NewTask("A-1", "Task A", NewProject("Project A"), 1, 2, false),
				NewTask("B-1", "Task A", NewProject("Project B"), 3, 4, false),
				NewTask("A-2", "Task B", NewProject("Project A"), 5, 6, true),
			},
			less: func(a Task, b Task) bool {
				return a.Key < b.Key
			},
			out: Set{
				NewTask("A-1", "Task A", NewProject("Project A"), 1, 2, false),
				NewTask("A-2", "Task B", NewProject("Project A"), 5, 6, true),
				NewTask("B-1", "Task A", NewProject("Project B"), 3, 4, false),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := require.New(t)
			assert.Equal(tt.out, tt.in.Sort(tt.less))
		})
	}

}

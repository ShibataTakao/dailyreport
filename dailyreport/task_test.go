package dailyreport

import (
	"cmp"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestTaskEquals(t *testing.T) {
	tests := []struct {
		name   string
		t1     Task
		t2     Task
		equals bool
	}{
		{
			name:   "TasksAreSame",
			t1:     newTask("Task A", NewProject("Project A"), 0, 1, false),
			t2:     newTask("Task A", NewProject("Project A"), 2, 3, true),
			equals: true,
		},
		{
			name:   "TaskNameIsDifferent",
			t1:     newTask("Task A", NewProject("Project A"), 0, 1, false),
			t2:     newTask("Task B", NewProject("Project A"), 2, 3, true),
			equals: false,
		},
		{
			name:   "ProjectIsDifferent",
			t1:     newTask("Task A", NewProject("Project A"), 0, 1, false),
			t2:     newTask("Task A", NewProject("Project B"), 2, 3, true),
			equals: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := require.New(t)
			assert.Equal(tt.equals, tt.t1.Equals(tt.t2))
		})
	}
}

func TestTaskMerge(t *testing.T) {
	tests := []struct {
		name   string
		t1     Task
		t2     Task
		merged Task
	}{
		{
			name:   "TasksAreSame",
			t1:     newTask("Task A", NewProject("Project A"), 1, 2, false),
			t2:     newTask("Task A", NewProject("Project A"), 3, 4, true),
			merged: newTask("Task A", NewProject("Project A"), 4, 6, true),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := require.New(t)
			merged, err := tt.t1.Merge(tt.t2)
			assert.NoError(err)
			assert.Equal(tt.merged, merged)
		})
	}
}

func TestTaskSetUnion(t *testing.T) {
	tests := []struct {
		name string
		in   TaskSet
		out  TaskSet
	}{
		{
			name: "Union",
			in: TaskSet{
				newTask("Task A", NewProject("Project A"), 1, 2, false),
				newTask("Task A", NewProject("Project B"), 3, 4, false),
				newTask("Task A", NewProject("Project A"), 5, 6, true),
			},
			out: TaskSet{
				newTask("Task A", NewProject("Project A"), 6, 8, true),
				newTask("Task A", NewProject("Project B"), 3, 4, false),
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

func TestTaskSetEstimate(t *testing.T) {
	tests := []struct {
		name     string
		set      TaskSet
		estimate time.Duration
	}{
		{
			name: "Estimate",
			set: TaskSet{
				newTask("Task A", NewProject("Project A"), 1, 2, false),
				newTask("Task A", NewProject("Project B"), 3, 4, false),
				newTask("Task A", NewProject("Project A"), 5, 6, true),
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

func TestTaskSetActual(t *testing.T) {
	tests := []struct {
		name   string
		set    TaskSet
		actual time.Duration
	}{
		{
			name: "Actual",
			set: TaskSet{
				newTask("Task A", NewProject("Project A"), 1, 2, false),
				newTask("Task A", NewProject("Project B"), 3, 4, false),
				newTask("Task A", NewProject("Project A"), 5, 6, true),
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

func TestTaskSetFilter(t *testing.T) {
	tests := []struct {
		name   string
		in     TaskSet
		filter func(Task) bool
		out    TaskSet
	}{
		{
			name: "Filter",
			in: TaskSet{
				newTask("Task A", NewProject("Project A"), 1, 2, false),
				newTask("Task A", NewProject("Project B"), 3, 4, false),
				newTask("Task A", NewProject("Project A"), 5, 6, true),
			},
			filter: func(task Task) bool {
				prj := NewProject("Project A")
				return task.Project.Equals(prj)
			},
			out: TaskSet{
				newTask("Task A", NewProject("Project A"), 1, 2, false),
				newTask("Task A", NewProject("Project A"), 5, 6, true),
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

func TestTaskSetSort(t *testing.T) {
	tests := []struct {
		name string
		in   TaskSet
		less func(Task, Task) int
		out  TaskSet
	}{
		{
			name: "Filter",
			in: TaskSet{
				newTask("Task A", NewProject("Project A"), 1, 2, false),
				newTask("Task A", NewProject("Project B"), 3, 4, false),
				newTask("Task B", NewProject("Project A"), 5, 6, true),
			},
			less: func(a Task, b Task) int {
				if !a.Project.Equals(b.Project) {
					return cmp.Compare(a.Project.Name, b.Project.Name)
				}
				return cmp.Compare(a.Summary, b.Summary)
			},
			out: TaskSet{
				newTask("Task A", NewProject("Project A"), 1, 2, false),
				newTask("Task B", NewProject("Project A"), 5, 6, true),
				newTask("Task A", NewProject("Project B"), 3, 4, false),
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

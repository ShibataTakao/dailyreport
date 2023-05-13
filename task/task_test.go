package task

import (
	"testing"

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
			t1:     NewTask("A-1", "Task A", NewProject("Project A"), 0, 1, false),
			t2:     NewTask("A-1", "Task A", NewProject("Project A"), 2, 3, true),
			equals: true,
		},
		{
			name:   "TaskNameIsDifferent",
			t1:     NewTask("A-1", "Task A", NewProject("Project A"), 0, 1, false),
			t2:     NewTask("A-2", "Task B", NewProject("Project A"), 2, 3, true),
			equals: false,
		},
		{
			name:   "ProjectIsDifferent",
			t1:     NewTask("A-1", "Task A", NewProject("Project A"), 0, 1, false),
			t2:     NewTask("B-1", "Task A", NewProject("Project B"), 2, 3, true),
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
			t1:     NewTask("A-1", "Task A", NewProject("Project A"), 1, 2, false),
			t2:     NewTask("A-1", "Task A", NewProject("Project A"), 3, 4, true),
			merged: NewTask("A-1", "Task A", NewProject("Project A"), 4, 6, true),
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

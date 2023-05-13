package backlog

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewGetTasksQuery(t *testing.T) {
	tests := []struct {
		name     string
		queryStr string
		query    GetTasksQuery
	}{
		{
			name: "NewGetTasksQuery",
			queryStr: `
[
    {
        "projectIds": [1],
        "assigneeIds": [2],
        "milestoneIds": [3],
        "statusIds": [1, 2]
    },
    {
        "projectIds": [4],
        "assigneeIds": [5],
        "milestoneIds": [6],
        "statusIds": [3, 4]
    }
]
			`,
			query: GetTasksQuery{
				{
					ProjectIDs:   []int{1},
					AssigneeIDs:  []int{2},
					MilestoneIDs: []int{3},
					StatusIDs:    []int{1, 2},
				},
				{
					ProjectIDs:   []int{4},
					AssigneeIDs:  []int{5},
					MilestoneIDs: []int{6},
					StatusIDs:    []int{3, 4},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := require.New(t)
			query, err := NewGetTasksQuery(tt.queryStr)
			assert.NoError(err)
			assert.Equal(tt.query, query)
		})
	}
}

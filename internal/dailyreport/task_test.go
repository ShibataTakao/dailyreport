package dailyreport

import (
	"reflect"
	"testing"
	"time"
)

func TestExpectWorktime(t *testing.T) {
	tests := []struct {
		name  string
		tasks taskItems
		out   time.Duration
	}{
		{
			name: "case01",
			tasks: taskItems{
				taskItem{expectTime: time.Duration(1) * time.Hour},
				taskItem{expectTime: time.Duration(1)*time.Hour + time.Duration(30)*time.Minute},
			},
			out: time.Duration(2)*time.Hour + time.Duration(30)*time.Minute,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.tasks.expectWorktime()
			if actual != tt.out {
				t.Errorf("Expected is %v but actual is %v", tt.out, actual)
			}
		})
	}
}

func TestActualWorktime(t *testing.T) {
	tests := []struct {
		name  string
		tasks taskItems
		out   time.Duration
	}{
		{
			name: "case01",
			tasks: taskItems{
				{actualTime: time.Duration(1) * time.Hour},
				{actualTime: time.Duration(1)*time.Hour + time.Duration(30)*time.Minute},
			},
			out: time.Duration(2)*time.Hour + time.Duration(30)*time.Minute,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.tasks.actualWorktime()
			if actual != tt.out {
				t.Errorf("Expected is %v but actual is %v", tt.out, actual)
			}
		})
	}
}

func TestFilteredByCategory(t *testing.T) {
	tests := []struct {
		name     string
		tasks    taskItems
		category string
		out      taskItems
	}{
		{
			name: "case01",
			tasks: taskItems{
				{
					category: "cat1",
					name:     "name1-1",
				},
				{
					category: "cat2",
					name:     "name2-1",
				},
				{
					category: "cat1",
					name:     "name1-2",
				},
			},
			category: "cat1",
			out: taskItems{
				{
					category: "cat1",
					name:     "name1-1",
				},
				{
					category: "cat1",
					name:     "name1-2",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.tasks.filteredByCategory(tt.category)
			if !reflect.DeepEqual(actual, tt.out) {
				t.Errorf("Expected is %v but actual is %v", tt.out, actual)
			}
		})
	}
}

func TestCategories(t *testing.T) {
	tests := []struct {
		name       string
		tasks      taskItems
		categories []string
	}{
		{
			name: "case01",
			tasks: taskItems{
				{
					category: "cat1",
					name:     "name1-1",
				},
				{
					category: "cat2",
					name:     "name2-1",
				},
				{
					category: "cat1",
					name:     "name1-2",
				},
			},
			categories: []string{"cat1", "cat2"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.tasks.categories()
			if !reflect.DeepEqual(actual, tt.categories) {
				t.Errorf("Expected is %v but actual is %v", tt.categories, actual)
			}
		})
	}
}

func TestAggregated(t *testing.T) {
	tests := []struct {
		name  string
		tasks taskItems
		out   taskItems
	}{
		{
			name: "case01",
			tasks: taskItems{
				{
					category:   "cat1",
					name:       "name1",
					expectTime: 1 * time.Hour,
					actualTime: 1 * time.Hour,
				},
				{
					category:   "cat1",
					name:       "name1",
					expectTime: 1 * time.Hour,
					actualTime: 1 * time.Hour,
				},
				{
					category:   "cat1",
					name:       "name2",
					expectTime: 1 * time.Hour,
					actualTime: 1 * time.Hour,
				},
			},
			out: taskItems{
				{
					category:   "cat1",
					name:       "name1",
					expectTime: 2 * time.Hour,
					actualTime: 2 * time.Hour,
				},
				{
					category:   "cat1",
					name:       "name2",
					expectTime: 1 * time.Hour,
					actualTime: 1 * time.Hour,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.tasks.aggregated()
			if !reflect.DeepEqual(tt.out, actual) {
				t.Errorf("Expected is %v but actual is %v", tt.out, actual)
			}
		})
	}
}

func TestMergeIssues(t *testing.T) {
	tests := []struct {
		name   string
		tasks  taskItems
		issues []issueItem
		out    taskItems
	}{
		{
			name: "case01",
			tasks: taskItems{
				{
					name:       "name1",
					category:   "cat1",
					expectTime: 1 * time.Hour,
					actualTime: 1 * time.Hour,
				},
				{
					name:       "name3",
					category:   "cat3",
					expectTime: 3 * time.Hour,
					actualTime: 3 * time.Hour,
				},
			},
			issues: []issueItem{
				{
					name:      "name1",
					status:    "status1",
					createdAt: newTime(0, 0, 1),
				},
				{
					name:      "name2",
					status:    "status2",
					createdAt: newTime(0, 0, 2),
				},
			},
			out: taskItems{
				{
					name:       "name1",
					category:   "cat1",
					expectTime: 1 * time.Hour,
					actualTime: 1 * time.Hour,
					done:       false,
					createdAt:  newTime(0, 0, 1),
				},
				{
					name:       "name2",
					category:   "",
					expectTime: 0,
					actualTime: 0,
					done:       false,
					createdAt:  newTime(0, 0, 2),
				},
				{
					name:       "name3",
					category:   "cat3",
					expectTime: 3 * time.Hour,
					actualTime: 3 * time.Hour,
					done:       false,
					createdAt:  time.Now(),
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.tasks.mergeIssues(tt.issues)
			if len(tt.out) >= 3 && len(actual) >= 3 {
				tt.out[2].createdAt = actual[2].createdAt
			}
			if !reflect.DeepEqual(tt.out, actual) {
				t.Errorf("Expected is %v but actual is %v", tt.out, actual)
			}
		})
	}
}

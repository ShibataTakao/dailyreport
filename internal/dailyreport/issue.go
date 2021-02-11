package dailyreport

import (
	"sort"
	"strings"
	"time"

	"github.com/adlio/trello"
)

type issueClient struct {
	client *trello.Client
	appKey string
	token  string
}

type issueItem struct {
	name      string
	status    string
	createdAt time.Time
}

type taskIssuePair struct {
	name       string
	status     string
	expectTime time.Duration
	actualTime time.Duration
	createdAt  time.Time

	task  *taskItem
	issue *issueItem
}

func newIssueClient(trelloAppKey, trelloToken string) (issueClient, error) {
	return issueClient{
		client: trello.NewClient(trelloAppKey, trelloToken),
		appKey: trelloAppKey,
		token:  trelloToken,
	}, nil
}

func (c issueClient) fetchIssuesbyQueries(queries []string) ([]issueItem, error) {
	issues := []issueItem{}
	for _, query := range queries {
		i, err := c.fetchIssues(query)
		if err != nil {
			return nil, err
		}
		issues = append(issues, i...)
	}
	return issues, nil
}

func (c issueClient) fetchIssues(query string) ([]issueItem, error) {
	cards, err := c.client.SearchCards(query, map[string]string{
		"modelType":   "cards",
		"cards_limit": "1000",
		"card_list":   "true",
	})
	if err != nil {
		return nil, err
	}
	issues := []issueItem{}
	for _, card := range cards {
		issues = append(issues, issueItem{
			name:      card.Name,
			status:    card.List.Name,
			createdAt: card.CreatedAt(),
		})
	}
	return issues, nil
}

func zipTasksAndIssues(tasks taskItems, issues []issueItem) ([]taskIssuePair, error) {
	pairs := []taskIssuePair{}
	isTaskPaired := make([]bool, len(tasks))
	for _, issue := range issues {
		pair := taskIssuePair{
			name:       issue.name,
			status:     issue.status,
			expectTime: 0,
			actualTime: 0,
			createdAt:  issue.createdAt,
			task:       nil,
			issue:      &issue,
		}
		for i, task := range tasks {
			if !isTaskPaired[i] && strings.Contains(issue.name, task.name) {
				pair.expectTime = task.expectTime
				pair.actualTime = task.actualTime
				pair.task = &tasks[i]
				isTaskPaired[i] = true
				break
			}
		}
		pairs = append(pairs, pair)
	}
	for i, task := range tasks {
		if !isTaskPaired[i] {
			pairs = append(pairs, taskIssuePair{
				name:       task.name,
				status:     "",
				expectTime: task.expectTime,
				actualTime: task.actualTime,
				createdAt:  time.Now(),
				task:       &tasks[i],
				issue:      nil,
			})

		}
	}
	sort.Slice(pairs, func(i, j int) bool { return pairs[i].createdAt.Before(pairs[j].createdAt) })
	return pairs, nil
}

func (t taskIssuePair) isDone() bool {
	return t.status == "Resolved"
}

package dailyreport

import (
	"encoding/json"
	"time"

	"github.com/kenzo0107/backlog"
)

type issueClient struct {
	client *backlog.Client
}

type issueItem struct {
	name      string
	status    string
	createdAt time.Time
}

func newIssueClient(backlogApiKey, backlogBaseUrl string) issueClient {
	return issueClient{
		client: backlog.New(backlogApiKey, backlogBaseUrl),
	}
}

func (c issueClient) fetchIssues(queries string) ([]issueItem, error) {
	var options []backlog.GetIssuesOptions
	if err := json.Unmarshal([]byte(queries), &options); err != nil {
		return nil, err
	}

	issues := []issueItem{}
	for _, option := range options {
		backlogIssues, err := c.client.GetIssues(&option)
		if err != nil {
			return nil, err
		}
		for _, bi := range backlogIssues {
			issues = append(issues, issueItem{
				name:      *bi.Summary,
				status:    *bi.Status.Name,
				createdAt: bi.Created.Time,
			})
		}
	}

	return issues, nil
}

func (i issueItem) isDone() bool {
	return i.status == "処理済み" || i.status == "完了"
}

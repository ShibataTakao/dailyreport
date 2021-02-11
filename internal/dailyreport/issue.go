package dailyreport

import (
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

func newIssueClient(trelloAppKey, trelloToken string) issueClient {
	return issueClient{
		client: trello.NewClient(trelloAppKey, trelloToken),
		appKey: trelloAppKey,
		token:  trelloToken,
	}
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

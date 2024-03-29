package ticket

import (
	"net/http"
	"strings"

	"github.com/andygrunwald/go-jira"

	"github.com/gfleury/squaas/config"
	"github.com/gfleury/squaas/log"
)

type JiraTicket struct {
	api      TicketApi
	issue    *jira.Issue
	watchers *[]jira.User
}

type JiraApi struct {
	client *jira.Client
}

func NewJiraApi() TicketApi {
	j := &JiraApi{}
	var err error
	var httpClient *http.Client

	jiraUrl := config.GetConfig().GetString("ticket.jira.url")

	if config.GetConfig().GetString("ticket.jira.username") != "" {
		tp := jira.BasicAuthTransport{
			Username: config.GetConfig().GetString("ticket.jira.username"),
			Password: config.GetConfig().GetString("ticket.jira.password"),
		}
		httpClient = tp.Client()
	}

	j.client, err = jira.NewClient(httpClient, jiraUrl)
	if err != nil {
		log.Fatalf("Jira client initialization failed: %s", err.Error())
	}

	return j
}

func (t *JiraTicket) Valid(username string) bool {
	if strings.Contains(username, "@") {
		username = strings.Split(username, "@")[0]
	}

	// Check if Creator or Assginee
	if t.issue.Fields.Creator != nil && t.issue.Fields.Creator.Name == username {
		return true
	}
	if t.issue.Fields.Assignee != nil && t.issue.Fields.Assignee.Name == username {
		return true
	}

	// Check if at least the watchers
	for _, watcher := range *t.watchers {
		if watcher.Name == username {
			return true
		}
	}
	return false
}

func (t *JiraTicket) AddComment(comment string) error {
	c := &jira.Comment{
		Body: comment,
		// Visibility: jira.CommentVisibility{},
	}

	_, _, err := t.api.(*JiraApi).client.Issue.AddComment(t.issue.ID, c)
	return err
}

func (t *JiraTicket) Issue() *jira.Issue {
	return t.issue
}

func (j *JiraApi) GetTicket(id string) (Ticket, error) {
	issue, _, err := j.client.Issue.Get(id, nil)
	if err != nil {
		return &JiraTicket{api: j, issue: nil, watchers: nil}, err
	}

	watchers, _, err := j.client.Issue.GetWatchers(id)

	return &JiraTicket{api: j, issue: issue, watchers: watchers}, err
}

func (j *JiraApi) GetCommentFormat() string {
	return `
%s

{code:sql}
%s
{code}

Check it here: [%s]
	`
}

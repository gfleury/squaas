package ticket

import (
	"fmt"
	"gopkg.in/check.v1"
	"strings"
)

func (s *Suite) TestJiraValidGetTicket(c *check.C) {
	jiraAPI := NewJiraApi()

	t, err := jiraAPI.GetTicket("MESOS-3325")
	c.Assert(err, check.IsNil)

	c.Assert(t.Valid("cfortier"), check.Equals, true)

	err = t.AddComment("Test")
	c.Assert(err, check.ErrorMatches, "You do not have the permission to comment on this issue.: Request failed. Please analyze the request body for more details. Status code: 400")
}

func (s *Suite) TestJiraValidateIssue(c *check.C) {
	jiraAPI := NewJiraApi()

	t, err := jiraAPI.GetTicket("MESOS-3325")
	c.Assert(err, check.IsNil)

	c.Assert(t.(*JiraTicket).Issue(), check.NotNil)
}

func (s *Suite) TestJiraApiGetComment(c *check.C) {
	jiraAPI := NewJiraApi()

	comment := fmt.Sprintf(jiraAPI.GetCommentFormat(), "Find this", "and this", "and also this")

	c.Assert(strings.Contains(comment, "Find this"), check.Equals, true)
	c.Assert(strings.Contains(comment, "and this"), check.Equals, true)
	c.Assert(strings.Contains(comment, "and also this"), check.Equals, true)
}

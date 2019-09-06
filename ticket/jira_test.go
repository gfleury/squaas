package ticket

import (
	"gopkg.in/check.v1"
)

func (s *Suite) TestJiraValidGetTicket(c *check.C) {
	jiraAPI := NewJiraApi()

	t, err := jiraAPI.GetTicket("MESOS-3325")
	c.Assert(err, check.IsNil)

	c.Assert(t.Valid("cfortier"), check.Equals, true)

	err = t.AddComment("Test")
	c.Assert(err, check.ErrorMatches, "You do not have the permission to comment on this issue.: Request failed. Please analyze the request body for more details. Status code: 400")
}

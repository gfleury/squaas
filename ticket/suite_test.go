package ticket

import (
	"testing"

	"gopkg.in/check.v1"

	"github.com/gfleury/squaas/config"
)

type Suite struct {
}

func (s *Suite) SetUpSuite(c *check.C) {
	config.GetConfig().Set("ticket.jira.url", "https://issues.apache.org/jira/")
}

func (s *Suite) TearDownSuite(c *check.C) {
}

var _ = check.Suite(&Suite{})

func Test(t *testing.T) {
	check.TestingT(t)
}

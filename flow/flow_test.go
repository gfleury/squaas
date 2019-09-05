package flow

import (
	"gopkg.in/check.v1"

	"github.com/gfleury/squaas/config"
)

func (s *Suite) TestGetReadyQueries(c *check.C) {

	queries, err := GetReadyQueries()
	c.Assert(err, check.IsNil)

	c.Assert(len(queries), check.Equals, 2)
}

func (s *Suite) TestShouldBeApproved(c *check.C) {

	config.GetConfig().Set("flow.minApproved", 2)
	config.GetConfig().Set("flow.maxDisapproved", 1)

	queries, err := GetReadyQueries()
	c.Assert(err, check.IsNil)

	for _, q := range queries {
		c.Assert(q.ShouldBeApproved(), check.Equals, false)
	}
}

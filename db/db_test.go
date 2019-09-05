package db

import (
	"gopkg.in/check.v1"
)

func (s *S) TestStorageInit(c *check.C) {

	err := DBStorage.Init()

	c.Assert(err, check.IsNil)

	names, err := DBStorage.Connection().Session.DB("squaastest").CollectionNames()
	c.Check(err, check.IsNil)
	c.Check(names, check.DeepEquals, []string{"queries"})
}

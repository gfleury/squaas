package worker

import (
	"gopkg.in/check.v1"
	"gopkg.in/mgo.v2/bson"

	"github.com/gfleury/squaas/config"
	"github.com/gfleury/squaas/db"
	"github.com/gfleury/squaas/models"
)

func (s *Suite) TestGetReadyQueries(c *check.C) {

	s.AddQueryFlow(c)
	queries, err := GetReadyQueries()
	c.Assert(err, check.IsNil)

	c.Assert(len(queries), check.Equals, 2)

	s.DropQueryFlow(c)
}

func (s *Suite) TestShouldBeApproved(c *check.C) {
	s.AddQueryFlow(c)

	config.GetConfig().Set("flow.minApproved", 2)
	config.GetConfig().Set("flow.maxDisapproved", 1)

	queries, err := GetReadyQueries()
	c.Assert(err, check.IsNil)

	for _, q := range queries {
		c.Assert(q.ShouldBeApproved(), check.Equals, false)
	}

	config.GetConfig().Set("flow.minApproved", 1)
	config.GetConfig().Set("flow.maxDisapproved", 10)

	for _, q := range queries {
		c.Assert(q.ShouldBeApproved(), check.Equals, true)
	}

	s.DropQueryFlow(c)
}

func (s *Suite) AddQueryFlow(c *check.C) {
	query := &models.Query{
		Query:  "SELECT * FROM XTABLE;",
		Status: models.StatusReady,
		Approvals: []models.Approvals{
			{User: &models.User{Name: "admin"}, Approved: true},
			{User: &models.User{Name: "admin"}, Approved: true},
			{User: &models.User{Name: "admin"}, Approved: true},
			{User: &models.User{Name: "admin"}, Approved: true},
			{User: &models.User{Name: "admin"}, Approved: true},
			{User: &models.User{Name: "admin"}, Approved: true},
			{User: &models.User{Name: "admin"}, Approved: false},
		},
	}

	QueryDB := db.DBStorage.Connection().Model("Query")

	err, _ := QueryDB.New(query)
	c.Assert(err, check.IsNil)

	err = query.Save()
	c.Assert(err, check.IsNil)

	query = &models.Query{
		Query:  "SELECT * FROM XTABLE;",
		Status: models.StatusReady,
		Approvals: []models.Approvals{
			{User: &models.User{Name: "admin"}, Approved: false},
			{User: &models.User{Name: "admin"}, Approved: false},
			{User: &models.User{Name: "admin"}, Approved: false},
			{User: &models.User{Name: "admin"}, Approved: false},
			{User: &models.User{Name: "admin"}, Approved: false},
			{User: &models.User{Name: "admin"}, Approved: false},
			{User: &models.User{Name: "admin"}, Approved: true},
		},
	}

	QueryDB = db.DBStorage.Connection().Model("Query")

	err, _ = QueryDB.New(query)
	c.Assert(err, check.IsNil)

	err = query.Save()
	c.Assert(err, check.IsNil)
}

func (s *Suite) DropQueryFlow(c *check.C) {
	_, err := db.DBStorage.Connection().Model("Query").RemoveAll(bson.M{})
	c.Assert(err, check.IsNil)
}

func (s *Suite) TestShouldBeApprovedExtended(c *check.C) {
	query := &models.Query{
		Query:      "SELECT * FROM XTABLE;",
		Status:     models.StatusReady,
		ServerName: "extendedgood",
		Approvals: []models.Approvals{
			{User: &models.User{Name: "test-user@blah.net"}, Approved: true},
			{User: &models.User{Name: "admin"}, Approved: true},
			{User: &models.User{Name: "admin"}, Approved: true},
		},
	}

	QueryDB := db.DBStorage.Connection().Model("Query")

	err, _ := QueryDB.New(query)
	c.Assert(err, check.IsNil)

	err = query.Save()
	c.Assert(err, check.IsNil)

	queries, err := GetReadyQueries()
	c.Assert(err, check.IsNil)

	for _, q := range queries {
		c.Assert(q.ShouldBeApproved(), check.Equals, true)
	}

	_, err = db.DBStorage.Connection().Model("Query").RemoveAll(bson.M{})
	c.Assert(err, check.IsNil)
}

func (s *Suite) TestShouldNotBeApprovedExtended(c *check.C) {
	query := &models.Query{
		Query:      "SELECT * FROM XTABLE;",
		Status:     models.StatusReady,
		ServerName: "extendedgood",
		Approvals: []models.Approvals{
			{User: &models.User{Name: "admin-que"}, Approved: true},
			{User: &models.User{Name: "admin"}, Approved: true},
			{User: &models.User{Name: "admin"}, Approved: true},
		},
	}

	QueryDB := db.DBStorage.Connection().Model("Query")

	err, _ := QueryDB.New(query)
	c.Assert(err, check.IsNil)

	err = query.Save()
	c.Assert(err, check.IsNil)

	queries, err := GetReadyQueries()
	c.Assert(err, check.IsNil)

	for _, q := range queries {
		c.Assert(q.ShouldBeApproved(), check.Equals, false)
	}

	_, err = db.DBStorage.Connection().Model("Query").RemoveAll(bson.M{})
	c.Assert(err, check.IsNil)
}

package flow

import (
	"bytes"
	"github.com/gfleury/squaas/models"
	"gopkg.in/mgo.v2/bson"
	"testing"

	"gopkg.in/check.v1"

	"github.com/gfleury/squaas/config"
	"github.com/gfleury/squaas/db"
)

func Test(t *testing.T) {
	check.TestingT(t)
}

type Suite struct {
}

var _ = check.Suite(&Suite{})

func (s *Suite) SetUpSuite(c *check.C) {
	config.Init()

	var yamlExample = []byte(`
mongo:
  url: "mongodb://127.0.0.1:27017/squaastest"
`)

	err := config.GetConfig().ReadConfig(bytes.NewBuffer(yamlExample))
	c.Assert(err, check.IsNil)

	db.InitStorage()

	err = db.DBStorage.Init()
	c.Assert(err, check.IsNil)
}

func (s *Suite) SetUpTest(c *check.C) {
	query := &models.Query{
		Query:  "SELECT * FROM XTABLE;",
		Status: models.StatusReady,
		Approvals: []models.Approvals{
			{&models.User{Name: "admin"}, true},
			{&models.User{Name: "admin"}, true},
			{&models.User{Name: "admin"}, true},
			{&models.User{Name: "admin"}, true},
			{&models.User{Name: "admin"}, true},
			{&models.User{Name: "admin"}, true},
			{&models.User{Name: "admin"}, false},
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
			{&models.User{Name: "admin"}, false},
			{&models.User{Name: "admin"}, false},
			{&models.User{Name: "admin"}, false},
			{&models.User{Name: "admin"}, false},
			{&models.User{Name: "admin"}, false},
			{&models.User{Name: "admin"}, true},
			{&models.User{Name: "admin"}, false},
		},
	}

	QueryDB = db.DBStorage.Connection().Model("Query")

	err, _ = QueryDB.New(query)
	c.Assert(err, check.IsNil)

	err = query.Save()
	c.Assert(err, check.IsNil)
}

func (s *Suite) TearDownTest(c *check.C) {
	_, err := db.DBStorage.Connection().Model("Query").RemoveAll(bson.M{})
	c.Assert(err, check.IsNil)
}

func (s *Suite) TearDownSuite(c *check.C) {
	c.Assert(db.DBStorage.Connection(), check.NotNil)

	err := db.DBStorage.Connection().Session.DB("squaastest").DropDatabase()
	c.Assert(err, check.IsNil)
}

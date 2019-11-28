package worker

import (
	"bytes"
	"testing"

	check "gopkg.in/check.v1"

	"github.com/gfleury/squaas/config"
	"github.com/gfleury/squaas/db"
)

type Suite struct {
}

func (s *Suite) SetUpSuite(c *check.C) {

	var yamlExample = []byte(`
mongo:
  url: "mongodb://127.0.0.1:27017/squaastest"
databases:
  broken: "postgres://postgres@localhost:1025/data?sslmode=disable"
  good: "postgres://postgres@localhost/data?sslmode=disable"
  extendedgood: 
    uri: "postgres://postgres@localhost/data?sslmode=disable"
    approval_rule:
      required_users:
        - "test-user@blah.net"
        - "admin@blah.net"
      min_approved: 3
      max_disapproved: 1
`)

	err := config.GetConfig().ReadConfig(bytes.NewBuffer(yamlExample))
	c.Check(err, check.IsNil)

	db.InitStorage()

	err = db.DBStorage.Init()
	c.Check(err, check.IsNil)
}
func (s *Suite) TearDownSuite(c *check.C) {
	c.Assert(db.DBStorage.Connection(), check.NotNil)

	err := db.DBStorage.Connection().Session.DB("squaastest").DropDatabase()
	c.Check(err, check.IsNil)
}

var _ = check.Suite(&Suite{})

func Test(t *testing.T) {
	check.TestingT(t)
}

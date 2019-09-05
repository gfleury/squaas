package db

import (
	"bytes"
	"testing"

	"github.com/gfleury/squaas/config"

	"gopkg.in/check.v1"
)

func Test(t *testing.T) { check.TestingT(t) }

type S struct {
}

var _ = check.Suite(&S{})

func (s *S) SetUpSuite(c *check.C) {
	config.Init()

	var yamlExample = []byte(`
mongo:
  url: "mongodb://127.0.0.1:27017/squaastest"
`)

	err := config.GetConfig().ReadConfig(bytes.NewBuffer(yamlExample))
	c.Check(err, check.IsNil)

	InitStorage()

}

func (s *S) SetUpTest(c *check.C) {
}

func (s *S) TearDownTest(c *check.C) {
}

func (s *S) TearDownSuite(c *check.C) {
	c.Assert(DBStorage.Connection(), check.NotNil)

	err := DBStorage.Connection().Session.DB("squaastest").DropDatabase()
	c.Check(err, check.IsNil)
}

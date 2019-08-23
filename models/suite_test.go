package models

import (
	// "fmt"
	// "net/http"
	// "net/http/httptest"
	"testing"

	check "gopkg.in/check.v1"
)

type Suite struct {
}

func (s *Suite) SetUpSuite(c *check.C) {

}
func (s *Suite) TearDownSuite(c *check.C) {
}

var _ = check.Suite(&Suite{})

func Test(t *testing.T) {
	check.TestingT(t)
}

package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gfleury/dbquerybench/config"
	"github.com/gfleury/dbquerybench/db"

	"github.com/gin-gonic/gin"
	check "gopkg.in/check.v1"
)

type Suite struct {
	router *gin.Engine
}

func (s *Suite) SetUpSuite(c *check.C) {
	config.Init()

	s.router = NewRouter()
	db.InitStorage()
	err := db.DBStorage.Init()
	c.Assert(err, check.IsNil)
}
func (s *Suite) TearDownSuite(c *check.C) {
	// Clean test database
	if db.DBStorage.Connection() != nil {
		Pipeline := db.DBStorage.Connection().Model("Query")
		err := Pipeline.DropCollection()
		if err != nil {
			fmt.Println(err)
		}
	}
}

var _ = check.Suite(&Suite{})

func Test(t *testing.T) {
	check.TestingT(t)
}

func (s *Suite) executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	s.router.ServeHTTP(rr, req)

	return rr
}

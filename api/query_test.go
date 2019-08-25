package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gfleury/dbquerybench/models"
	"io/ioutil"
	"net/http"
	// "net/url"
	// "strings"

	check "gopkg.in/check.v1"
)

func (s *Suite) TestCreateQuery(c *check.C) {
	query := &models.Query{
		TicketID: "BLEH-330",
		// Owner:    models.User{Name: "admin@boom.org"},
		Query:      "SELECT * FROM TABLE;",
		ServerName: "db1.blah.com",
	}

	queryBytes, err := query.Byte()
	c.Assert(err, check.IsNil)

	req, _ := http.NewRequest("POST", "/v1/query", bytes.NewReader(queryBytes))
	response := s.executeRequest(req)

	c.Assert(response.Code, check.Equals, http.StatusOK)

	var p []byte
	p, err = ioutil.ReadAll(response.Body)
	c.Assert(err, check.IsNil)

	responseQuery := &models.Query{}
	err = json.Unmarshal(p, responseQuery)
	c.Assert(err, check.IsNil)

	c.Assert(responseQuery.TicketID, check.Equals, query.TicketID)
	c.Assert(responseQuery.Status, check.Equals, "pending")
	c.Assert(responseQuery.Query, check.Equals, query.Query)
	c.Assert(responseQuery.ServerName, check.Equals, query.ServerName)

	// Check GetQuery

	req, _ = http.NewRequest("GET", fmt.Sprintf("/v1/query/%s", responseQuery.Id), nil)
	response = s.executeRequest(req)

	c.Assert(response.Code, check.Equals, http.StatusOK)

	p, err = ioutil.ReadAll(response.Body)
	c.Assert(err, check.IsNil)

	err = json.Unmarshal(p, responseQuery)
	c.Assert(err, check.IsNil)

	c.Assert(responseQuery.TicketID, check.Equals, query.TicketID)
	c.Assert(responseQuery.Status, check.Equals, "pending")
	c.Assert(responseQuery.Query, check.Equals, query.Query)

}

func (s *Suite) TestDeleteQuery(c *check.C) {
	query := &models.Query{
		TicketID: "pipelineDelete",
		Query:    "SELECT * FROM TABLE;",
	}

	queryBytes, err := query.Byte()
	c.Assert(err, check.IsNil)

	req, _ := http.NewRequest("POST", "/v1/query", bytes.NewReader(queryBytes))
	response := s.executeRequest(req)

	c.Assert(response.Code, check.Equals, http.StatusOK)

	req, _ = http.NewRequest("DELETE", fmt.Sprintf("/v1/query/%s", query.Id), nil)
	response = s.executeRequest(req)

	c.Assert(response.Code, check.Equals, http.StatusOK)

	req, _ = http.NewRequest("DELETE", fmt.Sprintf("/v1/query/%s", query.Id), nil)
	response = s.executeRequest(req)

	c.Assert(response.Code, check.Equals, http.StatusNotFound)
}

func (s *Suite) TestUpdateQuery(c *check.C) {
	query := &models.Query{
		TicketID: "pipelineUpdate",
		Query:    "SELECT * FROM TABLE;",
	}

	queryBytes, err := query.Byte()
	c.Assert(err, check.IsNil)

	req, _ := http.NewRequest("POST", "/v1/query", bytes.NewReader(queryBytes))
	response := s.executeRequest(req)

	c.Assert(response.Code, check.Equals, http.StatusOK)

	responseQuery := &models.Query{}
	var p []byte

	p, err = ioutil.ReadAll(response.Body)
	c.Assert(err, check.IsNil)

	err = json.Unmarshal(p, responseQuery)
	c.Assert(err, check.IsNil)
	// Update pipeline

	responseQuery.Query = "SELECT * FROM TABLE2;"

	queryBytes, err = responseQuery.Byte()
	c.Assert(err, check.IsNil)
	req, _ = http.NewRequest("PUT", "/v1/query", bytes.NewReader(queryBytes))
	response = s.executeRequest(req)

	c.Assert(response.Code, check.Equals, http.StatusOK)

	// Get Pipeline
	req, _ = http.NewRequest("GET", fmt.Sprintf("/v1/query/%s", responseQuery.Id), nil)
	response = s.executeRequest(req)

	c.Assert(response.Code, check.Equals, http.StatusOK)

	p, err = ioutil.ReadAll(response.Body)
	c.Assert(err, check.IsNil)

	responseQueryNew := &models.Query{}
	err = json.Unmarshal(p, responseQueryNew)
	c.Assert(err, check.IsNil)

	c.Assert(responseQuery.TicketID, check.Equals, responseQueryNew.TicketID)
	c.Assert(responseQuery.Status, check.Equals, responseQueryNew.Status)
	c.Assert(responseQuery.Owner, check.DeepEquals, responseQueryNew.Owner)
	c.Assert(responseQuery.Query, check.DeepEquals, responseQueryNew.Query)
}

package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	check "gopkg.in/check.v1"

	"github.com/gfleury/squaas/db"
	"github.com/gfleury/squaas/models"
)

func (s *Suite) TestCreateQuery(c *check.C) {
	query := &models.Query{
		TicketID:   "BLEH-330",
		Query:      "SELECT * FROM XTABLE;",
		ServerName: "server1",
		Status:     models.StatusReady,
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
	c.Assert(responseQuery.Status, check.Equals, models.StatusReady)
	c.Assert(responseQuery.Query, check.Equals, query.Query)
	c.Assert(responseQuery.ServerName, check.Equals, query.ServerName)
	c.Assert(responseQuery.Owner, check.Equals, models.User{Name: "admin"})

	// Check GetQuery

	req, _ = http.NewRequest("GET", fmt.Sprintf("/v1/query/%s", responseQuery.Id.Hex()), nil)
	response = s.executeRequest(req)

	c.Assert(response.Code, check.Equals, http.StatusOK)

	p, err = ioutil.ReadAll(response.Body)
	c.Assert(err, check.IsNil)

	err = json.Unmarshal(p, responseQuery)
	c.Assert(err, check.IsNil)

	c.Assert(responseQuery.TicketID, check.Equals, query.TicketID)
	c.Assert(responseQuery.Status, check.Equals, models.StatusReady)
	c.Assert(responseQuery.Query, check.Equals, query.Query)
	c.Assert(responseQuery.Owner, check.Equals, models.User{Name: "admin"})
}

func (s *Suite) TestCreateQueryInvalidStatus(c *check.C) {
	query := &models.Query{
		TicketID:   "BLEH-330",
		Query:      "SELECT * FROM XTABLE;",
		ServerName: "server1",
		Status:     models.StatusApproved,
	}

	queryBytes, err := query.Byte()
	c.Assert(err, check.IsNil)

	req, _ := http.NewRequest("POST", "/v1/query", bytes.NewReader(queryBytes))
	response := s.executeRequest(req)

	c.Assert(response.Code, check.Equals, http.StatusBadRequest)

}

func (s *Suite) TestDeleteQuery(c *check.C) {
	query := &models.Query{
		TicketID:   "pipelineDelete",
		Query:      "SELECT * FROM XTABLE;",
		ServerName: "server1",
		Status:     models.StatusReady,
	}

	queryBytes, err := query.Byte()
	c.Assert(err, check.IsNil)

	req, _ := http.NewRequest("POST", "/v1/query", bytes.NewReader(queryBytes))
	response := s.executeRequest(req)

	c.Assert(response.Code, check.Equals, http.StatusOK)

	err = query.Parse(response.Body)
	c.Assert(err, check.IsNil)

	req, _ = http.NewRequest("DELETE", fmt.Sprintf("/v1/query/%s", query.Id.Hex()), nil)
	response = s.executeRequest(req)

	c.Assert(response.Code, check.Equals, http.StatusOK)

	req, _ = http.NewRequest("DELETE", fmt.Sprintf("/v1/query/%s", query.Id.Hex()), nil)
	response = s.executeRequest(req)

	c.Assert(response.Code, check.Equals, http.StatusNotFound)
}

func (s *Suite) TestDeleteQueryInvalidStatus(c *check.C) {
	query := &models.Query{
		TicketID:   "BLEH-330",
		Query:      "SELECT * FROM XTABLE;",
		ServerName: "server1",
		Status:     models.StatusReady,
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

	QueryDB := db.DBStorage.Connection().Model("Query")
	err = QueryDB.FindId(responseQuery.Id).Exec(query)
	c.Assert(err, check.IsNil)

	query.Status = models.StatusApproved
	err = query.Save()
	c.Assert(err, check.IsNil)

	req, _ = http.NewRequest("DELETE", fmt.Sprintf("/v1/query/%s", query.Id.Hex()), nil)
	response = s.executeRequest(req)

	c.Assert(response.Code, check.Equals, http.StatusBadRequest)
}

func (s *Suite) TestUpdateQuery(c *check.C) {
	query := &models.Query{
		TicketID:   "pipelineUpdate",
		Query:      "SELECT * FROM XTABLE;",
		ServerName: "server1",
		Status:     models.StatusReady,
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
	// Update Query

	responseQuery.Query = "SELECT * FROM TABLEXXX;"

	queryBytes, err = responseQuery.Byte()
	c.Assert(err, check.IsNil)
	req, _ = http.NewRequest("PUT", "/v1/query", bytes.NewReader(queryBytes))
	response = s.executeRequest(req)

	c.Assert(response.Code, check.Equals, http.StatusOK)

	// Get Query
	req, _ = http.NewRequest("GET", fmt.Sprintf("/v1/query/%s", responseQuery.Id.Hex()), nil)
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

func (s *Suite) TestGetDatabases(c *check.C) {
	req, _ := http.NewRequest("GET", "/v1/databases", nil)
	response := s.executeRequest(req)

	c.Assert(response.Code, check.Equals, http.StatusOK)

	servers := &models.Servers{}

	err := servers.Parse(response.Body)
	c.Assert(err, check.IsNil)

	c.Assert(len(*servers), check.Equals, 2)
	c.Assert(*servers, check.DeepEquals, models.Servers{
		{Name: "server1"},
		{Name: "server2"},
	})
}

func (s *Suite) TestApproveQuery(c *check.C) {

	query := &models.Query{
		TicketID:   "BLEH-330",
		Query:      "SELECT * FROM XTABLE;",
		ServerName: "server1",
		Status:     models.StatusReady,
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

	req, _ = http.NewRequest("POST", fmt.Sprintf("/v1/query/%s/approve", responseQuery.Id.Hex()), strings.NewReader("approved"))
	response = s.executeRequest(req)

	c.Assert(response.Code, check.Equals, http.StatusOK)

	req, _ = http.NewRequest("GET", fmt.Sprintf("/v1/query/%s", responseQuery.Id.Hex()), nil)
	response = s.executeRequest(req)

	c.Assert(response.Code, check.Equals, http.StatusOK)

	p, err = ioutil.ReadAll(response.Body)
	c.Assert(err, check.IsNil)

	responseQuery = &models.Query{}
	err = json.Unmarshal(p, responseQuery)
	c.Assert(err, check.IsNil)

	c.Assert(responseQuery.TicketID, check.Equals, query.TicketID)
	c.Assert(responseQuery.Status, check.Equals, models.StatusReady)
	c.Assert(responseQuery.Query, check.Equals, query.Query)
	c.Assert(responseQuery.ServerName, check.Equals, query.ServerName)
	c.Assert(responseQuery.Owner, check.Equals, models.User{Name: "admin"})
	c.Assert(responseQuery.Approvals, check.DeepEquals, []models.Approvals{{User: &models.User{Name: "admin"}, Approved: true}})

	req, _ = http.NewRequest("DELETE", fmt.Sprintf("/v1/query/%s/approve", responseQuery.Id.Hex()), strings.NewReader("approved"))
	response = s.executeRequest(req)

	c.Assert(response.Code, check.Equals, http.StatusOK)

	req, _ = http.NewRequest("GET", fmt.Sprintf("/v1/query/%s", responseQuery.Id.Hex()), nil)
	response = s.executeRequest(req)

	c.Assert(response.Code, check.Equals, http.StatusOK)

	p, err = ioutil.ReadAll(response.Body)
	c.Assert(err, check.IsNil)

	responseQuery = &models.Query{}
	err = json.Unmarshal(p, responseQuery)
	c.Assert(err, check.IsNil)

	c.Assert(responseQuery.TicketID, check.Equals, query.TicketID)
	c.Assert(responseQuery.Status, check.Equals, models.StatusReady)
	c.Assert(responseQuery.Query, check.Equals, query.Query)
	c.Assert(responseQuery.ServerName, check.Equals, query.ServerName)
	c.Assert(responseQuery.Owner, check.Equals, models.User{Name: "admin"})
	c.Assert(responseQuery.Approvals, check.DeepEquals, []models.Approvals{{User: &models.User{Name: "admin"}, Approved: false}})
}

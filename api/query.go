/*
 * DBworkBench
 */

package api

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"

	"github.com/gfleury/squaas/db"
	"github.com/gfleury/squaas/models"
	"github.com/gfleury/squaas/ticket"
)

func Index(c *gin.Context) {
	w := c.Writer

	fmt.Fprintf(w, "Hello World!")
}

func GetQueries(c *gin.Context) {
	var queries []*models.Query

	QueryDB := db.DBStorage.Connection().Model("Query")

	err := QueryDB.Find(bson.M{"deleted": false}).Exec(&queries)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if len(queries) == 0 {
		c.JSON(http.StatusOK, []models.Query{})
		return
	}

	c.JSON(http.StatusOK, queries)
}

func AddQuery(c *gin.Context) {
	var query models.Query

	ownerUser := c.MustGet(gin.AuthUserKey).(string)

	err := c.BindJSON(&query)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("AddQuery: Parsing query: %#v", query)

	// Lint SQL
	err = query.LintSQLQuery()
	if err != nil {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": err.Error()})
		return
	}

	// Return if only linting SQL
	if query.Status == models.StatusParseOnly {
		c.JSON(http.StatusOK, query)
		return
	}

	// Ticket Validation
	if query.TicketID == "" {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "You must insert a ticketID to link your query to"})
		return
	}

	ticket, err := ticket.TicketServive.GetTicket(query.TicketID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(query.TicketID) < 1 || !ticket.Valid(ownerUser) {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid ticketID, not reporter/asignee or watcher")})
		return
	}

	// Status Validation
	if !query.Status.Valid() {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid status: %s", query.Status)})
		return
	}

	if query.Status != models.StatusReady && query.Status != models.StatusPending && query.Status != models.StatusParseOnly {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Queries can only be created in Pending and Ready status, invalid status: %s", query.Status)})
		return
	}

	query.Owner.Name = ownerUser

	if query.ServerName == "" {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "You must select a database"})
		return
	}

	// ADD To database
	QueryDB := db.DBStorage.Connection().Model("Query")

	var requestMap map[string]interface{}
	err, requestMap = QueryDB.New(&query)
	if err != nil {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": err.Error()})
		return
	}

	if valid, issues := query.Validate(requestMap); valid {
		err = query.Save()
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": err.Error()})
			return
		}
	} else {
		var errorString string
		for _, err := range issues {
			errorString = fmt.Sprintf("%s Error: %s\n", errorString, err.Error())
		}
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": errorString})
		return
	}

	// Ticket add Comment (silently fail(logging))
	hostURL := ""
	pURL, err := url.Parse(c.Request.Referer())
	if err == nil {
		hostURL = pURL.Scheme + "://" + pURL.Host
	}
	err = query.TicketCommentAdded(hostURL)
	if err != nil {
		log.Printf("Adding comment to ticket failed on query %#v: %s", query, err.Error())
	}

	c.JSON(http.StatusOK, query)
}

func DeleteQuery(c *gin.Context) {
	var query models.Query

	QueryID := c.Param("queryId")

	if !models.IsValidObjectId(QueryID) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query ID"})
		return
	}

	QueryDB := db.DBStorage.Connection().Model("Query")

	err := QueryDB.FindId(bson.ObjectIdHex(QueryID)).Exec(&query)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if query.Deleted {
		c.JSON(http.StatusNotFound, gin.H{"error": "Query not found"})
		return
	}

	if query.Status != models.StatusReady && query.Status != models.StatusPending {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Queries can only be deleted in Pending and Ready status, invalid status for deleting: %s", query.Status)})
		return
	}

	err = query.Delete()
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": "Successfully deleted"})
}

func ApproveQuery(c *gin.Context) {
	var query models.Query

	QueryID := c.Param("queryId")

	if !models.IsValidObjectId(QueryID) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query ID"})
		return
	}

	QueryDB := db.DBStorage.Connection().Model("Query")

	err := QueryDB.FindId(bson.ObjectIdHex(QueryID)).Exec(&query)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if query.Status != models.StatusReady {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("You can only approve while the query is in Ready status, not in %s", query.Status)})
		return
	}

	userApproving := c.MustGet(gin.AuthUserKey).(string)

	query.AddApproval(&models.User{Name: userApproving}, true)

	err = query.Save()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": "Approved with success"})
}

func DeleteApprovalQuery(c *gin.Context) {
	var query models.Query

	QueryID := c.Param("queryId")

	if !models.IsValidObjectId(QueryID) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query ID"})
		return
	}

	QueryDB := db.DBStorage.Connection().Model("Query")

	err := QueryDB.FindId(bson.ObjectIdHex(QueryID)).Exec(&query)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if query.Status != models.StatusReady {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("You can only disapprove while the query is in Ready status, not in %s", query.Status)})
		return
	}

	userApproving := c.MustGet(gin.AuthUserKey).(string)

	query.AddApproval(&models.User{Name: userApproving}, false)

	err = query.Save()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": "Disapproved with success"})
}

func FindQueryByOwner(c *gin.Context) {
	w := c.Writer
	// r := c.Request

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func FindQueryByStatus(c *gin.Context) {
	w := c.Writer
	// r := c.Request

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func GetQueryById(c *gin.Context) {
	var query models.Query

	QueryID := c.Param("queryId")

	if !models.IsValidObjectId(QueryID) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query ID"})
		return
	}

	QueryDB := db.DBStorage.Connection().Model("Query")

	err := QueryDB.FindId(bson.ObjectIdHex(QueryID)).Exec(&query)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, query)
}

func UpdateQuery(c *gin.Context) {
	var queryUpdated, queryOriginal models.Query

	requestingUser := c.MustGet(gin.AuthUserKey).(string)

	err := c.BindJSON(&queryUpdated)
	if err != nil {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": err.Error()})
		return
	}

	ticket, err := ticket.TicketServive.GetTicket(queryUpdated.TicketID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(queryUpdated.TicketID) < 1 && ticket.Valid(requestingUser) {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid ticketID, not reporter/asignee or watcher")})
		return
	}

	QueryDB := db.DBStorage.Connection().Model("Query")

	err = QueryDB.FindId(queryUpdated.GetId()).Exec(&queryOriginal)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if requestingUser != queryOriginal.Owner.Name {
		c.JSON(http.StatusForbidden, gin.H{"error": "You must have ownership to be able to update it"})
		return
	}

	err = queryOriginal.Merge(&queryUpdated)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = queryOriginal.Save()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, queryOriginal)
}

func GetDatabases(c *gin.Context) {
	servers := models.GetDatabases(false)

	c.JSON(http.StatusOK, servers)
}

/*
 * DBworkBench
 */

package api

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/gfleury/squaas/db"
	"github.com/gfleury/squaas/models"
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

	err := c.BindJSON(&query)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !query.Status.Valid() {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid status: %s", query.Status)})
		return
	}

	if query.Status != models.StatusReady && query.Status != models.StatusPending && query.Status != models.StatusParseOnly {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Queries can only be created in Pending and Ready status, invalid status: %s", query.Status)})
		return
	}

	ownerUser := c.MustGet(gin.AuthUserKey).(string)

	query.Owner.Name = ownerUser

	log.Printf("Parsing query owner: %s\n", query.Owner)
	log.Printf("Parsing query ticketID: %s\n", query.TicketID)
	log.Printf("Parsing query status: %s\n", query.Status)
	log.Printf("Parsing query serverName: %s\n", query.ServerName)
	log.Printf("Parsing query query: %s\n", query.Query)

	err = query.LintSQLQuery()
	if err != nil {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": err.Error()})
		return
	}

	if query.Status == models.StatusParseOnly {
		c.JSON(http.StatusOK, query)
		return
	}

	if query.ServerName == "" {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "You must select a database"})
		return
	}

	// Do Ticket Validation, TODO Check Ticket existence in JIRA
	if query.TicketID == "" {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "You must insert a ticketID to link your query to"})
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
			errorString = fmt.Sprintf("%sError: %s\n", errorString, err.Error())
		}
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": errorString})
		return
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

	err := c.BindJSON(&queryUpdated)
	if err != nil {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": err.Error()})
		return
	}

	QueryDB := db.DBStorage.Connection().Model("Query")

	err = QueryDB.FindId(queryUpdated.GetId()).Exec(&queryOriginal)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	requestingUser := c.MustGet(gin.AuthUserKey).(string)

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

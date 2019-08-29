/*
 * DBworkBench
 */

package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/gfleury/dbquerybench/config"
	"github.com/gfleury/dbquerybench/db"
	"github.com/gfleury/dbquerybench/models"
)

func Index(c *gin.Context) {
	w := c.Writer

	fmt.Fprintf(w, "Hello World!")
}

func GetQueries(c *gin.Context) {
	//w := c.Writer

	// w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// w.WriteHeader(http.StatusOK)

	c.JSON(http.StatusOK, []gin.H{{"id": "123123", "name": "Some Query", "row": "AM-123", "status": "OPEN", "owner": "George", "hasTransaction": "true"}})
}

func AddQuery(c *gin.Context) {
	// w := c.Writer
	// r := c.Request

	var query models.Query

	err := c.BindJSON(&query)
	if err != nil {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Parsing query ticketID: %s\n", query.TicketID)
	log.Printf("Parsing query status: %s\n", query.Status)
	log.Printf("Parsing query serverName: %s\n", query.ServerName)
	log.Printf("Parsing query query: %s\n", query.Query)

	err = query.LintSQLQuery()
	if err != nil {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": err.Error()})
		return
	}

	if query.Status == "PARSEONLY" {
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

	if err != nil {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, query)
}

func DeleteQuery(c *gin.Context) {
	w := c.Writer
	// r := c.Request

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func ApproveQuery(c *gin.Context) {
	w := c.Writer
	// r := c.Request

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func DeleteApprovalQuery(c *gin.Context) {
	w := c.Writer
	// r := c.Request

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
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
	w := c.Writer
	// r := c.Request

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func UpdateQuery(c *gin.Context) {
	w := c.Writer
	// r := c.Request

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func GetDatabases(c *gin.Context) {
	databases := config.GetConfig().GetStringMapString("databases")
	servers := models.Servers{}

	for server := range databases {
		servers = append(servers, models.Server{Name: server})
	}

	c.JSON(http.StatusOK, servers)
}

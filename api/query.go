/*
 * DBworkBench
 */

package api

import (
	"fmt"
	"github.com/gfleury/dbquerybench/config"
	"github.com/gfleury/dbquerybench/models"
	"net/http"

	"github.com/gin-gonic/gin"
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
	w := c.Writer
	// r := c.Request

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
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

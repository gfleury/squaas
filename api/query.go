/*
 * DBworkBench
 */

package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddQuery(c *gin.Context) {
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

func DeleteQuery(c *gin.Context) {
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

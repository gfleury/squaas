/*
 * DBworkBench
 */

package api

import (
	"fmt"
	"strings"

	"github.com/gfleury/dbquerybench/auth"
	"github.com/gin-gonic/gin"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc gin.HandlerFunc
}

type Routes []Route

func NewRouter() *gin.Engine {
	m := gin.Default()

	baseAuth := m.Group("/", auth.LdapBasicAuth())

	for _, route := range routes {
		baseAuth.Handle(route.Method, route.Pattern, route.HandlerFunc)
	}

	return m
}

func Index(c *gin.Context) {
	w := c.Writer
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/v1/",
		Index,
	},

	Route{
		"AddQuery",
		strings.ToUpper("Post"),
		"/v1/query",
		AddQuery,
	},

	Route{
		"ApproveQuery",
		strings.ToUpper("Post"),
		"/v1/query/approve/{queryId}",
		ApproveQuery,
	},

	Route{
		"DeleteQuery",
		strings.ToUpper("Delete"),
		"/v1/query/approve/{queryId}",
		DeleteQuery,
	},

	Route{
		"FindQueryByOwner",
		strings.ToUpper("Get"),
		"/v1/query/findByOwner",
		FindQueryByOwner,
	},

	Route{
		"FindQueryByStatus",
		strings.ToUpper("Get"),
		"/v1/query/findByStatus",
		FindQueryByStatus,
	},

	Route{
		"GetQueryById",
		strings.ToUpper("Get"),
		"/v1/query/{queryId}",
		GetQueryById,
	},

	Route{
		"UpdateQuery",
		strings.ToUpper("Put"),
		"/v1/query",
		UpdateQuery,
	},
}

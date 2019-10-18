/*
 * DBworkBench
 */

package api

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/gfleury/squaas/auth"
	"github.com/gfleury/squaas/config"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc gin.HandlerFunc
}

type Routes []Route

func NewRouter() *gin.Engine {
	var baseAuth *gin.RouterGroup

	m := gin.Default()

	m.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "HEAD", "PUT", "PATCH", "OPTIONS", "DELETE"},
		AllowHeaders:     []string{"Origin", "authorization", "content-type"},
		ExposeHeaders:    []string{"Origin", "authorization", "content-type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	authScheme := config.GetConfig().GetString("auth.scheme")

	if authScheme == "" || authScheme == "config" {
		baseAuth = m.Group("/", auth.ConfigBasicAuth())
	} else if authScheme == "ldap" {
		baseAuth = m.Group("/", auth.LdapBasicAuth())
	}

	baseAuth.Static("/frontend", "./frontend/build")

	for _, route := range routes {
		baseAuth.Handle(route.Method, route.Pattern, route.HandlerFunc)
	}

	// HealthCheck, unauthorized endpoint
	m.Handle(strings.ToUpper("Get"), "/ping", AppHealthCheck)

	return m
}

func AppIndex(c *gin.Context) {
	c.Redirect(http.StatusTemporaryRedirect, "/frontend")
	c.Abort()
}

func AppHealthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "OK",
	})
}

var routes = Routes{
	Route{
		"Index",
		strings.ToUpper("Get"),
		"/",
		AppIndex,
	},

	Route{
		"Index",
		strings.ToUpper("Get"),
		"/v1/",
		Index,
	},

	Route{
		"GetDatabases",
		strings.ToUpper("Get"),
		"/v1/databases",
		GetDatabases,
	},

	Route{
		"GetQueries",
		strings.ToUpper("Get"),
		"/v1/query",
		GetQueries,
	},

	Route{
		"GetQueryById",
		strings.ToUpper("Get"),
		"/v1/query/:queryId",
		GetQueryById,
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
		"/v1/query/:queryId/approve",
		ApproveQuery,
	},

	Route{
		"DeleteApprovalQuery",
		strings.ToUpper("Delete"),
		"/v1/query/:queryId/approve",
		DeleteApprovalQuery,
	},

	Route{
		"DeleteQuery",
		strings.ToUpper("Delete"),
		"/v1/query/:queryId",
		DeleteQuery,
	},

	Route{
		"FindQueryByOwner",
		strings.ToUpper("Get"),
		"/v1/search/query/findByOwner",
		FindQueryByOwner,
	},

	Route{
		"FindQueryByStatus",
		strings.ToUpper("Get"),
		"/v1/search/query/findByStatus",
		FindQueryByStatus,
	},

	Route{
		"UpdateQuery",
		strings.ToUpper("Put"),
		"/v1/query",
		UpdateQuery,
	},
}

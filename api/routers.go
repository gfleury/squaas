/*
 * DBworkBench
 */

package api

import (
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/gfleury/dbquerybench/auth"
	"github.com/gfleury/dbquerybench/config"
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
		AllowMethods:     []string{"GET", "POST", "HEAD", "PUT", "PATCH", "OPTIONS"},
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

	return m
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/v1/",
		Index,
	},

	Route{
		"GetQueries",
		strings.ToUpper("Get"),
		"/v1/query",
		GetQueries,
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

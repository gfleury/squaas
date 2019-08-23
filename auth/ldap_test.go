package auth

import (
	"encoding/base64"
	"net/http"
	"net/http/httptest"
	// "github.com/gfleury/dbquerybench/config"

	"github.com/gin-gonic/gin"
	"gopkg.in/check.v1"
)

// Check ldap_bind_test.go
// Valid credentials for ldap testing:
// user:      xaiza@avoid.com
// passoword: 123456

func (s *S) TestLdapBasicAuthForRealm(c *check.C) {
	router := gin.New()
	router.Use(LdapBasicAuthForRealm(""))
	router.GET("/login", func(c *gin.Context) {
		c.String(http.StatusOK, c.MustGet(AuthUserKey).(string))
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/login", nil)
	req.SetBasicAuth("xaiza@avoid.com", "123456")
	router.ServeHTTP(w, req)

	c.Check(http.StatusOK, check.Equals, w.Code)
	c.Check("xaiza@avoid.com", check.Equals, w.Body.String())
}

func (s *S) TestLdapBasicAuthForRealmWithCache(c *check.C) {
	router := gin.New()
	router.Use(LdapBasicAuthForRealm(""))
	router.GET("/login", func(c *gin.Context) {
		c.String(http.StatusOK, c.MustGet(AuthUserKey).(string))
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/login", nil)
	req.SetBasicAuth("xaiza@avoid.com", "123456")
	router.ServeHTTP(w, req)

	c.Check(http.StatusOK, check.Equals, w.Code)
	c.Check("xaiza@avoid.com", check.Equals, w.Body.String())

	w.Body.Reset()
	w.Flush()

	req, _ = http.NewRequest("GET", "/login", nil)
	req.SetBasicAuth("xaiza@avoid.com", "123456")
	router.ServeHTTP(w, req)

	c.Check(http.StatusOK, check.Equals, w.Code)
	c.Check("xaiza@avoid.com", check.Equals, w.Body.String())
	c.Check("true", check.Equals, w.Header().Get("X-Cached-Auth"))

}

func (s *S) TestBasicAuth401(c *check.C) {
	called := false

	router := gin.New()
	router.Use(LdapBasicAuth())
	router.GET("/login", func(c *gin.Context) {
		called = true
		c.String(http.StatusOK, c.MustGet(AuthUserKey).(string))
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/login", nil)
	// Set wrong authentication data (to force a 401)
	req.SetBasicAuth("admin", "3232")
	router.ServeHTTP(w, req)

	c.Check(called, check.Equals, false)
	c.Check(http.StatusUnauthorized, check.Equals, w.Code)
	c.Check("Basic realm=\"Authorization Required\"", check.Equals, w.Header().Get("WWW-Authenticate"))
}

func (s *S) TestBasicAuth401WithCustomRealm(c *check.C) {
	called := false

	router := gin.New()
	router.Use(LdapBasicAuthForRealm("My Custom \"Realm\""))
	router.GET("/login", func(c *gin.Context) {
		called = true
		c.String(http.StatusOK, c.MustGet(AuthUserKey).(string))
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/login", nil)
	req.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte("admin:password")))
	router.ServeHTTP(w, req)

	c.Check(called, check.Equals, false)
	c.Check(http.StatusUnauthorized, check.Equals, w.Code)
	c.Check("Basic realm=\"My Custom \\\"Realm\\\"\"", check.Equals, w.Header().Get("WWW-Authenticate"))
}

package auth

import (
	"crypto/sha512"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AuthUserKey is the cookie name for user credential in basic auth.
const AuthUserKey = "user"

type cachedCredentials map[string]string

func shazify(s string) (string, error) {
	h := sha512.New()
	_, err := h.Write([]byte(s))
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

func (a cachedCredentials) searchCredential(username, cleanPassword string) (string, bool) {
	password, err := shazify(cleanPassword)

	if err != nil || username == "" || password == "" {
		return "", false
	}

	if cachedPassword, ok := a[username]; ok {
		if cachedPassword == password {
			return username, true
		}
	}

	return "", false
}

func (a cachedCredentials) addCredentialCache(username, password string) {
	hashedPassword, err := shazify(password)
	if err == nil {
		a[username] = hashedPassword
	}
}

func LdapBasicAuthForRealm(realm string) gin.HandlerFunc {
	if realm == "" {
		realm = "Authorization Required"
	}
	realm = "Basic realm=" + strconv.Quote(realm)
	pairs := cachedCredentials{}

	return func(c *gin.Context) {
		found := false
		username, password, hasAuth := c.Request.BasicAuth()

		if hasAuth {
			// Search user in the slice of cached credentials
			_, found = pairs.searchCredential(username, password)

			if !found {
				err := login(username, password)
				if err == nil {
					found = true
					pairs.addCredentialCache(username, password)
				}
			} else {
				c.Header("X-Cached-Auth", "true")
			}
		}
		if !found {
			// Credentials doesn't match, we return 401 and abort handlers chain.
			c.Header("WWW-Authenticate", realm)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// The user credentials was found, set user's id to key AuthUserKey in this context, the user's id can be read later using
		// c.MustGet(gin.AuthUserKey).
		c.Set(AuthUserKey, username)
	}
}

// BasicAuth returns a Basic HTTP Authorization middleware. It takes as argument a map[string]string where
// the key is the user name and the value is the password.
func LdapBasicAuth() gin.HandlerFunc {
	return LdapBasicAuthForRealm("")
}

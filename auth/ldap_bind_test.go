package auth

import (
	"github.com/gfleury/dbquerybench/config"

	"gopkg.in/check.v1"
)

func (s *S) TestLDAPLogin(c *check.C) {
	user := "xaiza@avoid.com"
	password := "123456"
	err := login(user, password)
	c.Assert(err, check.IsNil)
}

func (s *S) TestNativeLoginWrongPassword(c *check.C) {
	user := "xaiza@avoid.com"
	password := "xxxxxx"
	err := login(user, password)
	c.Assert(err, check.ErrorMatches, "Authentication failed, wrong password")
}

func (s *S) TestNativeLoginServerUnreachable(c *check.C) {
	config.GetConfig().Set("auth.ldap.host", "127.0.0.1")
	config.GetConfig().Set("auth.ldap.port", "1234")
	user := "xaiza@avoid.com"
	password := "xxxxxx"
	err := login(user, password)
	c.Assert(err, check.ErrorMatches, "LDAP Result Code 200 \"Network Error\": dial tcp 127.0.0.1:1234: connect: connection refused")
	config.GetConfig().Set("auth.ldap.host", "127.0.0.1")
	config.GetConfig().Set("auth.ldap.port", "10000")
}

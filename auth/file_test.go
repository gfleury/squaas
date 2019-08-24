package auth

import (
	"gopkg.in/check.v1"

	"github.com/gfleury/dbquerybench/config"
)

func (s *S) TestGetConfigCredentials(c *check.C) {
	config.GetConfig().Set("auth.config.users", map[string]string{"admin": "wassss", "admin3": "wer"})

	creds := getConfigCredentials()
	c.Check(creds, check.DeepEquals, map[string]string{"admin": "wassss", "admin3": "wer"})
}

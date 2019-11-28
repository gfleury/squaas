package models

import (
	check "gopkg.in/check.v1"

	"github.com/gfleury/squaas/config"
)

func (s *Suite) TestModelServerGetDatabases(c *check.C) {

	info := GetDatabases(true)

	c.Assert(len(info), check.Equals, 4)

}

func (s *Suite) TestModelServerGetDatabasesBrokenConfig(c *check.C) {

	// Database config used for tests
	config.GetConfig().Set("databases", map[string]interface{}{
		"server1": "postgresql://localhost:34992/database2",
		"server2": "postgresql://caixaprego:93939/database333",
		"server3": map[string]interface{}{
			"uri":           "uri-ok",
			"approval_rule": "broken_approval_rule",
		},
	})

	info := GetDatabases(true)

	c.Assert(len(info), check.Equals, 2)

}

func (s *Suite) TestModelServerGetDatabasesBrokenURIConfig(c *check.C) {

	// Database config used for tests
	config.GetConfig().Set("databases", map[string]interface{}{
		"server3": map[string]interface{}{
			"approval_rule": map[string]interface{}{},
		},
	})

	info := GetDatabases(true)

	c.Assert(len(info), check.Equals, 0)

}

func (s *Suite) TestModelServerGetDatabasesExtendedConfigOnlyURI(c *check.C) {

	// Database config used for tests
	config.GetConfig().Set("databases", map[string]interface{}{
		"server3": map[string]interface{}{
			"uri": "uri://ip:port?options",
			"AnyotheroptionsWrong": map[string]interface{}{
				"thu": "",
			},
		},
	})

	info := GetDatabases(true)

	c.Assert(len(info), check.Equals, 1)

}

func (s *Suite) TestModelServerGetDatabasesExtendedConfig(c *check.C) {

	// Database config used for tests
	config.GetConfig().Set("databases", map[string]interface{}{
		"server3": map[string]interface{}{
			"uri": "uri://ip:port?options",
			"approval_rule": map[string]interface{}{
				"required_users": []string{"admin@blah.net", "saveus@blew.org"},
			},
		},
	})

	info := GetDatabases(true)

	c.Assert(len(info), check.Equals, 1)

	c.Assert(len(info[0].ApprovalRule.RequiredUsers), check.Equals, 2)

	c.Assert(info[0].ApprovalRule.RequiredUsers, check.DeepEquals, []string{"admin@blah.net", "saveus@blew.org"})
}

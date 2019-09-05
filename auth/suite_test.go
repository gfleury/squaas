package auth

import (
	"bytes"
	"strings"
	"testing"

	"github.com/gfleury/squaas/config"

	"github.com/bradleypeabody/godap"
	"gopkg.in/check.v1"
)

func Test(t *testing.T) { check.TestingT(t) }

type S struct {
}

var _ = check.Suite(&S{})

func runLdapServer(ldapServer *godap.LDAPServer) {
	err := ldapServer.ListenAndServe("127.0.0.1:10000")
	if err != nil {
		panic(err.Error())
	}
}

func (s *S) SetUpSuite(c *check.C) {
	config.Init()
	var yamlExample = []byte(`
auth:
  scheme: ldap
  ldap:
    host: "127.0.0.1"
    port: 10000
    basedn: "dc=dddd,dc=ddddd"
    skiptls: true
    usessl: false
    skiptls: true
    sslskipverify: true
    binddn: "cn=admin,dc=dddd,dc=coxxxxm"
    bindpassword: "dddd"
`)

	err := config.GetConfig().ReadConfig(bytes.NewBuffer(yamlExample))
	c.Check(err, check.IsNil)

	// Mockup LDAP server
	hs := make([]godap.LDAPRequestHandler, 0)

	// use a LDAPBindFuncHandler to provide a callback function to respond
	// to bind requests
	hs = append(hs, &godap.LDAPBindFuncHandler{LDAPBindFunc: func(binddn string, bindpw []byte) bool {
		if strings.HasPrefix(binddn, "cn=admin,") && string(bindpw) == "blew" {
			return true
		}
		if (strings.HasPrefix(binddn, "cn=xaiza,") ||
			strings.HasPrefix(binddn, "cn=para,") ||
			strings.HasPrefix(binddn, "cn=wolverine,")) && string(bindpw) == "123456" {
			return true
		}
		return false
	}})

	// use a LDAPSimpleSearchFuncHandler to reply to search queries
	hs = append(hs, &godap.LDAPSimpleSearchFuncHandler{LDAPSimpleSearchFunc: func(req *godap.LDAPSimpleSearchRequest) []*godap.LDAPSimpleSearchResultEntry {

		ret := make([]*godap.LDAPSimpleSearchResultEntry, 0, 1)

		// here we produce a single search result that matches whatever
		// they are searching for
		if req.FilterAttr == "email" {
			ret = append(ret, &godap.LDAPSimpleSearchResultEntry{
				DN: "cn=" + strings.Split(req.FilterValue, "@")[0] + "," + req.BaseDN,
				Attrs: map[string]interface{}{
					"sn":            strings.Split(req.FilterValue, "@")[0],
					"cn":            strings.Split(req.FilterValue, "@")[0],
					"uid":           strings.Split(req.FilterValue, "@")[0],
					"email":         req.FilterValue,
					"homeDirectory": "/home/" + req.FilterValue,
					"objectClass": []string{
						"top",
						"posixAccount",
						"inetOrgPerson",
					},
				},
			})
		} else if req.FilterAttr == "memberUid" {
			ret = append(ret, &godap.LDAPSimpleSearchResultEntry{})
		}

		return ret

	}})

	ldapServer := &godap.LDAPServer{
		Handlers: hs,
	}

	go runLdapServer(ldapServer)

	config.GetConfig().Set("auth.ldap.host", "127.0.0.1")
	config.GetConfig().Set("auth.ldap.port", "10000")
	config.GetConfig().Set("auth.ldap.usessl", false)
	config.GetConfig().Set("auth.ldap.skiptls", true)
	config.GetConfig().Set("auth.ldap.basedn", "dc=mamail,dc=ltd")
	config.GetConfig().Set("auth.ldap.binddn", "cn=admin,dc=mamail,dc=ltd")
	config.GetConfig().Set("auth.ldap.bindpassword", "blew")
}

func (s *S) SetUpTest(c *check.C) {
}

func (s *S) TearDownTest(c *check.C) {
}

func (s *S) TearDownSuite(c *check.C) {
}

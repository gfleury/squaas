package auth

import (
	"fmt"
	"log"

	"github.com/gfleury/squaas/config"
	"github.com/jtblin/go-ldap-client"
	"github.com/pkg/errors"

	ldapv2 "gopkg.in/ldap.v2"
)

var (
	// LDAP Wise vars
	ldapBaseDn             string
	ldapHost               string
	ldapPort               int
	ldapUseSSL             bool
	ldapSkipTLS            bool
	ldapInsecureSkipVerify bool
	ldapServerName         string
	ldapBindDN             string
	ldapBindPassword       string
	ldapUserFilter         string
	ldapGroupFilter        string
	ldapGroupMustBePresent string
)

func loadConfig() error {

	ldapBaseDn = config.GetConfig().GetString("auth.ldap.basedn")

	if ldapHost = config.GetConfig().GetString("auth.ldap.host"); ldapHost == "" {
		return errors.Errorf("You must set LDAP authentication Hostname, in auth:ldap:host")
	}

	if ldapPort = config.GetConfig().GetInt("auth.ldap.port"); ldapPort == 0 {
		ldapPort = 389
	}

	ldapUseSSL = config.GetConfig().GetBool("auth.ldap.usessl")
	ldapSkipTLS = config.GetConfig().GetBool("auth.ldap.skiptls")
	ldapInsecureSkipVerify = config.GetConfig().GetBool("auth.ldap.sslskipverify")

	if ldapServerName = config.GetConfig().GetString("auth.ldap.servername"); ldapServerName == "" {
		ldapServerName = ldapHost
	}

	ldapBindDN = config.GetConfig().GetString("auth.ldap.binddn")

	ldapBindPassword = config.GetConfig().GetString("auth.ldap.bindpassword")

	if ldapUserFilter = config.GetConfig().GetString("auth.ldap.userfilter"); ldapUserFilter == "" {
		ldapUserFilter = "(email=%s)"
	}
	if ldapGroupFilter = config.GetConfig().GetString("auth.ldap.groupfilter"); ldapGroupFilter == "" {
		ldapGroupFilter = "(memberUid=%s)"
	}
	ldapGroupMustBePresent = config.GetConfig().GetString("auth.ldap.groupmustbepresent")

	return nil
}

func login(u, password string) (err error) {
	if u == "" {
		return errors.New("User is empty")
	}
	if err = ldapBind(u, password); err != nil {
		ldapError, isLdapError := err.(*ldapv2.Error)
		if isLdapError {
			if ldapError.ResultCode == 49 {
				return errors.New("Authentication failed, wrong password")
			}
		}
		return err
	}
	return err
}

func ldapBind(uid, password string) error {
	err := loadConfig()
	if err != nil {
		panic(fmt.Sprintf("ERROR: %v", err.Error()))
	}
	client := &ldap.LDAPClient{
		Base:               ldapBaseDn,
		Host:               ldapHost,
		Port:               ldapPort,
		UseSSL:             ldapUseSSL,
		SkipTLS:            ldapSkipTLS,
		InsecureSkipVerify: ldapInsecureSkipVerify,
		ServerName:         ldapServerName,
		BindDN:             ldapBindDN,
		BindPassword:       ldapBindPassword,
		UserFilter:         ldapUserFilter,
		GroupFilter:        ldapGroupFilter,
		Attributes:         []string{"uidNumber", "cn", "email", "uid"},
	}

	// It is the responsibility of the caller to close the connection
	defer client.Close()

	ok, user, err := client.Authenticate(uid, password)
	if err != nil {
		return err
	}
	if !ok {
		return err
	}
	log.Printf("LDAP Authenticated user: %+v", user)

	groups, err := client.GetGroupsOfUser(user["uid"])
	if err != nil {
		return err
	}
	log.Printf("LDAP Authenticated user groups: %+v", groups)
	if ldapGroupMustBePresent != "" {
		err = errors.Errorf("LDAP user %s not present member on group %s.", user["uid"], ldapGroupMustBePresent)
		for _, group := range groups {
			if group == ldapGroupMustBePresent {
				return nil
			}
		}
	}
	return err
}

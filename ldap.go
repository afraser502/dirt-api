package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/viper"
	ldap "gopkg.in/ldap.v2"
)

type Users struct {
	CN          string `json:"cn"`
	Description string `json:"description"`
}

//UserList retrieves LDAP details
func UserList(w http.ResponseWriter, r *http.Request) {
	ldapSearch()
}

//LDAP Functions
func ldapSearch() {

	bindUser := viper.GetString("ldap.username") //bind username
	bindPass := viper.GetString("ldap.password") //bind password
	ldapHost := viper.GetString("ldap.host")     // ldap host address
	ldapPort := viper.GetInt("ldap.port")        // ldap port
	ldapBaseDN := viper.GetString("ldap.baseDN") // baseDN

	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", ldapHost, ldapPort))
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	err = l.Bind(bindUser, bindPass)
	if err != nil {
		log.Fatal(err)
	}

	searchRequest := ldap.NewSearchRequest(
		ldapBaseDN, // The base dn to search
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		"(&(objectClass=*))",          // The filter to apply
		[]string{"cn", "description"}, // A list attributes to retrieve
		nil,
	)

	sr, err := l.Search(searchRequest)
	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range sr.Entries {
		fmt.Printf("%s: %v\n", entry.GetAttributeValue("cn"), entry.GetAttributeValue("description"))
	}
}

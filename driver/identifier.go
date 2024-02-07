package driver

import (
	"regexp"
	"strconv"

	"github.com/SAP/go-hdb/driver/internal/rand/alphanum"
)

var reSimple = regexp.MustCompile("^[_A-Z][_#$A-Z0-9]*$")

// Identifier in hdb SQL statements like schema or table name.
type Identifier string

// RandomIdentifier returns a random Identifier prefixed by the prefix parameter.
// This function is used to generate database objects with random names for test and example code.
func RandomIdentifier(prefix string) Identifier {
	return Identifier(prefix + alphanum.ReadString(16))
}

func (i Identifier) String() string {
	s := string(i)
	if reSimple.MatchString(s) {
		return s
	}
	return strconv.Quote(s)
}

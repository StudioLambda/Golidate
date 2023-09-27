package rule

import (
	"regexp"

	"github.com/studiolambda/golidate"
)

var DomainRegex = regexp.MustCompile(`^(?:[a-z0-9-]+\.){1,}[a-z]{2,63}$`)

func Domain() golidate.Rule {
	return Regex(DomainRegex).Rename("domain")
}

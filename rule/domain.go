package rule

import (
	"regexp"

	"github.com/studiolambda/golidate"
)

// DomainRegex matches simple lowercase domain names with at least one dot.
var DomainRegex = regexp.MustCompile(`^(?:[a-z0-9-]+\.){1,}[a-z]{2,63}$`)

// Domain returns a rule that passes for simple lowercase domain names.
//
// The rule is intentionally regex-based and does not perform DNS lookup or full
// IDNA validation.
func Domain() golidate.Rule {
	return Regex(DomainRegex).Rename("domain")
}

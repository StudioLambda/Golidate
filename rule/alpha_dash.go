package rule

import (
	"regexp"

	"github.com/studiolambda/golidate"
)

// AlphaDashRegex matches ASCII letters, digits, underscores, and hyphens.
var AlphaDashRegex = regexp.MustCompile(`^[a-zA-Z0-9_\-]*$`)

// AlphaDash returns a rule for slug-like ASCII values.
//
// The rule allows letters, digits, underscores, and hyphens. Empty strings pass
// because the underlying regular expression accepts zero characters.
func AlphaDash() golidate.Rule {
	return Regex(AlphaDashRegex).Rename("alpha_dash")
}

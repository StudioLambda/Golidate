package rule

import (
	"regexp"

	"github.com/studiolambda/golidate"
)

// AlphaExtendedRegex matches ASCII letters, digits, underscores, hyphens, and dots.
var AlphaExtendedRegex = regexp.MustCompile(`^[a-zA-Z0-9_\-\.]*$`)

// AlphaExtended returns a rule for dotted slug-like ASCII values.
//
// The rule allows letters, digits, underscores, hyphens, and dots. Empty strings
// pass because the underlying regular expression accepts zero characters.
func AlphaExtended() golidate.Rule {
	return Regex(AlphaExtendedRegex).Rename("alpha_extended")
}

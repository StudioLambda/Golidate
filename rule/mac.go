package rule

import (
	"regexp"

	"github.com/studiolambda/golidate"
)

// MacRegex matches colon-separated six-byte MAC addresses.
var MacRegex = regexp.MustCompile(`^(([a-fA-F0-9]{2}[:]){5}[a-fA-F0-9]{2})$`)

// Mac returns a rule that passes for colon-separated MAC addresses.
//
// Hex digits may be uppercase or lowercase. Other common separators are not
// accepted by the built-in expression.
func Mac() golidate.Rule {
	return Regex(MacRegex).Code("mac")
}

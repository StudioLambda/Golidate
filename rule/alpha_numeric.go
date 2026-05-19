package rule

import (
	"regexp"

	"github.com/studiolambda/golidate"
)

// AlphaNumericRegex matches strings that contain only ASCII letters and digits.
var AlphaNumericRegex = regexp.MustCompile(`^[a-zA-Z0-9]*$`)

// AlphaNumeric returns a rule that passes for ASCII letters and digits only.
//
// Empty strings pass because the underlying regular expression accepts zero
// characters.
func AlphaNumeric() golidate.Rule {
	return Regex(AlphaNumericRegex).Rename("alpha_numeric")
}

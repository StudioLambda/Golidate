package rule

import (
	"regexp"

	"github.com/studiolambda/golidate"
)

// AlphaRegex matches strings that contain only ASCII letters.
var AlphaRegex = regexp.MustCompile(`^[a-zA-Z]*$`)

// Alpha returns a rule that passes when the formatted value has only letters.
//
// The rule uses AlphaRegex and therefore allows an empty string. Non-string
// values are formatted with fmt.Sprintf by Regex before matching.
func Alpha() golidate.Rule {
	return Regex(AlphaRegex).Rename("alpha")
}

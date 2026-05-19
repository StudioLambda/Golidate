package rule

import (
	"regexp"

	"github.com/studiolambda/golidate"
)

// AsciiRegex matches strings containing only single-byte ASCII characters.
var AsciiRegex = regexp.MustCompile(`^[[:ascii:]]*$`)

// Ascii returns a rule that passes when the formatted value is all ASCII.
//
// Unicode characters outside the ASCII range fail. Empty strings pass.
func Ascii() golidate.Rule {
	return Regex(AsciiRegex).Rename("ascii")
}

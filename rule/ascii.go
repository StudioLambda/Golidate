package rule

import (
	"regexp"

	"github.com/studiolambda/golidate"
)

var AsciiRegex = regexp.MustCompile(`^[[:ascii:]]*$`)

func Ascii() golidate.Rule {
	return Regex(AsciiRegex).Code("ascii")
}

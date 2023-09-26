package rule

import (
	"regexp"

	"github.com/studiolambda/golidate"
)

var AlphaDashRegex = regexp.MustCompile(`^[a-zA-Z0-9_\-]*$`)

func AlphaDash() golidate.Rule {
	return Regex(AlphaDashRegex).Code("alpha_dash")
}

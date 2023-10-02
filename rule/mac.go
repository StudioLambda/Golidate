package rule

import (
	"regexp"

	"github.com/studiolambda/golidate"
)

var MacRegex = regexp.MustCompile(`^(([a-fA-F0-9]{2}[:]){5}[a-fA-F0-9]{2})$`)

func Mac() golidate.Rule {
	return Regex(MacRegex).Code("mac")
}

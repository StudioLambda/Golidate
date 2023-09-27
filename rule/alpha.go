package rule

import (
	"regexp"

	"github.com/studiolambda/golidate"
)

var AlphaRegex = regexp.MustCompile(`^[a-zA-Z]*$`)

func Alpha() golidate.Rule {
	return Regex(AlphaRegex).Rename("alpha")
}

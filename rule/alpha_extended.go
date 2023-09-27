package rule

import (
	"regexp"

	"github.com/studiolambda/golidate"
)

var AlphaExtendedRegex = regexp.MustCompile(`^[a-zA-Z0-9_\-\.]*$`)

func AlphaExtended() golidate.Rule {
	return Regex(AlphaExtendedRegex).Rename("alpha_extended")
}

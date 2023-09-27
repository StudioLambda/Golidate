package rule

import (
	"regexp"

	"github.com/studiolambda/golidate"
)

var AlphaNumericRegex = regexp.MustCompile(`^[a-zA-Z0-9]*$`)

func AlphaNumeric() golidate.Rule {
	return Regex(AlphaNumericRegex).Rename("alpha_numeric")
}

package rule

import (
	"regexp"

	"github.com/studiolambda/golidate"
)

var EmailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

func Email() golidate.Rule {
	return Regex(EmailRegex).Rename("email")
}

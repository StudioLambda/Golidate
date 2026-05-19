package rule

import (
	"regexp"

	"github.com/studiolambda/golidate"
)

// EmailRegex matches a pragmatic ASCII email address shape.
var EmailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

// Email returns a rule that passes for a pragmatic email address shape.
//
// The rule is regex-based and does not attempt full RFC mailbox validation.
func Email() golidate.Rule {
	return Regex(EmailRegex).Rename("email")
}

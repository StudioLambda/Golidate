package rule

import (
	"regexp"

	"github.com/studiolambda/golidate"
)

// UrlRegex matches a pragmatic URL or hostname-like string.
var UrlRegex = regexp.MustCompile(`^([a-zA-Z]+:\/\/.)?[-a-zA-Z0-9@:%._\+~#=]{2,256}\.[a-z]{2,63}\b([-a-zA-Z0-9@:%_\+.~#?&//=]*)$`)

// Url returns a rule that passes for a pragmatic URL shape.
//
// The scheme is optional, but the value must contain a dotted host-like segment
// with a lowercase top-level domain.
func Url() golidate.Rule {
	return Regex(UrlRegex).Rename("url")
}

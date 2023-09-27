package rule

import (
	"regexp"

	"github.com/studiolambda/golidate"
)

var UrlRegex = regexp.MustCompile(`^([a-zA-Z]+:\/\/.)?[-a-zA-Z0-9@:%._\+~#=]{2,256}\.[a-z]{2,63}\b([-a-zA-Z0-9@:%_\+.~#?&//=]*)$`)

func Url() golidate.Rule {
	return Regex(UrlRegex).Rename("url")
}

package rule

import "github.com/studiolambda/golidate"

func Optional(rule golidate.Rule) golidate.Rule {
	return Or(Nil(), rule).Rename("optional")
}

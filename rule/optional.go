package rule

import "github.com/studiolambda/golidate"

// Optional returns a rule that passes for nil or for values accepted by rule.
//
// The composed rule is implemented as Or(Nil(), rule) and renamed to
// "optional". Non-nil values still need to satisfy the supplied rule.
func Optional(rule golidate.Rule) golidate.Rule {
	return Or(Nil(), rule).Rename("optional")
}

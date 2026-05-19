package rule

import "github.com/studiolambda/golidate"

// Required returns a rule that requires a non-nil value and then applies rule.
//
// The composed rule is implemented as And(Not(Nil()), rule) and renamed to
// "required" so translation can present a single required-field message.
func Required(rule golidate.Rule) golidate.Rule {
	return And(Not(Nil()), rule).Rename("required")
}

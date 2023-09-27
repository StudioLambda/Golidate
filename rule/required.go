package rule

import "github.com/studiolambda/golidate"

func Required(rule golidate.Rule) golidate.Rule {
	return And(Not(Nil()), rule).Rename("required")
}

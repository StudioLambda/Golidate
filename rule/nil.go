package rule

import (
	"github.com/studiolambda/golidate"
)

func Nil() golidate.Rule {
	return func(value any) golidate.Result {
		result := golidate.Uncertain(value, "nil")

		if value != nil {
			return result.Fail()
		}

		return result.Pass()
	}
}

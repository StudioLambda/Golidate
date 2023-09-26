package rule

import (
	"github.com/studiolambda/golidate"
)

func Not(rule golidate.Rule) golidate.Rule {
	return func(value any) golidate.Result {
		operation := rule(value)
		result := golidate.
			Uncertain(value, "not").
			With("operation", operation).
			OnRename(golidate.OnRenameSingle("operation"))

		if operation.Passes() {
			return result.Fail()
		}

		return result.Pass()
	}
}

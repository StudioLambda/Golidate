package rule

import "github.com/studiolambda/golidate"

func And(rules ...golidate.Rule) golidate.Rule {
	return func(value any) golidate.Result {
		operations := make(golidate.Results, len(rules))

		for i, rule := range rules {
			operations[i] = rule(value)
		}

		result := golidate.
			Uncertain(value, "and").
			With("operations", operations).
			OnRename(golidate.OnRenameMany("operations"))

		if operations.PassesAll() {
			return result.Pass()
		}

		return result.Fail()
	}
}

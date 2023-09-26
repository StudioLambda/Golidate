package rule

import "github.com/studiolambda/golidate"

func Or(rules ...golidate.Rule) golidate.Rule {
	return func(value any) golidate.Result {
		operations := golidate.Results{}

		for _, rule := range rules {
			operations = append(operations, rule(value))
		}

		result := golidate.
			Uncertain(value, "or").
			With("operations", operations).
			OnRename(golidate.OnRenameMany("operations"))

		if operations.PassesAny() {
			return result.Pass()
		}

		return result.Fail()
	}
}

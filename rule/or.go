package rule

import "github.com/studiolambda/golidate"

// Or returns a rule that passes when any child rule passes.
//
// Child rules are evaluated in order and evaluation stops at the first passing
// result. The evaluated child results are stored in metadata as "operations".
func Or(rules ...golidate.Rule) golidate.Rule {
	return func(value any) golidate.Result {
		operations := make(golidate.Results, 0, len(rules))

		result := golidate.
			Uncertain(value, "or").
			OnRename(golidate.OnRenameMany("operations"))

		for _, rule := range rules {
			current := rule(value)
			operations = append(operations, current)

			if current.PassesAll() {
				return result.
					With("operations", operations).
					Pass()
			}
		}

		return result.
			With("operations", operations).
			Fail()
	}
}

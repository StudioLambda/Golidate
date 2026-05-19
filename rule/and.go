package rule

import (
	"github.com/studiolambda/golidate"
)

// And returns a rule that passes only when every child rule passes.
//
// Child rules are evaluated in order and evaluation stops at the first failure.
// The evaluated child results are stored in metadata as "operations" for
// translators and for callers that want to inspect the failing branch.
func And(rules ...golidate.Rule) golidate.Rule {
	return func(value any) golidate.Result {
		operations := make(golidate.Results, 0, len(rules))
		result := golidate.
			Uncertain(value, "and").
			OnRename(golidate.OnRenameMany("operations"))

		for _, rule := range rules {
			current := rule(value)
			operations = append(operations, current)

			if !current.PassesAll() {
				return result.
					With("operations", operations).
					Fail()
			}
		}

		return result.
			With("operations", operations).
			Pass()
	}
}

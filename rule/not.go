package rule

import (
	"github.com/studiolambda/golidate"
)

// Not returns a rule that inverts another rule's direct pass state.
//
// The wrapped operation result is stored in metadata as "operation". Translation
// dictionaries can use that nested operation to build a human-readable negated
// message.
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

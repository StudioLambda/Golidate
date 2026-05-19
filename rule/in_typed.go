package rule

import (
	"github.com/studiolambda/golidate"
)

// InTyped returns a rule that passes when value is one of the typed values.
//
// The input must type-assert to T before comparable equality is checked. This is
// stricter than numeric conversion and avoids accidental matches across types.
func InTyped[T comparable](values ...T) golidate.Rule {
	return func(value any) golidate.Result {
		result := golidate.
			Uncertain(value, "in_typed").
			With("values", values)

		val, ok := value.(T)
		if !ok {
			return result.Fail()
		}

		for _, v := range values {
			if val == v {
				return result.Pass()
			}
		}

		return result.Fail()
	}
}

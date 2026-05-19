package rule

import "github.com/studiolambda/golidate"

// MaxLen returns a rule that passes when a value length is at most max.
//
// Arrays, channels, maps, slices, and strings are supported. Nil and unsupported
// values fail without panic recovery.
func MaxLen(max int) golidate.Rule {
	return func(value any) golidate.Result {
		result := golidate.
			Uncertain(value, "max_len").
			With("max", max)

		length, ok := lengthOf(value)
		if !ok || length > max {
			return result.Fail()
		}

		return result.Pass()
	}
}

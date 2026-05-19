package rule

import "github.com/studiolambda/golidate"

// MinLen returns a rule that passes when a value length is at least min.
//
// Arrays, channels, maps, slices, and strings are supported. Nil and unsupported
// values fail without panic recovery.
func MinLen(min int) golidate.Rule {
	return func(value any) golidate.Result {
		result := golidate.
			Uncertain(value, "min_len").
			With("min", min)

		length, ok := lengthOf(value)
		if !ok || length < min {
			return result.Fail()
		}

		return result.Pass()
	}
}

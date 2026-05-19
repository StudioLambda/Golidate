package rule

import "github.com/studiolambda/golidate"

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

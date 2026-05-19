package rule

import "github.com/studiolambda/golidate"

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

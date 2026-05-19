package rule

import "github.com/studiolambda/golidate"

func Max(max int64) golidate.Rule {
	return func(value any) golidate.Result {
		result := golidate.
			Uncertain(value, "max").
			With("max", max)

		val, ok := numberValue(value)

		if !ok || val > float64(max) {
			return result.Fail()
		}

		return result.Pass()
	}
}

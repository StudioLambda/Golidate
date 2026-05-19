package rule

import "github.com/studiolambda/golidate"

// Max returns a rule that passes for numeric values less than or equal to max.
//
// Signed integers, unsigned integers, and floats are accepted. Values are
// compared as float64 so decimal inputs can satisfy integer limits. Non-numeric
// values and nil fail without panicking.
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

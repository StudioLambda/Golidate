package rule

import (
	"github.com/spf13/cast"
	"github.com/studiolambda/golidate"
)

func Min(min int64) golidate.Rule {
	return func(value any) golidate.Result {
		result := golidate.
			Uncertain(value, "min").
			With("min", min)

		val, err := cast.ToInt64E(value)

		if err != nil || val < min {
			return result.Fail()
		}

		return result.Pass()
	}
}

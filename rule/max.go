package rule

import (
	"github.com/spf13/cast"
	"github.com/studiolambda/golidate"
)

func Max(max int64) golidate.Rule {
	return func(value any) golidate.Result {
		result := golidate.
			Uncertain(value, "max").
			With("max", max)

		val, err := cast.ToInt64E(value)

		if err != nil || val > max {
			return result.Fail()
		}

		return result.Pass()
	}
}

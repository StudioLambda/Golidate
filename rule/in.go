package rule

import (
	"github.com/spf13/cast"
	"github.com/studiolambda/golidate"
)

func In(values ...any) golidate.Rule {
	return func(value any) golidate.Result {
		result := golidate.
			Uncertain(value, "in").
			With("values", values)

		val, err := cast.ToStringE(value)

		if err != nil {
			return result.Fail()
		}

		for _, v := range values {
			current, err := cast.ToStringE(v)

			if err != nil {
				return result.Fail()
			}

			if val == current {
				return result.Pass()
			}
		}

		return result.Fail()
	}
}

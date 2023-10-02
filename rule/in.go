package rule

import (
	"fmt"

	"github.com/studiolambda/golidate"
)

func In(values ...any) golidate.Rule {
	return func(value any) golidate.Result {
		result := golidate.
			Uncertain(value, "in").
			With("values", values)

		val := fmt.Sprintf("%+v", value)

		for _, v := range values {
			if val == fmt.Sprintf("%+v", v) {
				return result.Pass()
			}
		}

		return result.Fail()
	}
}
